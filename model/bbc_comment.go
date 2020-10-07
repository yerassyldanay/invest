package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"invest/utils"
)

type Comment struct {
	Id					uint64					`json:"id" gorm:"primary key"`
	Body				string					`json:"body" gorm:"default:''"`

	UserId				uint64					`json:"user_id" gorm:"foreignkey:users.id"`
	ProjectId			uint64					`json:"project_id" gorm:"foreignkey:projects.id"`

	Status				string					`json:"status" gorm:"not null"`
}

func (Comment) TableName() string {
	return "comments"
}

type SpkComment struct {
	Comment					Comment							`json:"comment"`
	DocStatuses				[]DocumentUserStatus			`json:"doc_statuses"`
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

func (sc *SpkComment) OnlyCreateStatuses(tx *gorm.DB) (err error) {
	// store statuses of each document on db
	//var wg = sync.WaitGroup{}
	var errorChan = make(chan error, 1)

	for _, docStatus := range sc.DocStatuses {
		docStatus := docStatus

		//defer wg.Add(1)
		func(docStatus *DocumentUserStatus, trans *gorm.DB) {
			//defer wg.Done()
			err := docStatus.OnlyCreate(trans)
			if err != nil {
				select {
				case errorChan <- err:
				default:
					// pass if the chan is full
				}
			}
		}(&docStatus, tx)
	}
	//wg.Wait()

	select {
	case err := <- errorChan:
		return err
	default:
	}

	close(errorChan)
	return nil
}