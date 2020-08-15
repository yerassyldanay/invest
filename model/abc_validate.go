package model

import (
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
func Validate_password(p string) bool {
	upper := regexp.MustCompile("[A-Z]+")
	lower := regexp.MustCompile("[a-z]+")
	number := regexp.MustCompile("[0-9]+")

	var ok = upper.FindString(p) != "" && lower.FindString(p) != "" && number.FindString(p) != ""
	return ok
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