package model

import (
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

/*
	the phone number must meet certain pattern requirement
*/
func (p *Phone) Is_valid() bool {
	number, err := regexp.Compile("[0-9]+")
	ccode, err2 := regexp.Compile("\\+[0-9]{1}")
	if err != nil || err2 != nil {
		return false
	}

	return number.Match([]byte(p.Number)) && ccode.Match([]byte(p.Ccode))
}
