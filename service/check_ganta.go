package service

import (
	"invest/model"
	"invest/utils"
)

func (is *InvestService) Ganta_check_permission_to_read_ganta(project_id uint64) (utils.Msg) {
	/*
		users:
			* with administrate privileges
			* assigned to the project
			* investor of the project
	*/
	var permission = model.Permission{}
	var project = model.Project{
		Id: project_id,
		OfferedById: is.UserId,
	}

	switch {
	case permission.Check_permission_by_role_id(is.RoleId, "administrate", model.GetDB()) == nil:
		// nothing to do - this is a user with admin privileges
	case project.OnlyCheckInvestorByProjectAndInvestorId(model.GetDB()) == nil:
		// nothing to do - this is an investor of the project
	case project.OnlyCheckUserByProjectAndUserId(project_id, is.UserId, model.GetDB()) == nil:
		// this user is assigned to the project
	default:
		return model.ReturnMethodNotAllowed("not admin, investor of the project or spk user, who is assigned to it")
	}

	return model.ReturnNoError()
}

func (is *InvestService) Check_who_can_shift(project_id uint64) (utils.Msg) {
	return utils.Msg{}
}