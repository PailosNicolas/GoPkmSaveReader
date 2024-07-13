package helpers

var ExperienceType = map[string]string{
	"Bulbasaur":  "Medium-Slow",
	"Ivysaur":    "Medium-Slow",
	"Venusaur":   "Medium-Slow",
	"Charmander": "Medium-Slow",
	"Charmeleon": "Medium-Slow",
	"Charizard":  "Medium-Slow",
	"Squirtle":   "Medium-Slow",
	"Wartortle":  "Medium-Slow",
	"Blastoise":  "Medium-Slow",
	"Caterpie":   "Medium-Fast",
	"Metapod":    "Medium-Fast",
	"Butterfree": "Medium-Fast",
	"Weedle":     "Medium-Fast",
	"Kakuna":     "Medium-Fast",
	"Beedrill":   "Medium-Fast",
	"Pidgey":     "Medium-Slow",
	"Pidgeotto":  "Medium-Slow",
	"Pidgeot":    "Medium-Slow",
	"Rattata":    "Medium-Fast",
	"Raticate":   "Medium-Fast",
	"Spearow":    "Medium-Fast",
	"Fearow":     "Medium-Fast",
	"Ekans":      "Medium-Fast",
	"Arbok":      "Medium-Fast",
	"Pikachu":    "Medium-Fast",
	"Raichu":     "Medium-Fast",
	"Sandshrew":  "Medium-Fast",
	"Sandslash":  "Medium-Fast",
	"Nidoran-f":  "Medium-Slow",
	"Nidorina":   "Medium-Slow",
	"Nidoqueen":  "Medium-Slow",
	"Nidoran-m":  "Medium-Slow",
	"Nidorino":   "Medium-Slow",
	"Nidoking":   "Medium-Slow",
	"Clefairy":   "Fast",
	"Clefable":   "Fast",
	"Vulpix":     "Medium-Fast",
	"Ninetales":  "Medium-Fast",
	"Jigglypuff": "Fast",
	"Wigglytuff": "Fast",
	"Zubat":      "Medium-Fast",
	"Golbat":     "Medium-Fast",
	"Oddish":     "Medium-Slow",
	"Gloom":      "Medium-Slow",
	"Vileplume":  "Medium-Slow",
	"Paras":      "Medium-Fast",
	"Parasect":   "Medium-Fast",
	"Venonat":    "Medium-Fast",
	"Venomoth":   "Medium-Fast",
	"Diglett":    "Medium-Fast",
	"Dugtrio":    "Medium-Fast",
	"Meowth":     "Medium-Fast",
	"Persian":    "Medium-Fast",
	"Psyduck":    "Medium-Fast",
	"Golduck":    "Medium-Fast",
	"Mankey":     "Medium-Fast",
	"Primeape":   "Medium-Fast",
	"Growlithe":  "Slow",
	"Arcanine":   "Slow",
	"Poliwag":    "Medium-Slow",
	"Poliwhirl":  "Medium-Slow",
	"Poliwrath":  "Medium-Slow",
	"Abra":       "Medium-Slow",
	"Kadabra":    "Medium-Slow",
	"Alakazam":   "Medium-Slow",
	"Machop":     "Medium-Slow",
	"Machoke":    "Medium-Slow",
	"Machamp":    "Medium-Slow",
	"Bellsprout": "Medium-Slow",
	"Weepinbell": "Medium-Slow",
	"Victreebel": "Medium-Slow",
	"Tentacool":  "Slow",
	"Tentacruel": "Slow",
	"Geodude":    "Medium-Slow",
	"Graveler":   "Medium-Slow",
	"Golem":      "Medium-Slow",
	"Ponyta":     "Medium-Fast",
	"Rapidash":   "Medium-Fast",
	"Slowpoke":   "Medium-Fast",
	"Slowbro":    "Medium-Fast",
	"Magnemite":  "Medium-Fast",
	"Magneton":   "Medium-Fast",
	"Farfetchd":  "Medium-Fast",
	"Doduo":      "Medium-Fast",
	"Dodrio":     "Medium-Fast",
	"Seel":       "Medium-Fast",
	"Dewgong":    "Medium-Fast",
	"Grimer":     "Medium-Fast",
	"Muk":        "Medium-Fast",
	"Shellder":   "Slow",
	"Cloyster":   "Slow",
	"Gastly":     "Medium-Slow",
	"Haunter":    "Medium-Slow",
	"Gengar":     "Medium-Slow",
	"Onix":       "Medium-Fast",
	"Drowzee":    "Medium-Fast",
	"Hypno":      "Medium-Fast",
	"Krabby":     "Medium-Fast",
	"Kingler":    "Medium-Fast",
	"Voltorb":    "Medium-Fast",
	"Electrode":  "Medium-Fast",
	"Exeggcute":  "Slow",
	"Exeggutor":  "Slow",
	"Cubone":     "Medium-Fast",
	"Marowak":    "Medium-Fast",
	"Hitmonlee":  "Medium-Fast",
	"Hitmonchan": "Medium-Fast",
	"Lickitung":  "Medium-Fast",
	"Koffing":    "Medium-Fast",
	"Weezing":    "Medium-Fast",
	"Rhyhorn":    "Slow",
	"Rhydon":     "Slow",
	"Chansey":    "Fast",
	"Tangela":    "Medium-Fast",
	"Kangaskhan": "Medium-Fast",
	"Horsea":     "Medium-Fast",
	"Seadra":     "Medium-Fast",
	"Goldeen":    "Medium-Fast",
	"Seaking":    "Medium-Fast",
	"Staryu":     "Slow",
	"Starmie":    "Slow",
	"Mr-mime":    "Medium-Fast",
	"Scyther":    "Medium-Fast",
	"Jynx":       "Medium-Fast",
	"Electabuzz": "Medium-Fast",
	"Magmar":     "Medium-Fast",
	"Pinsir":     "Slow",
	"Tauros":     "Slow",
	"Magikarp":   "Slow",
	"Gyarados":   "Slow",
	"Lapras":     "Slow",
	"Ditto":      "Medium-Fast",
	"Eevee":      "Medium-Fast",
	"Vaporeon":   "Medium-Fast",
	"Jolteon":    "Medium-Fast",
	"Flareon":    "Medium-Fast",
	"Porygon":    "Medium-Fast",
	"Omanyte":    "Medium-Fast",
	"Omastar":    "Medium-Fast",
	"Kabuto":     "Medium-Fast",
	"Kabutops":   "Medium-Fast",
	"Aerodactyl": "Slow",
	"Snorlax":    "Slow",
	"Articuno":   "Slow",
	"Zapdos":     "Slow",
	"Moltres":    "Slow",
	"Dratini":    "Slow",
	"Dragonair":  "Slow",
	"Dragonite":  "Slow",
	"Mewtwo":     "Slow",
	"Mew":        "Medium-Slow",
	"Chikorita":  "Medium-Slow",
	"Bayleef":    "Medium-Slow",
	"Meganium":   "Medium-Slow",
	"Cyndaquil":  "Medium-Slow",
	"Quilava":    "Medium-Slow",
	"Typhlosion": "Medium-Slow",
	"Totodile":   "Medium-Slow",
	"Croconaw":   "Medium-Slow",
	"Feraligatr": "Medium-Slow",
	"Sentret":    "Medium-Fast",
	"Furret":     "Medium-Fast",
	"Hoothoot":   "Medium-Fast",
	"Noctowl":    "Medium-Fast",
	"Ledyba":     "Fast",
	"Ledian":     "Fast",
	"Spinarak":   "Fast",
	"Ariados":    "Fast",
	"Crobat":     "Medium-Fast",
	"Chinchou":   "Slow",
	"Lanturn":    "Slow",
	"Pichu":      "Medium-Fast",
	"Cleffa":     "Fast",
	"Igglybuff":  "Fast",
	"Togepi":     "Fast",
	"Togetic":    "Fast",
	"Natu":       "Medium-Fast",
	"Xatu":       "Medium-Fast",
	"Mareep":     "Medium-Slow",
	"Flaaffy":    "Medium-Slow",
	"Ampharos":   "Medium-Slow",
	"Bellossom":  "Medium-Slow",
	"Marill":     "Fast",
	"Azumarill":  "Fast",
	"Sudowoodo":  "Medium-Fast",
	"Politoed":   "Medium-Slow",
	"Hoppip":     "Medium-Slow",
	"Skiploom":   "Medium-Slow",
	"Jumpluff":   "Medium-Slow",
	"Aipom":      "Fast",
	"Sunkern":    "Medium-Slow",
	"Sunflora":   "Medium-Slow",
	"Yanma":      "Medium-Fast",
	"Wooper":     "Medium-Fast",
	"Quagsire":   "Medium-Fast",
	"Espeon":     "Medium-Fast",
	"Umbreon":    "Medium-Fast",
	"Murkrow":    "Medium-Slow",
	"Slowking":   "Medium-Fast",
	"Misdreavus": "Fast",
	"Unown":      "Medium-Fast",
	"Wobbuffet":  "Medium-Fast",
	"Girafarig":  "Medium-Fast",
	"Pineco":     "Medium-Fast",
	"Forretress": "Medium-Fast",
	"Dunsparce":  "Medium-Fast",
	"Gligar":     "Medium-Slow",
	"Steelix":    "Medium-Fast",
	"Snubbull":   "Fast",
	"Granbull":   "Fast",
	"Qwilfish":   "Medium-Fast",
	"Scizor":     "Medium-Fast",
	"Shuckle":    "Medium-Slow",
	"Heracross":  "Slow",
	"Sneasel":    "Medium-Slow",
	"Teddiursa":  "Medium-Fast",
	"Ursaring":   "Medium-Fast",
	"Slugma":     "Medium-Fast",
	"Magcargo":   "Medium-Fast",
	"Swinub":     "Slow",
	"Piloswine":  "Slow",
	"Corsola":    "Fast",
	"Remoraid":   "Medium-Fast",
	"Octillery":  "Medium-Fast",
	"Delibird":   "Fast",
	"Mantine":    "Slow",
	"Skarmory":   "Slow",
	"Houndour":   "Slow",
	"Houndoom":   "Slow",
	"Kingdra":    "Medium-Fast",
	"Phanpy":     "Medium-Fast",
	"Donphan":    "Medium-Fast",
	"Porygon2":   "Medium-Fast",
	"Stantler":   "Slow",
	"Smeargle":   "Fast",
	"Tyrogue":    "Medium-Fast",
	"Hitmontop":  "Medium-Fast",
	"Smoochum":   "Medium-Fast",
	"Elekid":     "Medium-Fast",
	"Magby":      "Medium-Fast",
	"Miltank":    "Slow",
	"Blissey":    "Fast",
	"Raikou":     "Slow",
	"Entei":      "Slow",
	"Suicune":    "Slow",
	"Larvitar":   "Slow",
	"Pupitar":    "Slow",
	"Tyranitar":  "Slow",
	"Lugia":      "Slow",
	"Ho-oh":      "Slow",
	"Celebi":     "Medium-Slow",
	"Treecko":    "Medium-Slow",
	"Grovyle":    "Medium-Slow",
	"Sceptile":   "Medium-Slow",
	"Torchic":    "Medium-Slow",
	"Combusken":  "Medium-Slow",
	"Blaziken":   "Medium-Slow",
	"Mudkip":     "Medium-Slow",
	"Marshtomp":  "Medium-Slow",
	"Swampert":   "Medium-Slow",
	"Poochyena":  "Medium-Fast",
	"Mightyena":  "Medium-Fast",
	"Zigzagoon":  "Medium-Fast",
	"Linoone":    "Medium-Fast",
	"Wurmple":    "Medium-Fast",
	"Silcoon":    "Medium-Fast",
	"Beautifly":  "Medium-Fast",
	"Cascoon":    "Medium-Fast",
	"Dustox":     "Medium-Fast",
	"Lotad":      "Medium-Slow",
	"Lombre":     "Medium-Slow",
	"Ludicolo":   "Medium-Slow",
	"Seedot":     "Medium-Slow",
	"Nuzleaf":    "Medium-Slow",
	"Shiftry":    "Medium-Slow",
	"Taillow":    "Medium-Slow",
	"Swellow":    "Medium-Slow",
	"Wingull":    "Medium-Fast",
	"Pelipper":   "Medium-Fast",
	"Ralts":      "Slow",
	"Kirlia":     "Slow",
	"Gardevoir":  "Slow",
	"Surskit":    "Medium-Fast",
	"Masquerain": "Medium-Fast",
	"Shroomish":  "Fluctuating",
	"Breloom":    "Fluctuating",
	"Slakoth":    "Slow",
	"Vigoroth":   "Slow",
	"Slaking":    "Slow",
	"Nincada":    "Erratic",
	"Ninjask":    "Erratic",
	"Shedinja":   "Erratic",
	"Whismur":    "Medium-Slow",
	"Loudred":    "Medium-Slow",
	"Exploud":    "Medium-Slow",
	"Makuhita":   "Fluctuating",
	"Hariyama":   "Fluctuating",
	"Azurill":    "Fast",
	"Nosepass":   "Medium-Fast",
	"Skitty":     "Fast",
	"Delcatty":   "Fast",
	"Sableye":    "Medium-Slow",
	"Mawile":     "Fast",
	"Aron":       "Slow",
	"Lairon":     "Slow",
	"Aggron":     "Slow",
	"Meditite":   "Medium-Fast",
	"Medicham":   "Medium-Fast",
	"Electrike":  "Slow",
	"Manectric":  "Slow",
	"Plusle":     "Medium-Fast",
	"Minun":      "Medium-Fast",
	"Volbeat":    "Erratic",
	"Illumise":   "Fluctuating",
	"Roselia":    "Medium-Slow",
	"Gulpin":     "Fluctuating",
	"Swalot":     "Fluctuating",
	"Carvanha":   "Slow",
	"Sharpedo":   "Slow",
	"Wailmer":    "Fluctuating",
	"Wailord":    "Fluctuating",
	"Numel":      "Medium-Fast",
	"Camerupt":   "Medium-Fast",
	"Torkoal":    "Medium-Fast",
	"Spoink":     "Fast",
	"Grumpig":    "Fast",
	"Spinda":     "Fast",
	"Trapinch":   "Medium-Slow",
	"Vibrava":    "Medium-Slow",
	"Flygon":     "Medium-Slow",
	"Cacnea":     "Medium-Slow",
	"Cacturne":   "Medium-Slow",
	"Swablu":     "Erratic",
	"Altaria":    "Erratic",
	"Zangoose":   "Erratic",
	"Seviper":    "Fluctuating",
	"Lunatone":   "Fast",
	"Solrock":    "Fast",
	"Barboach":   "Medium-Fast",
	"Whiscash":   "Medium-Fast",
	"Corphish":   "Fluctuating",
	"Crawdaunt":  "Fluctuating",
	"Baltoy":     "Medium-Fast",
	"Claydol":    "Medium-Fast",
	"Lileep":     "Erratic",
	"Cradily":    "Erratic",
	"Anorith":    "Erratic",
	"Armaldo":    "Erratic",
	"Feebas":     "Erratic",
	"Milotic":    "Erratic",
	"Castform":   "Medium-Fast",
	"Kecleon":    "Medium-Slow",
	"Shuppet":    "Fast",
	"Banette":    "Fast",
	"Duskull":    "Fast",
	"Dusclops":   "Fast",
	"Tropius":    "Slow",
	"Chimecho":   "Fast",
	"Absol":      "Medium-Slow",
	"Wynaut":     "Medium-Fast",
	"Snorunt":    "Medium-Fast",
	"Glalie":     "Medium-Fast",
	"Spheal":     "Medium-Slow",
	"Sealeo":     "Medium-Slow",
	"Walrein":    "Medium-Slow",
	"Clamperl":   "Erratic",
	"Huntail":    "Erratic",
	"Gorebyss":   "Erratic",
	"Relicanth":  "Slow",
	"Luvdisc":    "Fast",
	"Bagon":      "Slow",
	"Shelgon":    "Slow",
	"Salamence":  "Slow",
	"Beldum":     "Slow",
	"Metang":     "Slow",
	"Metagross":  "Slow",
	"Regirock":   "Slow",
	"Regice":     "Slow",
	"Registeel":  "Slow",
	"Latias":     "Slow",
	"Latios":     "Slow",
	"Kyogre":     "Slow",
	"Groudon":    "Slow",
	"Rayquaza":   "Slow",
	"Jirachi":    "Slow",
	"Deoxys":     "Slow",
}
