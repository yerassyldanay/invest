package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/utils/constants"
	"github.com/yerassyldanay/invest/utils/errormsg"
	"github.com/yerassyldanay/invest/utils/message"
)

func (is *InvestService) Service_create_project(projectWithFinTable *model.ProjectWithFinanceTables) (message.Msg){
	defer func() message.Msg {
		if err := recover(); err != nil {
			fmt.Println("CreateProject - could not send email: ", err)
			return model.ReturnInternalDbError("the function createProject failed")
		}
		return model.ReturnNoError()
	}()

	/*
		LevelDefault IsolationLevel = iota
		LevelReadUncommitted
		LevelReadCommitted
		LevelWriteCommitted
		LevelRepeatableRead
		LevelSnapshot
		LevelSerializable
		LevelLinearizable
	 */
	var trans = model.GetDB().BeginTx(context.Background(), &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	})
	defer func() { if trans != nil { trans.Rollback() } }()

	var msg = message.Msg{}

	// set fields
	projectWithFinTable.Project.Status = constants.ProjectStatusPendingAdmin
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
		fmt.Println(err.Error())
	}

	// assign all experts to the project
	var pu = model.ProjectsUsers{ProjectId: projectWithFinTable.Project.Id}
	if err := pu.OnlyAssignExpertsToProject(projectWithFinTable.Project.Id, model.GetDB()); err != nil {
		fmt.Println("could not assign experts to project: ", err)
	}

	// send notification
	nps := model.NotifyProjectCreation{
		ProjectId: projectWithFinTable.Project.Id,
		UserId:    is.UserId,
	}

	// in case
	select {
	case model.GetMailerQueue().NotificationChannel <- &nps:
	default:
	}

	// notify an investor
	nipc := model.NotifyOnlyInvestorProjectCreation{
		ProjectId: projectWithFinTable.Project.Id,
		Project:   projectWithFinTable.Project,
		UserId:    is.UserId,
	}

	// send notification
	select {
	case model.GetMailerQueue().NotificationChannel <- &nipc:
	default:
	}

	return model.ReturnNoError()
}

// get all project info
func (is *InvestService) Project_get_by_id(project_id uint64) (message.Msg) {

	var projectWithFinTables = struct {
		model.Project
		model.Finance
		model.Cost
	}{
		Project: model.Project{
			Id: project_id,
		},
		Finance: model.Finance{
			ProjectId: project_id,
		},
		Cost: model.Cost{
			ProjectId: project_id,
		},
	}

	// get project
	if err := projectWithFinTables.Project.Get_this_project_by_project_id(); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// get finance table
	if err := projectWithFinTables.Finance.OnlyGetByProjectId(model.GetDB()); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// get costs table
	if err := projectWithFinTables.Cost.OnlyGetByProjectId(model.GetDB()); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// convert to map
	var resp = errormsg.NoErrorFineEverthingOk
	resp["info"] = model.Struct_to_map(projectWithFinTables)

	return model.ReturnNoErrorWithResponseMessage(resp)
}
