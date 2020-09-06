package model

import "invest/utils"

var GantaDefaultSteps = []Ganta{
	{
		Name: GantaName {
			Kaz:   "",
			Rus:   "",
			Eng:   "An admin is assigning users to the project",
		},
		IsDefault: true,
		Start:     utils.GetCurrentTime(),
	},
	{
		Name: GantaName{
			Kaz:   "",
			Rus:   "",
			Eng:   "The first step | Manager confirmation",
		},
		IsDefault: true,
		Start:     utils.GetCurrentTime(),
	},
	{
		Name: GantaName{
			Kaz:   "",
			Rus:   "",
			Eng:   "Confirmation from a lawyer and financier",
		},
		IsDefault: true,
		Start:     utils.GetCurrentTime(),
	},
	{
		Name: GantaName{
			Kaz:   "",
			Rus:   "",
			Eng:   "Second Step | Confirmation from a manager",
		},
		IsDefault: true,
		Start:     utils.GetCurrentTime(),
	},
	{
		Name: GantaName{
			Kaz:   "",
			Rus:   "",
			Eng:   "Confirmation from a board",
		},
		IsDefault: true,
		Start:     utils.GetCurrentTime(),
	},
}
