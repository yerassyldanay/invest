package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

/*
	error messages for validation
*/
var errorInvalidSubject = errors.New("invalid subject name")
var errorInvalidBody = errors.New("invalid body of the email")
var errorInvalidProjectId = errors.New("invalid project id")
var errorInvalidUserId = errors.New("invalid user id")

/*
	comment documents must be stored on disk beforehand
		comment can have no docs attached
*/
func (c *Comment) Validate() error {
	switch {
	case c.Subject == "":
		return errorInvalidSubject
	case c.Body == "":
		return errorInvalidBody
	case c.ProjectId == 0:
		return errorInvalidProjectId
	case c.UserId == 0:
		return errorInvalidUserId
	}

	return nil
}

func (c *Comment) Only_create(trans *gorm.DB) error {
	return trans.Create(c).Error
}

func (c *Comment) Only_get_comments_by_project_id(offset interface{}, tx *gorm.DB) (comments []Comment, err error) {
	err = tx.Offset(offset).Find(&comments, "project_id = ?", c.ProjectId).Error
	return comments, err
}

func (c *Comment) Only_get_comment_by_comment_id(tx *gorm.DB) (err error) {
	return tx.First(c, "id = ?", c.Id).Error
}



