package model

import (
	"fmt"
	"invest/utils"
)

type NotifyGantaDeadline struct {
	ProjectId					uint64					`json:"project_id"`
	Project						Project					`json:"project"`
}

var MapNotifyGantaDeadline = map[string]string{
	"subject": "Ақырғы мерзім келіп қалды. Крайний срок для гантт стадии. Deadline is coming.",
	"html": "Өтініш, гантт кестесіне өзгерістер енгізіңіз! Жоба атуы: %s. Саты атауы: %s. Ақтық мерзім: %s\n\n\n" +
		"Пожалуйста, внесите изменения в таблице гантт. Название проекта: %s. Название стадии: %s. Крайний срок: %s\n\n\n " +
		"Please, enter some changes. The name of the project: %s. The name of the stage: %s. Deadline: %s\n\n\n",
}

// get map
func (n *NotifyGantaDeadline) GetMap() map[string]string {
	return MapNotifyGantaDeadline
}

// sender
func (n *NotifyGantaDeadline) GetFrom() (string) {
	return utils.BaseEmailAddress
}

// get the list of users, who has connection to project
func (n *NotifyGantaDeadline) GetToList() []string {
	if n.Project.CurrentStep.Id < 1 {
		n.Project.Id = n.ProjectId
		if err := n.Project.GetAndUpdateStatusOfProject(GetDB()); err != nil {
			return []string{}
		}
	}

	var user = User{}
	users, err := user.OnlyGetSpkUsersByProjectIdAndRoleName(n.ProjectId, n.Project.CurrentStep.Responsible, GetDB())
	if err != nil {
		return []string{}
	}

	var emails = []string{}
	for i, _ := range users {
		_ = users[i].OnlyGetByIdPreloaded(GetDB())
		emails = append(emails, users[i].Email.Address)
	}

	return emails
}

// get subject
func (n *NotifyGantaDeadline) GetSubject() string {
	return MapNotifyGantaDeadline[utils.KeyEmailSubject]
}

// body in html
func (n *NotifyGantaDeadline) GetHtml() string {
	if n.Project.Id < 1 {
		if err := n.Project.OnlyGetById(GetDB()); err != nil {
			return ""
		}
	}

	if n.Project.CurrentStep.Id < 1 {
		if err := n.Project.GetAndUpdateStatusOfProject(GetDB()); err != nil {
			return ""
		}
	}

	deadline := n.Project.CurrentStep.Deadline.Format("2006-01-02")

	// prepare template
	// A new password: %s
	body := n.GetMap()[utils.KeyEmailHtml]
	body = fmt.Sprintf(body, n.Project.Name, n.Project.CurrentStep.Kaz, deadline,
		n.Project.Name, n.Project.CurrentStep.Rus, deadline,
		n.Project.Name, n.Project.CurrentStep.Eng, deadline)

	return body
}

func (n *NotifyGantaDeadline) GetPlainText() string {
	return n.GetHtml()
}

func (n *NotifyGantaDeadline) GetProjectId() uint64 {
	return n.ProjectId
}
