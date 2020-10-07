package service

import (
	"invest/model"
	"invest/utils"
)

func (is InvestService) Comment_on_project_documents(spkComment model.SpkComment) (utils.Msg) {
	var trans = model.GetDB().Begin()
	defer func() { if trans != nil { trans.Rollback() } }()

	// check ganta child steps
	var documentIds = []uint64{}
	for i, comment := range spkComment.DocStatuses {
		documentIds = append(documentIds, comment.DocumentId)
		spkComment.DocStatuses[i].UserId = is.UserId
	}

	/*
		check gantt ids are valid
		note: it does not check whether all documents are covered
	 */
	var du = model.DocumentUserStatus{}
	if ok := du.AreAllValidDocumentIds(documentIds, spkComment.Comment.ProjectId, trans); !ok {
		//fmt.Println("not all")
		return model.ReturnMethodNotAllowed("includes invalid document ids")
	}

	// create comment
	spkComment.Comment.UserId = is.UserId
	if err := spkComment.Comment.Validate(); err != nil {
		return model.ReturnInvalidParameters(err.Error())
	}

	if err := spkComment.Comment.OnlyCreate(trans); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// create document-user statuses
	if err := spkComment.OnlyCreateStatuses(trans); err != nil {
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

	// upgrade the status
	//var project = model.Project{
	//	Id: ganta.ProjectId,
	//}
	//_ = project.GetAndUpdateStatusOfProject(model.GetDB())

	trans = nil
	return model.ReturnNoError()
}

func (is *InvestService) Get_comments_of_project(project_id uint64) (utils.Msg) {
	var comment = model.Comment{
		ProjectId: project_id,
	}
	return comment.Get_all_comments_of_the_project_by_project_id(is.Offset)
}

