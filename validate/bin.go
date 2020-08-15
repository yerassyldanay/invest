package validate

import (
	"regexp"
)

/*
	this function validates the bin number of the organization
		this is an additional function for govalidator
 */
func Validate_bin(i interface{}, o interface{}) (bool) {
	var bin = i.(string)
	r := regexp.MustCompile("[0-9]+")
	if ok := r.Match([]byte(bin)); !ok {
		return false
	}
	return true
}

