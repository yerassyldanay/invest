package model

import "time"

type ProjectStatus struct {
	ProjectId				uint64				`json:"project_id"`
	Project					Project				`json:"project" gorm:"foreignkey:ProjectId"`

	UserId					uint64				`json:"user_id"`
	User					User				`json:"user" gorm:"foreignkey:UserId"`

	Status						string				`json:"status" gorm:"default:'not confirmed'"`
	Modified					time.Time			`json:"date" gorm:"default:now()"`
	
	//IsScoped					bool				`json:"is_scoped"`
	Deadline					time.Time			`json:"deadline" gorm:"default:null"`
	Notified					time.Time			`json:"notified" gorm:"default:null"`
}

func (ps ProjectStatus) TableName() string {
	return "projects_statuses"
}
