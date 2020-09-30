package model

import (
	"errors"
	"invest/utils"
	"regexp"
)

func (p *Phone) Is_verified() (map[string]interface{}, error) {
	var err error
	if err = GetDB().Model(&Phone{}).Where("ccode=? and number=?", p.Ccode, p.Number).First(p).Error; err == nil {
		if p.Verified {
			return nil, nil
		}
	}

	return utils.ErrorPhoneNumberIsNotVerified, err
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
