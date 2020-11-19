package model

import (
	"fmt"
	"invest/utils/constants"
)

type NotifyCreateProfile struct {
	UserId				uint64
	User				User
	CreatedById			uint64
	CrUser				User

	RawPassword			string
}

var MapNotifyCreateProfile = map[string]string{
	"subject": "Жаңа аккаунт. Новый аккаунт. A new account. ",
	"html": "Сіз Қарағады қ. ӘКК-сының порталына тіркелдіңіз. Аты-жөні: %s. Рөлі: %s. Логин: %s | %s. Құпия сөз: %s \n\n\n" +
		"Вы были зарегистрированы на веб-портале СПК Караганды. ФИО: %s. Роль: %s. Логин: %s | %s. Пароль: %s \n\n\n" +
		"You have been registered on the web-platform of SEC Karaganda. Name: %s. Role: %s. Login: %s | %s. Password: %s \n",
}

// get map
func (n *NotifyCreateProfile) GetMap() map[string]string {
	return MapNotifyCreateProfile
}

// sender
func (n *NotifyCreateProfile) GetFrom() (string) {
	return constants.BaseEmailAddress
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
	return n.GetMap()[constants.KeyEmailSubject]
}

// body in html
func (n *NotifyCreateProfile) GetHtml() string {

	if n.User.Id < 1 || n.User.Email.Id < 1 || n.User.Phone.Id < 1 {
		n.User.Id = n.UserId
		_ = n.User.OnlyGetByIdPreloaded(GetDB())
	}

	body := n.GetMap()[constants.KeyEmailHtml]

	// prepare email address & phone number with country code (+7 for KZ)
	email := n.User.Email.Address
	phoneNumber := n.User.Phone.Ccode + n.User.Phone.Number

	// You have ... Name: %s. Role: %s. Login: %s | %s. Password: %s
	body = fmt.Sprintf(body, n.User.Fio, constants.MapRole[n.User.Role.Name]["kaz"], email, phoneNumber, n.RawPassword,
		n.User.Fio, constants.MapRole[n.User.Role.Name]["rus"], email, phoneNumber, n.RawPassword,
		n.User.Fio, constants.MapRole[n.User.Role.Name]["eng"], email, phoneNumber, n.RawPassword)

	return body
}

func (n *NotifyCreateProfile) GetPlainText() string {
	return n.GetHtml()
}

func (n *NotifyCreateProfile) GetProjectId() uint64 {
	return uint64(0)
}
