package model

import (
	"errors"
	"net/http"
	"regexp"
)

/*
	password is valid if:
		* 20 >= len >= 8
		* A-Z
		* a-z
		* 0-9
*/
var errorPasswordInvalidLength = errors.New("invalid length: must be 8-20 characters")
var errorPasswordNoUpperLetter = errors.New("no upper letter characters")
var errorPasswordNoLowerLetter = errors.New("no lower letter characters")
var errorPasswordNoDigits = errors.New("no digits")

func Validate_password(val string) error {
	upper := regexp.MustCompile("[A-Z]+")
	lower := regexp.MustCompile("[a-z]+")
	number := regexp.MustCompile("[0-9]+")

	switch {
	case upper.FindString(val) == "":
		return errorPasswordNoUpperLetter
	case lower.FindString(val) == "":
		return errorPasswordNoLowerLetter
	case number.FindString(val) == "":
		return errorPasswordNoDigits
	case len(val) > 20:
		return errorPasswordInvalidLength
	case len(val) < 8:
		return errorPasswordInvalidLength
	}

	return nil
}

/*
	validate bin
 */
func Validate_bin(s string) bool {
	ok, _ := regexp.MatchString("^[0-9]+$", s)
	return ok
}

/*
	validate that all fields contain at least one character
 */
func Validate_all_contain_at_least_one_char(ss... string) bool {
	for _, s := range ss {
		if s == "" {
			return false
		}
	}
	return true
}

/*
	DRY - get value from query
 */
func Get_value_from_query(r *http.Request, key string) string {
	l := r.URL.Query()[key]
	if len(l) == 0 {
		return ""
	}
	return l[0]
}
