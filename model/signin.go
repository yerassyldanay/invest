package model

import "errors"

// error messages for validation
var errorInvalidSignInKey = errors.New("invalid key")
var errorInvalidSignInPassword = errors.New("invalid password")
var errorInvalidSignInValue = errors.New("invalid value")

func (sis *SignIn) Validate() (error) {
	switch {
	case sis.Value == "":
		return errorInvalidSignInValue
	case sis.KeyUsername == "":
		return errorInvalidSignInKey
	case sis.Password == "":
		return errorInvalidSignInPassword
	}

	return nil
}
