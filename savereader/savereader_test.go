package savereader

import (
	"testing"
)

func TestReadDataFromSave(t *testing.T) {

	// Happy Path
	path := "../testfiles/testfile"

	save, err := ReadDataFromSave(path)

	if err != nil {
		t.Fatal(err.Error())
	}

	if save.Trainer.Name() != "Fry" {
		t.Error("name should be Fry but is:", save.Trainer.Name())
	}

	if save.Trainer.Gender() != "Boy" {
		t.Error("gender should be Boy but is:", save.Trainer.Gender())
	}

	// Replace pokemon in team
	save, err = save.ReplacePokemonInTeam(save.Trainer.Team()[0], 5)

	if err != nil {
		t.Error("Error replacing pokemon in team")
	}

	firstTeam := save.Trainer.Team()[0]

	fifthTeam := save.Trainer.Team()[4]

	if firstTeam.PersonalityValue() != fifthTeam.PersonalityValue() {
		t.Error("Error replacing pokemon in team")
	}

	// Replace pokemon in team wrong index
	save, err = save.ReplacePokemonInTeam(save.Trainer.Team()[0], 20)

	if err != ErrIncorrectIndex {
		t.Error("I should error because of wrong index")
	}

	// Replace pokemon in team with boxed
	b, _ := save.Trainer.pc.Box(1)
	save, err = save.ReplacePokemonInTeam(b.BoxedPokemon()[0], 5)

	if err != ErrIncorrectLenght {
		t.Error("I should error because of boxed pokemon")
	}

	// Replace pokemon in box with team member
	save, err = save.ReplacePokemonInPC(firstTeam, 3)

	if err != nil {
		t.Error("Error replacing pokemon in box")
	}

	pc := save.Trainer.Pc()

	box, _ := pc.Box(1)

	pkmList := box.BoxedPokemon()

	if firstTeam.PersonalityValue() != pkmList[2].PersonalityValue() {
		t.Error("Error replacing pokemon in box")
	}

	// Error Short file
	pathShort := "../testfiles/Lola.pkm"

	_, err = ReadDataFromSave(pathShort)

	if err != ErrShortFile {
		t.Fatalf("it shold have been ErrShortFile")
	}

}
