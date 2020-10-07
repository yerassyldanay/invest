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

	// get document & change uri & save it
	var tempDoc = model.Document{Id: document.Id}
	err := tempDoc.OnlyGetDocumentById(trans)
	if err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// check if the document is already uploaded
	if tempDoc.Uri != "" {
		return model.ReturnMethodNotAllowed("document is already uploaded. first delete")
	}

	// set new fields
	tempDoc.Modified = utils.GetCurrentTime()
	tempDoc.Uri = document.Uri
	tempDoc.Deadline = utils.GetCurrentTime().Add(time.Hour * (-24))

	// save changes
	if err = tempDoc.OnlyUpdateUriById(trans); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	/*
		get gantt table

		here we need to shift the step at gantt table, in case:
			* investor has uploaded all documents
			* investor has reconsidered documents, which are sent back
	 */
	var ganta = model.Ganta{ProjectId: document.ProjectId}
	if err := ganta.OnlyGetCurrentStepByProjectId(trans); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	/*
		Investor:
			case: when investor is responsible & nothing to upload (all documents are uploaded)
			case: when investor is responsible to reconsider documents
	 */
	var countNumberOfEmptyDocuments, countNumberOfDocumentsWithUndesirableStatus int
	if ganta.Responsible == is.RoleName && is.RoleName == utils.RoleInvestor {
		// number of documents with empty uri
		countNumberOfEmptyDocuments = document.OnlyCountNumberOfEmptyDocuments(is.RoleName, ganta.Step, trans)

		// if there is any document with undesirable status then do not move to the next gantt step
		countNumberOfDocumentsWithUndesirableStatus = document.OnlyCountNumberOfNotAcceptedDocuments(ganta.ProjectId, ganta.Step, trans)

		var dus = model.DocumentUserStatus{
			DocumentId: document.Id,
			Status: utils.ProjectStatusNewOne,
		}
		if err := dus.OnlyUpdate(trans); err != nil {
			return model.ReturnInternalDbError(err.Error())
		}

		// count the number of documents, which are to be reconsidered or uploaded
		if countNumberOfEmptyDocuments + countNumberOfDocumentsWithUndesirableStatus != 0 {
			// there are still documents
		} else if err = ganta.OnlyChangeStatusToDoneById(trans); err != nil {
			return model.ReturnInternalDbError(err.Error())
		}
	} else if utils.Does_a_slice_contain_element([]string{utils.RoleSpk, utils.RoleManager, utils.RoleExpert}, ganta.Responsible) &&
		utils.Does_a_slice_contain_element([]string{utils.RoleSpk, utils.RoleManager, utils.RoleExpert}, is.RoleName) {
		fmt.Println(is.RoleName + " | " + ganta.Responsible)
	}

	if err = trans.Commit().Error; err != nil {
		return model.ReturnInternalDbError(err.Error())
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
	document.Modified = utils.GetCurrentTime()

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

	return model.ReturnNoError()
}
