package service

import (
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/utils/constants"
	"github.com/yerassyldanay/invest/utils/helper"
	"github.com/yerassyldanay/invest/utils/message"
	"strings"
	"time"
)

func (is *InvestService) Ganta_can_user_change_current_status(project_id uint64) (ganta model.Ganta, msg message.Msg) {

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
	case ganta.Responsible == constants.RoleInvestor:
		return ganta, model.ReturnMethodNotAllowed("this step cannot be passed manually")
	case is.RoleName == constants.RoleAdmin:
		// pass
	case ganta.Responsible == constants.RoleSpk:
		// this must not happen
	case ganta.Responsible == is.RoleName:
		// pass
	default:
		return ganta, model.ReturnMethodNotAllowed("you cannot change the status of the project")
	}

	// a manager or expert is trying to change the status
	// while he/she has not yet checked documents (IsDocCheck == true means there are documents to check)
	if (is.RoleName == constants.RoleManager || is.RoleName == constants.RoleExpert) && ganta.IsDocCheck {
		return ganta, model.ReturnMethodNotAllowed("a manager or expert is trying to change the status, while not having documents checked")
	}

	msg = model.ReturnNoError()
	return ganta, msg
}

/*
	there are several cases that might happen after any of the users changes the status of the project
 */
func (is *InvestService) Ganta_change_the_status_of_project(project_id uint64, status string) (message.Msg) {
	// validate status
	switch {
	case status == constants.ProjectStatusAccept:
	case status == constants.ProjectStatusReject:
	case status == constants.ProjectStatusReconsider:
	default:
		return model.ReturnInvalidParameters("invalid status, status is " + status)
	}

	// get the current gantt step
	var currentGanta = model.Ganta{ProjectId: project_id}
	if err := currentGanta.OnlyGetCurrentStepByProjectId(model.GetDB()); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	var lastGantaStep = currentGanta

	// there two cases when nobody is responsible
	if currentGanta.Responsible == constants.RoleNobody {
		return model.ReturnMethodNotAllowed("the project is either rejected or there is no more gantt step")
	}

	// create transaction
	var trans = model.GetDB().Begin()
	defer func() { if trans != nil {trans.Rollback()} }()

	var err error
	status = strings.ToLower(status)

	// whose is changing the status
	switch {
	case is.RoleName == constants.RoleInvestor:
		// investor never can change status manually
		return model.ReturnMethodNotAllowed("investor cannot change status")
	case is.RoleName == constants.RoleAdmin || is.RoleName == constants.RoleManager || is.RoleName == constants.RoleExpert:
		/*
			choices:
				* reject: a new gantt step will be created and put in front of all (with status of 'reject')
						thus the status of the project will be the status of this gantt step
				* reconsider: a new gantt step will be created (status: pending_investor),
						the current gantt step will be put after this step
				* accept: move to the next step
		 */
		switch {
		case status == constants.ProjectStatusReject:
			// create a new gantt step, which indicates that the project has been rejected
			// it helps when dealing with notifications
			newGanta := model.Ganta{
				IsAdditional:   false,
				ProjectId:      project_id,
				Kaz:            "Жоба қабылданбады",
				Rus:            "Проект отклонен",
				Eng:            "Project has been rejected",
				StartDate:      helper.GetCurrentTruncatedDate(),
				DurationInDays: 3,
				Deadline:       time.Time{}, // to avoid sending notifications
				Step:           4,
				Status:         constants.ProjectStatusReject,
				IsDone:         false,
				Responsible:    constants.RoleNobody,
				NotToShow:      true,
			}

			// creates this gantt step
			err = newGanta.OnlyCreate(trans)
			if err != nil {
				return model.ReturnInternalDbError(err.Error())
			}

			// set current status to done
			// it must not appear at the top after ordering by start date
			err = currentGanta.OnlyChangeStatusToDoneAndUpdateDeadlineById(trans)
			if err != nil {
				return model.ReturnInternalDbError(err.Error())
			}

		case status == constants.ProjectStatusAccept:
			// set that the current step is done
			if err = currentGanta.OnlyChangeStatusToDoneAndUpdateDeadlineById(trans); err != nil {
				return model.ReturnInternalDbError(err.Error())
			}

			/*
				if the current step is done, we need to shift all gantt steps to left
					this is what we are doing here
			 */

			// now the current step is
			// at the end this step will the default final step
			err = currentGanta.OnlyGetCurrentStepByProjectId(trans);
			if err != nil {
				return model.ReturnInternalDbError(err.Error())
			}

			// difference in hour between the current
			var difference = int(helper.GetCurrentTruncatedDate().Sub(currentGanta.StartDate).Hours())

			// shift all gantt step to the current date (to left or to right)
			// they must start from this date
			err = currentGanta.OnlyUpdateStartDatesOfAllUndoneGantaStepsByProjectId(difference, trans)
			if err != nil {
				return model.ReturnInternalDbError(err.Error())
			}

		case status == constants.ProjectStatusReconsider:
			daysGivenToInvestor := time.Duration(15)

			// prepare gantt step
			newGanta := model.Ganta{
				IsAdditional:   true,
				ProjectId:      project_id,
				Kaz:            "Жоба ұсынушының қарауында",
				Rus:            "Доработка инициатором проекта",
				Eng:            "Review by a project initiator",
				DurationInDays: daysGivenToInvestor,
				Step:           currentGanta.Step,
				Status:         constants.ProjectStatusPendingInvestor,
				StartDate:      helper.GetCurrentTruncatedDate(),
				Deadline:       helper.GetCurrentTruncatedDate().Add(time.Hour * 24 * daysGivenToInvestor),
				IsDone:         false,
				Responsible:    constants.RoleInvestor,
			}

			// shift all gantt steps to right
			// calculate shift hour
			hoursToShift := int(helper.GetCurrentTruncatedDate().Sub(currentGanta.StartDate).Hours())
			hoursToShift = hoursToShift + int(daysGivenToInvestor * 24)

			// shift all gantt steps
			if err = currentGanta.OnlyUpdateStartDatesOfAllUndoneGantaStepsByProjectId(hoursToShift, trans); err != nil {
				return model.ReturnInternalDbError(err.Error())
			}

			// create a new gantt step
			if err = newGanta.OnlyCreate(trans); err != nil {
				return model.ReturnInternalDbError(err.Error())
			}
		}

	default:
		return model.ReturnMethodNotAllowed("role is not supported. role is " + is.RoleName)
	}

	if err := trans.Commit().Error; err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	var project = model.Project{Id: project_id}
	project.Id = project_id
	_ = project.GetAndUpdateStatusOfProject(model.GetDB())

	// send notification
	notifyStatusChangeMessage := &model.NotifyProjectStatus{
		UserId:        is.UserId,
		ProjectId:     project_id,
		Project:       project,
		LastGantaStep: lastGantaStep,
	}

	// send message (handles everything: stores on database, prepares smtp message,
	// gets smtp server credentials, dials to smtp server & sends message )
	select {
	case model.GetMailerQueue().NotificationChannel <- notifyStatusChangeMessage:
	default:
	}

	return model.ReturnNoError()
}

func (is *InvestService) Ganta_change_time(ganta model.Ganta) (message.Msg) {

	// validate
	if ganta.Start < 1 {
		return model.ReturnInvalidParameters("gantt start time is not valid")
	}

	// set start time
	ganta.StartDate = time.Unix(ganta.Start, 0)

	// update gantt step time
	if err := ganta.OnlyUpdateStartDateById(model.GetDB()); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	return model.ReturnNoError()
}

