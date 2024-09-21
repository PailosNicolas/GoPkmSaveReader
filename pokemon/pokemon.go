package pokemon

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math"
	"os"

	"github.com/PailosNicolas/GoPkmSaveReader/helpers"
)

var orders = [24]string{"GAEM", "GAME", "GEAM", "GEMA", "GMAE", "GMEA", "AGEM", "AGME", "AEGM", "AEMG", "AMGE", "AMEG", "EGAM", "EGMA", "EAGM", "EAMG", "EMGA", "EMAG", "MGAE", "MGEA", "MAGE", "MAEG", "MEGA", "MEAG"}

/*
Basic Pokemon Gen3 information.
*/
type Pokemon struct {
	raw              []byte // added in case it may be needed later
	personalityValue int
	oTPublicId       int
	oTSecretId       int
	oTName           string
	nickname         string
	species          string
	speciesIndex     int
	itemHeld         Item
	experience       int
	friendship       int
	moves            [4]Move
	level            int
	evs              EvsAndIV
	iVs              EvsAndIV
	stats            Stats
	metLocation      string
	metAtLevel       int
	gameOfOrigin     string
	pokeBall         string
	trainerGender    string
	isEgg            bool
	secondAbility    bool
	language         string
	unencryptedData  [][]byte
	nature           Nature
	experienceType   string
}

func (p *Pokemon) PersonalityValue() int {
	return p.personalityValue
}

func (p *Pokemon) OTPublicId() int {
	return p.oTPublicId
}

func (p *Pokemon) OTSecretId() int {
	return p.oTSecretId
}

func (p *Pokemon) OTName() string {
	return p.oTName
}

func (p *Pokemon) Nickname() string {
	return p.nickname
}

func (p *Pokemon) Species() string {
	return p.species
}

func (p *Pokemon) SpeciesIndex() int {
	return p.speciesIndex
}

func (p *Pokemon) ItemHeld() Item {
	return p.itemHeld
}

func (p *Pokemon) Experience() int {
	return p.experience
}

func (p *Pokemon) Friendship() int {
	return p.friendship
}

func (p *Pokemon) Moves() [4]Move {
	return p.moves
}

func (p *Pokemon) Level() int {
	return p.level
}

func (p *Pokemon) Evs() EvsAndIV {
	return p.evs
}

func (p *Pokemon) IVs() EvsAndIV {
	return p.iVs
}

func (p *Pokemon) Stats() Stats {
	return p.stats
}

func (p *Pokemon) MetLocation() string {
	return p.metLocation
}

func (p *Pokemon) MetAtLevel() int {
	return p.metAtLevel
}

func (p *Pokemon) GameOfOrigin() string {
	return p.gameOfOrigin
}

func (p *Pokemon) PokeBall() string {
	return p.pokeBall
}

func (p *Pokemon) TrainerGender() string {
	return p.trainerGender
}

func (p *Pokemon) IsEgg() bool {
	return p.isEgg
}

func (p *Pokemon) SecondAbility() bool {
	return p.secondAbility
}

func (p *Pokemon) Language() string {
	return p.language
}

func (p *Pokemon) Nature() Nature {
	return p.nature
}

func (p *Pokemon) ExperienceType() string {
	return p.experienceType
}

func (p *Pokemon) Raw() []byte {
	return p.raw
}

/*
Returns displayable name for the pokemon, it will be it's nickname if it has one else it will be it's species
*/
func (p *Pokemon) Display() string {
	if p.nickname != "" {
		return p.nickname
	} else {
		return p.species
	}
}

/*
Compares two Pokemons returns true if they are the same
*/
func (p *Pokemon) Compare(pokemon Pokemon) bool {
	return bytes.Equal(p.raw[:80], pokemon.raw[:80])
}

