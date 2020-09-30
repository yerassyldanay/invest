package model

import (
	"github.com/jinzhu/gorm"
)

func (fi *Finance) Validate() error {
	if fi.ProjectId < 1 {
		return errorInvalidProjectId
	}

	return nil
}

func (fi *Finance) OnlyCreate(tx *gorm.DB) error {
	return tx.Create(fi).Error
}

func (fi *Finance) OnlySave(tx *gorm.DB) error {
	return tx.Save(fi).Error
}

func (fi *Finance) OnlyUpdateAll(tx *gorm.DB) error {
	return tx.Updates(*fi).Error
}
