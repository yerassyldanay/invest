package utils

import "time"

/*
	to handle it in one place
 */
func GetCurrentTime() time.Time {
	return time.Now().UTC()
}