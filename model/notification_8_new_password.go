package model

import (
	"fmt"
	"invest/utils/constants"
)

type NotifyNewPassword struct {
	UserId					uint64				`json:"user_id"`
	User					User				`json:"user"`
	RawNewPassword			string				`json:"raw_new_password"`
}

var MapNotifyNewPassword = map[string]string{
	"subject": "Жаңа құпия сөз. Новый пароль. A new password.",
	"html": "Жаңа құпия сөз: %s\n\n\n " +
		"Новый пароль: %s\n\n\n " +
		"A new password: %s\n\n\n",
}

// get map
func (n *NotifyNewPassword) GetMap() map[string]string {
	return MapNotifyNewPassword
}

// sender
func (n *NotifyNewPassword) GetFrom() (string) {
	return constants.BaseEmailAddress
}

// get the list of users, who has connection to project
func (n *NotifyNewPassword) GetToList() []string {
	if n.User.Id < 1 {
		n.User.Id = n.UserId
		if err := n.User.OnlyGetByIdPreloaded(GetDB()); err != nil {
			return []string{}
		}
	}

	return []string{n.User.Email.Address}
}

// get subject
func (n *NotifyNewPassword) GetSubject() string {
	return MapNotifyNewPassword[constants.KeyEmailSubject]
}

// body in html
func (n *NotifyNewPassword) GetHtml() string {
	if n.User.Id < 1 {
		n.User.Id = n.UserId
		if err := n.User.OnlyGetByIdPreloaded(GetDB()); err != nil {
			return ""
		}
	}

	// prepare template
	// A new password: %s
	body := n.GetMap()[constants.KeyEmailHtml]
	body = fmt.Sprintf(body, n.RawNewPassword, n.RawNewPassword, n.RawNewPassword)

	return body
}

func (n *NotifyNewPassword) GetPlainText() string {
	return n.GetHtml()
}

func (n *NotifyNewPassword) GetProjectId() uint64 {
	return 0
}
