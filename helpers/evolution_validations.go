package helpers

const MethodLvlUp string = "levelup"
const MethodFriendship string = "friendship"
const MethodTrade string = "trade"
const MethodTradeHolding string = "tradeHolding"
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
	"Cleffa": {
		"Clefairy": {
			"method":     MethodFriendship,
			"friendship": "220",
			"id":         "35",
		},
	},
	"Clefairy": {
		"Clefable": {
			"method": MethodStone,
			"id":     "36",
		},
	},
	"Vulpix": {
		"Ninetales": {
			"method": MethodStone,
			"id":     "38",
		},
	},
	"Igglybuff": {
		"Jigglypuff": {
			"method":     MethodFriendship,
			"friendship": "220",
			"id":         "39",
		},
	},
	"Jigglypuff": {
		"Wigglytuff": {
			"method": MethodStone,
			"id":     "40",
		},
	},
	"Zubat": {
		"Golbat": {
			"method": MethodLvlUp,
			"level":  "22",
			"id":     "42",
		},
	},
	"Golbat": {
		"Crobat": {
			"method":     MethodFriendship,
			"friendship": "220",
			"id":         "169",
		},
	},
	"Oddish": {
		"Gloom": {
			"method": MethodLvlUp,
			"level":  "21",
			"id":     "44",
		},
	},
	"Gloom": {
		"Vileplume": {
			"method": MethodStone,
			"id":     "45",
		},
		"Bellossom": {
			"method": MethodStone,
			"id":     "182",
		},
	},
	"Paras": {
		"Parasect": {
			"method": MethodLvlUp,
			"level":  "24",
			"id":     "47",
		},
	},
	"Venonat": {
		"Venomoth": {
			"method": MethodLvlUp,
			"level":  "31",
			"id":     "49",
		},
	},
	"Diglett": {
		"Dugtrio": {
			"method": MethodLvlUp,
			"level":  "26",
			"id":     "51",
		},
	},
	"Meowth": {
		"Persian": {
			"method": MethodLvlUp,
			"level":  "28",
			"id":     "53",
		},
	},
	"Psyduck": {
		"Golduck": {
			"method": MethodLvlUp,
			"level":  "33",
			"id":     "55",
		},
	},
	"Mankey": {
		"Primeape": {
			"method": MethodLvlUp,
			"level":  "28",
			"id":     "57",
		},
	},
	"Growlithe": {
		"Arcanine": {
			"method": MethodStone,
			"id":     "59",
		},
	},
	"Poliwag": {
		"Poliwhirl": {
			"method": MethodLvlUp,
			"level":  "25",
			"id":     "61",
		},
	},
	"Poliwhirl": {
		"Poliwrath": {
			"method": MethodStone,
			"id":     "62",
		},
		"Politoed": {
			"method": MethodTradeHolding,
			"item":   "King's Rock",
			"id":     "186",
		},
	},
	"Abra": {
		"Kadabra": {
			"method": MethodLvlUp,
			"level":  "16",
			"id":     "64",
		},
	},
	"Kadabra": {
		"Alakazam": {
			"method": MethodTrade,
			"id":     "65",
		},
	},
	"Machop": {
		"Machoke": {
			"method": MethodLvlUp,
			"level":  "28",
			"id":     "67",
		},
	},
	"Machoke": {
		"Machamp": {
			"method": MethodTrade,
			"id":     "68",
		},
	},
	"Bellsprout": {
		"Weepinbell": {
			"method": MethodLvlUp,
			"level":  "21",
			"id":     "70",
		},
	},
	"Weepinbell": {
		"Victreebel": {
			"method": MethodStone,
			"id":     "71",
		},
	},
	"Tentacool": {
		"Tentacruel": {
			"method": MethodLvlUp,
			"level":  "30",
			"id":     "73",
		},
	},
	"Geodude": {
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
	"Ponyta": {
		"Rapidash": {
			"method": MethodLvlUp,
			"level":  "40",
			"id":     "78",
		},
	},
	"Slowpoke": {
		"Slowbro": {
			"method": MethodLvlUp,
			"level":  "37",
			"id":     "80",
		},
		"Slowking": {
			"method": MethodTradeHolding,
			"item":   "King's Rock",
			"id":     "199",
		},
	},
}
