package pokemon

import (
	"encoding/binary"

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
	TotalHP          int
	CurrentHP        int
	Attack           int
	Defense          int
	Speed            int
	SpecialAttack    int
	SpecialDefense   int
	Level            int
}

/*
Represents the information of a pokemon move
*/
type Move struct {
	Id   int
	Name string
	PP   int
}

type Item struct {
	Id   int
	Name string
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
	// var evsAndCondition []byte
	// var misc []byte

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
	pkm.CurrentHP = int(binary.LittleEndian.Uint16(pkmData[86:88]))
	pkm.TotalHP = int(binary.LittleEndian.Uint16(pkmData[88:90]))
	pkm.Attack = int(binary.LittleEndian.Uint16(pkmData[90:92]))
	pkm.Defense = int(binary.LittleEndian.Uint16(pkmData[92:94]))
	pkm.Speed = int(binary.LittleEndian.Uint16(pkmData[94:96]))
	pkm.SpecialAttack = int(binary.LittleEndian.Uint16(pkmData[96:98]))
	pkm.SpecialDefense = int(binary.LittleEndian.Uint16(pkmData[98:100]))

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
			_ = helpers.DecryptData(subdata[start:start+12], key)
		}
		if order[i] == 'M' {
			_ = helpers.DecryptData(subdata[start:start+12], key)
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

	return pkm
}
