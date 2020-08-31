package model

import (
	"errors"
	"invest/utils"
	"strings"
)

/*
	this is to assign user to project
 */
func (pu *ProjectsUsers) Assign_user_to_project() (map[string]interface{}, error) {
	if pu.ProjectId == 0 || pu.UserId == 0 {
		return utils.ErrorInvalidParameters, errors.New("invalid parameters passed. assgin project")
	}

	var role = Role{}
	if err := GetDB().Table(User{}.TableName()).Select("roles.*").Joins("join roles on roles.id = users.role_id").Where("roles.id = ?", pu.UserId).First(&role).Error;
		err != nil || strings.ToLower(role.Name) == utils.RoleInvestor {
			return utils.ErrorMethodNotAllowed, err
	}

	var main_query = "insert into " + pu.TableName() + "(project_id, user_id) values(?, ?);"
	if n := GetDB().Exec(main_query, pu.ProjectId, pu.UserId).RowsAffected; n == 1 {
		return utils.NoErrorFineEverthingOk, nil
	}

	return utils.ErrorInternalDbError, errors.New("could not insert values into table")
}

/*
	remove user from project
 */
func (pu *ProjectsUsers) Remove_user_from_project() (map[string]interface{}, error) {
	//var main_query = ` delete from projects_users where project_id = ? and user_id = ? ; `

	err := GetDB().Delete(&ProjectsUsers{}, " project_id = ? and user_id = ? ", pu.ProjectId, pu.UserId).Error
	if err != nil {
		return utils.ErrorInternalDbError, err
	}

	return utils.NoErrorFineEverthingOk, nil
}

