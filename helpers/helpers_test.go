package helpers

import (
	"testing"
)

func TestExperienceTypes(t *testing.T) {
	for _, specie := range Species {
		if specie != "?" && specie != "??????????" {
			_, ok := ExperienceType[specie]

			if !ok {
				t.Error("There is no experience type for:", specie)
			}
		}
	}
}
