package model

import (
	"invest/utils"
)

func (pu *ProjectsUsers) Assign_user_after_check() (utils.Msg) {
	// check whether a project exists
	var project = Project{Id: pu.ProjectId}
	err := project.OnlyGetById(GetDB())
	if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	// check whether a user exists (preloaded means gets also role, email & phone)
	var user = User{Id: pu.UserId}
	err = user.OnlyGetByIdPreloaded(GetDB())
	if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	// check that user in (manager or expert)
	if user.Role.Name != utils.RoleExpert && user.Role.Name != utils.RoleManager {
		return ReturnMethodNotAllowed("wrong user, user is neither expert or manager")
	}

	// after a long pondering, assign
	if err = pu.OnlyCreate(GetDB()); err != nil {
		return ReturnInternalDbError(err.Error())
	}

	return ReturnNoError()
}

func (pu *ProjectsUsers) Remove_relation() (utils.Msg) {
	if err := pu.OnlyDelete(GetDB()); err  != nil {
		return ReturnInternalDbError(err.Error())
	}

	return ReturnNoError()
}
