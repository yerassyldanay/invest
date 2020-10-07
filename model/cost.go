package model

import "github.com/jinzhu/gorm"

func (cost *Cost) Validate() error {
	if cost.ProjectId < 1 {
		return errorInvalidProjectId
	}

	return nil
}

func (cost *Cost) OnlyCreate(tx *gorm.DB) error {
	return tx.Create(cost).Error
}

func (cost *Cost) OnlySave(tx *gorm.DB) error {
	return tx.Save(cost).Error
}

func (cost *Cost) OnlyUpdate(tx *gorm.DB) error {
	return tx.Updates(*cost).Error
}

func (cost *Cost) OnlyGetByProjectId(tx *gorm.DB) (err error) {
	err = tx.First(cost, "project_id = ?", cost.ProjectId).Error
	return err
}
