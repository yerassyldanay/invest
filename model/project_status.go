package model

import (
	"bytes"
	"invest/utils"
	"net/http"
	"time"
)

/*
	this helps to assign users to project
		users are those, who can set a status to a project
 */
func (ps *ProjectStatus) Create_project_user_connection_to_set_status() (*utils.Msg) {
	if ps.ProjectId	== 0 || ps.UserId == 0 {
		return &utils.Msg{
			utils.ErrorInvalidParameters, http.StatusBadRequest, "", "project is or/and user id is not correct",
		}
	}

	ps.Status = utils.ProjectStatusNotConfirmed
	ps.Modified = time.Now()
	ps.Deadline = time.Now().Add(time.Hour * utils.ProjectStatusChangeTimeInHours)

	if err := GetDB().Create(ProjectStatus{}).Error; err != nil {
		return &utils.Msg{
			utils.ErrorInternalDbError, http.StatusExpectationFailed, "", err.Error(),
		}
	}

	return &utils.MsgNoErrorMessageOk
}

/*

 */
func (*ProjectStatus) Create_a_bulk_of_project_status_rows(pss []ProjectStatus) (*utils.Msg) {
	var errmsg bytes.Buffer
	for _, ps := range pss {
		if msg := ps.Create_project_user_connection_to_set_status(); msg.ErrMsg != "" {
			errmsg.WriteString(msg.ErrMsg)
			errmsg.WriteString(" ")
		}
	}

	return &utils.Msg{
		Message: utils.NoErrorFineEverthingOk,
		Status:  utils.If_condition_then(errmsg.String() != "", http.StatusMultiStatus, http.StatusOK).(int),
		ErrMsg:  errmsg.String(),
	}
}

/*
	deadline:
		case 1:
			'not confirmed' - a manager or other users have a certain amount of time to review & take action on it
			'returned for revision' - an investor will get nottifications in that case
		case 2:
			'blocked' or 'confirmed' - nobody gets notifications
 */
func (ps *ProjectStatus) Update_status_by_project_and_user_id() (*utils.Msg) {
	var deadline = time.Now().Add(time.Hour * utils.ProjectStatusChangeTimeInHours)
	if 	ps.Status == utils.ProjectStatusConfirmed || ps.Status == utils.ProjectStatusBlocked {
		deadline = time.Time{}
	}

	/*
		statuses are limited and predefined
		if none of predefined statuses is provided then set one of them
	 */
	if ps.Status != utils.ProjectStatusConfirmed && ps.Status != utils.ProjectStatusBlocked &&
			ps.Status != utils.ProjectStatusReturnedToChange && ps.Status != utils.ProjectStatusNotConfirmed {
		ps.Status = utils.ProjectStatusNotConfirmed
	}

	if err := GetDB().Table(ProjectStatus{}.TableName()).
		Where("project_id=$1 and user_id=$2", ps.ProjectId, ps.UserId).
		Update(ProjectStatus{
			Status:   ps.Status,
			Modified: time.Now(),
			Deadline: deadline,
		}).Error;
			err != nil {
				return &utils.Msg{
					utils.ErrorInternalDbError, http.StatusExpectationFailed, "", err.Error(),
				}
	}

	return &utils.MsgNoErrorMessageOk
}
