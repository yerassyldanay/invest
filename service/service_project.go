package service

import (
	"fmt"
	"invest/model"
	"invest/utils"
)

func (is *InvestService) Service_create_project(projectWithFinTable *model.ProjectWithFinanceTables) (utils.Msg){
	defer func() utils.Msg {
		if err := recover(); err != nil {
			fmt.Println("CreateProject - could not send email: ", err)
			return model.ReturnInternalDbError("the function createProject failed")
		}
		return model.ReturnNoError()
	}()

	_ = model.Update_sequence_id_thus_avoid_duplicate_primary_key_error( model.GetDB(), "finances")
	_ = model.Update_sequence_id_thus_avoid_duplicate_primary_key_error( model.GetDB(), "costs")
	_ = model.Update_sequence_id_thus_avoid_duplicate_primary_key_error( model.GetDB(), "projects")
	_ = model.Update_sequence_id_thus_avoid_duplicate_primary_key_error( model.GetDB(), "users")
	_ = model.Update_sequence_id_thus_avoid_duplicate_primary_key_error( model.GetDB(), "organizations")
	_ = model.Update_sequence_id_thus_avoid_duplicate_primary_key_error( model.GetDB(), "documents")

	var trans = model.GetDB().Begin()
	defer func() { if trans != nil {trans.Rollback()} }()

	var msg = utils.Msg{}

	// set fields
	projectWithFinTable.Project.Status = utils.ProjectStatusPendingAdmin
	projectWithFinTable.Project.OfferedById = is.UserId
	projectWithFinTable.Project.Lang = is.Lang

	/*
		create a project
	 */
	msg = projectWithFinTable.Project.Create_project(trans)
	if msg.ErrMsg != "" {
		return msg
	}

	// create finance table
	projectWithFinTable.Finance.ProjectId = projectWithFinTable.Project.Id
	if err := projectWithFinTable.Finance.OnlyCreate(trans); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// create cost table
	projectWithFinTable.Cost.ProjectId = projectWithFinTable.Project.Id
	if err := projectWithFinTable.Cost.OnlyCreate(trans); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	/*
		create:
			* Ganta table (parent)
			* Parent steps - will be shown for other
	 */
	msg = projectWithFinTable.Project.Create_ganta_table_for_this_project(trans)
	if msg.ErrMsg != "" {
		return msg
	}

	// create default documents with deadline, but empty uri
	var document = model.Document{}
	msg = document.Create_default_documents(projectWithFinTable.Project.Id, trans)
	if msg.IsThereAnError() {
		return msg
	}

	// commit changes
	err := trans.Commit().Error
	if err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	trans = nil

	// update the status of the project
	if err := projectWithFinTable.Project.GetAndUpdateStatusOfProject(model.GetDB()); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// send notification
	nps := model.NotifyProjectCreation{
		ProjectId: projectWithFinTable.Project.Id,
		UserId:    is.UserId,
	}

	//
	select {
	case model.GetMailerQueue().NotificationChannel <- &nps:
	default:
	}

	// assign all experts to the project
	var pu = model.ProjectsUsers{ProjectId: projectWithFinTable.Project.Id}
	if err := pu.OnlyAssignExpertsToProject(projectWithFinTable.Project.Id, model.GetDB()); err != nil {
		fmt.Println("could not assign experts to project: ", err)
	}

	return model.ReturnNoError()
}
