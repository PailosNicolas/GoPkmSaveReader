package pokemon

import (
	"testing"
)

func TestParsePokemons(t *testing.T) {

	// Happy path
	path := "../testfiles/Lola.pkm"

	pkm, err := ReadPokemonFromFile(path)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if pkm.Nickname != "Lola" {
		t.Error("nickname should be Lola but is:", pkm.Nickname)
	}

	if pkm.IsValid() != true {
		t.Error("The pokemon should be valid")
	}

	// Pkm is not valid
	pkm.raw[28] = uint8(12)
	_, err = ParsePokemon(pkm.raw)
	if err != ErrPkmNotValid {
		t.Fatalf("it should be invalid")
	}

	// File does not exist
	pathNonExist := "../testfiles/thisfiledoesnotexist"

	_, err = ReadPokemonFromFile(pathNonExist)

	if err == nil {
		t.Fatalf("file does not exist it shold error out")
	}

	// File too short
	pathTooShort := "../testfiles/dummyfile"

	_, err = ReadPokemonFromFile(pathTooShort)

	if err != ErrFileShort {
		t.Fatalf("it shold have been ErrFileShort")
	}
}
