package service

import (
	"invest/model"
	"invest/utils"
	"strings"
	"time"
)

func (is *InvestService) Ganta_can_user_change_current_status(project_id uint64) (ganta model.Ganta, msg utils.Msg) {

	var project = model.Project{Id: project_id}

	// get project with an updated status
	err := project.GetAndUpdateStatusOfProject(model.GetDB())
	if err != nil {
		return model.Ganta{}, model.ReturnInternalDbError(err.Error())
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

/*
	there are several cases that might happen after any of the users changes the status of the project
 */
func (is *InvestService) Ganta_change_the_status_of_project(project_id uint64, status string) (utils.Msg) {

	// get the current gantt step
	var ganta = model.Ganta{ProjectId: project_id}
	if err := ganta.OnlyGetCurrentStepByProjectId(model.GetDB()); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	var trans = model.GetDB().Begin()
	defer func() { if trans != nil {trans.Rollback()} }()

	var err error
	var project = model.Project{Id: project_id}
	status = strings.ToLower(status)

	// whose is changing the status
	switch {
	case is.RoleName == utils.RoleInvestor:
		// investor never can change status manually
		return model.ReturnMethodNotAllowed("investor cannot change status")
	case is.RoleName == utils.RoleAdmin:
		/*
			choices of an admin:
				* reject - status of the project will be reject
				* reconsider - create new gantt step & add status - pending_investor
										responsible - investor
				* accept - move to the next step
		 */
		switch {
		case status == utils.ProjectStatusReject:
			// status will change to 'reject'
			project.Status = utils.ProjectStatusReject
			err = project.OnlyUpdateStatusById(trans)
		case status == utils.ProjectStatusAccept:
			// go to the next step
			err = ganta.OnlyChangeStatusToDoneById(trans)
		case status == utils.ProjectStatusReconsider:
			// prepare gantt step
			newGanta := model.Ganta{
				IsAdditional:   true,
				ProjectId:      project_id,
				Kaz:            "Доработка инициатором проекта",
				Rus:            "Доработка инициатором проекта",
				Eng:            "Доработка инициатором проекта",
				DurationInDays: 3,
				Step:           ganta.Step,
				Status:         utils.ProjectStatusPendingInvestor,
				StartDate: 		ganta.StartDate.Add(time.Hour * (-1)),
				IsDone:         false,
				Responsible:    utils.RoleInvestor,
			}

			if err = newGanta.OnlyCreate(trans); err != nil {
				return model.ReturnInternalDbError(err.Error())
			}

			// add new step to the top
			if err = ganta.OnlyChangeStatusToDoneById(trans); err != nil {
				return model.ReturnInternalDbError(err.Error())
			}

			// change the status of the user
			project.Status = status
			err = project.OnlyUpdateStatusById(trans)
		}

	case is.RoleName == utils.RoleManager || is.RoleName == utils.RoleExpert:
		//var project = model.Project{Id: project_id}
		switch {
		case status == utils.ProjectStatusReject:
			// set pre-reject status as this must be confirmed by admin
			project.Status = utils.ProjectStatusPreliminaryReject
			err = project.OnlyUpdateStatusById(trans)
		case status == utils.ProjectStatusReconsider:
			// set pre-reject status as this must be confirmed by admin
			project.Status = utils.ProjectStatusPreliminaryReconsider
			err = project.OnlyUpdateStatusById(trans)
		case status == utils.ProjectStatusAccept:

		}

		if err != nil {
			return model.ReturnInternalDbError(err.Error())
		}

		// in any case we have to switch to the next gantt step
		err = ganta.OnlyChangeStatusToDoneById(trans)
	}

	if err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	if err := trans.Commit().Error; err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	project.Id = project_id
	_ = project.GetAndUpdateStatusOfProject(model.GetDB())

	return model.ReturnNoError()
}
