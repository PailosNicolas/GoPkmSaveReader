package pokemon

import (
	"encoding/binary"

	"github.com/PailosNicolas/GoPkmSaveReader/helpers"
)

type Pokemon struct {
	raw              []byte // added in case it may be needed later
	PersonalityValue int
	OTPublicId       int
	OTSecretId       int
	OTName           string
	Nickname         string
	Species          string
	SpeciesIndex     int
	ItemHeld         string
	ItemHeldIndex    int
	Experience       int
	Friendship       int
	Move1Index       int
	Move2Index       int
	Move3Index       int
	Move4Index       int
	Move1            string
	Move2            string
	Move3            string
	Move4            string
	TotalHP          int
	CurrentHP        int
	Attack           int
	Defense          int
	Speed            int
	SpecialAttack    int
	SpecialDefense   int
}

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
	pkm.ItemHeldIndex = int(binary.LittleEndian.Uint16(growth[2:4]))
	pkm.ItemHeld = helpers.ItemIndex[pkm.ItemHeldIndex]
	pkm.Experience = int(binary.LittleEndian.Uint32(growth[4:8]))
	pkm.Friendship = int(growth[9])

	// Attacks
	pkm.Move1Index = int(binary.LittleEndian.Uint16(attacks[0:2]))
	pkm.Move2Index = int(binary.LittleEndian.Uint16(attacks[2:4]))
	pkm.Move3Index = int(binary.LittleEndian.Uint16(attacks[4:6]))
	pkm.Move4Index = int(binary.LittleEndian.Uint16(attacks[6:8]))

	pkm.Move1 = helpers.MovesIndex[pkm.Move1Index]
	pkm.Move2 = helpers.MovesIndex[pkm.Move2Index]
	pkm.Move3 = helpers.MovesIndex[pkm.Move3Index]
	pkm.Move4 = helpers.MovesIndex[pkm.Move4Index]

	return pkm
}
