package helper

import (
	"errors"
	"regexp"
)

func Is_it_free_from_sql_injection(str string) bool {
	var onlyCharRegexp, err = regexp.Compile("[a-zA-Z0-9]+")
	if err != nil {
		return false
	}

	return onlyCharRegexp.Match([]byte(str))
}

func OnlyCheckSqlInjection(str string) (error) {
	for _, ch := range str {
		switch {
		case ch == 32:
			// space is ok
		case ch == 45:
			// pass
			// this is -
		case ch >= 48 && ch <= 57:
			// numbers are ok
		case ch >= 65 && ch <= 90:
			// A-Z
		case ch >= 97 && ch <= 122:
			// a-z
		default:
			return errors.New("check: invalid string")
		}
	}

	return nil
}