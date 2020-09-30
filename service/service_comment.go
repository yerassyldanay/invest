package service

import (
	"github.com/jinzhu/gorm"
	"invest/model"
	"invest/utils"
)

func (is InvestService) Comment_on_project_documents(spkComment model.SpkComment) (utils.Msg) {
	var trans = model.GetDB().Begin()
	defer func() { if trans != nil { trans.Rollback() } }()

	// check ganta child steps
	var childIds = []uint64{}
	for i, comment := range spkComment.DocStatuses {
		childIds = append(childIds, comment.GantaId)
		spkComment.DocStatuses[i].UserId = is.UserId
	}

	/*
		check ganta ids are valid
		note: it does not check whether all documents are covered
	 */
	var du = model.DocumentUserStatus{}
	if ok := du.AreAllValidChildGantaIds(childIds, spkComment.Comment.ProjectId, trans); !ok {
		//fmt.Println("not all")
		return model.ReturnMethodNotAllowed("includes invalid ganta child ids")
	}

	// create comment
	spkComment.Comment.UserId = is.UserId
	if err := spkComment.Comment.Validate(); err != nil {
		return model.ReturnInvalidParameters(err.Error())
	}

	if err := spkComment.Comment.OnlyCreate(trans); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// store statuses of each document on db
	//var wg = sync.WaitGroup{}
	var errorChan = make(chan error, 1)

	for _, docStatus := range spkComment.DocStatuses {
		docStatus := docStatus

		//defer wg.Add(1)
		func(docStatus *model.DocumentUserStatus, trans *gorm.DB) {
			//defer wg.Done()
			err := docStatus.OnlyCreate(trans)
			if err != nil {
				select {
				case errorChan <- err:
				default:
					// pass if the chan is full
				}
			}
		}(&docStatus, trans)
	}
	//wg.Wait()

	select {
	case err := <- errorChan:
		return model.ReturnInternalDbError(err.Error())
	default:
	}

	close(errorChan)

	var err error
	var ganta = model.Ganta{
		ProjectId: spkComment.Comment.ProjectId,
	}

	// change the status of project & ganta step
	switch {
	case spkComment.Comment.Status == utils.ProjectStatusReject:
		err = ganta.OnlyUpdateRejectStatusByProjectId(true, trans)
	case spkComment.Comment.Status == utils.ProjectStatusReconsider:
		err = ganta.OnlyUpdateReconsiderStatusByProjectId(true, trans)
	case spkComment.Comment.Status == utils.ProjectStatusAccept:
		err = ganta.OnlyChangeStatusToDoneById(trans)
	default:
		return model.ReturnMethodNotAllowed("invalid status. status is " + spkComment.Comment.Status)
	}

	if err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// commit
	if err := trans.Commit().Error; err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// upgrade the status
	var project = model.Project{
		Id: ganta.ProjectId,
	}
	_ = project.GetAndUpdateStatusOfProject(model.GetDB())

	trans = nil
	return model.ReturnNoError()
}

func (is *InvestService) Get_comments_of_project(project_id uint64) (utils.Msg) {
	var comment = model.Comment{
		ProjectId: project_id,
	}
	return comment.Get_all_comments_of_the_project_by_project_id(is.Offset)
}

