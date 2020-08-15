package utils

import "regexp"

func Is_it_free_from_sql_injection(str string) bool {
	var onlyCharRegexp, err = regexp.Compile("[a-zA-Z0-9]+")
	if err != nil {
		return false
	}

	return onlyCharRegexp.Match([]byte(str))
}