/*
Represents the pokemon nature and which stats does it change
*/
type Nature struct {
	ID        int
	Name      string
	Increases string
	Decreases string
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

// Errors
var ErrFileShort = errors.New("file is too short to be a pokemon raw data")
var ErrPkmNotValid = errors.New("the pokemon is not valid")

/*
Reads the pokemon data and returns a Pokemon with it's information.
*/
func ParsePokemon(pkmData []byte) (Pokemon, error) {
	pkm := Pokemon{}
	pkm.raw = pkmData
	var growth []byte
	var attacks []byte
	var evsAndCondition []byte
	var misc []byte

	//personality value
	personalityValue := pkmData[0:4]
	pkm.personalityValue = int(binary.LittleEndian.Uint32(personalityValue))

	//nature
	pkm.nature = natures[pkm.personalityValue%25]

	//trianer full id
	trainerId := pkmData[4:8]
	pkm.oTPublicId = int(binary.LittleEndian.Uint16(pkmData[6:8]))
	pkm.oTSecretId = int(binary.LittleEndian.Uint16(pkmData[4:6]))

	//order
	order := orders[pkm.personalityValue%24]

	//getting ot
	pkm.oTName = helpers.ReadString(pkmData[20:27])

	//getting nickname
	pkm.nickname = helpers.ReadString(pkmData[8:18])

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

	pkm.unencryptedData = [][]byte{growth, attacks, evsAndCondition, misc}

	if !pkm.IsValid() {
		return Pokemon{}, ErrPkmNotValid
	}

	// Growth
	pkm.speciesIndex = int(binary.LittleEndian.Uint16(growth[0:2]))
	pkm.species = helpers.Species[pkm.speciesIndex]
	pkm.experienceType = helpers.ExperienceType[pkm.species]
	pkm.itemHeld = Item{
		Id:   int(binary.LittleEndian.Uint16(growth[2:4])),
		Name: helpers.ItemIndex[int(binary.LittleEndian.Uint16(growth[2:4]))],
	}
	pkm.experience = int(binary.LittleEndian.Uint32(growth[4:8]))
	pkm.friendship = int(growth[9])

	// Attacks
	pkm.moves[0].Id = int(binary.LittleEndian.Uint16(attacks[0:2]))
	pkm.moves[1].Id = int(binary.LittleEndian.Uint16(attacks[2:4]))
	pkm.moves[2].Id = int(binary.LittleEndian.Uint16(attacks[4:6]))
	pkm.moves[3].Id = int(binary.LittleEndian.Uint16(attacks[6:8]))

	pkm.moves[0].Name = helpers.MovesIndex[pkm.moves[0].Id]
	pkm.moves[1].Name = helpers.MovesIndex[pkm.moves[1].Id]
	pkm.moves[2].Name = helpers.MovesIndex[pkm.moves[2].Id]
	pkm.moves[3].Name = helpers.MovesIndex[pkm.moves[3].Id]

	pkm.moves[0].PP = int(attacks[8])
	pkm.moves[1].PP = int(attacks[9])
	pkm.moves[2].PP = int(attacks[10])
	pkm.moves[3].PP = int(attacks[11])

	// Ev & condition
	pkm.evs.Hp = int(evsAndCondition[0])
	pkm.evs.Attack = int(evsAndCondition[1])
	pkm.evs.Defense = int(evsAndCondition[2])
	pkm.evs.Speed = int(evsAndCondition[3])
	pkm.evs.SpecialAttack = int(evsAndCondition[4])
	pkm.evs.SpecialDefense = int(evsAndCondition[5])

	// Misc
	pkm.metLocation = helpers.LocationIndexes[misc[1]]
	originInfo := helpers.Uint16ToBits(binary.LittleEndian.Uint16(misc[2:4]))
	pkm.metAtLevel = helpers.BitsToInt(originInfo[0:7])
	pkm.gameOfOrigin = gamesOfOrigin[helpers.BitsToInt(originInfo[7:11])]
	pkm.pokeBall = pokeBalls[helpers.BitsToInt(originInfo[11:15])]
	if helpers.BitsToInt(originInfo[15:16]) == 1 {
		pkm.trainerGender = "Female"
	} else {
		pkm.trainerGender = "Male"
	}

	ivsEggAbility := helpers.Uint32ToBits(binary.LittleEndian.Uint32(misc[4:8]))
	pkm.iVs.Hp = helpers.BitsToInt(ivsEggAbility[0:5])
	pkm.iVs.Attack = helpers.BitsToInt(ivsEggAbility[5:10])
	pkm.iVs.Defense = helpers.BitsToInt(ivsEggAbility[10:15])
	pkm.iVs.Speed = helpers.BitsToInt(ivsEggAbility[15:20])
	pkm.iVs.SpecialAttack = helpers.BitsToInt(ivsEggAbility[20:25])
	pkm.iVs.SpecialDefense = helpers.BitsToInt(ivsEggAbility[25:30])

	if helpers.BitsToInt(ivsEggAbility[30:31]) == 1 {
		pkm.isEgg = true
	} else {
		pkm.isEgg = false
	}

	if helpers.BitsToInt(ivsEggAbility[31:32]) == 1 {
		pkm.secondAbility = true
	} else {
		pkm.secondAbility = false
	}

	pkm.language = laguageOfOrigin[int(pkmData[18])]

	if len(pkmData) == 80 {
		// Boxed pokemon need to calculate stats

		//Level
		expType := helpers.ExperienceType[pkm.species]
		pkm.level = helpers.GetLevelFromExperience(helpers.ExperienceThresholds[expType], pkm.experience)

		pkm.stats.TotalHP = (((2*helpers.BaseStatsGenIII[pkm.species][helpers.HpString] + pkm.iVs.Hp + (pkm.evs.Hp / 4.0)) * pkm.level) / 100) + pkm.level + 10
		pkm.stats.CurrentHP = pkm.stats.TotalHP
		pkm.stats.Attack = pkm.caculateStats("Attack")
		pkm.stats.Defense = pkm.caculateStats("Defense")
		pkm.stats.SpecialAttack = pkm.caculateStats("SpAtk")
		pkm.stats.SpecialDefense = pkm.caculateStats("SpDef")
		pkm.stats.Speed = pkm.caculateStats("Speed")

	} else {
		// Stats
		pkm.stats.CurrentHP = int(binary.LittleEndian.Uint16(pkmData[86:88]))
		pkm.stats.TotalHP = int(binary.LittleEndian.Uint16(pkmData[88:90]))
		pkm.stats.Attack = int(binary.LittleEndian.Uint16(pkmData[90:92]))
		pkm.stats.Defense = int(binary.LittleEndian.Uint16(pkmData[92:94]))
		pkm.stats.Speed = int(binary.LittleEndian.Uint16(pkmData[94:96]))
		pkm.stats.SpecialAttack = int(binary.LittleEndian.Uint16(pkmData[96:98]))
		pkm.stats.SpecialDefense = int(binary.LittleEndian.Uint16(pkmData[98:100]))

		//Level
		pkm.level = int(pkmData[84])
	}

	return pkm, nil
}

/*
Exports Pokemon raw data to a file in the path directory
path should be only the directory with no file name.
*/
func (pkm *Pokemon) ExportPokemonToFile(path string) error {
	if path[len(path)-1] != '/' {
		return errors.New("wrong path file, it should end in '/'")
	}
	if pkm.nickname != "" {
		path += pkm.nickname + ".pkm"
	} else {
		path += pkm.species + ".pkm"
	}
	err := os.WriteFile(path, pkm.raw, 0444)
	if err != nil {
		return err
	}
	return nil
}

/*
Evolves a pokemon.
*/
func (pkm *Pokemon) EvolvePokemon() (Pokemon, error) {
	var newRaw []byte
	var growth []byte
	var growthIndex int = 32
	var newCheckSum uint16

	if len(pkm.raw) > 80 {
		newRaw = pkm.raw[:80]
	} else {
		newRaw = pkm.raw
	}

	order := orders[pkm.personalityValue%24]

	key := helpers.XorBytes(newRaw[0:4], newRaw[4:8])

	for i := range 4 {
		start := i * 12
		if order[i] == 'G' {
			growthIndex += start
			break
		}
	}

	decrypetdData := pkm.unencryptedData

	growth = decrypetdData[0]

	newSpecIndex := int(binary.LittleEndian.Uint16(growth[0:2])) + 1

	binary.LittleEndian.PutUint16(growth[0:2], uint16(newSpecIndex))

	encryptedGrowth := helpers.DecryptData(growth, key)

	copy(newRaw[growthIndex:growthIndex+12], encryptedGrowth)

	for index, data := range pkm.unencryptedData {
		if index == 0 {
			data = growth
		}
		for i := 0; i < len(data); i += 2 {
			var value uint16
			if i+1 < len(data) {
				value = uint16(data[i]) | (uint16(data[i+1]) << 8)
			} else {
				value = uint16(data[i])
			}
			newCheckSum += value
		}
	}

	binary.LittleEndian.PutUint16(newRaw[28:30], newCheckSum)

	evolution, _ := ParsePokemon(newRaw)

	return evolution, nil // TOOD: implement
}

/*
Checks if a pokemon is valid, otherwise it will be displayed as a bad egg in-game.
*/
func (pkm *Pokemon) IsValid() bool {
	valid := false
	var sum uint16
	for _, data := range pkm.unencryptedData {
		for i := 0; i < len(data); i += 2 {
			var value uint16
			if i+1 < len(data) {
				value = uint16(data[i]) | (uint16(data[i+1]) << 8)
			} else {
				value = uint16(data[i])
			}
			sum += value
		}
	}

	checksum := binary.LittleEndian.Uint16(pkm.raw[28:30])

	if sum == checksum {
		valid = true
	}

	return valid
}

/*
Reads a raw pokemon file from the path.
The file should be the raw 100 bytes from a Gen3 save.
*/
func ReadPokemonFromFile(path string) (Pokemon, error) {
	file, err := os.Open(path)
	if err != nil {
		return Pokemon{}, err
	}
	defer file.Close()

	raw := make([]byte, 100)

	bytesRead, err := file.Read(raw)

	if err != nil {
		return Pokemon{}, err
	}

	if 80 > bytesRead {
		return Pokemon{}, ErrFileShort
	}

	pkm, err := ParsePokemon(raw)

	if err != nil {
		return Pokemon{}, err
	}

	return pkm, nil
}

func (pkm *Pokemon) caculateStats(statName string) int {
	/*
		Calculate the stats of gen3 pokemon except for hp.
	*/
	var iv float64
	var ev float64
	var nature float64
	switch statName {
	case "Attack":
		iv = float64(pkm.iVs.Attack)
		ev = float64(pkm.evs.Attack)
	case "Defense":
		iv = float64(pkm.iVs.Defense)
		ev = float64(pkm.evs.Defense)
	case "SpAtk":
		iv = float64(pkm.iVs.SpecialAttack)
		ev = float64(pkm.evs.SpecialAttack)
	case "SpDef":
		iv = float64(pkm.iVs.SpecialDefense)
		ev = float64(pkm.evs.SpecialDefense)
	case "Speed":
		iv = float64(pkm.iVs.Speed)
		ev = float64(pkm.evs.Speed)
	}

	if pkm.nature.Increases == statName {
		nature = 1.1
	} else if pkm.nature.Decreases == statName {
		nature = 0.9
	} else {
		nature = 1.0
	}
	result := (math.Floor((((float64(2*helpers.BaseStatsGenIII[pkm.species][statName]) + iv + math.Floor((ev / 4.0))) * float64(pkm.level)) / 100.0)) + 5) * nature
	return int(result)
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
