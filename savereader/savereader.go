package savereader

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"os"

	"github.com/PailosNicolas/GoPkmSaveReader/helpers"
	"github.com/PailosNicolas/GoPkmSaveReader/pokemon"
)

type Save struct {
	saveRaw    [57344]byte
	Trainer    Trainer
	game       string
	gameCode   int
	TimePlayed TimePlayed
}

func (s *Save) Game() string {
	return s.game
}

func (s *Save) GameCode() int {
	return s.gameCode
}

type Trainer struct {
	name     string
	gender   string
	publicID int
	secretID int
	teamSize int
	team     [6]pokemon.Pokemon
	money    int
}

func (t *Trainer) Name() string {
	return t.name
}

func (t *Trainer) Gender() string {
	return t.gender
}

func (t *Trainer) PublicID() int {
	return t.publicID
}

func (t *Trainer) SecretID() int {
	return t.secretID
}

func (t *Trainer) TeamSize() int {
	return t.teamSize
}

func (t *Trainer) Team() [6]pokemon.Pokemon {
	return t.team
}

func (t *Trainer) Money() int {
	return t.money
}

/* Represents the time played in a save */
type TimePlayed struct {
	hours   int
	minutes int
	seconds int
	frames  int
}

func (t *TimePlayed) Hours() int {
	return t.hours
}

func (t *TimePlayed) Minutes() int {
	return t.minutes
}

func (t *TimePlayed) Seconds() int {
	return t.seconds
}

func (t *TimePlayed) Frames() int {
	return t.frames
}

// Errors
var ErrShortFile = errors.New("file to short to be a save")
var ErrReadingFile = errors.New("unable to read file")

/*
Reads the save file in the path and returns a Save file with Trainer info.
Currenly only supports Gen3
*/
func ReadDataFromSave(path string) (Save, error) {
	var primarySave [57344]byte
	var sections map[int]helpers.Section
	var save Save
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return Save{}, err
	}
	defer file.Close()

	buffer := make([]byte, 131072)

	bytesRead, err := file.Read(buffer)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return Save{}, err
	}

	if bytesRead < 131072 {
		return Save{}, ErrShortFile
	}

	// First 8kb contains primary and backup save
	saveA := buffer[:57344]
	saveB := buffer[57344 : 57344*2]

	if sectionsA := helpers.CreateSectionsMap(saveA); len(sectionsA) == 14 {
		copy(primarySave[:], saveA)
		sections = sectionsA
	}
	if sectionsB := helpers.CreateSectionsMap(saveB); len(sectionsB) == 14 {
		if len(sections) == 14 {
			if sections[0].Index < sectionsB[0].Index {
				copy(primarySave[:], saveB)
				sections = sectionsB
			}
		} else {
			copy(primarySave[:], saveB)
			sections = sectionsB
		}
	}

	save.saveRaw = primarySave

	// Getting trainer's name
	if len(sections[0].Contents) < 7 {
		return Save{}, ErrReadingFile
	}
	save.Trainer.name = helpers.ReadString(sections[0].Contents[:7])

	// Getting trainer's gender
	if gender := sections[0].Contents[8:9]; bytes.Equal(gender, []byte{0}) {
		save.Trainer.gender = "Boy"
	} else {
		save.Trainer.gender = "Girl"
	}

	// getting gamecode
	if code := int(sections[0].Contents[172]); code == 0 {
		save.game = "R/S"
		save.gameCode = code
		//getting money
		save.Trainer.money = int(binary.LittleEndian.Uint32(sections[1].Contents[1168:1172]))
	} else if code == 1 {
		save.game = "FR/LG"
		save.gameCode = code
		//getting money
		save.Trainer.money = int(binary.LittleEndian.Uint32(helpers.XorBytes(sections[1].Contents[656:660], sections[0].Contents[2808:2812])))
	} else {
		save.game = "E"
		save.gameCode = code
		//getting money
		save.Trainer.money = int(binary.LittleEndian.Uint32(helpers.XorBytes(sections[1].Contents[1168:1172], sections[0].Contents[172:176])))
	}

	//getting time played
	save.TimePlayed.hours, save.TimePlayed.minutes, save.TimePlayed.seconds, save.TimePlayed.frames = parseTimePlayed(sections[0].Contents[14:19])

	//getting trainer's ids
	save.Trainer.publicID = int(binary.LittleEndian.Uint16(sections[0].Contents[10:12]))
	save.Trainer.secretID = int(binary.LittleEndian.Uint16(sections[0].Contents[12:14]))

	// getting team size:
	if save.gameCode != 1 {
		save.Trainer.teamSize = int(binary.LittleEndian.Uint16(sections[1].Contents[564:568]))
	} else {
		save.Trainer.teamSize = int(binary.LittleEndian.Uint16(sections[1].Contents[52:56]))
	}

	// getting team
	for i := range save.Trainer.teamSize {
		var pkm pokemon.Pokemon
		var start int
		if save.gameCode != 1 {
			start = 568 + 100*i
		} else {
			start = 56 + 100*i
		}

		pkm, err = pokemon.ParsePokemon(sections[1].Contents[start : start+100])

		if err != nil {
			return Save{}, err
		}

		save.Trainer.team[i] = pkm
	}

	return save, nil
}

func parseTimePlayed(bytes []byte) (int, int, int, int) {
	bits := helpers.BytesToBits(bytes)

	hours := 0
	for i := 0; i < 16; i++ {
		hours = (hours << 1) | bits[i]
	}

	minutes := 0
	for i := 16; i < 24; i++ {
		minutes = (minutes << 1) | bits[i]
	}

	seconds := 0
	for i := 24; i < 32; i++ {
		seconds = (seconds << 1) | bits[i]
	}

	frames := 0
	for i := 32; i < 40; i++ {
		frames = (frames << 1) | bits[i]
	}

	return hours, minutes, seconds, frames
}
