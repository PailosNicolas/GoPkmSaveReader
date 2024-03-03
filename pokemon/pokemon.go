package pokemon

import "github.com/PailosNicolas/GoPkmSaveReader/helpers"

type Pokemon struct {
	OTName string
}

func ParsePokemon(pkmData []byte) Pokemon {
	pkm := Pokemon{}

	//getting ot
	pkm.OTName = helpers.ReadString(pkmData[20:27])

	return pkm
}
