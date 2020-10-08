package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"gopkg.in/validator.v2"
	"invest/utils"
	"time"
)

/*

 */
type Email struct {
	Id						uint64				`json:"id" gorm:"AUTO_INCREMENT;primary key"`

	Address					string				`json:"address" gorm:"UNIQUE; size: 255" validate:"required,email"`
	Verified				bool				`json:"verified"  gorm:"default:false"`

	SentCode				string				`json:"sent_code" gorm:"size: 10"`
	SentHash				string				`json:"sent_hash"`

	Deadline				time.Time			`json:"deadline" gorm:"default:null"`
}

func (Email) TableName() string {
	return "emails"
}

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

/*
	is it verified?
 */
func (e *Email) IsVerified() (map[string]interface{}, error) {
	var err error
	if err = GetDB().Model(&Email{}).Where("address", e.Address).First(e).Error; err == nil {
		if e.Verified {
			return nil, nil
		}
	}

	return utils.ErrorEmailIsNotVerified, err
}

/*
	* it validates an email address
	* sends a code & link
	* stores on db
 */
func (e *Email) Create_email() (map[string]interface{}, error) {
	trans := GetDB().Begin()

	if err := validator.Validate(e); err != nil {
		trans.Rollback()
		return utils.ErrorInvalidParameters, err
	}

	if err := trans.Create(e).Error; err != nil {
		trans.Rollback()
		return utils.ErrorInternalDbError, err
	}

	return utils.NoErrorFineEverthingOk, trans.Commit().Error
}


