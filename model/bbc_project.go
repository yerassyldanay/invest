package model

type Project struct {
	Id					uint64					`json:"id" gorm:"primary key"`

	Name				string						`json:"name" gorm:"index:unique_project" validate:"required"`
	Description			string						`json:"description" gorm:"default:''; index:unique_project" validate:"required"`

	Info				string						`json:"info" default:"''"`
	InfoSent			map[string]interface{}		`json:"info_sent" gorm:"-"`
	
	EmployeeCount			uint					`json:"employee_count" validate:"required"`

	Email					string						`json:"email" gorm:"default:''"`
	PhoneNumber				string						`json:"phone_number" default:"''"`

	OrganizationId			uint64					`json:"organization_id"`
	Organization			Organization			`json:"organization" gorm:"foreignkey:OrganizationId"`

	Users					[]User					`json:"user" gorm:"many2many:projects_users;"`
	Documents				[]Document				`json:"documents" gorm:"-"`
	Categors				[]Categor				`json:"categors" gorm:"many2many:projects_categors"`

	OfferedById					uint64					`json:"offered_by_id" gorm:"not null"`
	OfferedByPosition			string					`json:"offered_by_position" gorm:"not null"`

	Reject						bool					`json:"-" gorm:"default:false"`
	Reconsider					bool					`json:"-" gorm:"default:false"`
	Completed					bool					`json:"-" gorm:"default:false"`
	
	Status						string					`json:"status" gorm:"default:'pending_admin'"`
	Step						int						`json:"stage"`
	
	LandPlotFrom				string					`json:"land_plot_from" gorm:"default:'investor'"`
	LandArea 						int						`json:"land_area"`
	LandAddress					string					`json:"land_address"`
	
	CurrentStep						Ganta					`json:"current_ganta_step" gorm:"-"`
	
	AddInfo
}

type ProjectWithFinanceTables struct {
	Project						Project					`json:"project"`
	Cost						Cost					`json:"cost"`
	Finance						Finance					`json:"finance"`
}

func (Project) TableName() string {
	return "projects"
}

/*
	status:

 */
