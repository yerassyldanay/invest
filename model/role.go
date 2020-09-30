package model

import "github.com/jinzhu/gorm"

func (r *Role) OnlyGetByName(tx *gorm.DB) (err error) {
	err = tx.First(r, "name = ?", r.Name).Error
	return err
}

func (r *Role) OnlyGetById(tx *gorm.DB) (err error) {
	err = tx.First(r, "id = ?", r.Id).Error
	return err
}

func (r *Role) OnlyCreate(tx *gorm.DB) (err error) {
	err = tx.Create(r).Error
	return err
}

func (r *Role) OnlySave(tx *gorm.DB) (err error) {
	err = tx.Save(r).Error
	return err
}


