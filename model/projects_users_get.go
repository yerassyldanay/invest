package model

import "github.com/jinzhu/gorm"

func (pu *ProjectsUsers) Chech_whether_user_is_assigned_to_project(tx *gorm.DB) (error) {
	return tx.Table("projects_users").Where("project_id = ? and user_id = ?", pu.ProjectId, pu.UserId).Error
}

//func (pu *ProjectsUsers)
