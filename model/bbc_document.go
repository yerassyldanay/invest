package model

import (
	"time"
)

type Document struct {
	Id							uint64				`json:"id" gorm:"AUTO_INCREMENT; primary_key"`

	Created						time.Time			`json:"date" gorm:"default:now()"`
	Uri							string				`json:"uri"`

	ProjectId					uint64 				`json:"project_id" gorm:"foreignkey:projects.id"`
	GantaId						uint64				`json:"ganta_id"`
}

func (Document) TableName() string {
	return "documents"
}
