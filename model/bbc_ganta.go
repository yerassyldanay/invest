package model

import "time"

/*
	table of Ganta
 */
type Ganta struct {
	Id					uint64					`json:"id" gorm:"primary key"`

	ProjectId			uint64					`json:"project_id"`
	Project				Project					`json:"project" gorm:"foreignkey:ProjectId"`

	Kaz					string					`json:"kaz"`
	Rus					string					`json:"rus"`
	Eng					string					`json:"eng"`

	Start				time.Time				`json:"start" gorm:"not null"`
}

func (Ganta) TableName() string {
	return "gantas"
}

type GantaUpDate struct {
	Day					time.Duration				`json:"day" validate:"max=10,min=-10"`
	Hour				time.Duration				`json:"hour" validate:"max=24,min=-24"`
	//Minute				time.Duration				`json:"minute" validate:"max=60,min=-60"`

	Ganta
}