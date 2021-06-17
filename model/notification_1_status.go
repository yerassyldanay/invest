package model

import (
	"fmt"
	"github.com/yerassyldanay/invest/utils/constants"
)

type NotifyProjectStatus struct {
	UserId						uint64
	ChangedBy					User

	ProjectId					uint64
	Project						Project

	LastGantaStep				Ganta
}

/*
	// this indicates the sender address
	GetFrom() string
	// indicates email addresses of receivers
	GetToList() []string
	// the subject of the email message
	GetSubject() string
	// body of the message in html format
	GetHtml() string
	// body of the message in plain text format
	GetPlainText() string
 */

var MapNotifyProjectStatus = map[string]string{
	"subject": "Жобаға өзгерістер енгізілді." +
		" Внесены изменения в проект. " +
		" Changes has been made to the project",
	"html": "%s (%s) жобаны '%s' кезеңінен '%s' кезеңіне өзгертті \n\n\n " +
		"%s (%s) изменил стадию проекта с '%s' на '%s' \n\n\n " +
		"%s (%s) has changed the stage from '%s' to '%s'",
}

// get map
func (n *NotifyProjectStatus) GetMap() map[string]string {
	return MapNotifyProjectStatus
}

// sender
func (n *NotifyProjectStatus) GetFrom() (string) {
	return constants.BaseEmailAddress
}

// get the list of users, who has connection to project
func (n *NotifyProjectStatus) GetToList() []string {
	var email = Email{}
	emails, _ := email.OnlyGetEmailsHasConnectionToProject(n.ProjectId, GetDB())

	var emailsString = []string{}
	for _, email := range emails {
		emailsString = append(emailsString, email.Address)
	}

	return emailsString
}

// get subject
func (n *NotifyProjectStatus) GetSubject() string {
	return MapNotifyProjectStatus[constants.KeyEmailSubject]
}

// body in html
func (n *NotifyProjectStatus) GetHtml() string {
	// get the presone, who changed the stage
	if n.ChangedBy.Id <= 0 {
		n.ChangedBy.Id = n.UserId
		if err := n.ChangedBy.OnlyGetByIdPreloaded(GetDB()); err != nil {
			fmt.Println("notification 1: ", err)
			return ""
		}
	}

	// get current gantt step
	if n.Project.CurrentStep.Id <= 0 {
		if err := n.Project.GetAndUpdateStatusOfProject(GetDB()); err != nil {
			fmt.Println("notification 1: ", err)
			return ""
		}
	}

	// get the body of html
	body := n.GetMap()[constants.KeyEmailHtml]

	// %s (%s) has changed the stage from '%s' to '%s'
	role := n.ChangedBy.Role.Name
	resp := fmt.Sprintf(body,
		n.ChangedBy.Fio, constants.MapRole[role]["kaz"], n.LastGantaStep.Kaz, n.Project.CurrentStep.Kaz,
		n.ChangedBy.Fio, constants.MapRole[role]["rus"], n.LastGantaStep.Rus, n.Project.CurrentStep.Rus,
		n.ChangedBy.Fio, constants.MapRole[role]["eng"], n.LastGantaStep.Eng, n.Project.CurrentStep.Eng)

	return resp
}

func (n *NotifyProjectStatus) GetPlainText() string {
	return n.GetHtml()
}

func (n *NotifyProjectStatus) GetProjectId() uint64 {
	return n.ProjectId
}
