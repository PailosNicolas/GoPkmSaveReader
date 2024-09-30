package helpers

const MethodLvlUp string = "levelup"
const MethodFriendship string = "friendship"

var EvolutionValidation = map[string]map[string]string{ // TODO: For now only level up and direct evolutions, check for stones and well... eevee
	"Bulbasaur": {
		"method": MethodLvlUp,
		"level":  "16",
	},
	"Ivysaur": {
		"method": MethodLvlUp,
		"level":  "32",
	},
	"Charmander": {
		"method": MethodLvlUp,
		"level":  "16",
	},
	"Charmeleon": {
		"method": MethodLvlUp,
		"level":  "36",
	},
	"Squirtle": {
		"method": MethodLvlUp,
		"level":  "16",
	},
	"Wartortle": {
		"method": MethodLvlUp,
		"level":  "36",
	},
	"Caterpie": {
		"method": MethodLvlUp,
		"level":  "7",
	},
	"Metapod": {
		"method": MethodLvlUp,
		"level":  "10",
	},
	"Weedle": {
		"method": MethodLvlUp,
		"level":  "7",
	},
	"Kakuna": {
		"method": MethodLvlUp,
		"level":  "10",
	},
	"Pidgey": {
		"method": MethodLvlUp,
		"level":  "18",
	},
	"Pidgeotto": {
		"method": MethodLvlUp,
		"level":  "36",
	},
	"Rattata": {
		"method": MethodLvlUp,
		"level":  "20",
	},
	"Spearow": {
		"method": MethodLvlUp,
		"level":  "20",
	},
	"Ekans": {
		"method": MethodLvlUp,
		"level":  "22",
	},
	"Pichu": {
		"method":     MethodFriendship,
		"friendship": "220",
	},
	"Sandshrew": {
		"method": MethodLvlUp,
		"level":  "22",
	},
	"Nidoran-f": {
		"method": MethodLvlUp,
		"level":  "16",
	},
}
