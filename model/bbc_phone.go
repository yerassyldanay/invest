package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/yerassyldanay/invest/utils/errormsg"
	"regexp"
)

type Phone struct {
	Id				uint64				`json:"id" gorm:"AUTO_INCREMENT;primary key"`

	Ccode			string				`json:"ccode"`
	Number			string				`json:"number"`

	SentCode		string				`json:"sent_code"`
	Verified		bool				`json:"verified"`
}

func (Phone) TableName() string {
	return "phones"
}

/*
	validate info
		there will be cases when we need to create a phone number without
			a code
*/
func (p *Phone) IsNumberValid() bool {
	return p.Number != "" && p.Ccode != ""
}

func (p *Phone) IsSentCodeValid() bool {
	return p.SentCode != ""
}

/*
	create:
		consider two cases
*/
func (p *Phone) OnlyCreate(trans *gorm.DB) error {
	return trans.Create(p).Error
}

func (p *Phone) CreateAfterFullValidation(trans *gorm.DB) error {
	ok := p.IsSentCodeValid() && p.IsNumberValid()
	if !ok {
		return errors.New("phone info (ccode, number or sentCode) is not valid")
	}

	return trans.Create(p).Error
}

func (p *Phone) CreateAfterValidation(trans *gorm.DB) error {
	if !p.IsNumberValid() {
		return errors.New("phone info (ccode or number) is not valid")
	}

	return trans.Create(p).Error
}

/*
	delete phone number
*/
func (p *Phone) OnlyDeleteById(trans *gorm.DB) (err error) {
	err = trans.Delete(p, "id = ?", p.Id).Error
	return err
}

func (p *Phone) OnlySave(tx *gorm.DB) (err error) {
	err = tx.Save(p).Error
	return err
}

// get
func (p *Phone) OnlyGetByCcodeAndNumber(tx *gorm.DB) (error) {
	err := tx.First(p, "ccode = ? and number = ?", p.Ccode, p.Number).Error
	return err
}

func (p *Phone) IsVerified() (map[string]interface{}, error) {
	var err error
	if err = GetDB().Model(&Phone{}).Where("ccode=? and number=?", p.Ccode, p.Number).First(p).Error; err == nil {
		if p.Verified {
			return nil, nil
		}
	}

	return errormsg.ErrorPhoneNumberIsNotVerified, err
}

// errors
var errorPhoneInvalidCodeOrNumber = errors.New("invalid phone number")

/*
	the phone number must meet certain pattern requirement
*/
func (p *Phone) Validate() (err error) {
	number, err := regexp.Compile("[0-9]+")
	ccode, err2 := regexp.Compile("\\+[0-9]{1}")

	switch {
	case err != nil:
		return err
	case err2 != nil:
		return err2
	}

	ok := number.Match([]byte(p.Number)) && ccode.Match([]byte(p.Ccode))
	if !ok {
		return errorPhoneInvalidCodeOrNumber
	}

	return nil
}
