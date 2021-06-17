package model

import (
	"fmt"
	"github.com/yerassyldanay/invest/utils/constants"
)

type NotifyAddDoc struct {
	UserId				uint64					`json:"user_id"`
	User				User

	Document			Document				`json:"document"`
}

var MapNotifyAddDoc = map[string]string{
	"subject": "Құжат қосылды." +
		"  Документ добавлен." +
		" A document has been added",
	"html": "'%s' келесі '%s' құжатты жүктеуді сұрастырады. '%s' атаулы құжатты жүктеуі қажет \n\n\n" +
		"'%s' запросил документ '%s'. '%s' должен загрузите документ '%s' \n\n\n" +
		"'%s' has requested the document called '%s'. '%s' should upload the document",
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
	email := Email{}
	emails, err := email.OnlyGetEmailsHasConnectionToProject(n.Document.ProjectId, GetDB())
	if err != nil {
		return []string{}
	}

	emailsString := []string{}
	for _, email := range emails {
		emailsString = append(emailsString, email.Address)
	}

	return emailsString
}

// get subject
func (n *NotifyAddDoc) GetSubject() string {
	return n.GetMap()[constants.KeyEmailSubject]
}

// body in html
func (n *NotifyAddDoc) GetHtml() string {

	// prepare template
	body := n.GetMap()[constants.KeyEmailHtml]

	// '%s' has requested the document called '%s'. '%s' should upload the document
	requestedBy := n.User.Role.Name
	responsible := n.Document.Responsible

	resp := fmt.Sprintf(body,
		constants.MapRole[requestedBy]["kaz"], n.Document.Kaz, constants.MapRole[responsible]["kaz"],
		constants.MapRole[requestedBy]["rus"], n.Document.Kaz, constants.MapRole[responsible]["rus"],
		constants.MapRole[requestedBy]["eng"], n.Document.Kaz, constants.MapRole[responsible]["eng"],
	)

	return resp
}

func (n *NotifyAddDoc) GetPlainText() string {
	return n.GetHtml()
}

func (n *NotifyAddDoc) GetProjectId() uint64 {
	return n.Document.ProjectId
}
