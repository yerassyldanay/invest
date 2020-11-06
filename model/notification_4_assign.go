package model

import (
	"fmt"
	"invest/utils/constants"
)

type NotifyAssign struct {
	UserId						uint64				`json:"user_id"`
	User						User				`json:"user"`
	ProjectId					uint64				`json:"project_id"`
	Project						Project				`json:"project"`
}

var MapNotifyAssign = map[string]string{
	"subject": "Сіз жобаға қосылдыңыз. Вас добавили в проект. You have been added to the project",
	"html": "Менеджер '%s' аталатын жобаға қосылды. Жоба сипаттамасы: %s \n\n\n" +
		"Менеджер был назначен на проект '%s'. Описание проекта: %s \n\n\n" +
		"A manager has been assigned to the project '%s'. The description of the project: %s \n",
}

// get map
func (n *NotifyAssign) GetMap() map[string]string {
	return MapNotifyAssign
}

// sender
func (n *NotifyAssign) GetFrom() (string) {
	return constants.BaseEmailAddress
}

// get the list of users, who has connection to project
func (n *NotifyAssign) GetToList() []string {
	if n.Project.Id <= 0 {
		if err := n.Project.OnlyGetById(GetDB()); err != nil {
			return []string{}
		}
	}

	// get emails of those, who has connection to the project
	email := Email{}
	emails, err := email.OnlyGetEmailsHasConnectionToProject(n.Project.Id, GetDB())
	if err != nil {
		return []string{}
	}

	// get only email addresses
	emailAddresses := []string{}
	for _, email := range emails {
		emailAddresses = append(emailAddresses, email.Address)
	}

	return emailAddresses
}

// get subject
func (n *NotifyAssign) GetSubject() string {
	return n.GetMap()[constants.KeyEmailSubject]
}

// body in html
func (n *NotifyAssign) GetHtml() string {
	body := n.GetMap()[constants.KeyEmailHtml]

	if n.Project.Id < 0 {
		// to escape doing the same request more than once
		n.Project.Id = n.ProjectId
		if err := n.Project.OnlyGetByIdPreloaded(GetDB()); err != nil {
			return ""
		}
	}

	// A manager has been assigned to the project '%s'. The description of the project: %s
	resp := fmt.Sprintf(body,
		n.Project.Name, n.Project.Description,
		n.Project.Name, n.Project.Description,
		n.Project.Name, n.Project.Description)

	return resp
}

func (n *NotifyAssign) GetPlainText() string {
	return n.GetHtml()
}

func (n *NotifyAssign) GetProjectId() uint64 {
	return uint64(0)
}
