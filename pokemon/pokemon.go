package pokemon

import (
	"encoding/binary"
	"errors"
	"os"

	"github.com/PailosNicolas/GoPkmSaveReader/helpers"
)

/*
Basic Pokemon Gen3 information.
*/
type Pokemon struct {
	raw              []byte // added in case it may be needed later
	PersonalityValue int
	OTPublicId       int
	OTSecretId       int
	OTName           string
	Nickname         string
	Species          string
	SpeciesIndex     int
	ItemHeld         Item
	Experience       int
	Friendship       int
	Moves            [4]Move
	Level            int
	Evs              EvsAndIV
	IVs              EvsAndIV
	Stats
	MetLocation   string
	MetAtLevel    int
	GameOfOrigin  string
	PokeBall      string
	TrainerGender string
	IsEgg         bool
	SecondAbility bool
	Language      string
}

/*
Represents the information of a pokemon move
*/
type Move struct {
	Id   int
	Name string
	PP   int
}

/*
Game items, can also be HeldItems
*/
type Item struct {
	Id   int
	Name string
}

/*
Evs of a pokemon
*/
type EvsAndIV struct {
	Hp             int
	Attack         int
	Defense        int
	Speed          int
	SpecialAttack  int
	SpecialDefense int
}

/*
Pokemon stats
*/
type Stats struct {
	TotalHP        int
	CurrentHP      int
	Attack         int
	Defense        int
	Speed          int
	SpecialAttack  int
	SpecialDefense int
	Level          int
}

