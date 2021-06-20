package model

import "C"
import "errors"

const GetLimit = 20

func (c *User) RemoveAllUsersWithNotConfirmedEmail() (map[string]interface{}, error) {
	return nil, nil
}

var ErrorValidateFioPhoneEmail_InvalidPassword = errors.New("invalid password is provided")
var ErrorValidateFioPhoneEmail_InvalidFio = errors.New("invalid FIO is provided")
var ErrorValidateFioPhoneEmail_InvalidPhone = errors.New("invalid phone is provided")
var ErrorValidateFioPhoneEmail_InvalidEmail = errors.New("invalid email is provided")

func (c *User) ValidateFioPhoneEmail() error {
	if len(c.Password) < 8 || len(c.Password) > 50 {
		return ErrorValidateFioPhoneEmail_InvalidPassword
	} else if len(c.Fio) == 0 {
		return ErrorValidateFioPhoneEmail_InvalidFio
	} else if len(c.Phone.Ccode)*len(c.Phone.Number) == 0 {
		return ErrorValidateFioPhoneEmail_InvalidPhone
	} else if len(c.Email.Address) == 0 {
		return ErrorValidateFioPhoneEmail_InvalidEmail
	}

	return nil
}
