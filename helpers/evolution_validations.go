package helpers

const MethodLvlUp string = "levelup"
const MethodLvlUpAndStat string = "levelupAndStat"
const MethodFriendship string = "friendship"
const MethodTrade string = "trade"
const MethodTradeHolding string = "tradeHolding"
const MethodStone string = "stone"
const MethodWurmple string = "wurmple"
const MethodBeauty string = "beauty"

const ConditionHigher string = "higher"
const ConditionLower string = "lower"
const ConditionEqual string = "equal"

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
	"Magnemite": {
		"Magneton": {
			"method": MethodLvlUp,
			"level":  "30",
			"id":     "82",
		},
	},
	"Doduo": {
		"Dodrio": {
			"method": MethodLvlUp,
			"level":  "30",
			"id":     "85",
		},
	},
	"Seel": {
		"Dewgong": {
			"method": MethodLvlUp,
			"level":  "34",
			"id":     "87",
		},
	},
	"Grimer": {
		"Muk": {
			"method": MethodLvlUp,
			"level":  "38",
			"id":     "89",
		},
	},
	"Shellder": {
		"Cloyster": {
			"method": MethodStone,
			"id":     "91",
		},
	},
	"Gastly": {
		"Haunter": {
			"method": MethodLvlUp,
			"level":  "25",
			"id":     "93",
		},
	},
	"Haunter": {
		"Gengar": {
			"method": MethodTrade,
			"id":     "94",
		},
	},
	"Onix": {
		"Steelix": {
			"method": MethodTradeHolding,
			"item":   "King's Rock",
			"id":     "208",
		},
	},
	"Drowzee": {
		"Hypno": {
			"method": MethodLvlUp,
			"level":  "26",
			"id":     "97",
		},
	},
	"Krabby": {
		"Kingler": {
			"method": MethodLvlUp,
			"level":  "28",
			"id":     "99",
		},
	},
	"Voltorb": {
		"Electrode": {
			"method": MethodLvlUp,
			"level":  "30",
			"id":     "101",
		},
	},
	"Exeggcute": {
		"Exeggutor": {
			"method": MethodStone,
			"id":     "103",
		},
	},
	"Cubone": {
		"Marowak": {
			"method": MethodLvlUp,
			"level":  "28",
			"id":     "105",
		},
	},
	"Tyrogue": {
		"Hitmonlee": {
			"method":    MethodLvlUpAndStat,
			"level":     "20",
			"condition": ConditionHigher,
			"id":        "106",
		},
		"Hitmonchan": {
			"method":    MethodLvlUpAndStat,
			"level":     "20",
			"condition": ConditionLower,
			"id":        "106",
		},
		"Hitmontop": {
			"method":    MethodLvlUpAndStat,
			"level":     "20",
			"condition": ConditionEqual,
			"id":        "237",
		},
	},
	"Koffing": {
		"Weezing": {
			"method": MethodLvlUp,
			"level":  "35",
			"id":     "110",
		},
	},
	"Rhyhorn": {
		"Rhydon": {
			"method": MethodLvlUp,
			"level":  "42",
			"id":     "112",
		},
	},
	"Chansey": {
		"Blissey": {
			"method":     MethodFriendship,
			"friendship": "220",
			"id":         "242",
		},
	},
	"Horsea": {
		"Seadra": {
			"method": MethodLvlUp,
			"level":  "32",
			"id":     "117",
		},
	},
	"Seadra": {
		"Kingdra": {
			"method": MethodTradeHolding,
			"item":   "Dragon Scale",
			"id":     "230",
		},
	},
	"Goldeen": {
		"Seaking": {
			"method": MethodLvlUp,
			"level":  "33",
			"id":     "119",
		},
	},
	"Staryu": {
		"Starmie": {
			"method": MethodStone,
			"id":     "121",
		},
	},
	"Scyther": {
		"Scizor": {
			"method": MethodTradeHolding,
			"item":   "Metal Coat",
			"id":     "212",
		},
	},
	"Smoochum": {
		"Jynx": {
			"method": MethodLvlUp,
			"level":  "30",
			"id":     "124",
		},
	},
	"Elekid": {
		"Electabuzz": {
			"method": MethodLvlUp,
			"level":  "30",
			"id":     "125",
		},
	},
	"Magby": {
		"Magmar": {
			"method": MethodLvlUp,
			"level":  "30",
			"id":     "126",
		},
	},
	"Magikarp": {
		"Gyarados": {
			"method": MethodLvlUp,
			"level":  "20",
			"id":     "130",
		},
	},
	"Eevee": {
		"Vaporeon": {
			"method": MethodStone,
			"id":     "134",
		},
		"Jolteon": {
			"method": MethodStone,
			"id":     "135",
		},
		"Flareon": {
			"method": MethodStone,
			"id":     "136",
		},
		"Espeon": {
			"method":     MethodFriendship,
			"friendship": "220",
			"id":         "196",
		},
		"Umbreon": {
			"method":     MethodFriendship,
			"friendship": "220",
			"id":         "197",
		},
	},
	"Porygon": {
		"Porygon2": {
			"method": MethodTradeHolding,
			"item":   "Up-Grade",
			"id":     "233",
		},
	},
	"Omanyte": {
		"Omastar": {
			"method": MethodLvlUp,
			"level":  "40",
			"id":     "139",
		},
	},
	"Kabuto": {
		"Kabutops": {
			"method": MethodLvlUp,
			"level":  "40",
			"id":     "141",
		},
	},
	"Dratini": {
		"Dragonair": {
			"method": MethodLvlUp,
			"level":  "30",
			"id":     "148",
		},
	},
	"Dragonair": {
		"Dragonite": {
			"method": MethodLvlUp,
			"level":  "55",
			"id":     "149",
		},
	},
	"Chikorita": {
		"Bayleef": {
			"method": MethodLvlUp,
			"level":  "16",
			"id":     "153",
		},
	},
	"Bayleef": {
		"Meganium": {
			"method": MethodLvlUp,
			"level":  "32",
			"id":     "154",
		},
	},
	"Cyndaquil": {
		"Quilava": {
			"method": MethodLvlUp,
			"level":  "14",
			"id":     "156",
		},
	},
	"Quilava": {
		"Typhlosion": {
			"method": MethodLvlUp,
			"level":  "36",
			"id":     "157",
		},
	},
	"Totodile": {
		"Croconaw": {
			"method": MethodLvlUp,
			"level":  "18",
			"id":     "159",
		},
	},
	"Croconaw": {
		"Feraligatr": {
			"method": MethodLvlUp,
			"level":  "30",
			"id":     "160",
		},
	},
	"Sentret": {
		"Furret": {
			"method": MethodLvlUp,
			"level":  "15",
			"id":     "162",
		},
	},
	"Hoothoot": {
		"Noctowl": {
			"method": MethodLvlUp,
			"level":  "20",
			"id":     "164",
		},
	},
	"Ledyba": {
		"Ledian": {
			"method": MethodLvlUp,
			"level":  "18",
			"id":     "166",
		},
	},
	"Spinarak": {
		"Ariados": {
			"method": MethodLvlUp,
			"level":  "22",
			"id":     "168",
		},
	},
	"Chinchou": {
		"Lanturn": {
			"method": MethodLvlUp,
			"level":  "27",
			"id":     "171",
		},
	},
	"Togepi": {
		"Togetic": {
			"method":     MethodFriendship,
			"friendship": "220",
			"id":         "176",
		},
	},
	"Natu": {
		"Xatu": {
			"method": MethodLvlUp,
			"level":  "25",
			"id":     "178",
		},
	},
	"Mareep": {
		"Flaaffy": {
			"method": MethodLvlUp,
			"level":  "15",
			"id":     "180",
		},
	},
	"Flaaffy": {
		"Ampharos": {
			"method": MethodLvlUp,
			"level":  "30",
			"id":     "181",
		},
	},
	"Azurill": {
		"Marill": {
			"method":     MethodFriendship,
			"friendship": "220",
			"id":         "183",
		},
	},
	"Azumarill": {
		"Xatu": {
			"method": MethodLvlUp,
			"level":  "18",
			"id":     "184",
		},
	},
	"Hoppip": {
		"Skiploom": {
			"method": MethodLvlUp,
			"level":  "18",
			"id":     "188",
		},
	},
	"Skiploom": {
		"Jumpluff": {
			"method": MethodLvlUp,
			"level":  "27",
			"id":     "189",
		},
	},
	"Sunkern": {
		"Sunflora": {
			"method": MethodStone,
			"id":     "192",
		},
	},
	"Wooper": {
		"Quagsire": {
			"method": MethodLvlUp,
			"level":  "20",
			"id":     "195",
		},
	},
	"Wynaut": {
		"Wobbuffet": {
			"method": MethodLvlUp,
			"level":  "15",
			"id":     "202",
		},
	},
	"Pineco": {
		"Forretress": {
			"method": MethodLvlUp,
			"level":  "31",
			"id":     "205",
		},
	},
	"Snubbull": {
		"Granbull": {
			"method": MethodLvlUp,
			"level":  "23",
			"id":     "210",
		},
	},
	"Teddiursa": {
		"Ursaring": {
			"method": MethodLvlUp,
			"level":  "30",
			"id":     "217",
		},
	},
	"Slugma": {
		"Magcargo": {
			"method": MethodLvlUp,
			"level":  "38",
			"id":     "219",
		},
	},
	"Swinub": {
		"Piloswine": {
			"method": MethodLvlUp,
			"level":  "33",
			"id":     "221",
		},
	},
	"Remoraid": {
		"Octillery": {
			"method": MethodLvlUp,
			"level":  "25",
			"id":     "224",
		},
	},
	"Houndour": {
		"Houndoom": {
			"method": MethodLvlUp,
			"level":  "24",
			"id":     "229",
		},
	},
	"Phanpy": {
		"Donphan": {
			"method": MethodLvlUp,
			"level":  "25",
			"id":     "232",
		},
	},
	"Larvitar": {
		"Pupitar": {
			"method": MethodLvlUp,
			"level":  "30",
			"id":     "247",
		},
	},
	"Pupitar": {
		"Tyranitar": {
			"method": MethodLvlUp,
			"level":  "55",
			"id":     "248",
		},
	},
	"Treecko": {
		"Grovyle": {
			"method": MethodLvlUp,
			"level":  "16",
			"id":     "277",
		},
	},
	"Grovyle": {
		"Sceptile": {
			"method": MethodLvlUp,
			"level":  "36",
			"id":     "278",
		},
	},
	"Torchic": {
		"Combusken": {
			"method": MethodLvlUp,
			"level":  "16",
			"id":     "281",
		},
	},
	"Combusken": {
		"Blaziken": {
			"method": MethodLvlUp,
			"level":  "36",
			"id":     "282",
		},
	},
	"Mudkip": {
		"Marshtomp": {
			"method": MethodLvlUp,
			"level":  "16",
			"id":     "284",
		},
	},
	"Marshtomp": {
		"Swampert": {
			"method": MethodLvlUp,
			"level":  "36",
			"id":     "285",
		},
	},
	"Poochyena": {
		"Mightyena": {
			"method": MethodLvlUp,
			"level":  "18",
			"id":     "287",
		},
	},
	"Zigzagoon": {
		"Linoone": {
			"method": MethodLvlUp,
			"level":  "20",
			"id":     "289",
		},
	},
	"Wurmple": {
		"Silcoon": {
			"method":    MethodWurmple,
			"level":     "7",
			"condition": ConditionLower,
			"id":        "291",
		},
		"Cascoon": {
			"method":    MethodWurmple,
			"level":     "7",
			"condition": ConditionHigher,
			"id":        "293",
		},
	},
	"Silcoon": {
		"Beautifly": {
			"method": MethodLvlUp,
			"level":  "10",
			"id":     "292",
		},
	},
	"Cascoon": {
		"Dustox": {
			"method": MethodLvlUp,
			"level":  "10",
			"id":     "294",
		},
	},
	"Lotad": {
		"Lombre": {
			"method": MethodLvlUp,
			"level":  "14",
			"id":     "296",
		},
	},
	"Lombre": {
		"Ludicolo": {
			"method": MethodStone,
			"id":     "297",
		},
	},
	"Seedot": {
		"Nuzleaf": {
			"method": MethodLvlUp,
			"level":  "14",
			"id":     "300",
		},
	},
	"Nuzleaf": {
		"Shiftry": {
			"method": MethodStone,
			"id":     "301",
		},
	},
	"Taillow": {
		"Swellow": {
			"method": MethodLvlUp,
			"level":  "22",
			"id":     "305",
		},
	},
	"Wingull": {
		"Pelipper": {
			"method": MethodLvlUp,
			"level":  "25",
			"id":     "310",
		},
	},
	"Ralts": {
		"Kirlia": {
			"method": MethodLvlUp,
			"level":  "20",
			"id":     "393",
		},
	},
	"Kirlia": {
		"Gardevoir": {
			"method": MethodLvlUp,
			"level":  "30",
			"id":     "394",
		},
	},
	"Surskit": {
		"Masquerain": {
			"method": MethodLvlUp,
			"level":  "22",
			"id":     "312",
		},
	},
	"Shroomish": {
		"Breloom": {
			"method": MethodLvlUp,
			"level":  "23",
			"id":     "307",
		},
	},
	"Slakoth": {
		"Vigoroth": {
			"method": MethodLvlUp,
			"level":  "18",
			"id":     "365",
		},
	},
	"Vigoroth": {
		"Slaking": {
			"method": MethodLvlUp,
			"level":  "36",
			"id":     "366",
		},
	},
	"Nincada": {
		"Slaking": {
			"method": MethodLvlUp,
			"level":  "20",
			"id":     "302",
		},
		"Shedinja": { // Can't find a way to properly do Shedinja
			"method": MethodLvlUp,
			"level":  "20",
			"id":     "303",
		},
	},
	"Whismur": {
		"Loudred": {
			"method": MethodLvlUp,
			"level":  "20",
			"id":     "371",
		},
	},
	"Loudred": {
		"Exploud": {
			"method": MethodLvlUp,
			"level":  "40",
			"id":     "372",
		},
	},
	"Makuhita": {
		"Hariyama": {
			"method": MethodLvlUp,
			"level":  "24",
			"id":     "335",
		},
	},
	"Skitty": {
		"Delcatty": {
			"method": MethodStone,
			"id":     "316",
		},
	},
	"Aron": {
		"Lairon": {
			"method": MethodLvlUp,
			"level":  "32",
			"id":     "383",
		},
	},
	"Lairon": {
		"Aggron": {
			"method": MethodLvlUp,
			"level":  "42",
			"id":     "384",
		},
	},
	"Meditite": {
		"Medicham": {
			"method": MethodLvlUp,
			"level":  "37",
			"id":     "357",
		},
	},
	"Electrike": {
		"Manectric": {
			"method": MethodLvlUp,
			"level":  "26",
			"id":     "338",
		},
	},
	"Gulpin": {
		"Swalot": {
			"method": MethodLvlUp,
			"level":  "26",
			"id":     "368",
		},
	},
	"Carvanha": {
		"Sharpedo": {
			"method": MethodLvlUp,
			"level":  "30",
			"id":     "331",
		},
	},
	"Wailmer": {
		"Wailord": {
			"method": MethodLvlUp,
			"level":  "40",
			"id":     "314",
		},
	},
	"Numel": {
		"Camerupt": {
			"method": MethodLvlUp,
			"level":  "33",
			"id":     "340",
		},
	},
	"Spoink": {
		"Grumpig": {
			"method": MethodLvlUp,
			"level":  "32",
			"id":     "352",
		},
	},
	"Trapinch": {
		"Vibrava": {
			"method": MethodLvlUp,
			"level":  "35",
			"id":     "333",
		},
	},
	"Vibrava": {
		"Flygon": {
			"method": MethodLvlUp,
			"level":  "45",
			"id":     "334",
		},
	},
	"Cacnea": {
		"Cacturne": {
			"method": MethodLvlUp,
			"level":  "32",
			"id":     "345",
		},
	},
	"Swablu": {
		"Altaria": {
			"method": MethodLvlUp,
			"level":  "35",
			"id":     "359",
		},
	},
	"Barboach": {
		"Whiscash": {
			"method": MethodLvlUp,
			"level":  "30",
			"id":     "324",
		},
	},
	"Corphish": {
		"Crawdaunt": {
			"method": MethodLvlUp,
			"level":  "30",
			"id":     "327",
		},
	},
	"Baltoy": {
		"Claydol": {
			"method": MethodLvlUp,
			"level":  "36",
			"id":     "319",
		},
	},
	"Lileep": {
		"Cradily": {
			"method": MethodLvlUp,
			"level":  "40",
			"id":     "389",
		},
	},
	"Anorith": {
		"Armaldo": {
			"method": MethodLvlUp,
			"level":  "40",
			"id":     "391",
		},
	},
	"Feebas": {
		"Milotic": {
			"method": MethodBeauty,
			"id":     "329",
		},
	},
}
