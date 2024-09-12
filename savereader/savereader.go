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

// Errors
var ErrShortFile = errors.New("file to short to be a save")
var ErrReadingFile = errors.New("unable to read file")
var ErrReadingPokemonFromBox = errors.New("unable to read pokemon from box")
var ErrIdNotAllowed = errors.New("id must be 1~15")
var ErrIncorrectLenght = errors.New("the raw of the pokemon is not the right lenght")

type Save struct {
	saveRaw    [57344]byte
	Trainer    Trainer
	game       string
	gameCode   int
	TimePlayed TimePlayed
	primaryA   bool
	fullRaw    []byte
	sections   map[int]helpers.Section
}

func (s *Save) Game() string {
	return s.game
}

func (s *Save) GameCode() int {
	return s.gameCode
}

/*
Replaces a team member with the given pokemon (It can not be a boxed one, validate raw lenght of 100).
Index are 1-6, and returns a new Save object with the change, in case of error returns unmodified save.
*/
func (s *Save) ReplacePokemonInTeam(pkm pokemon.Pokemon, teamIndex int) (Save, error) {
	var start int
	var saveStart int
	var startingSave Save = *s

	if len(pkm.Raw()) < 100 { // TODO: Maybe look for a reconstruction of the boxed pokemon
		return startingSave, ErrIncorrectLenght
	}

	//Getting team member section index
	if s.gameCode != 1 {
		start = 568 + 100*(teamIndex-1)
	} else {
		start = 56 + 100*(teamIndex-1)
	}

	//Getting raw save index
	if s.primaryA {
		saveStart = 0
	} else {
		saveStart = 57344
	}

	finalIndex := saveStart + s.sections[1].SectionIndex + start
	checksumIndex := saveStart + s.sections[1].SectionIndex + 4086

	copy(s.fullRaw[finalIndex:finalIndex+100], pkm.Raw())

	newCheckSum := helpers.CalculateChecksum(s.fullRaw[s.sections[1].SectionIndex : s.sections[1].SectionIndex+3968])

	newCheckSumBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(newCheckSumBytes, newCheckSum)

	copy(s.fullRaw[checksumIndex:checksumIndex+2], newCheckSumBytes)

	sAux, err := ReadDataFromMemory(s.fullRaw)

	if err != nil {
		return startingSave, err
	}

	return sAux, nil

}

/*
Replaces a pokemon in PC with the given pokemon.
Index are 1-420, and returns a new Save object with the change, in case of error returns unmodified save.
*/
func (s *Save) ReplacePokemonInPC(pkm pokemon.Pokemon, pcIndex int) (Save, error) {
	var start int
	var saveStart int
	var startingSave Save = *s

	startingSection := 5

	// Getting raw save index
	if s.primaryA {
		saveStart = 0
	} else {
		saveStart = 57344
	}

	sectionPosition := ((80 * (pcIndex - 1)) + 4) / 3968

	startingSection += sectionPosition

	positionAdjustment := ((80 * (pcIndex - 1)) + 4) - (3968 * sectionPosition)

	start = s.sections[startingSection].SectionIndex + positionAdjustment

	checksumIndex := saveStart + s.sections[startingSection].SectionIndex + 4086

	raw := pkm.Raw()

	if len(raw) == 100 {
		copy(s.fullRaw[start:start+80], raw[:80])
	} else if len(raw) == 80 {
		copy(s.fullRaw[start:start+80], raw)
	} else {
		return *s, ErrIncorrectLenght
	}

	newCheckSum := helpers.CalculateChecksum(s.fullRaw[s.sections[startingSection].SectionIndex : s.sections[startingSection].SectionIndex+3968])

	newCheckSumBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(newCheckSumBytes, newCheckSum)

	copy(s.fullRaw[checksumIndex:checksumIndex+2], newCheckSumBytes)

	sAux, err := ReadDataFromMemory(s.fullRaw)

	if err != nil {
		return startingSave, err
	}

	return sAux, nil
}

