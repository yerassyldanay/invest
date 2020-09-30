package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

func (pu *ProjectsUsers) Chech_whether_user_is_assigned_to_project(tx *gorm.DB) (error) {
	var count int
	err := tx.Table("projects_users").Where("project_id = ? and user_id = ?", pu.ProjectId, pu.UserId).Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("there is no such user & project relation")
	}

	return nil
}

func (pu *ProjectsUsers) OnlyCreate(tx *gorm.DB) (err error) {
	err = tx.Create(pu).Error
	return err
}

func (pu *ProjectsUsers) OnlyDelete(tx *gorm.DB) (err error) {
	err = tx.Delete(pu, "project_id = ? and user_id = ?", pu.ProjectId, pu.UserId).Error
	return err
}

func (pu *ProjectsUsers) OnlyDeleteRelation (tx *gorm.DB) (err error) {
	err = tx.Delete(pu, "project_di = ? and user_id = ?", pu.ProjectId, pu.UserId).Error
	return err
}
