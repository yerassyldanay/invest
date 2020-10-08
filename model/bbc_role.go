package model

import "github.com/jinzhu/gorm"

type Role struct {
	Id				uint64				`json:"id" gorm:"AUTO_INCREMENT;primary key"`
	Name			string				`json:"name" validate:"required"`
	Description		string				`json:"description"`
	Permissions		[]Permission		`json:"permissions" gorm:"many2many:roles_permissions"`
	PermissionsSent		[]uint64			`json:"permissions_sent" gorm:"-"`
}

func (Role) TableName() string {
	return "roles"
}

func (r *Role) OnlyGetByName(tx *gorm.DB) (err error) {
	err = tx.First(r, "name = ?", r.Name).Error
	return err
}

func (r *Role) OnlyGetById(tx *gorm.DB) (err error) {
	err = tx.First(r, "id = ?", r.Id).Error
	return err
}

func (r *Role) OnlyCreate(tx *gorm.DB) (err error) {
	err = tx.Create(r).Error
	return err
}

func (r *Role) OnlySave(tx *gorm.DB) (err error) {
	err = tx.Save(r).Error
	return err
}

/*
	get permissions by id of role table
*/
func (r *Role) Get_permissions_by_role_id() error {
	return GetDB().Exec("select p.* from roles r inner join roles_permissions rp on r.id = rp.role_id " +
		"inner join permissions p on p.id = rp.permission_id where r.id =?", r.Id).Find(&r.Permissions).Error
}
