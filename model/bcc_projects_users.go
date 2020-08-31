package model

import (
	"github.com/jinzhu/gorm"
	"invest/utils"
)

type ProjectsUsers struct {
	ProjectId				uint64		`json:"project_id"`
	UserId					uint64			`json:"user_id"`
}

func (ProjectsUsers) TableName() string {
	return "projects_users"
}

/*
	do not allow to delete default users
 */
func (pu ProjectsUsers) BeforeDelete(tx *gorm.DB) error {

	if pu.UserId <= utils.DefaultNotAllowedUserToDelete {
		return errorDafultUsersAreBeingAltered
	}

	return nil
}
