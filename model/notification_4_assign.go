package model

import (
	"fmt"
	"invest/utils"
)

type NotifyAssign struct {
	UserId						uint64				`json:"user_id"`
	User						User				`json:"user"`
	ProjectId					uint64				`json:"project_id"`
	Project						Project				`json:"project"`
}

var MapNotifyAssign = map[string]string{
	"subject": "Сіз жобаға қосылдыңыз. Вас добавили в проект. You have been added to the project",
	"html": "Жоба тақырыбы: %s. Жоба сипаттамасы: %s. Ұйым атауы: %s. \n\n\n" +
		"Название проекта: %s. Описание проекта: %s. Название организации: %s. \n\n\n" +
		"The project: %s. The description: %s. The name of the organization: %s\n",
}

// get map
func (n *NotifyAssign) GetMap() map[string]string {
	return MapNotifyAssign
}

// sender
func (n *NotifyAssign) GetFrom() (string) {
	return utils.BaseEmailAddress
}

// get the list of users, who has connection to project
func (n *NotifyAssign) GetToList() []string {
	if n.User.Email.Id < 0 {
		// one request is enough
		n.User.Id = n.UserId
		if err := n.User.OnlyGetByIdPreloaded(GetDB()); err != nil {
			return []string{}
		}
	}

	return []string{n.User.Email.Address}
}

// get subject
func (n *NotifyAssign) GetSubject() string {
	return n.GetMap()[utils.KeyEmailSubject]
}

// body in html
func (n *NotifyAssign) GetHtml() string {
	body := n.GetMap()[utils.KeyEmailHtml]

	if n.Project.Id < 0 {
		// to escape doing the same request more than once
		n.Project.Id = n.ProjectId
		if err := n.Project.OnlyGetByIdPreloaded(GetDB()); err != nil {
			return ""
		}
	}

	// The project: %s. The description: %s. The name of the organization: %s
	resp := fmt.Sprintf(body, n.Project.Name, n.Project.Description, n.Project.Organization.Name,
		n.Project.Name, n.Project.Description, n.Project.Organization.Name,
		n.Project.Name, n.Project.Description, n.Project.Organization.Name)

	return resp
}

func (n *NotifyAssign) GetPlainText() string {
	return n.GetHtml()
}

func (n *NotifyAssign) GetProjectId() uint64 {
	return uint64(0)
}
