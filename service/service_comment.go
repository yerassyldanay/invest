package service

import (
	"invest/model"
	"invest/utils"
)

func (is InvestService) Comment_on_project_documents(comment model.Comment) (utils.Msg) {
	/*
		check whether this is a user, who is assigned to the project
	 */
	var pu = model.ProjectsUsers {
		ProjectId: comment.ProjectId,
		UserId:    comment.UserId,
	}
	if err := pu.Chech_whether_user_is_assigned_to_project(model.GetDB()); err != nil {
		return utils.Msg{utils.ErrorMethodNotAllowed, 405, "", err.Error()}
	}

	var trans = model.GetDB().Begin()
	defer func() {if trans != nil { trans.Rollback() }} ()

	/*
		check validity of a comment
	*/
	if err := comment.Validate(); err != nil {
		return utils.Msg{utils.ErrorInvalidParameters, 400, "", err.Error()}
	}

	/*
		update statuses of a document
	*/
	var document = model.Document{
		ProjectId: comment.ProjectId,
	}
	err := document.Only_update_statuses(comment.DocStatuses, trans)
	if err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	/*
		just create the comment
	 */
	if err := comment.Only_create(trans); err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	/*
		commit changes & catch an error
	 */
	if err := trans.Commit().Error; err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	trans = nil
	return utils.Msg{utils.NoErrorFineEverthingOk, 200, "", ""}
}
