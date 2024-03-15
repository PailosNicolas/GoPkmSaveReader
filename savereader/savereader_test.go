package savereader

import (
	"testing"
)

func TestReadDataFromSave(t *testing.T) {
	path := "../testfiles/testfile"

	save, err := ReadDataFromSave(path)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if save.Trainer.Name != "Fry" {
		t.Error("name should be Fry but is:", save.Trainer.Name)
	}

	if save.Trainer.Gender != "Boy" {
		t.Error("gender should be Boy but is:", save.Trainer.Gender)
	}

}
