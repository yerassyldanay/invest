package model

import (
	"invest/utils"
	"time"
)

var DefaultGantaParent1 = Ganta {
	IsAdditional:  false,
	ProjectId:     0,  // will be set
	Kaz:           "", // will be set as An admin is assigning users to the project
	Rus:           "", // the same
	Eng:           "Step 1",
	StartDate:     utils.GetCurrentTime(),
	Duration:      10,
	GantaParentId: 0,
	Status:        utils.ProjectStatusInprogress,
}

var DefaultGantaParent2 = Ganta{
	IsAdditional:  false,
	Kaz:           "",
	Rus:           "",
	Eng:           "Step 2",
	StartDate:     utils.GetCurrentTime().Add(time.Hour * 24 * 10),
	Duration:      10,
	GantaParentId: 0,
	Status:        utils.ProjectStatusInprogress,
}

var DefaultGantaChildren1 = []Ganta{
	{
		IsAdditional:  false,
		Kaz:           "",
		Rus:           "",
		Eng:           "Child 1 Step 1",
		StartDate:     utils.GetCurrentTime(),
		Duration:      3,
		GantaParentId: 0,			// will be set
		DocumentId:    0,
		Document:      Document{},
		Status:        "",			// this status will not be considered instead consider the status of the child
	},
	{
		IsAdditional:  false,
		Kaz:           "",
		Rus:           "",
		Eng:           "Child 2 Step 1",
		StartDate:     utils.GetCurrentTime(),
		Duration:      3,
		GantaParentId: 0,
		Status:        "",
	},
}

var DefaultGantaChildren2 = []Ganta{
	{
		IsAdditional:  false,
		Kaz:           "",
		Rus:           "",
		Eng:           "Child 1 Step 2",
		StartDate:     utils.GetCurrentTime(),
		Duration:      3,
		GantaParentId: 0,			// will be set
		DocumentId:    0,
		Document:      Document{},
		Status:        "",			// this status will not be considered instead consider the status of the child
	},
	{
		IsAdditional:  false,
		Kaz:           "",
		Rus:           "",
		Eng:           "Child 2 Step 2",
		StartDate:     utils.GetCurrentTime(),
		Duration:      3,
		GantaParentId: 0,
		Status:        "",
	},
}
