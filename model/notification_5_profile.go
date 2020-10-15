package model

import (
	"fmt"
	"invest/utils"
)

type NotifyCreateProfile struct {
	UserId				uint64
	User				User
	CreatedById			uint64
	CrUser				User
}

var MapNotifyCreateProfile = map[string]string{
	"subject": "Жаңа аккаунт. Новый аккаунт. ",
	"html": "Сіздің электрондық почтаңыз жаңа аккаунтқа тіркелді. Аты-жөні: %s. Рөлі: %s.\n" +
		"Ваша электронная почта привязана к аккаунту на платформе. ФИО: %s. Роль: %s.\n" +
		"This email address was assigned to a user account on platform. Name: %s. Role: %s",
}

// get map
func (n *NotifyCreateProfile) GetMap() map[string]string {
	return MapNotifyCreateProfile
}

// sender
func (n *NotifyCreateProfile) GetFrom() (string) {
	return utils.BaseEmailAddress
}

// get the list of users, who has connection to project
func (n *NotifyCreateProfile) GetToList() []string {
	// get the email address of the user
	if n.User.Email.Id < 1 {
		n.User.Id = n.UserId

		// get all user info
		if err := n.User.OnlyGetByIdPreloaded(GetDB()); err != nil {
			return []string{}
		}
	}

	return []string{n.User.Email.Address}
}

// get subject
func (n *NotifyCreateProfile) GetSubject() string {
	return n.GetMap()[utils.KeyEmailSubject]
}

// body in html
func (n *NotifyCreateProfile) GetHtml() string {

	if n.User.Id < 1 {
		n.User.Id = n.UserId
		_ = n.User.OnlyGetByIdPreloaded(GetDB())
	}

	body := n.GetMap()[utils.KeyEmailHtml]
	// This email address was assigned to a user account on platform. Name: %s. Role: %s
	body = fmt.Sprintf(body, n.User.Fio, utils.MapRole[n.User.Role.Name]["kaz"],
		n.User.Fio, utils.MapRole[n.User.Role.Name]["rus"],
		n.User.Fio, utils.MapRole[n.User.Role.Name]["eng"])

	return body
}

func (n *NotifyCreateProfile) GetPlainText() string {
	return n.GetHtml()
}

func (n *NotifyCreateProfile) GetProjectId() uint64 {
	return uint64(0)
}
