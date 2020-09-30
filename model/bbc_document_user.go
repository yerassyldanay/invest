package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"invest/utils"
	"time"
)

type DocumentUserStatus struct {
	GantaId					uint64				`json:"ganta_id" gorm:"gantas.id"`
	UserId					uint64				`json:"user_id" gorm:"users.id"`
	Status					string				`json:"status" gorm:"not null"`
	Modified				time.Time				`json:"modified" gorm:"default:now()"`
}

func (du *DocumentUserStatus) TableName() string {
	return "document_user_statuses"
}

// errors
var errorDocUserInvalidStatus = errors.New("an invalid status for a document / ganta child step")

func (du *DocumentUserStatus) Validate() (error) {
	switch du.Status {
	case utils.ProjectStatusReject:
	case utils.ProjectStatusAccept:
	case utils.ProjectStatusReconsider:
	default:
		return errorDocUserInvalidStatus
	}

	du.Modified = utils.GetCurrentTime()

	return nil
}

func (du *DocumentUserStatus) OnlyCreate(tx *gorm.DB) (err error) {
	du.Modified = utils.GetCurrentTime()
	err = tx.Create(du).Error
	return err
}

func (du *DocumentUserStatus) OnlyDelete(tx *gorm.DB) (err error) {
	err = tx.Delete(du, "ganta_id = ? and user_id = ?", du.GantaId, du.UserId).Error
	return err
}

func (du *DocumentUserStatus) AreAllValidChildGantaIds(ids []uint64, project_id uint64, tx *gorm.DB) (bool) {
	var count int
	err := tx.Table(Ganta{}.TableName()).
		Where("ganta_parent_id != 0 and id in (?) and project_id = ?", ids, project_id).
		Count(&count).Error

	ok := err != nil && count == len(ids)
	return ok
}

func (du *DocumentUserStatus) AreAllValidParentGantaIds(ids []uint64, project_id uint64, tx *gorm.DB) (bool) {
	var count int
	err := tx.Table(Ganta{}.TableName()).
		Where("ganta_parent_id != 0 and id in (?) and project_id = ?", ids, project_id).
		Count(&count).Error

	ok := err != nil && count == len(ids)
	return ok
}

func (du *DocumentUserStatus) AreAllValidGantaIds(ids []uint64, project_id uint64, tx *gorm.DB) (bool) {
	var count int
	err := tx.Table(Ganta{}.TableName()).
		Where("id in (?) and project_id = ?", ids, project_id).
		Count(&count).Error

	ok := err != nil && count == len(ids)
	return ok
}

