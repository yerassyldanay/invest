package model

import "time"

type Document struct {
	Id							uint64				`json:"id" gorm:"AUTO_INCREMENT; primary_key"`

	Name						string				`json:"name" gorm:"unique_index:unique_doc_per_project" validate:"required"`
	//Format						string				`json:"format"`
	//Directory					string				`json:"directory"`
	
	Info						string				`json:"info" gorm:"default:'{}'"`
	InfoSent					map[string]interface{}		`json:"info_sent" gorm:"-"`

	Created						time.Time			`json:"date" gorm:"default:now()"`
	Url							string				`json:"url"`

	ProjectId					uint64 				`json:"project_id" gorm:"unique_index:unique_doc_per_project"`
	Project						Project				`json:"project" gorm:"foreignkey:ProjectId"`

	Type						string				`json:"status" gorm:"default:'docs'"`
	Deleted						time.Time			`json:"deleted" gorm:"default:null"`
}

func (Document) TableName() string {
	return "documents"
}
