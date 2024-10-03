package helpers

const MethodLvlUp string = "levelup"
const MethodFriendship string = "friendship"

var EvolutionValidation = map[string]map[string]map[string]string{ // TODO: For now only level up and direct evolutions, check for stones and well... eevee
	"Bulbasaur": {
		"Ivysaur": {
			"method": MethodLvlUp,
			"level":  "16",
		},
	},
	"Ivysaur": {
		"Venusaur": {
			"method": MethodLvlUp,
			"level":  "32",
		},
	},
	"Charmander": {
		"Charmeleon": {
			"method": MethodLvlUp,
			"level":  "16",
		},
	},
	"Charmeleon": {
		"Charizard": {
			"method": MethodLvlUp,
			"level":  "36",
		},
	},
	"Squirtle": {
		"Wartortle": {
			"method": MethodLvlUp,
			"level":  "16",
		},
	},
	"Wartortle": {
		"Blastoise": {
			"method": MethodLvlUp,
			"level":  "36",
		},
	},
	"Caterpie": {
		"Metapod": {
			"method": MethodLvlUp,
			"level":  "7",
		},
	},
	"Metapod": {
		"Butterfree": {
			"method": MethodLvlUp,
			"level":  "10",
		},
	},
	"Weedle": {
		"Kakuna": {
			"method": MethodLvlUp,
			"level":  "7",
		},
	},
	"Kakuna": {
		"Beedrill": {
			"method": MethodLvlUp,
			"level":  "10",
		},
	},
	"Pidgey": {
		"Pidgeotto": {
			"method": MethodLvlUp,
			"level":  "18",
		},
	},
	"Pidgeotto": {
		"Pidgeot": {
			"method": MethodLvlUp,
			"level":  "36",
		},
	},
	"Rattata": {
		"Raticate": {
			"method": MethodLvlUp,
			"level":  "20",
		},
	},
	"Spearow": {
		"Fearow": {
			"method": MethodLvlUp,
			"level":  "20",
		},
	},
	"Ekans": {
		"Arbok": {
			"method": MethodLvlUp,
			"level":  "22",
		},
	},
	"Pichu": {
		"Pikachu": {
			"method":     MethodFriendship,
			"friendship": "220",
		},
	},
	"Sandshrew": {
		"Sandslash": {
			"method": MethodLvlUp,
			"level":  "22",
		},
	},
}
