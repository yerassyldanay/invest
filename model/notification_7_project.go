package model

import (
	"fmt"
	"invest/utils"
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
	"html": "Жаңа жоба қосылды. Жоба атауы: %s. Жоба ұсынушы: %s. \n" +
		"Проект был добавлен. Название проекта: %s. ФИО инициатора проекта: %s. \n" +
		"A new project has been added. The name of the project: %s. The name of an initiator: %s",
}

// get map
func (n *NotifyProjectCreation) GetMap() map[string]string {
	return MapNotifyProjectCreation
}

// sender
func (n *NotifyProjectCreation) GetFrom() (string) {
	return utils.BaseEmailAddress
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
	return MapNotifyProjectCreation[utils.KeyEmailSubject]
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
	// A new project has been added. The name of the project: %s. The name of an initiator: %s
	body := n.GetMap()[utils.KeyEmailHtml]
	body = fmt.Sprintf(body, n.Project.Name, user.Fio)

	return body
}

func (n *NotifyProjectCreation) GetPlainText() string {
	return n.GetHtml()
}

func (n *NotifyProjectCreation) GetProjectId() uint64 {
	return n.ProjectId
}
