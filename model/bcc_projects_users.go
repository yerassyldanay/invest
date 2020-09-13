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
	//Step1Confirmed					bool			`json:"step_1_confirmed" gorm:"column:step_1_confirmed;default:false"`
	//Step2Confirmed					bool			`json:"step_2_confirmed" gorm:"column:step_2_confirmed;default:false"`
	Created							time.Time			`json:"created" gorm:"default:now()"`
}

func (ProjectsUsers) TableName() string {
	return "projects_users"
}

/*
	* do not allow to delete default users
	* do not allow assign investor to the project
 */
func (pu ProjectsUsers) BeforeDelete(tx *gorm.DB) error {

	if pu.UserId <= utils.DefaultNotAllowedUserToDelete {
		return errorDafultUsersAreBeingAltered
	}

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
