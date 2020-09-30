package model

import "time"

/*
	there are two main steps
			and many sub-steps

	there will be two statuses:
		* status of the ganta = status of the parent
		* status of the document, which will be considered as the status of the child ganta step

	ManuallyChangeable:
		* 0 - unchangeable
		* 1 - only admins can change ~
		* 2 - admins & users, who are responsible, can change ~
		* 3 - all users can change ~
				~ the status of the ganta step
 */

type Ganta struct {
	Id								uint64					`json:"id" gorm:"primary key"`

	IsAdditional					bool					`json:"is_additional" gorm:"default:false"`
	ProjectId						uint64					`json:"project_id" gorm:"foreignkey:projects.id"`
	//Project							Project					`json:"project" gorm:"foreignkey:ProjectId"`

	Kaz								string					`json:"kaz" gorm:"default:''"`
	Rus								string					`json:"rus" gorm:"default:''"`
	Eng								string					`json:"eng" gorm:"default:''"`

	Start 							int64					`json:"start" gorm:"-"`
	StartDate						time.Time				`json:"start_date" gorm:"default:now()"`
	DurationInDays					time.Duration			`json:"duration_in_days"`
	
	Document						Document				`json:"document"`
	
	GantaParentId					uint64					`json:"ganta_parent_id"`
	GantaChildren					[]Ganta					`json:"ganta_children" gorm:"-"`

	//IsHidden						bool					`json:"is_hidden" gorm:"default:true"`
	Step							int						`json:"step" gorm:"default:1"`
	Status							string					`json:"status" gorn:"default:'prending_manager'"`

	IsDone							bool					`json:"is_done" gorm:"default:false"`
	Responsible						string					`json:"responsible" gorm:"default:'spk'"`
	IsDocCheck						bool					`json:"-" gorm:"default:false"`
}

func (Ganta) TableName() string {
	return "gantas"
}

