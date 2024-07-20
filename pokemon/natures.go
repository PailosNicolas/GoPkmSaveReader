package pokemon

import "github.com/PailosNicolas/GoPkmSaveReader/helpers"

var natures = []Nature{
	{
		ID: 0,

		Name: "Hardy",

		Increases: helpers.AttackString,

		Decreases: helpers.AttackString,
	},
	{
		ID: 1,

		Name: "Lonely",

		Increases: helpers.AttackString,

		Decreases: helpers.DefenseString,
	},
	{
		ID: 2,

		Name: "Brave",

		Increases: helpers.AttackString,

		Decreases: helpers.SpeedString,
	},
	{
		ID: 3,

		Name: "Adamant",

		Increases: helpers.AttackString,

		Decreases: helpers.SpecialAttackString,
	},
	{
		ID: 4,

		Name: "Naughty",

		Increases: helpers.AttackString,

		Decreases: helpers.SpecialDefenseString,
	},
	{
		ID: 5,

		Name: "Bold",

		Increases: helpers.DefenseString,

		Decreases: helpers.AttackString,
	},
	{
		ID: 6,

		Name: "Docile",

		Increases: helpers.DefenseString,

		Decreases: helpers.DefenseString,
	},
	{
		ID: 7,

		Name: "Relaxed",

		Increases: helpers.DefenseString,

		Decreases: helpers.SpeedString,
	},
	{
		ID: 8,

		Name: "Impish",

		Increases: helpers.DefenseString,

		Decreases: helpers.SpecialAttackString,
	},
	{
		ID: 9,

		Name: "Lax",

		Increases: helpers.DefenseString,

		Decreases: helpers.SpecialDefenseString,
	},
	{
		ID: 10,

		Name: "Timid",

		Increases: helpers.SpeedString,

		Decreases: helpers.AttackString,
	},
	{
		ID: 11,

		Name: "Hasty",

		Increases: helpers.SpeedString,

		Decreases: helpers.DefenseString,
	},
	{
		ID: 12,

		Name: "Serious",

		Increases: helpers.SpeedString,

		Decreases: helpers.SpeedString,
	},
	{
		ID: 13,

		Name: "Jolly",

		Increases: helpers.SpeedString,

		Decreases: helpers.SpecialAttackString,
	},
	{
		ID: 14,

		Name: "Naive",

		Increases: helpers.SpeedString,

		Decreases: helpers.SpecialDefenseString,
	},
	{
		ID: 15,

		Name: "Modest",

		Increases: helpers.SpecialAttackString,

		Decreases: helpers.AttackString,
	},
	{
		ID: 16,

		Name: "Mild",

		Increases: helpers.SpecialAttackString,

		Decreases: helpers.DefenseString,
	},
	{
		ID: 17,

		Name: "Quiet",

		Increases: helpers.SpecialAttackString,

		Decreases: helpers.SpeedString,
	},
	{
		ID: 18,

		Name: "Bashful",

		Increases: helpers.SpecialAttackString,

		Decreases: helpers.SpecialAttackString,
	},
	{
		ID: 19,

		Name: "Rash",

		Increases: helpers.SpecialAttackString,

		Decreases: helpers.SpecialDefenseString,
	},
	{
		ID: 20,

		Name: "Calm",

		Increases: helpers.SpecialDefenseString,

		Decreases: helpers.AttackString,
	},
	{
		ID: 21,

		Name: "Gentle",

		Increases: helpers.SpecialDefenseString,

		Decreases: helpers.DefenseString,
	},
	{
		ID: 22,

		Name: "Sassy",

		Increases: helpers.SpecialDefenseString,

		Decreases: helpers.SpeedString,
	},
	{
		ID: 23,

		Name: "Careful",

		Increases: helpers.SpecialDefenseString,

		Decreases: helpers.SpecialAttackString,
	},
	{
		ID: 24,

		Name: "Quirky",

		Increases: helpers.SpecialDefenseString,

		Decreases: helpers.SpecialDefenseString,
	},
}
