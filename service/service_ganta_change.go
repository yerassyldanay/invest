package service

import (
	"invest/model"
	"invest/utils"
	"strings"
)

func (is *InvestService) Ganta_can_user_change_current_status(project_id uint64) (ganta model.Ganta, msg utils.Msg) {

	var project = model.Project{Id: project_id}

	// get project with an updated status
	_ = project.GetAndUpdateStatusOfProject(model.GetDB())

	// check for validity
	switch {
	case project.Id < 1:
		return model.Ganta{}, model.ReturnInvalidParameters("project_id is invalid")
	case (project.Reject || project.Reconsider) && is.RoleName != utils.RoleAdmin:
		/*
			in case if it is rejected, let admins change the status of the project
				cases:
					* manager claims rejected, but admin might not agree
		*/
		return model.Ganta{}, model.ReturnMethodNotAllowed("the project is either rejected or in consideration by an investor")
	}

	ganta = project.CurrentStep

	/*
		Security check, to pass:
			* responsible != investor
			* role-name == responsible
			* role-name == admin & responsible - spk
			* admin
	*/
	switch {
	case ganta.Responsible == utils.RoleInvestor:
		return ganta, model.ReturnMethodNotAllowed("this step cannot be passed manually")
	case is.RoleName == utils.RoleAdmin:
		// pass
	case ganta.Responsible == utils.RoleSpk:
		// this must not happen
	case ganta.Responsible == is.RoleName:
		// pass
	default:
		return ganta, model.ReturnMethodNotAllowed("you cannot change the status of the project")
	}

	// a manager or expert is trying to change the status
	// while he/she has not yet checked documents (IsDocCheck == true means there are documents to check)
	if (is.RoleName == utils.RoleManager || is.RoleName == utils.RoleExpert) && ganta.IsDocCheck {
		return ganta, model.ReturnMethodNotAllowed("a manager or expert is trying to change the status, while not having documents checked")
	}

	msg = model.ReturnNoError()
	return ganta, msg
}

func (is *InvestService) Ganta_change_the_status_of_project(project_id uint64, status string) (utils.Msg) {
	// check permission
	ganta, msg := is.Ganta_can_user_change_current_status(project_id)
	if msg.IsThereAnError() {
		return msg
	}

	var err error
	status = strings.ToLower(status)

	switch {
	case is.RoleName == utils.RoleInvestor:
		// investor never can change status manually
		return model.ReturnMethodNotAllowed("investor cannot change status")
	case is.RoleName == utils.RoleAdmin:
		/*
			admin can change status:
				* reject - then nobody can change it back
				* reconsider - investor need make changes in documentation | admin can change status back
				* accept - next step
		 */
		switch strings.ToLower(status) {
		case utils.ProjectStatusReject:
			err = ganta.OnlySetRejectStatusForProjectByProjectId(model.GetDB())
		case utils.ProjectStatusReconsider:
			err = ganta.OnlySetReconsiderStatusForProjectByProjectId(model.GetDB())
		case utils.ProjectStatusAccept:
			_ = ganta.OnlyUpdateRejectStatusByProjectId(false, model.GetDB())
			_ = ganta.OnlyUpdateReconsiderStatusByProjectId(false, model.GetDB())
			err = ganta.OnlyChangeStatusToDoneById(model.GetDB())
		default:
			return model.ReturnMethodNotAllowed("status is invalid. it is " + status)
		}

	case is.RoleName == utils.RoleManager || is.RoleName == utils.RoleExpert:
		/*
			manager & expert:
				* reject - 'reject' field is set to true
				* reconsider - 'reconsider' field is set to true
				* accept - move to the next step
		 */
		switch status {
		case utils.ProjectStatusReject:
			err = ganta.OnlyUpdateRejectStatusByProjectId(true, model.GetDB())
		case utils.ProjectStatusReconsider:
			err = ganta.OnlyUpdateReconsiderStatusByProjectId(true, model.GetDB())
		case utils.ProjectStatusAccept:
			err = ganta.OnlyChangeStatusToDoneById(model.GetDB())
		default:
			return model.ReturnMethodNotAllowed("status is invalid. it is " + status)
		}
	}

	if err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	return model.ReturnNoError()
}
