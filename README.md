![tests_status](https://github.com/PailosNicolas/GoPkmSaveReader/actions/workflows/tests.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/PailosNicolas/GoPkmSaveReader.svg)](https://pkg.go.dev/github.com/PailosNicolas/GoPkmSaveReader)
# GoPkmSaveReader
This is a Pokemon save reader, it's far from complete and currently only supports Gen III. The goal is to be able to get
all the save data, export pokemons for backup purposes and import them into saves. Depending on progress thinking
about trades and trade evolution wouldn't be far fetched.

## Usage
Get it with:
```bash
go get github.com/PailosNicolas/GoPkmSaveReader@v0.2.0
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
- Time played
- Import pokemon to save (A bit finicky due to boxed vs team pokemons, can be used to simulate a trade and trade evolution with the `EvolvePokemon` method)
- Export save
#### Trainer info:
- Public/secret ID
- Name
- Gender
- Team size
- Team
- PC boxes
- Money
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
- Nature
- Export pokemon
- Evolve pokemon

## What is coming ?
My focus has shifted significantly toward creating a feature-rich package, but it currently lacks sufficient testing, which may result in bugs. For now, I will refrain from adding new features and concentrate on writing as many tests as possible, applying hotfixes when necessary.

- Improve testing (Priority number one !)
- Complete pokemon data
- Complete Trainer information