type Trainer struct {
	name     string
	gender   string
	publicID int
	secretID int
	teamSize int
	team     [6]pokemon.Pokemon
	money    int
	pc       PC
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

func (t *Trainer) Pc() PC {
	return t.pc
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

type Box struct {
	id          int
	name        string
	pokemonList [30]pokemon.Pokemon
}

func (b *Box) BoxedPokemon() [30]pokemon.Pokemon {
	return b.pokemonList
}

type PC struct {
	boxes [14]Box
}

/*
Returns the box select by id (1~15)
*/
func (pc *PC) Box(id int) (Box, error) {
	if id < 1 || id > 15 {
		return Box{}, ErrIdNotAllowed
	} else {
		return pc.boxes[id-1], nil
	}
}

/*
Reads the save file in the path and returns a Save file with Trainer info.
Currenly only supports Gen3
*/
func ReadDataFromSave(path string) (Save, error) {
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

	return ReadDataFromMemory(buffer)

}

/*
Reads the save file from a []byte in memory and returns a Save file with Trainer info.
Currenly only supports Gen3
*/
func ReadDataFromMemory(buffer []byte) (Save, error) {
	var primarySave [57344]byte
	var sections map[int]helpers.Section
	var save Save

	if len(buffer) < 131072 {
		return Save{}, ErrShortFile
	}

	save.fullRaw = make([]byte, len(buffer))
	copy(save.fullRaw, buffer)

	// First 8kb contains primary and backup save
	saveA := buffer[:57344]
	saveB := buffer[57344 : 57344*2]

	if sectionsA := helpers.CreateSectionsMap(saveA); len(sectionsA) == 14 {
		save.primaryA = true
		copy(primarySave[:], saveA)
		sections = sectionsA
	}
	if sectionsB := helpers.CreateSectionsMap(saveB); len(sectionsB) == 14 {
		if len(sections) == 14 {
			if sections[0].Index < sectionsB[0].Index {
				copy(primarySave[:], saveB)
				sections = sectionsB
				save.primaryA = false
			}
		} else {
			copy(primarySave[:], saveB)
			sections = sectionsB
			save.primaryA = false
		}
	}

	save.saveRaw = primarySave
	save.sections = sections

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

		pkm, err := pokemon.ParsePokemon(sections[1].Contents[start : start+100])

		if err != nil {
			return Save{}, err
		}

		save.Trainer.team[i] = pkm
	}

	//PC
	pcBoxedList := sections[5].Contents[4:]

	for i := 6; i <= 14; i++ {
		pcBoxedList = append(pcBoxedList, sections[i].Contents...)
	}

	nameBytes := pcBoxedList[33600:33726]

	for i := 0; i <= 13; i++ { // TODO: Analyze a paralel implementation
		var err error
		save.Trainer.pc.boxes[i].pokemonList, err = makeBoxList(pcBoxedList[2400*i : 2400*(i+1)])
		if err != nil {
			return Save{}, err
		}
		save.Trainer.pc.boxes[i].id = i + 1
		save.Trainer.pc.boxes[i].name = helpers.ReadString(nameBytes[9*i : 9*(i+1)])
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

func makeBoxList(data []byte) ([30]pokemon.Pokemon, error) {
	pkmList := []pokemon.Pokemon{}
	for i := 0; i <= 2320; i += 80 {
		if int(binary.LittleEndian.Uint32(data[i:i+4])) == 0 {
			// Empty space in box
			pkmList = append(pkmList, pokemon.Pokemon{})
		} else {
			pkm, err := pokemon.ParsePokemon(data[i : i+80])

			if err != nil {
				return [30]pokemon.Pokemon{}, ErrReadingPokemonFromBox
			}

			pkmList = append(pkmList, pkm)
		}

	}

	return [30]pokemon.Pokemon(pkmList), nil
}
