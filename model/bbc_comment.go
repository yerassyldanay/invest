package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"invest/utils"
	"time"
)

type Comment struct {
	Id					uint64					`json:"id" gorm:"primary key"`
	Body				string					`json:"body" gorm:"default:''"`

	UserId				uint64					`json:"user_id" gorm:"foreignkey:users.id"`
	ProjectId			uint64					`json:"project_id" gorm:"foreignkey:projects.id"`

	Status				string					`json:"status" gorm:"not null"`
	Created				time.Time				`json:"created" gorm:"default:now()"`
}

func (Comment) TableName() string {
	return "comments"
}

type SpkComment struct {
	Comment					Comment							`json:"comment"`
	Documents				[]Document						`json:"documents"`
}

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

	switch c.Status {
	case utils.ProjectStatusReconsider:
	case utils.ProjectStatusAccept:
	case utils.ProjectStatusReject:
	default:
		return errorInvalidStatus
	}

	c.Created = utils.GetCurrentTime()

	return nil
}

func (c *Comment) OnlyCreate(trans *gorm.DB) error {
	return trans.Create(c).Error
}

func (c *Comment) OnlyGetCommentsByProjectId(offset interface{}, tx *gorm.DB) (comments []Comment, err error) {
	err = tx.Offset(offset).Find(&comments, "project_id = ?", c.ProjectId).Error
	return comments, err
}

func (c *Comment) OnlyGetById(tx *gorm.DB) (err error) {
	return tx.First(c, "id = ?", c.Id).Error
}

func (sc *SpkComment) OnlyUpdateDocumentStatusesByIdAndProjectId(project_id uint64, documents []Document, tx *gorm.DB) (err error) {
	// update one by one
	for _, document := range documents {
		err = tx.Model(&Document{}).Where("id = ? and project_id = ?", document.Id, project_id).
			Updates(map[string]interface{}{
				"status": document.Status,
				"modified": utils.GetCurrentTime(),
		}).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (sc *SpkComment) OnlyValidateStatusesOfDocuments() (err error) {
	// check documents - can be escaped
	for _, document := range sc.Documents {
		// validate status
		switch {
		case document.Status == utils.ProjectStatusAccept:
		case document.Status == utils.ProjectStatusReconsider:
		case document.Status == utils.ProjectStatusReject:
		default:
			return errors.New("invalid document status. status: " + document.Status)
		}
	}
	return nil
}