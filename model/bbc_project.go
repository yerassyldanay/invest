package model

import "time"

type Project struct {
	Id					uint64					`json:"id" gorm:"primary key"`

	Name				string					`json:"name" gorm:"unique" validate:"required"`
	Created				time.Time				`json:"date" gorm:"default:now()"`

	Description			string					`json:"description" default:"''" validate:"required"`
	Info				string					`json:"info" default:"''"`
	InfoSent			map[string]interface{}	`json:"info_sent" gorm:"-"`

	EmployeeCount			uint					`json:"employee_count" validate:"required"`

	Email				string					`json:"email" gorm:"default:''"`
	Ccode				string					`json:"ccode" gorm:"default:''"`
	Phone				string					`json:"phone" gorm:"default:''"`

	OrganizationId		uint64					`json:"organization_id"`
	Organization		Organization			`json:"organization" gorm:"foreignkey:OrganizationId"`

	CreatedBy				uint64					`json:"-" gorm:"-"`
	User				[]User					`json:"user" gorm:"many2many:projects_users;"`

	Documents			[]Document				`json:"documents" gorm:"-"`
	Deleted				string					`json:"deleted" gorm:"default:null"`

	Categors			[]Categor				`json:"categors" gorm:"many2many:projects_categors"`

	Status				string					`json:"status" gorm:"default:'not confirmed'"`
	//ApprovedBy			[]User					`json:"approved_by" gorm:"many2many:projects_statuses;"`
}

func (Project) TableName() string {
	return "projects"
}

