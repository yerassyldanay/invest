package service

import (
	"fmt"
	"invest/model"
	"invest/utils"
)

func Service_create_project(project *model.Project) (*utils.Msg){
	var msg = &utils.Msg{
		Fname: "Service_create_project",
	}

	project.Status = ""

	/*
		create a project
	 */
	msg = project.Create_project()
	if msg.ErrMsg != "" {
		return msg
	}

	/*
		create a table of ganta for this project
	*/
	msg = project.Create_ganta_table_for_this_project()

	/*
		create finance table for this project
	*/
	var finance = model.Finance{
		ProjectId: project.Id,
	}
	if err := finance.Create_this_table(); err != nil {
		fmt.Println("created fin table: ", err)
	}

	/*
		create finresult table for this project
	*/
	var finresult = model.Finresult{
		ProjectId: project.Id,
	}
	if err := finresult.Create_this_table(); err != nil {
		fmt.Println("create fin table: ", err)
	}

	/*
		inform administrators about it
	 */
	var user = model.User{Id: project.OfferedById}
	err := user.Get_user_by_id()
	if err != nil {
		user.Fio = "инвестор"
	}

	/*
		create a template & set values
	*/
	var t = model.Template{}
	sm := t.Template_prepare_notify_users_about_changes_in_project("kaz", project.Name, user.Fio)

	/*
		set default values such as an email address of a message sender
	*/
	var sms = model.SendgridMessageStore{}
	sms.Set_default_values()

	/*
		set template
	*/
	sms.SendgridMessage = sm

	/*
		to who
	*/
	var addressers = []model.EmailAddresser{}
	var admins = user.Get_admins_only_user_info()
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
		return &utils.Msg{utils.ErrorCouldNotSendEmail, 204, "", err.Error()}
	}

	return &utils.Msg{utils.NoErrorFineEverthingOk, 200, "", ""}
}
