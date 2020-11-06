package model

import (
	"fmt"
	"invest/utils/constants"
)

type NotifyOnlyInvestorProjectCreation struct {
	ProjectId				uint64				`json:"project_id"`
	Project					Project				`json:"project"`

	UserId					uint64				`json:"user_id"`
	User					User				`json:"user"`
}

var MapNotifyOnlyInvestorProjectCreation = map[string]string{
	"subject": "Жоба қосылды. Проект добавлен. A project has been added",
	"html": "'%s' жобасы ұсыныс жасалды. Жоба қаралымға өтуі үшін құжаттарды %s дейін жүктеңіз \n\n\n" +
		"Заявка по проекту '%s' создана. Для начала рассмотрения - загрузите документы по проекту до %s \n\n\n" +
		"An application for the project called '%s' has been created. To start the consideration by SEC, upload documents until %s \n",
}

// get map
func (n *NotifyOnlyInvestorProjectCreation) GetMap() map[string]string {
	return MapNotifyOnlyInvestorProjectCreation
}

// sender
func (n *NotifyOnlyInvestorProjectCreation) GetFrom() (string) {
	return constants.BaseEmailAddress
}

// get the list of users, who has connection to project
func (n *NotifyOnlyInvestorProjectCreation) GetToList() []string {
	if n.UserId <= 0 || n.User.Id <= 0 {
		n.UserId = n.Project.OfferedById
		n.User.Id = n.Project.OfferedById

		if err := n.User.OnlyGetByIdPreloaded(GetDB()); err != nil {
			return []string{}
		}
	}

	return []string{ n.User.Email.Address }
}

// get subject
func (n *NotifyOnlyInvestorProjectCreation) GetSubject() string {
	return MapNotifyOnlyInvestorProjectCreation[constants.KeyEmailSubject]
}

// body in html
func (n *NotifyOnlyInvestorProjectCreation) GetHtml() string {

	// we need project name
	if n.Project.Id < 1 {
		n.Project.Id = n.ProjectId
		if err := n.Project.OnlyGetById(GetDB()); err != nil {
			return ""
		}
	}

	// we need deadline
	if n.Project.CurrentStep.Id <= 0 {
		if err := n.Project.GetAndUpdateStatusOfProject(GetDB()); err != nil {
			return ""
		}
	}

	// prepare deadline
	deadline := n.Project.CurrentStep.Deadline.Format("2006-01-02")

	// prepare template
	// The project '%s' has been applied.
	// Please, to start the consideration by SEC, upload documents until %s
	body := n.GetMap()[constants.KeyEmailHtml]
	body = fmt.Sprintf(body,
		n.Project.Name, deadline,
		n.Project.Name, deadline,
		n.Project.Name, deadline)

	return body
}

func (n *NotifyOnlyInvestorProjectCreation) GetPlainText() string {
	return n.GetHtml()
}

func (n *NotifyOnlyInvestorProjectCreation) GetProjectId() uint64 {
	return n.ProjectId
}
