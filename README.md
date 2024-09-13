![tests_status](https://github.com/PailosNicolas/GoPkmSaveReader/actions/workflows/tests.yml/badge.svg)
# GoPkmSaveReader
This is a Pokemon save reader, it's far from complete and currently only supports Gen III. The goal is to be able to get
all the save data, export pokemons for backup purposes and import them into saves. Depending on progress thinking
about trades and trade evolution wouldn't be far fetched.

## Usage
Get it with:
```bash
go get github.com/PailosNicolas/GoPkmSaveReader@v0.1.0
```

Example:
```go
package main

import (
	"github.com/PailosNicolas/GoPkmSaveReader/savereader"
)

func main(){
	path := "../myPKMNsave"

	save, err := savereader.ReadDataFromSave(path)

	if err != nil {
		println("error.")
		return 
	}

	println(save.Trainer.Name())
}
```

## Features
### Gen III:
Right now you are able to access the following data of a save:
#### General save data:
- Game code
- Time played (Done coming in next release)
#### Trainer info:
- Public/secret ID
- Name
- Gender
- Team size
- Team
- Money (Done coming in next release)
#### Pokemon data:
- Nickname
- Species
- Level
- Held item
- Experience
- Friendship
- Moves
- Evs
- Ivs
- Stats
- Met location
- Met level
- Game of origin
- Captured ball
- Nature (Done coming in next release)

Also you are able to export a pokemon to a file.

## What is coming ?
Currently my focus is in getting all the pokemon data, next step would be getting and completing data of the boxed pokemons.
- Complete pokemon data
- Boxed pokemons (Done coming in next release)
- Complete Trainer information
- Improve testing (it is really bare bones right now)
- Import pokemon to a save file (Done coming in next release)
- Implement trade between save files (possibly trade evolution)

In the far away future if all of the above mentioned features are implemented I'll start research on the Gen IV save files.
