package model

import (
	"fmt"
	"invest/utils/constants"
)

type NotifyGantaDeadline struct {
	ProjectId					uint64					`json:"project_id"`
	Project						Project					`json:"project"`
}

var MapNotifyGantaDeadline = map[string]string{
	"subject": "Ақырғы мерзім келіп қалды. Крайний срок для гантт стадии. Running out of time.",
	"html": "'%s' атаулы жоба кезеңінің аяқталуына аз уақыт қалды. Өтініш, кезеңді %s дейін аяқтаңыз \n\n\n" +
		"Осталось мало времени до завершения стадии. Пожалуйста, завершите стадию по проекту '%s' до %s \n\n\n " +
		"The stage is running out of time until it must be complete. Please, complete the stage of the project called '%s' until %s\n\n\n",
}

// get map
func (n *NotifyGantaDeadline) GetMap() map[string]string {
	return MapNotifyGantaDeadline
}

// sender
func (n *NotifyGantaDeadline) GetFrom() (string) {
	return constants.BaseEmailAddress
}

// get the list of users, who has connection to project
func (n *NotifyGantaDeadline) GetToList() []string {
	if n.Project.CurrentStep.Id < 1 {
		n.Project.Id = n.ProjectId
		if err := n.Project.GetAndUpdateStatusOfProject(GetDB()); err != nil {
			return []string{}
		}
	}

	var emails []Email
	var email = Email{}
	var err error

	switch {
	case n.Project.CurrentStep.Responsible == constants.RoleInvestor:
		err = email.OnlyGetEmailOfInvestorByProjectId(n.Project.Id, GetDB())
		emails = append(emails, email)
	case n.Project.CurrentStep.Responsible == constants.RoleManager:
		emails, err = email.OnlyGetEmailOfManagerByProjectId(n.Project.Id, GetDB())
	case n.Project.CurrentStep.Responsible == constants.RoleExpert:
		emails, err = email.OnlyGetAllEmailsByRole(constants.RoleExpert, GetDB())
	default:
		emails, err = email.OnlyGetAllEmailsByRole(constants.RoleAdmin, GetDB())
	}

	if err != nil {
		return []string{}
	}

	emailAddresses := []string{}
	for _, email := range emails {
		emailAddresses = append(emailAddresses, email.Address)
	}

	return emailAddresses
}

// get subject
func (n *NotifyGantaDeadline) GetSubject() string {
	return MapNotifyGantaDeadline[constants.KeyEmailSubject]
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
	// You are kindly asked to make changes to the project '%s' until %s. Please, refer to the Gantt table
	body := n.GetMap()[constants.KeyEmailHtml]
	body = fmt.Sprintf(body, n.Project.Name, deadline,
		n.Project.Name, deadline,
		n.Project.Name, deadline)

	return body
}

func (n *NotifyGantaDeadline) GetPlainText() string {
	return n.GetHtml()
}

func (n *NotifyGantaDeadline) GetProjectId() uint64 {
	return n.ProjectId
}
