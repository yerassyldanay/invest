package model

import (
	"fmt"
	"invest/utils/constants"
)

type NotifyAddDoc struct {
	Document			Document					`json:"document"`
	UserId				uint64					`json:"user_id"`
}

var MapNotifyAddDoc = map[string]string{
	"subject": "Құжат қосылды." +
		"  Документ добавлен." +
		" A document has been added",
	"html": "Құжат: %s. Жауапты: %s. Құжатты қосқан тұлға: %s.\n\n\n" +
		"Документ: %s. Ответственный за документ: %s. Кто добавил: %s\n\n\n" +
		"Document: %s. Responsible: %s. Added by: %s",
}

// get map
func (n *NotifyAddDoc) GetMap() map[string]string {
	return MapNotifyAddDoc
}

// sender
func (n *NotifyAddDoc) GetFrom() (string) {
	return constants.BaseEmailAddress
}

// get the list of users, who has connection to project
func (n *NotifyAddDoc) GetToList() []string {
	var email = Email{}
	emails, err := email.OnlyGetEmailsHasConnectionToProject(n.Document.ProjectId, GetDB())
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
	return MapNotifyAddDoc[constants.KeyEmailSubject]
}

// body in html
func (n *NotifyAddDoc) GetHtml() string {
	// get user, who added
	var user = User{Id: n.UserId}
	if err := user.OnlyGetByIdPreloaded(GetDB()); err != nil {
		return ""
	}

	// prepare template
	body := n.GetMap()[constants.KeyEmailHtml]

	// Document: %s. Responsible: %s. Added by: %s
	resp := fmt.Sprintf(body, n.Document.Kaz, constants.MapRole[n.Document.Responsible]["kaz"], user.Fio,
		n.Document.Rus, constants.MapRole[n.Document.Responsible]["rus"], user.Fio,
		n.Document.Eng, constants.MapRole[n.Document.Responsible]["eng"], user.Fio)

	return resp
}

func (n *NotifyAddDoc) GetPlainText() string {
	return n.GetHtml()
}

func (n *NotifyAddDoc) GetProjectId() uint64 {
	return n.Document.ProjectId
}
