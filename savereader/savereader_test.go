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

	trainerBag := save.Trainer.Bag()
	itemTests := []struct {
		index            int
		expectedName     string
		expectedQuantity int
	}{
		{
			0,
			"Potion",
			8,
		},
		{
			1,
			"Repel",
			4,
		},
		{
			2,
			"Escape Rope",
			2,
		},
		{
			3,
			"Super Potion",
			14,
		},
		{
			4,
			"X Attack",
			4,
		},
		{
			5,
			"Parlyz Heal",
			1,
		},
		{
			6,
			"Ether",
			2,
		},
		{
			7,
			"Antidote",
			3,
		},
	}

	itemPocket := trainerBag.ItemsPocket()
	for _, test := range itemTests {
		item := itemPocket[test.index]
		if item.name != test.expectedName {
			t.Errorf("Index %d expected name %s but got %s", test.index, test.expectedName, item.name)
		}
		if item.quantity != test.expectedQuantity {
			t.Errorf("Index %d expected quantity %d but got %d", test.index, test.expectedQuantity, item.quantity)
		}

	}
	keyItemTests := []struct {
		index            int
		expectedName     string
		expectedQuantity int
	}{
		{
			0,
			"Wailmer Pail",
			1,
		}, {
			1,
			"Old Rod",
			1,
		}, {
			2,
			"Powder Jar",
			1,
		}, {
			3,
			"Itemfinder",
			1,
		}, {
			4,
			"Acro Bike",
			1,
		}, {
			5,
			"Soot Sack",
			1,
		}, {
			6,
			"Meteorite",
			1,
		},
	}

	keyItemPocket := trainerBag.KeyItemsPocket()
	for _, test := range keyItemTests {
		item := keyItemPocket[test.index]
		if item.name != test.expectedName {
			t.Errorf("Index %d expected name %s but got %s", test.index, test.expectedName, item.name)
		}
		if item.quantity != test.expectedQuantity {
			t.Errorf("Index %d expected quantity %d but got %d", test.index, test.expectedQuantity, item.quantity)
		}

	}

	// Replace pokemon in team
	save, err = save.ReplacePokemonInTeam(save.Trainer.Team()[0], 5)

	if err != nil {
		t.Error("Error replacing pokemon in team")
	}

	firstTeam := save.Trainer.Team()[0]

	fifthTeam := save.Trainer.Team()[4]

	if !firstTeam.Compare(fifthTeam) {
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

	if !firstTeam.Compare(pkmList[2]) {
		t.Error("Error replacing pokemon in box")
	}

	// Replace pokemon in box with boxed pokemon
	save, err = save.ReplacePokemonInPC(pkmList[0], 3)

	if err != nil {
		t.Error("Error replacing pokemon in box")
	}

	pc = save.Trainer.Pc()

	box, _ = pc.Box(1)

	pkmList = box.BoxedPokemon()

	if !pkmList[0].Compare(pkmList[2]) {
		t.Error("Error replacing pokemon in box")
	}

	// Replace pokemon in box with boxed pokemon
	save, err = save.ReplacePokemonInPC(pkmList[0], 8943)

	if err != ErrIncorrectIndex {
		t.Error("I should error because of wrong index")
	}

	// Error Short file
	pathShort := "../testfiles/Lola.pkm"

	_, err = ReadDataFromSave(pathShort)

	if err != ErrShortFile {
		t.Fatalf("it shold have been ErrShortFile")
	}
	// TODO: add more tests

}
