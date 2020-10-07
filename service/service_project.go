package service

import (
	"fmt"
	"invest/model"
	"invest/utils"
)

func (is *InvestService) Service_create_project(projectWithFinTable *model.ProjectWithFinanceTables) (utils.Msg){
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("CreateProject - could not send email: ", err)
		}
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

	// update project status
	//_ = projectWithFinTable.Project.GetAndUpdateStatusOfProject(trans)

	/*
		NOTIFICATION:
			* inform administrators about it
	 */
	var user = model.User{Id: is.UserId}
	err = user.OnlyGetUserById(model.GetDB())
	if err != nil {
		user.Fio = "инвестор"
	}

	/*
		create a template & set values
	*/
	var t = model.Template{}
	sm := t.Template_prepare_notify_users_about_changes_in_project(projectWithFinTable.Project.Lang, projectWithFinTable.Project.Name, user.Fio)

	/*
		set default values such as an email address of a message sender
	*/
	var sms = model.SendgridMessageStore{}
	sms.Set_default_values()

	/*
		set template
	*/
	sms.SendgridMessage = sm

	// to who
	var addressers = []model.EmailAddresser{}
	var admins, _ = user.OnlyGetPreloadedUsersByRole(utils.RoleAdmin, trans)
	for _, admin := range admins {
		addressers = append(addressers, model.EmailAddresser{
			Name:    admin.Fio,
			Address: admin.Email.Address,
		})
	}

	/*
		set receivers & send
	*/
	sms.ToAddresser = addressers
	_, err = sms.SendMessageToList()
	if err != nil {
		return utils.Msg{utils.ErrorCouldNotSendEmail, 204, "", err.Error()}
	}

	return model.ReturnNoError()
}
