package service

import (
	"invest/model"
	"invest/utils"
)

// comment on project with document statuses
func (is *InvestService) Comment_on_project_documents(spkComment model.SpkComment) (utils.Msg) {
	var trans = model.GetDB().Begin()
	defer func() { if trans != nil { trans.Rollback() } }()

	// validate statuses of documents
	if err := spkComment.OnlyValidateStatusesOfDocuments(); err != nil {
		return model.ReturnInvalidParameters(err.Error())
	}

	// validate
	spkComment.Comment.UserId = is.UserId
	if err := spkComment.Comment.Validate(); err != nil {
		return model.ReturnInvalidParameters(err.Error())
	}

	// create
	if err := spkComment.Comment.OnlyCreate(trans); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// update document statuses
	if err := spkComment.OnlyUpdateDocumentStatusesByIdAndProjectId(spkComment.Comment.ProjectId, spkComment.Documents, trans); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// change the status of project & gantt step
	msg := is.Ganta_change_the_status_of_project(spkComment.Comment.ProjectId, spkComment.Comment.Status)
	if msg.IsThereAnError() {
		return msg
	}

	// commit
	if err := trans.Commit().Error; err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	trans = nil

	// send notification
	nc := model.NotifyAddComment {
		CommentedById: 	is.UserId,
		ProjectId:     	spkComment.Comment.ProjectId,
		CommentBody:   	spkComment.Comment.Body,
		Status:        	spkComment.Comment.Status,
	}

	select {
	case model.GetMailerQueue().NotificationChannel <- &nc:
	default:
	}

	return model.ReturnNoError()
}

func (is *InvestService) Get_comments_of_project(project_id uint64) (utils.Msg) {
	var comment = model.Comment{
		ProjectId: project_id,
	}
	return comment.Get_all_comments_of_the_project_by_project_id(is.Offset)
}

