package model

type Role struct {
	Id				uint64				`json:"id" gorm:"AUTO_INCREMENT;primary key"`
	Name			string				`json:"name" validate:"required"`
	Description		string				`json:"description" validate:"required"`
	Permissions		[]Permission		`json:"permissions" gorm:"many2many:roles_permissions"`
}

func (Role) TableName() string {
	return "roles"
}
