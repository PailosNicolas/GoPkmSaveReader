package helpers

import (
	"testing"
)

func TestExperienceTypes(t *testing.T) {
	for _, specie := range Species {
		if specie != "?" && specie != "??????????" && specie != "Pokémon Egg" {
			_, ok := ExperienceType[specie]

			if !ok {
				t.Error("There is no experience type for:", specie)
			}
		}
	}
}

func TestBaseStats(t *testing.T) {
	for _, specie := range Species {
		if specie != "?" && specie != "??????????" && specie != "Pokémon Egg" {
			_, ok := BaseStatsGenIII[specie]

			if !ok {
				t.Error("There is no base stat type for:", specie)
			}
		}
	}
}

func TestLevelFromExperience(t *testing.T) {
	exp := 20822

	tre := GetExperience()
	level := GetLevelFromExperience(tre["Medium-Fast"], exp)

	if level != 27 {
		t.Error("It should be lvl 27 but it is:", level)
	}
}
