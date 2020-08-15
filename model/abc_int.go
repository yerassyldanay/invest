package model

import "strconv"

/*
	to prettify offset
 */
func Offset(o string) (int) {
	if i, err := strconv.Atoi(o); err == nil {
		if i > 0 {
			return i
		}
	}
	return 0
}