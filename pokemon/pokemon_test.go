package pokemon

import (
	"testing"
)

func TestReadDataFromSave(t *testing.T) {

	// Happy path
	path := "../testfiles/Lola.pkm"

	pkm, err := ReadPokemonFromFile(path)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if pkm.Nickname != "Lola" {
		t.Error("nickname should be Lola but is:", pkm.Nickname)
	}

	// File does not exist
	pathNonExist := "../testfiles/thisfiledoesnotexist"

	_, err = ReadPokemonFromFile(pathNonExist)

	if err == nil {
		t.Fatalf("file does not exist it shold error out")
	}

	// File too big
	pathTooShort := "../testfiles/dummyfile"

	_, err = ReadPokemonFromFile(pathTooShort)

	if err != ErrFileShort {
		t.Fatalf("it shold have been ErrFileShort")
	}
}
