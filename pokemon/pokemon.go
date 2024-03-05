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
}

func ParsePokemon(pkmData []byte) Pokemon {
	orders := [24]string{"GAEM", "GAME", "GEAM", "GEMA", "GMAE", "GMEA", "AGEM", "AGME", "AEGM", "AEMG", "AMGE", "AMEG", "EGAM", "EGMA", "EAGM", "EAMG", "EMGA", "EMAG", "MGAE", "MGEA", "MAGE", "MAEG", "MEGA", "MEAG"}
	pkm := Pokemon{}
	pkm.raw = pkmData
	var growth []byte
	// var attacks []byte
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

	//subdata
	key := helpers.XorBytes(personalityValue, trainerId)
	subdata := pkmData[32:80]

	for i := range 4 {
		start := i * 12
		if order[i] == 'G' {
			growth = helpers.DecryptData(subdata[start:start+12], key)
		}
		if order[i] == 'A' {
			_ = helpers.DecryptData(subdata[start:start+12], key)
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

	return pkm
}
