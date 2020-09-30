package model

import "github.com/jinzhu/gorm"

/*
	check whether there is nobody is assigned to the project
 */
func (pu *ProjectsUsers) Nobody_is_assigned_to_project(trans *gorm.DB) (bool) {
	err := trans.First(pu, "project_id = ?", pu.ProjectId).Error
	if err == gorm.ErrRecordNotFound {
		return true
	} else if err != nil {
		return false
	}

	return pu.UserId == 0
}
