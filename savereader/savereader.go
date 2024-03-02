package savereader

import (
	"fmt"
	"os"

	"github.com/PailosNicolas/GoPkmSaveReader/helpers"
)

type Save struct {
	Trainer Trainer
}

func ReadDataFromSave(path string) (Save, error) {
	var primarySave []byte
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
		primarySave = saveB
	} else if saveA[0] > saveB[0] {
		primarySave = saveA
	} else {
		if saveA[1] < saveB[1] {
			primarySave = saveB
		} else {
			primarySave = saveA
		}
	}

	return Save{
		Trainer: Trainer{
			Name: helpers.ReadString(primarySave[:7]),
		},
	}, nil
}

type Trainer struct {
	Name string
}
