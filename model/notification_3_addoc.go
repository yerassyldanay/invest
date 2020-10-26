package model

import (
	"fmt"
	"invest/utils"
	"time"
)

type NotifyAddDoc struct {
	Name				string					`json:"name"`
	Deadline			time.Time				`json:"deadline"`
	Responsible			string					`json:"responsible"`
	UserId				uint64					`json:"user_id"`
	ProjectId			uint64					`json:"project_id"`
}

var MapNotifyAddDoc = map[string]string{
	"subject": "Құжат қосылды." +
		"  Документ добавлен." +
		" A document has been added",
	"html": "Құжат: %s. Жауапты: %s. Өткізілуі тиіс уақыт: %s. Құжатты қосқан тұлға: %s.\n\n\n" +
		"Документ: %s. Ответственный за документ: %s. Крайний срок для загрузки документа: %s. Кто добавил: %s\n\n\n" +
		"Document: %s. Responsible: %s. Deadline: %s. Added by: %s",
}

// get map
func (n *NotifyAddDoc) GetMap() map[string]string {
	return MapNotifyAddDoc
}

// sender
func (n *NotifyAddDoc) GetFrom() (string) {
	return utils.BaseEmailAddress
}

// get the list of users, who has connection to project
func (n *NotifyAddDoc) GetToList() []string {
	var email = Email{}
	emails, err := email.OnlyGetEmailsHasConnectionToProject(n.ProjectId, GetDB())
	if err != nil {
		return []string{}
	}

	var emailList = []string{}
	for _, email := range emails {
		emailList = append(emailList, email.Address)
	}

	return emailList
}

// get subject
func (n *NotifyAddDoc) GetSubject() string {
	return MapNotifyAddDoc[utils.KeyEmailSubject]
}

// body in html
func (n *NotifyAddDoc) GetHtml() string {
	// get user, who added
	var user = User{Id: n.UserId}
	if err := user.OnlyGetByIdPreloaded(GetDB()); err != nil {
		return ""
	}

	// prepare template
	body := n.GetMap()[utils.KeyEmailHtml]
	// Document: %s. Responsible: %s. Deadline: %s. Added by: %s
	deadline := fmt.Sprintf("%d.%d.%d", n.Deadline.Day(), n.Deadline.Month(), n.Deadline.Year())
	resp := fmt.Sprintf(body, n.Name, utils.MapRole[n.Responsible]["kaz"], deadline, user.Fio,
		n.Name, utils.MapRole[n.Responsible]["rus"], deadline, user.Fio,
		n.Name, utils.MapRole[n.Responsible]["eng"], deadline, user.Fio,)

	return resp
}

func (n *NotifyAddDoc) GetPlainText() string {
	return n.GetHtml()
}

func (n *NotifyAddDoc) GetProjectId() uint64 {
	return n.ProjectId
}
