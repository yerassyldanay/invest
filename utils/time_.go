package utils

import "time"

/*
	to handle it in one place
 */
func GetCurrentTime() time.Time {
	return time.Now()
}

// prettify time
func OnlyPrettifyTime(t time.Time) (time.Time) {
	return t
}