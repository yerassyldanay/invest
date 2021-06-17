package model

import (
	"fmt"
	"github.com/yerassyldanay/invest/utils/constants"
)

type NotifyProjectCreation struct {
	ProjectId				uint64				`json:"project_id"`
	Project					Project				`json:"project"`

	UserId					uint64				`json:"user_id"`
	User					User				`json:"user"`
}

var MapNotifyProjectCreation = map[string]string{
	"subject": "Жоба қосылды." +
		"  Проект добавлен." +
		" A project has been added",
	"html": "%s платформада '%s' атаулы жоба ұсынып отыр \n\n\n" +
		"%s подал(-а) заявку на рассмотрение проекта '%s' \n\n\n" +
		"%s has submitted the project called '%s' \n",
}

// get map
func (n *NotifyProjectCreation) GetMap() map[string]string {
	return MapNotifyProjectCreation
}

// sender
func (n *NotifyProjectCreation) GetFrom() (string) {
	return constants.BaseEmailAddress
}

// get the list of users, who has connection to project
func (n *NotifyProjectCreation) GetToList() []string {
	// get only admins
	var email = Email{}
	emails, err := email.OnlyGetEmailsOfSpkUsersAndAdmins(n.ProjectId, GetDB())
	if err != nil {
		return []string{}
	}

	var emailList = []string{}
	for _, email = range emails {
		emailList = append(emailList, email.Address)
	}

	return emailList
}

// get subject
func (n *NotifyProjectCreation) GetSubject() string {
	return MapNotifyProjectCreation[constants.KeyEmailSubject]
}

// body in html
func (n *NotifyProjectCreation) GetHtml() string {
	// get user, who added
	var user = User{Id: n.UserId}
	if err := user.OnlyGetByIdPreloaded(GetDB()); err != nil {
		return ""
	}

	// get project
	if n.Project.Id < 1 {
		n.Project.Id = n.ProjectId
		if err := n.Project.OnlyGetById(GetDB()); err != nil {
			return ""
		}
	}

	// prepare template
	// %s has submitted the project called '%s'
	body := n.GetMap()[constants.KeyEmailHtml]
	body = fmt.Sprintf(body, user.Fio, n.Project.Name,
		user.Fio, n.Project.Name,
		user.Fio, n.Project.Name)

	return body
}

func (n *NotifyProjectCreation) GetPlainText() string {
	return n.GetHtml()
}

func (n *NotifyProjectCreation) GetProjectId() uint64 {
	return n.ProjectId
}
