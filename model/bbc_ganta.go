package model

import "time"

/*
	table of Ganta
 */
type Ganta struct {
	Id							uint64					`json:"id" gorm:"primary key"`
	
	IsDefault					bool					`json:"is_default" gorm:"default:false"`
	
	ProjectId					uint64					`json:"project_id"`
	Project						Project					`json:"project" gorm:"foreignkey:ProjectId"`

	NameId						uint64						`json:"name_id"`
	Name						GantaName					`json:"name" gorm:"foreignkey:GantaNameId"`
	
	Start						time.Time				`json:"start" gorm:"not null"`
	Deadline					int						`json:"deadline" gorm:"-"`
}

type GantaName struct {
	Kaz					string					`json:"kaz"`
	Rus					string					`json:"rus"`
	Eng					string					`json:"eng"`
}

func (Ganta) TableName() string {
	return "gantas"
}

type GantaUpDate struct {
	Day					time.Duration				`json:"day" validate:"max=10,min=-10"`
	Hour				time.Duration				`json:"hour" validate:"max=24,min=-24"`

	Ganta
}

