package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

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
func (p *Phone) CreateOnly(trans *gorm.DB) error {
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
func (p *Phone) DeleteById(trans *gorm.DB) error {
	return trans.Delete(p, "id = ?", p.Id).Error
}

