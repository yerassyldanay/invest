package model

import "time"

/*
	there are two main steps
			and many sub-steps

	there will be two statuses:
		* status of the ganta = status of the parent
		* status of the document, which will be considered as the status of the child ganta step
 */
type Ganta struct {
	Id								uint64					`json:"id" gorm:"primary key"`

	IsAdditional					bool					`json:"is_default" gorm:"default:false"`
	ProjectId						uint64					`json:"project_id" gorm:"foreignkey:projects.id"`
	//Project							Project					`json:"project" gorm:"foreignkey:ProjectId"`

	Kaz								string					`json:"kaz" gorm:"default:''"`
	Rus								string					`json:"rus" gorm:"default:''"`
	Eng								string					`json:"eng" gorm:"default:''"`

	Start 							int64					`json:"start;omitempty" gorm:"-"`
	StartDate						time.Time				`json:"start_date" gorm:"default:now()"`
	DurationInDays						int						`json:"duration_in_days"`

	GantaParentId					uint64					`json:"ganta_parent_id"`

	//DocumentId						uint64					`json:"document_id"`
	Document					Document				`json:"document" foreignkey:"DocumentId"`

	Status							string					`json:"status" gorn:"default:'newone'"`
}

func (Ganta) TableName() string {
	return "gantas"
}

type GantaUpDate struct {
	Day					time.Duration				`json:"day" validate:"max=10,min=-10"`
	Hour				time.Duration				`json:"hour" validate:"max=24,min=-24"`
	
	UserId				uint64						`json:"user_id"`
	
	Ganta
}

