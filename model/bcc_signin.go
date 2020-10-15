package model

import "errors"

type SignIn struct {
	KeyUsername			string				`json:"key"`
	Value				string				`json:"value"`
	Password			string				`json:"password"`
	Id					uint64				`json:"id"`
	TokenCompound				string				`json:"-,omitempty"`
}

// error messages for validation
var ErrorInvalidSignInKey = errors.New("invalid key")
var ErrorInvalidSignInPassword = errors.New("invalid password")
var ErrorInvalidSignInValue = errors.New("invalid value")

func (sis *SignIn) Validate() (error) {
	switch {
	case sis.Value == "":
		return ErrorInvalidSignInValue
	case sis.KeyUsername == "":
		return ErrorInvalidSignInKey
	case sis.Password == "":
		return ErrorInvalidSignInPassword
	}

	return nil
}
