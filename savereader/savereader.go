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
	saveRaw  [57344]byte
	Trainer  Trainer
	Game     string
	GameCode int
}

type Trainer struct {
	Name     string
	Gender   string
	PublicID int
	SecretID int
	TeamSize int
	Team     [6]pokemon.Pokemon
}

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
		return Save{}, errors.New("file to short to be a save")
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
		return Save{}, errors.New("unable to read file")
	}
	save.Trainer.Name = helpers.ReadString(sections[0].Contents[:7])

	// Getting trainer's gender
	if gender := sections[0].Contents[8:9]; bytes.Equal(gender, []byte{0}) {
		save.Trainer.Gender = "Boy"
	} else {
		save.Trainer.Gender = "Girl"
	}

	// getting gamecode
	if code := int(sections[0].Contents[172]); code == 0 {
		save.Game = "R/S"
		save.GameCode = code
	} else if code == 1 {
		save.Game = "FR/LG"
		save.GameCode = code
	} else {
		save.Game = "E"
		save.GameCode = code
	}

	//getting trainer's ids
	save.Trainer.PublicID = int(binary.LittleEndian.Uint16(sections[0].Contents[10:12]))
	save.Trainer.SecretID = int(binary.LittleEndian.Uint16(sections[0].Contents[12:14]))

	// getting team size:
	if save.GameCode != 1 {
		save.Trainer.TeamSize = int(binary.LittleEndian.Uint16(sections[1].Contents[564:568]))
	} else {
		save.Trainer.TeamSize = int(binary.LittleEndian.Uint16(sections[1].Contents[52:56]))
	}

	// getting team
	for i := range save.Trainer.TeamSize {
		var pkm pokemon.Pokemon
		var start int
		if save.GameCode != 1 {
			start = 568 + 100*i
		} else {
			start = 56 + 100*i
		}

		pkm = pokemon.ParsePokemon(sections[1].Contents[start : start+100])

		save.Trainer.Team[i] = pkm
	}

	return save, nil
}
