package service

import (
	"github.com/jinzhu/gorm"
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
	var currentGanta = model.Ganta{ProjectId: project_id}
	if err := currentGanta.OnlyGetCurrentStepByProjectId(model.GetDB()); err == gorm.ErrRecordNotFound {
		var project = model.Project{Id: project_id}
		_ = project.OnlyGetById(model.GetDB())
		project.CurrentStep = model.DefaultGantaFinalStep

		var resp = utils.NoErrorFineEverthingOk
		resp["info"] = model.Struct_to_map(project)

		return model.ReturnNoErrorWithResponseMessage(resp)
	} else if err != nil {
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
			err = currentGanta.OnlyChangeStatusToDoneAndUpdateDeadlineById(trans)
			// update the status of the project
			project.Status = utils.ProjectStatusAccept
			err = project.OnlyUpdateStatusById(trans)
		case status == utils.ProjectStatusReconsider:
			// prepare gantt step
			newGanta := model.Ganta{
				IsAdditional:   true,
				ProjectId:      project_id,
				Kaz:            "Доработка инициатором проекта",
				Rus:            "Доработка инициатором проекта",
				Eng:            "Доработка инициатором проекта",
				DurationInDays: 3,
				Step:           currentGanta.Step,
				Status:         utils.ProjectStatusPendingInvestor,
				StartDate: 		utils.GetCurrentTime(),
				Deadline: 		utils.GetCurrentTime().Add(time.Hour * 24 * 3),
				IsDone:         false,
				Responsible:    utils.RoleInvestor,
			}

			// create a new gantt step
			if err = newGanta.OnlyCreate(trans); err != nil {
				return model.ReturnInternalDbError(err.Error())
			}

			// add a new step to the top
			// by indicating that current one is done
			if err = currentGanta.OnlyChangeStatusToDoneAndUpdateDeadlineById(trans); err != nil {
				return model.ReturnInternalDbError(err.Error())
			}

			// change the status of the project
			project.Status = utils.ProjectStatusReconsider
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

		// anyway case we have to switch to the next gantt step
		err = currentGanta.OnlyChangeStatusToDoneAndUpdateDeadlineById(trans)
	default:
		return model.ReturnMethodNotAllowed("role is not supported. role is " + is.RoleName)
	}

	if err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	/*
		there is a problem, the project status must be updated
	 */
	nextGanta := model.Ganta{ProjectId: project_id}
	if err = nextGanta.OnlyGetCurrentStepByProjectId(trans); err == gorm.ErrRecordNotFound	{
		// pass
	} else if err != nil {
		return model.ReturnInternalDbError(err.Error())
	} else {
		nextGanta.Deadline.Add(time.Hour * 24 * nextGanta.DurationInDays)
	}

	if err := trans.Commit().Error; err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	project.Id = project_id
	_ = project.GetAndUpdateStatusOfProject(model.GetDB())

	return model.ReturnNoError()
}

func (is *InvestService) Ganta_change_time(ganta model.Ganta) (utils.Msg) {

	// validate
	if ganta.Start < 1 {
		return model.ReturnInvalidParameters("gantt start time is not valid")
	}

	// set start time
	ganta.StartDate = time.Unix(ganta.Start, 0)

	// update gantt step time
	if err := ganta.OnlyUpdateTimeByIdAndProjectId(model.GetDB()); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	return model.ReturnNoError()
}

