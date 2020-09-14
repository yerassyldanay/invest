package model

import (
	"invest/utils"
)

/*
	this is to assign user to project
 */
func (pu *ProjectsUsers) Assign_user_to_project() (utils.Msg) {

	var trans = GetDB().Begin()
	defer func() {if trans != nil {trans.Rollback()}}()

	if pu.Is_right_that_nobody_is_assigned_to_project(trans) {
		var project = Project{Id: pu.ProjectId}
		err := project.Change_by_project_id_status_to(utils.ProjectStatusInprogress, trans)
		if err != nil {
			return utils.Msg{utils.ErrorInternalDbError, 417, "",  err.Error()}
		}
	}

	var main_query = "insert into " + pu.TableName() + " (project_id, user_id) values(?, ?);"
	n := trans.Exec(main_query, pu.ProjectId, pu.UserId).RowsAffected
	if n != 1 {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", "could not insert values into table"}
	}

	trans.Commit()
	trans = nil

	return utils.Msg{utils.NoErrorFineEverthingOk, 200, "", ""}

}

/*
	remove user from project
 */
func (pu *ProjectsUsers) Remove_user_from_project() (utils.Msg) {
	//var main_query = ` delete from projects_users where project_id = ? and user_id = ? ; `

	var trans = GetDB().Begin()
	defer func() {if trans != nil {trans.Rollback()}}()

	err := trans.Delete(pu, " project_id = ? and user_id = ? ", pu.ProjectId, pu.UserId).Error
	if err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	if pu.Is_right_that_nobody_is_assigned_to_project(trans) {
		var project = Project{Id: pu.ProjectId}
		err := project.Change_by_project_id_status_to(utils.ProjectStatusNewone, trans)
		if err != nil {
			return utils.Msg{utils.ErrorInternalDbError, 417, "",  err.Error()}
		}
	}

	trans.Commit()
	trans = nil

	return utils.Msg{utils.NoErrorFineEverthingOk, 200, "", ""}
}

