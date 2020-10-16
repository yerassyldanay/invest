package utils

import "time"

/*
	to handle it in one place
 */
func GetCurrentTime() time.Time {
	return time.Now()
}

func GetCurrentTruncatedDate() time.Time {
	return time.Now().Truncate(time.Hour).Add(time.Minute)
}

// prettify time
func OnlyPrettifyTime(t time.Time) (time.Time) {
	return t
}