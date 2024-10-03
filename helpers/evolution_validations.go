package helpers

const MethodLvlUp string = "levelup"
const MethodFriendship string = "friendship"
const MethodTrade string = "trade"
const MethodStone string = "stone"

var EvolutionValidation = map[string]map[string]map[string]string{
	"Bulbasaur": {
		"Ivysaur": {
			"method": MethodLvlUp,
			"level":  "16",
			"id":     "2",
		},
	},
	"Ivysaur": {
		"Venusaur": {
			"method": MethodLvlUp,
			"level":  "32",
			"id":     "3",
		},
	},
	"Charmander": {
		"Charmeleon": {
			"method": MethodLvlUp,
			"level":  "16",
			"id":     "5",
		},
	},
	"Charmeleon": {
		"Charizard": {
			"method": MethodLvlUp,
			"level":  "36",
			"id":     "6",
		},
	},
	"Squirtle": {
		"Wartortle": {
			"method": MethodLvlUp,
			"level":  "16",
			"id":     "8",
		},
	},
	"Wartortle": {
		"Blastoise": {
			"method": MethodLvlUp,
			"level":  "36",
			"id":     "9",
		},
	},
	"Caterpie": {
		"Metapod": {
			"method": MethodLvlUp,
			"level":  "7",
			"id":     "11",
		},
	},
	"Metapod": {
		"Butterfree": {
			"method": MethodLvlUp,
			"level":  "10",
			"id":     "12",
		},
	},
	"Weedle": {
		"Kakuna": {
			"method": MethodLvlUp,
			"level":  "7",
			"id":     "14",
		},
	},
	"Kakuna": {
		"Beedrill": {
			"method": MethodLvlUp,
			"level":  "10",
			"id":     "15",
		},
	},
	"Pidgey": {
		"Pidgeotto": {
			"method": MethodLvlUp,
			"level":  "18",
			"id":     "17",
		},
	},
	"Pidgeotto": {
		"Pidgeot": {
			"method": MethodLvlUp,
			"level":  "36",
			"id":     "18",
		},
	},
	"Rattata": {
		"Raticate": {
			"method": MethodLvlUp,
			"level":  "20",
			"id":     "20",
		},
	},
	"Spearow": {
		"Fearow": {
			"method": MethodLvlUp,
			"level":  "20",
			"id":     "22",
		},
	},
	"Ekans": {
		"Arbok": {
			"method": MethodLvlUp,
			"level":  "22",
			"id":     "24",
		},
	},
	"Pichu": {
		"Pikachu": {
			"method":     MethodFriendship,
			"friendship": "220",
			"id":         "25",
		},
	},
	"Pikachu": {
		"Raichu": {
			"method": MethodStone,
			"id":     "26",
		},
	},
	"Sandshrew": {
		"Sandslash": {
			"method": MethodLvlUp,
			"level":  "22",
			"id":     "28",
		},
	},
	"Geodude": { // Added early for test
		"Graveler": {
			"method": MethodLvlUp,
			"level":  "25",
			"id":     "75",
		},
	},
	"Graveler": {
		"Golem": {
			"method": MethodTrade,
			"id":     "76",
		},
	},
	"Nidoran-f": {
		"Nidorina": {
			"method": MethodLvlUp,
			"level":  "16",
			"id":     "30",
		},
	},
	"Nidorina": {
		"Nidoqueen": {
			"method": MethodStone,
			"id":     "31",
		},
	},
	"Nidoran-m": {
		"Nidorino": {
			"method": MethodLvlUp,
			"level":  "16",
			"id":     "33",
		},
	},
	"Nidorino": {
		"Nidoking": {
			"method": MethodStone,
			"id":     "34",
		},
	},
}
