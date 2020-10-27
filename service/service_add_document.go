package service

import (
	"fmt"
	"invest/model"
	"invest/utils"
	"time"
)

/*
	Adding documents:
		* investor - uploading documents based on the list
		* investor - uploading documents rejected or reconsider
		* spk - uploading documents through the process
 */
func (is *InvestService) Upload_documents_to_project(document *model.Document) (utils.Msg) {

	var trans = model.GetDB().Begin()
	defer func() { if trans != nil { trans.Rollback() }} ()

	// set new fields
	document.Modified = utils.GetCurrentTruncatedDate()
	document.Uri = document.Uri
	document.Deadline = utils.GetCurrentTruncatedDate()

	// update fields: uri, modified & deadline
	// changes saved
	if err := document.OnlyUpdateUriAndDeadlineByIdAndEmptyUri(trans); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	/*
		Automatic shift - has nothing to do with the very process of document upload

		get gantt table

		here we need to shift the step at gantt table, in case:
			* an investor has uploaded all documents
			* an investor has reconsidered documents, which are sent back
	 */
	var currentGanta = model.Ganta{ProjectId: document.ProjectId}
	if err := currentGanta.OnlyGetCurrentStepByProjectId(trans); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	/*
		Automatically shifts to the next stage:
			* when an investor uploaded all documents
			* when an investor reconsidered (removed & uploaded) documents with an undesirable status
	 */
	var countNumberOfEmptyDocuments int
	if currentGanta.Responsible == is.RoleName && is.RoleName == utils.RoleInvestor {
		// number of documents with empty uri
		// which does not allow to move to the next step
		countNumberOfEmptyDocuments = document.OnlyCountNumberOfEmptyDocuments(is.RoleName, currentGanta.Step, trans)

		// if there is any document with undesirable status then do not move to the next gantt step
		countNumberOfDocumentsWithUndesirableStatus, err := document.OnlyCountNumberOfDocumentsWithUndesirableStatus(is.RoleName, currentGanta.ProjectId, currentGanta.Step, trans)
		if err != nil {
			return model.ReturnInternalDbError(err.Error())
		}

		// count the number of documents, which are to be reconsidered or uploaded
		if countNumberOfEmptyDocuments + countNumberOfDocumentsWithUndesirableStatus == 0 {
			if err = currentGanta.OnlyChangeStatusToDoneAndUpdateDeadlineById(trans); err != nil {
				// if there is no document to change or upload -> shift gantt step
				return model.ReturnInternalDbError(err.Error())
			}

			if err = currentGanta.OnlyGetCurrentStepByProjectId(trans); err != nil {
				// update current step
				return model.ReturnInternalDbError(err.Error())
			}

			hours := int(utils.GetCurrentTruncatedDate().Sub(currentGanta.StartDate).Hours())
			if err := currentGanta.OnlyUpdateStartDatesOfAllUndoneGantaStepsByProjectId(hours, trans); err != nil {
				// could not shift all gantt steps
				return model.ReturnInternalDbError(err.Error())
			}
		}

	} else if utils.Does_a_slice_contain_element([]string{utils.RoleSpk, utils.RoleManager, utils.RoleExpert}, currentGanta.Responsible) &&
		utils.Does_a_slice_contain_element([]string{utils.RoleSpk, utils.RoleManager, utils.RoleExpert}, is.RoleName) {
		fmt.Println(is.RoleName + " | " + currentGanta.Responsible)
	}

	if err := trans.Commit().Error; err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// update the status of the project
	var project = model.Project{Id: document.ProjectId}
	err := project.GetAndUpdateStatusOfProject(model.GetDB())
	if err != nil {
		fmt.Println(err)
	}

	return model.ReturnNoError()
}

// add box to upload a document
func (is *InvestService) Add_box_to_upload_document(document model.Document) (utils.Msg) {

	// set a due date
	if document.SetDeadline > 0 {
		document.Deadline = time.Unix(document.SetDeadline, 0)
	}

	document.IsAdditional = true
	document.Uri = ""
	document.Modified = utils.GetCurrentTruncatedDate()

	// get current gantt step to set project step
	var project = model.Project{Id: document.ProjectId}
	if err := project.GetAndUpdateStatusOfProject(model.GetDB()); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// set step
	document.Step = project.Step

	// valid
	if err := document.Validate(); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// create
	if err := document.OnlyCreate(model.GetDB()); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// send notification
	na := model.NotifyAddDoc{
		Name:        document.Kaz,
		Deadline:    document.Deadline,
		Responsible: document.Responsible,
		UserId:      is.UserId,
		ProjectId:   document.ProjectId,
	}

	// handles everything
	select {
	case model.GetMailerQueue().NotificationChannel <- &na:
	default:
	}

	return model.ReturnNoError()
}
