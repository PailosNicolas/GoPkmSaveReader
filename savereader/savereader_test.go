package savereader

import (
	"testing"
)

func TestReadDataFromSave(t *testing.T) {
	path := "../testfiles/testfile"

	save, err := ReadDataFromSave(path)

	if err != nil {
		t.Fatalf("could not load testfile")
	}

	if save.Trainer.Name != "Fry" {
		t.Error("name should be Fry but got:", save.Trainer.Name)
	}

}