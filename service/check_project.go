package service

import (
	"invest/model"
	"invest/utils/constants"
	"invest/utils/message"
)

// check whether this user has privileges to get project data
func (is *InvestService) Check_whether_this_user_can_get_access_to_project_info(project_id uint64)(message.Msg) {
	var err error
	var user = model.User{Id: is.UserId}

	switch {
	case is.RoleName == constants.RoleAdmin:
		return model.ReturnNoError()
	case is.RoleName == constants.RoleInvestor:
		err = user.DoesOwnThisProjectById(project_id, model.GetDB())
	default:
		err = user.IsAssignedToThisProjectById(project_id, model.GetDB())
	}

	if err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	return model.ReturnNoError()
}

/*
	Is the user is not responsible for current step, he/she is not allowed to make changes
	for example:
		* remove documents
 */
func (is *InvestService) Check_whether_this_user_responsible_for_current_step(project_id uint64) (message.Msg) {
	var project = model.Project{Id: project_id}
	err := project.GetAndUpdateStatusOfProject(model.GetDB())
	if err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	var responsible = project.CurrentStep.Responsible

	if is.RoleName == constants.RoleAdmin {
		// pass
	} else if responsible != is.RoleName {
		return model.ReturnMethodNotAllowed("responsible: " + responsible + " | your role: " + is.RoleName)
	}

	return model.ReturnNoError()
}

func (is *InvestService) Does_project_exist(project_id uint64) (message.Msg) {
	// check whether a project exists
	var project = model.Project{Id: project_id}
	err := project.OnlyGetById(model.GetDB())
	if err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	return model.ReturnNoError()
}
