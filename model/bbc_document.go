package model

import (
	"time"
)

type Document struct {
	Id							uint64				`json:"id" gorm:"AUTO_INCREMENT; primary_key"`
	Name						string				`json:"name" validate:"required"`
	
	//Info						string				`json:"info" gorm:"default:'{}'"`
	//InfoSent					map[string]interface{}		`json:"info_sent" gorm:"-"`

	Created						time.Time			`json:"date" gorm:"default:now()"`
	Uri							string				`json:"uri"`

	ProjectId					uint64 				`json:"project_id" gorm:"foreignkey:projects.id"`
	GantaId						uint64				`json:"ganta_id"`
	
	Status						string				`json:"status" gorm:"default:'inprogress'"`
	ChangesMadeById						uint64				`json:"-" gorm:"-"`
}

func (Document) TableName() string {
	return "documents"
}
