package utils

import "time"

/*
	to handle it in one place
 */
func GetCurrentTime() time.Time {
	return time.Now()
}

func GetCurrentTruncatedDate() time.Time {
	minute := time.Duration(GetCurrentTime().Hour()) * 10 + time.Duration(GetCurrentTime().Minute())
	return time.Now().Truncate(time.Hour).Add(minute)
}

// prettify time
func OnlyPrettifyTime(t time.Time) (time.Time) {
	return t
}