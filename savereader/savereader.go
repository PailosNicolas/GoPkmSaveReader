package savereader

import (
	"fmt"
	"os"

	"github.com/PailosNicolas/GoPkmSaveReader/helpers"
)

type Save struct {
	saveRaw [57344]byte
	Trainer Trainer
}

func ReadDataFromSave(path string) (Save, error) {
	var primarySave [57344]byte
	var save Save
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return Save{}, err
	}
	defer file.Close()

	buffer := make([]byte, 131072)

	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return Save{}, err
	}

	// First 8kb contains primary and backup save
	saveA := buffer[:57344]
	saveB := buffer[57344 : 57344*2]

	if saveA[0] < saveB[0] {
		copy(primarySave[:], saveB)
	} else if saveA[0] > saveB[0] {
		copy(primarySave[:], saveA)
	} else {
		if saveA[1] < saveB[1] {
			copy(primarySave[:], saveB)
		} else {
			copy(primarySave[:], saveA)
		}
	}

	save.saveRaw = primarySave

	save.Trainer.Name = helpers.ReadString(primarySave[:7])

	return save, nil
}

type Trainer struct {
	Name string
}
