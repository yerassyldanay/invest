package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

/*
	error messages for validation
*/
var errorInvalidBody = errors.New("invalid body of the email")
var errorInvalidProjectId = errors.New("invalid project id")
var errorInvalidUserId = errors.New("invalid user id")
var errorInvalidStatus = errors.New("invalid status of a project")

/*
	comment documents must be stored on disk beforehand
		comment can have no docs attached
*/
func (c *Comment) Validate() error {
	switch {
	case c.Body == "":
		return errorInvalidBody
	case c.ProjectId == 0:
		return errorInvalidProjectId
	case c.UserId == 0:
		return errorInvalidUserId
	}

	//if c.Status != utils.ProjectStatusDone &&
	//	c.Status != utils.ProjectStatusPendingAdmin &&
	//	c.Status != utils.ProjectStatusRejected {
	//		return errorInvalidStatus
	//}

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

