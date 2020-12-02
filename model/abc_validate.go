package model

import (
	"errors"
	"net/mail"
	"regexp"
)

/*
	password is valid if:
		* 20 >= len >= 8
		* A-Z
		* a-z
		* 0-9
*/
var errorPasswordInvalidLength = errors.New("invalid length: must be 8-50 characters")
var errorPasswordNoUpperLetter = errors.New("no upper letter characters")
var errorPasswordNoLowerLetter = errors.New("no lower letter characters")
var errorPasswordNoDigits = errors.New("no digits")
var errorNonWordCharacter = errors.New("contains non-word character. must contain only [a-zA-Z0-9]{n+}")
var errorNonDigitCharacter = errors.New("contains non digit character")

var containUpperLetter = regexp.MustCompile("[A-Z]+")
var containLowerLetter = regexp.MustCompile("[a-z]+")
var containDigit = regexp.MustCompile("[0-9]+")
var containOneNonDigit = regexp.MustCompile("[^0-9]{1}")
var containNonWordCharacter = regexp.MustCompile("[^0-9A-Za-z]{1}")

// validate password
func OnlyValidatePassword(val string) error {

	switch {
	case containUpperLetter.FindString(val) == "":
		return errorPasswordNoUpperLetter
	case containLowerLetter.FindString(val) == "":
		return errorPasswordNoLowerLetter
	case containDigit.FindString(val) == "":
		return errorPasswordNoDigits
	case len(val) > 50:
		return errorPasswordInvalidLength
	case len(val) < 8:
		return errorPasswordInvalidLength
	}

	return nil
}

// validate org bin
func OnlyValidateBin(bin string) error {
	switch {
	case containOneNonDigit.FindString(bin) != "":
		return errorNonDigitCharacter
	}

	return nil
}

// email
func OnlyValidateEmailAddress(emailAddress string) error {
	// parse name & address
	addr, err := mail.ParseAddress(emailAddress)
	if err != nil {
		return err
	}

	// check address
	if addr.Address == "" {
		return errors.New("invalid email address")
	}

	return nil
}