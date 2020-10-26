package model

import (
	"fmt"
	"invest/utils"
)

type NotifyDocDeadline struct {
	ProjectId			uint64				`json:"project_id"`
	Project				Project				`json:"project"`
	DocumentId			uint64				`json:"document_id"`
	Document			Document			`json:"document"`
	emails				[]string
}

var MapNotifyDocDeadline = map[string]string{
	"subject": "[DOC] Ақырғы мерзім келіп қалды. Крайний срок. Deadline is coming.",
	"html": "Сізге платформада құжат жүктеу қажет. Жоба атауы: %s. Құжат атауы: %s. Ақтық мерзім: %s\n\n\n" +
		"Вы должны загрузить документ на платформе. Название проекта: %s. Название документа: %s. Крайний срок: %s\n\n\n" +
		"You need to upload a document on the platform. The project: %s. The document: %s. Deadline: %s\n\n\n",
}

// get map
func (n *NotifyDocDeadline) GetMap() map[string]string {
	return MapNotifyDocDeadline
}

// sender
func (n *NotifyDocDeadline) GetFrom() (string) {
	return utils.BaseEmailAddress
}

func (n *NotifyDocDeadline) GetToList() []string {
	if len(n.emails) < 1 {
		return n._getToList()
	}

	return n.emails
}

// get the list of users, who has connection to project
func (n *NotifyDocDeadline) _getToList() []string {
	if n.Document.Id < 1 {
		n.Document.Id = n.DocumentId
		if err := n.Document.OnlyGetDocumentById(GetDB()); err != nil {
			return []string{}
		}
	}

	if n.Project.Id < 1 {
		n.Project.Id = n.Document.ProjectId
		if err := n.Project.OnlyGetById(GetDB()); err != nil {
			return []string{}
		}
	}

	emails := []string{}
	user := User{}
	if n.Document.Responsible == utils.RoleSpk || n.Document.Responsible == utils.RoleManager ||
		n.Document.Responsible == utils.RoleExpert {
		roles := []string{ utils.RoleSpk, utils.RoleManager, utils.RoleExpert }
		users, err := user.OnlyGetSpkUsersByProjectIdAndRoles(n.Document.ProjectId, roles, GetDB())
		if err != nil {
			return []string{}
		}

		for i, _ := range users {
			if err := users[i].OnlyGetByIdPreloaded(GetDB()); err != nil {
				fmt.Println(err)
			}
			emails = append(emails, user.Email.Address)
		}
	} else {
		if err := user.OnlyGetInvestorByProjectId(n.Document.ProjectId, GetDB()); err != nil {
			fmt.Println("notify deadline. doc: ", err)
		}

		if err := user.OnlyGetByIdPreloaded(GetDB()); err != nil {
			fmt.Println("notify deadline. doc: ", err)
		}

		emails = append(emails, user.Email.Address)
	}

	n.emails = emails
	return emails
}

// get subject
func (n *NotifyDocDeadline) GetSubject() string {
	return MapNotifyDocDeadline[utils.KeyEmailSubject]
}

// body in html
func (n *NotifyDocDeadline) GetHtml() string {
	if n.Project.Id < 1 {
		if err := n.Project.OnlyGetById(GetDB()); err != nil {
			return ""
		}
	}

	deadline := n.Document.Deadline.Format("2006-01-02")

	// prepare template
	// You need to upload a document on the platform.
	// The project: %s. The document: %s. Deadline: %s
	body := n.GetMap()[utils.KeyEmailHtml]
	body = fmt.Sprintf(body, n.Project.Name, n.Document.Kaz, deadline,
		n.Project.Name, n.Document.Rus, deadline,
		n.Project.Name, n.Document.Eng, deadline)

	return body
}

func (n *NotifyDocDeadline) GetPlainText() string {
	return n.GetHtml()
}

func (n *NotifyDocDeadline) GetProjectId() uint64 {
	return n.ProjectId
}
