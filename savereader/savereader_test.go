package savereader

import (
	"testing"
)

func TestReadDataFromSave(t *testing.T) {

	// Happy Path
	path := "../testfiles/testfile"

	save, err := ReadDataFromSave(path)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if save.Trainer.Name() != "Fry" {
		t.Error("name should be Fry but is:", save.Trainer.Name())
	}

	if save.Trainer.Gender() != "Boy" {
		t.Error("gender should be Boy but is:", save.Trainer.Gender())
	}

	// Error Short file
	pathShort := "../testfiles/Lola.pkm"

	_, err = ReadDataFromSave(pathShort)

	if err != ErrShortFile {
		t.Fatalf("it shold have been ErrShortFile")
	}

}
