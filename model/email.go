package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

/*

 */
func (e *Email) OnlyDeleteById(trans *gorm.DB) error {
	return trans.Delete(Email{}, "id = ?", e.Id).Error
}

/*
	this function creates an email address
		other fields must be set
 */
func (e *Email) OnlyCreate(trans *gorm.DB) error {
	return trans.Create(e).Error
}

/*
	create new email with hash & code
 */
func (e *Email) CreateEmailWithHashAfterValidation(trans *gorm.DB) error {
	if e.SentCode == "" || e.SentHash == "" {
		return errors.New("code or hash is empty")
	}
	return trans.Create(e).Error
}

/*
	pay attention to transaction:
		refer to documentation
 */
func (e *Email) OnlyGetById(trans *gorm.DB) error {
	return trans.First(e, "id = ?", e.Id).Error
}

func (e *Email) OnlySave (tx *gorm.DB) (error) {
	return tx.Save(e).Error
}
