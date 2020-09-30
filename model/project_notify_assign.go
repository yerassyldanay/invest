package model

import (
	"errors"
	"fmt"
	"invest/templates"
	"invest/utils"

)

func (pu *ProjectsUsers) Notify_user(lang string) (map[string]interface{}, error) {
	if pu.UserId == 0 || pu.ProjectId == 0 {
		return utils.ErrorInvalidParameters, errors.New("invalid aparameters passed")
	}

	//var main_query = "select  from projects   where projects_users.user_id=?;"

	var both = struct {
		Project
		User
	}{}

	err := GetDB().Table(Project{}.TableName()).Select("projects.*, users.*").Joins("join projects_users on projects.id = projects_users.project_id").
		Joins("join users on projects_users.user_id = users.id").Where("projects_users.user_id=?", pu.UserId).First(&both).Error
	if err != nil {
		return utils.ErrorInternalDbError, err
	}

	//err = rows.Scan(&boths)
	//if len(boths) == 0 || err != nil {
	//	fmt.Println(err)
	//	return utils.ErrorInvalidParameters, errors.New("the number of scanned rows are 0")
	//}
	//
	//var both = boths[0]

	if err := GetDB().Table(Email{}.TableName()).Where("id=?", both.User.EmailId).First(&both.User.Email).Error;
		err != nil {
			return utils.ErrorInternalDbError, err
	}

	var subject, html, page string
	switch lang {
	case "kaz":
		subject = templates.Base_email_notification_subject_kaz
		html = templates.Base_email_notification_page_kaz
		page = templates.Base_email_notification_page_kaz
 	case "rus":
		subject = templates.Base_email_notification_subject_rus
		html = templates.Base_email_notification_page_rus
		page = templates.Base_email_notification_page_rus
	default:
		subject = templates.Base_email_notification_subject_eng
		html = templates.Base_email_notification_page_eng
		page = templates.Base_email_notification_page_eng
	}

	html = fmt.Sprintf(html, both.Project.Name, both.Project.Description)
	page = fmt.Sprintf(page, both.Project.Name, both.Project.Description)

	var sendgmsg = SendgridMessageStore{
		From:     utils.BaseEmailAddress,
		To:       both.User.Email.Address,
		FromName: utils.BaseEmailName,
		ToName:   both.User.Fio,
		SendgridMessage:   SendgridMessage{
			Subject:   		subject,
			PlainText: 		page,
			HTML:      		html,
		},
		Created:      		utils.GetCurrentTime(),
	}

	resp, _ := sendgmsg.Send_message()

	return resp, nil
}
