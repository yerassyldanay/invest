package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"invest/utils"

	"time"
)

type ProjectsUsers struct {
	ProjectId				uint64					`json:"project_id" gorm:"foreignkey:projects.id"`
	UserId					uint64					`json:"user_id" gorm:"foreignkey:users.id"`
	Created							time.Time			`json:"created" gorm:"default:now()"`
}

func (ProjectsUsers) TableName() string {
	return "projects_users"
}

/*
	* do not allow to delete default users
	* do not allow assign investor to the project
 */
func (pu *ProjectsUsers) BeforeDelete(tx *gorm.DB) error {

	if pu.UserId <= utils.DefaultNotAllowedUserToDelete {
		return errorDafultUsersAreBeingAltered
	}

	return nil
}

func (pu *ProjectsUsers) BeforeCreate(tx *gorm.DB) error {

	var user = User{}
	err := GetDB().Preload("Role").First(&user, "id = ?", pu.UserId).Error
	if err != nil {
		return err
	}

	if user.Role.Name == utils.RoleInvestor {
		return errors.New("cannot assign investor to the project")
	}

	return nil
}
