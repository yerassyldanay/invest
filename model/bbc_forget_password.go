package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"net/mail"
	"time"
)

type ForgetPassword struct {
	NewPassword			string						`json:"new_password" gorm:"-"`
	EmailAddress		string						`json:"email_address" gorm:"primaryKey"`

	Code				string				`json:"code"`
	Deadline			time.Time			`json:"deadline"`
	
	Lang 				string				`json:"lang" gorm:"-"`
}

func (fp *ForgetPassword) TableName() string {
	return "forget_passwords"
}

// errors
var errorForgetPasswordInvalidEmail = errors.New("invalid email address is provided")

func (fp *ForgetPassword) Validate() (error) {
	_, err := mail.ParseAddress(fp.EmailAddress)
	return err
}

// create
func (fp *ForgetPassword) OnlyCreate(tx *gorm.DB) (err error) {
	err = tx.Create(fp).Error
	return err
}

// update
func (fp *ForgetPassword) OnlyUpdateByEmailAddress(tx *gorm.DB, fields... string) (err error) {
	err = tx.Model(&ForgetPassword{EmailAddress: fp.EmailAddress}).Select(fields).
		Updates(map[string]interface{}{
			"code": fp.Code,
			"deadline": fp.Deadline,
	}).Error

	return err
}

// delete
func (fp *ForgetPassword) OnlyDelete(tx *gorm.DB) (err error) {
	err = tx.Delete(fp, "email_address = ?", fp.EmailAddress).Error
	return err
}

func (fp *ForgetPassword) OnlyGet(tx *gorm.DB) (err error) {
	err = tx.First(fp, "email_address = ?", fp.EmailAddress).Error
	return err
}

func (fp *ForgetPassword) OnlyGetByCode(tx *gorm.DB) (err error) {
	err = tx.First(fp, "code = ?", fp.Code).Error
	return err
}

