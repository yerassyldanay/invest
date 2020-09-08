package model

type Project struct {
	Id					uint64					`json:"id" gorm:"primary key"`

	Name				string						`json:"name" gorm:"unique" validate:"required"`
	Description			string						`json:"description" default:"''" validate:"required"`

	Info				string						`json:"info" default:"''"`
	InfoSent			map[string]interface{}		`json:"info_sent" gorm:"-"`

	EmployeeCount			uint					`json:"employee_count" validate:"required"`

	Email					string						`json:"email" gorm:"default:''"`
	PhoneNumber				string						`json:"phone_number" default:"''"`

	OrganizationId			uint64					`json:"organization_id"`
	Organization			Organization			`json:"organization" gorm:"foreignkey:OrganizationId"`

	User				[]User					`json:"user" gorm:"many2many:projects_users;"`
	Documents			[]Document				`json:"documents" gorm:"-"`
	Categors			[]Categor				`json:"categors" gorm:"many2many:projects_categors"`

	OfferedById					uint64					`json:"offered_by_id" gorm:"not null"`
	OfferedByPosition			string					`json:"offered_by_position" gorm:"not null"`

	Status				string					`json:"status" gorm:"default:'not confirmed'"`

	AddInfo
}

func (Project) TableName() string {
	return "projects"
}

