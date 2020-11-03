package model

import (
	"fmt"
	"invest/utils/constants"
)

type NotifyProjectStatus struct {
	UserId					uint64
	ChangedByFio			string
	StatusBefore			string
	StatusAfter				string
	ProjectId				uint64
	Step					int
	Lang					string
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
	"html": "Жоба '%s' мәртебесінен '%s' мәртебесіне өзгертілді. Өзгерткен %s \n\n\n " +
		"Статус проекта изменен с '%s' на '%s'. Внес изменения: %s \n\n\n " +
		"The project status has been changed from '%s' to '%s'. Changes were made by %s",
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
	var user = User{Id: n.UserId}
	if err := user.OnlyGetUserById(GetDB()); err != nil {
		return ""
	}

	body := n.GetMap()[constants.KeyEmailHtml]

	// The project status has been changed from '%s' to '%s'. Changes were made by %s
	resp := fmt.Sprintf(body, constants.MapProjectStatusFirstStatusThenLang[n.StatusBefore]["kaz"], constants.MapProjectStatusFirstStatusThenLang[n.StatusBefore]["kaz"], user.Fio,
		constants.MapProjectStatusFirstStatusThenLang[n.StatusBefore]["rus"], constants.MapProjectStatusFirstStatusThenLang[n.StatusBefore]["rus"], user.Fio,
		constants.MapProjectStatusFirstStatusThenLang[n.StatusBefore]["eng"], constants.MapProjectStatusFirstStatusThenLang[n.StatusBefore]["eng"], user.Fio)

	return resp
}

func (n *NotifyProjectStatus) GetPlainText() string {
	return n.GetHtml()
}

func (n *NotifyProjectStatus) GetProjectId() uint64 {
	return n.ProjectId
}
