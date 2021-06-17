package model

import (
	"fmt"
	"github.com/yerassyldanay/invest/utils/constants"
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
	"subject": "Жобаға түсініктеме жасалды. Добавлен комментарий к проекту. A comment is added to a project.",
	"html": "'%s' атаулы жоба %s. %s түсініктемені енгізді: %s \n\n\n" +
		"Проект '%s' был %s . %s оставил(-а) комментарий: %s \n\n\n" +
		"The project '%s' has been %s. %s has left the comment: %s \n",
}

// get map
func (n *NotifyAddComment) GetMap() map[string]string {
	return MapNotifyAddComment
}

// sender
func (n *NotifyAddComment) GetFrom() (string) {
	return constants.BaseEmailAddress
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
	return n.GetMap()[constants.KeyEmailSubject]
}

// body in html
func (n *NotifyAddComment) GetHtml() string {
	body := n.GetMap()[constants.KeyEmailHtml]

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

	// The project '%s' has been %s. %s has left the following comment: %s
	body = fmt.Sprintf(body,
		n.Project.Name, constants.MapConvertStatusToHumanReadableWordStatusThenLang[n.Status]["kaz"], n.CommentedBy.Fio, n.CommentBody,
		n.Project.Name, constants.MapConvertStatusToHumanReadableWordStatusThenLang[n.Status]["rus"], n.CommentedBy.Fio, n.CommentBody,
		n.Project.Name, constants.MapConvertStatusToHumanReadableWordStatusThenLang[n.Status]["eng"], n.CommentedBy.Fio, n.CommentBody)

	return body
}

func (n *NotifyAddComment) GetPlainText() string {
	return n.GetHtml()
}

func (n *NotifyAddComment) GetProjectId() uint64 {
	return uint64(0)
}