/*
Reads the pokemon data and returns a Pokemon with it's information.
*/
func ParsePokemon(pkmData []byte) Pokemon {
	orders := [24]string{"GAEM", "GAME", "GEAM", "GEMA", "GMAE", "GMEA", "AGEM", "AGME", "AEGM", "AEMG", "AMGE", "AMEG", "EGAM", "EGMA", "EAGM", "EAMG", "EMGA", "EMAG", "MGAE", "MGEA", "MAGE", "MAEG", "MEGA", "MEAG"}
	pkm := Pokemon{}
	pkm.raw = pkmData
	var growth []byte
	var attacks []byte
	var evsAndCondition []byte
	var misc []byte

	//personality value
	personalityValue := pkmData[0:4]
	pkm.PersonalityValue = int(binary.LittleEndian.Uint32(personalityValue))

	//trianer full id
	trainerId := pkmData[4:8]
	pkm.OTPublicId = int(binary.LittleEndian.Uint16(pkmData[6:8]))
	pkm.OTSecretId = int(binary.LittleEndian.Uint16(pkmData[4:6]))

	//order
	order := orders[pkm.PersonalityValue%24]

	//getting ot
	pkm.OTName = helpers.ReadString(pkmData[20:27])

	//getting nickname
	pkm.Nickname = helpers.ReadString(pkmData[8:18])

	// Stats
	pkm.Stats.CurrentHP = int(binary.LittleEndian.Uint16(pkmData[86:88]))
	pkm.Stats.TotalHP = int(binary.LittleEndian.Uint16(pkmData[88:90]))
	pkm.Stats.Attack = int(binary.LittleEndian.Uint16(pkmData[90:92]))
	pkm.Stats.Defense = int(binary.LittleEndian.Uint16(pkmData[92:94]))
	pkm.Stats.Speed = int(binary.LittleEndian.Uint16(pkmData[94:96]))
	pkm.Stats.SpecialAttack = int(binary.LittleEndian.Uint16(pkmData[96:98]))
	pkm.Stats.SpecialDefense = int(binary.LittleEndian.Uint16(pkmData[98:100]))

	//Level
	pkm.Level = int(pkmData[84])

	//subdata
	key := helpers.XorBytes(personalityValue, trainerId)
	subdata := pkmData[32:80]

	for i := range 4 {
		start := i * 12
		if order[i] == 'G' {
			growth = helpers.DecryptData(subdata[start:start+12], key)
		}
		if order[i] == 'A' {
			attacks = helpers.DecryptData(subdata[start:start+12], key)
		}
		if order[i] == 'E' {
			evsAndCondition = helpers.DecryptData(subdata[start:start+12], key)
		}
		if order[i] == 'M' {
			misc = helpers.DecryptData(subdata[start:start+12], key)
		}
	}

	// Growth
	pkm.SpeciesIndex = int(binary.LittleEndian.Uint16(growth[0:2]))
	pkm.Species = helpers.Species[pkm.SpeciesIndex]
	pkm.ItemHeld = Item{
		Id:   int(binary.LittleEndian.Uint16(growth[2:4])),
		Name: helpers.ItemIndex[int(binary.LittleEndian.Uint16(growth[2:4]))],
	}
	pkm.Experience = int(binary.LittleEndian.Uint32(growth[4:8]))
	pkm.Friendship = int(growth[9])

	// Attacks
	pkm.Moves[0].Id = int(binary.LittleEndian.Uint16(attacks[0:2]))
	pkm.Moves[1].Id = int(binary.LittleEndian.Uint16(attacks[2:4]))
	pkm.Moves[2].Id = int(binary.LittleEndian.Uint16(attacks[4:6]))
	pkm.Moves[3].Id = int(binary.LittleEndian.Uint16(attacks[6:8]))

	pkm.Moves[0].Name = helpers.MovesIndex[pkm.Moves[0].Id]
	pkm.Moves[1].Name = helpers.MovesIndex[pkm.Moves[1].Id]
	pkm.Moves[2].Name = helpers.MovesIndex[pkm.Moves[2].Id]
	pkm.Moves[3].Name = helpers.MovesIndex[pkm.Moves[3].Id]

	pkm.Moves[0].PP = int(attacks[8])
	pkm.Moves[1].PP = int(attacks[9])
	pkm.Moves[2].PP = int(attacks[10])
	pkm.Moves[3].PP = int(attacks[11])

	// Ev & condition
	pkm.Evs.Hp = int(evsAndCondition[0])
	pkm.Evs.Attack = int(evsAndCondition[1])
	pkm.Evs.Defense = int(evsAndCondition[2])
	pkm.Evs.Speed = int(evsAndCondition[3])
	pkm.Evs.SpecialAttack = int(evsAndCondition[4])
	pkm.Evs.SpecialDefense = int(evsAndCondition[5])

	// Misc
	pkm.MetLocation = helpers.LocationIndexes[misc[1]]
	originInfo := helpers.Uint16ToBits(binary.LittleEndian.Uint16(misc[2:4]))
	pkm.MetAtLevel = helpers.BitsToInt(originInfo[0:7])
	pkm.GameOfOrigin = gamesOfOrigin[helpers.BitsToInt(originInfo[7:11])]
	pkm.PokeBall = pokeBalls[helpers.BitsToInt(originInfo[11:15])]
	if helpers.BitsToInt(originInfo[15:16]) == 1 {
		pkm.TrainerGender = "Female"
	} else {
		pkm.TrainerGender = "Male"
	}

	ivsEggAbility := helpers.Uint32ToBits(binary.LittleEndian.Uint32(misc[4:8]))
	pkm.IVs.Hp = helpers.BitsToInt(ivsEggAbility[0:5])
	pkm.IVs.Attack = helpers.BitsToInt(ivsEggAbility[5:10])
	pkm.IVs.Defense = helpers.BitsToInt(ivsEggAbility[10:15])
	pkm.IVs.Speed = helpers.BitsToInt(ivsEggAbility[15:20])
	pkm.IVs.SpecialAttack = helpers.BitsToInt(ivsEggAbility[20:25])
	pkm.IVs.SpecialDefense = helpers.BitsToInt(ivsEggAbility[25:30])

	if helpers.BitsToInt(ivsEggAbility[30:31]) == 1 {
		pkm.IsEgg = true
	} else {
		pkm.IsEgg = false
	}

	if helpers.BitsToInt(ivsEggAbility[31:32]) == 1 {
		pkm.SecondAbility = true
	} else {
		pkm.SecondAbility = false
	}

	pkm.Language = laguageOfOrigin[int(pkmData[18])]

	return pkm
}

/*
Exports Pokemon raw data to a file in the path directory
path should be only the directory with no file name.
*/
func (pkm *Pokemon) ExportPokemonToFile(path string) error {
	if path[len(path)-1] != '/' {
		return errors.New("wrong path file, it should end in '/'")
	}
	if pkm.Nickname != "" {
		path += pkm.Nickname + ".pkm"
	} else {
		path += pkm.Species + ".pkm"
	}
	err := os.WriteFile(path, pkm.raw, 0444)
	if err != nil {
		return err
	}
	return nil
}

var gamesOfOrigin = map[int]string{
	1:  "Sapphire",
	2:  "Ruby",
	3:  "Emerald",
	4:  "FireRed",
	5:  "LeafGreen",
	15: "Colosseum or XD",
}

var pokeBalls = map[int]string{
	1:  "Master Ball",
	2:  "Ultra Ball",
	3:  "Great Ball",
	4:  "Poke Ball",
	5:  "Safari Ball",
	6:  "Net Ball",
	7:  "Dive Ball",
	8:  "Nest Ball",
	9:  "Repeat Ball",
	10: "Timer Ball",
	11: "Luxury Ball",
	12: "Premier Ball",
}

var laguageOfOrigin = map[int]string{
	1: "Japanese",
	2: "English",
	3: "French",
	4: "Italian",
	5: "German",
	6: "unused",
	7: "Spanish",
}
