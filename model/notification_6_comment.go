package model

import (
	"fmt"
	"invest/utils"
)

type NotifyAddComment struct {
	CommentedById				uint64			`json:"commented_by_id"`
	CommentedBy					User

	ProjectId					uint64			`json:"project_id"`
	Project						Project

	CommentBody					string			`json:"comment_body"`
	Status						string			`json:"status"`
}

var MapNotifyAddComment = map[string]string{
	"subject": "Жобаға түсініктеме жасалды. Добавлен комментарий к проекту.  ",
	"html": "Түсініктемес жасаушы: %s. Түсініктеме: %s. Статус: %s. Жоба атауы: \n" +
		"Комментарий от: %s. Коммертарий: %s. Статус: %s. Название проекта: %s. \n" +
		"Comment from: %s. Comment body: %s. Status: %s. The name of the project: %s",
}

// get map
func (n *NotifyAddComment) GetMap() map[string]string {
	return MapNotifyAddComment
}

// sender
func (n *NotifyAddComment) GetFrom() (string) {
	return utils.BaseEmailAddress
}

// get the list of users, who has connection to project
func (n *NotifyAddComment) GetToList() []string {
	// get all users, who has connection to the project
	var email = Email{}
	emails, err := email.OnlyGetEmailsHasConnectionToProject(n.ProjectId, GetDB())
	if err != nil {
		return []string{}
	}

	var emailList = []string{}
	for _, email = range emails {
		emailList = append(emailList, email.Address)
	}

	return emailList
}

// get subject
func (n *NotifyAddComment) GetSubject() string {
	return n.GetMap()[utils.KeyEmailSubject]
}

// body in html
func (n *NotifyAddComment) GetHtml() string {
	body := n.GetMap()[utils.KeyEmailHtml]

	if n.CommentedBy.Id < 1 {
		// get user, who commented
		n.CommentedBy.Id = n.CommentedById
		if err := n.CommentedBy.OnlyGetByIdPreloaded(GetDB()); err != nil {
			return ""
		}
	}

	if n.Project.Id < 1 {
		// get project, on which commented
		n.Project.Id = n.ProjectId
		if err := n.Project.OnlyGetByIdPreloaded(GetDB()); err != nil {
			return ""
		}
	}

	// Comment from: %s. Comment itself: %s. Status: %s. The name of the project: %s
	body = fmt.Sprintf(body, n.CommentedBy.Fio, n.CommentBody, n.Status, n.Project.Name,
		n.CommentedBy.Fio, n.CommentBody, n.Status, n.Project.Name,
		n.CommentedBy.Fio, n.CommentBody, n.Status, n.Project.Name)

	return body
}

func (n *NotifyAddComment) GetPlainText() string {
	return n.GetHtml()
}

func (n *NotifyAddComment) GetProjectId() uint64 {
	return uint64(0)
}

