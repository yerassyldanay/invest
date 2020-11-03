package model

import (
	"fmt"
	"invest/utils/constants"
)

/*
	NotifyCode:
		this is to send code to an email address
 */
type NotifyCode struct {
	//Hash				string				`json:"hash"`
	Code				string				`json:"code"`
	Address				string				`json:"address"`
}

var MapNotifyCode = map[string]string{
	"subject": "Операцияны растауға арналған құпия сөз." +
		"  Код для валидации." +
		" Confirmation code.",
	"html": "Сіздің растауға арналған құпия сөзіңіз: %s. \n\n\n Ваш код для валидации: %s. \n\n\n Confirmation code: %s. \n\n\n",
}

// get map
func (n *NotifyCode) GetMap() map[string]string {
	return MapNotifyCode
}

// sender
func (n *NotifyCode) GetFrom() (string) {
	return constants.BaseEmailAddress
}

// get the list of users, who has connection to project
func (n *NotifyCode) GetToList() []string {
	return []string{ n.Address }
}

// get subject
func (n *NotifyCode) GetSubject() string {
	return MapNotifyCode[constants.KeyEmailSubject]
}

// body in html
func (n *NotifyCode) GetHtml() string {
	body := n.GetMap()[constants.KeyEmailHtml]

	resp := fmt.Sprintf(body, n.Code, n.Code, n.Code)

	return resp
}

func (n *NotifyCode) GetPlainText() string {
	return n.GetHtml()
}

func (n *NotifyCode) GetProjectId() uint64 {
	return uint64(0)
}

