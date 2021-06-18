package randomer

import (
	"math/rand"
	"strings"
	"time"
)

// variables needed
const lowerCaseLetters = "abcdefghijklmnopqrstuvwxyz"
const upperCaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const allLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const allDigits = "0123456789"

// helps generate random different numbers
// without it, numbers will not be random
func init() {
	rand.Seed(time.Now().UnixNano())
}

// generate random integer
func RandomInt(minv, maxv int32) int32 {
	return minv + rand.Int31n(maxv - minv + 1)
}

// generate lower case letters
func RandomOnlyLowerCaseString(n int) string {
	var sb strings.Builder
	k := len(lowerCaseLetters)

	for i := 0; i < n; i++ {
		c := lowerCaseLetters[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// generate upper case letters
func RandomOnlyUpperCaseLetters(n int) string {
	var sb strings.Builder
	k := len(lowerCaseLetters)

	for i := 0; i < n; i++ {
		c := upperCaseLetters[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// generate a string (which consists uppercase & lowercase letters)
func RandomString(n int) string {
	var sb strings.Builder
	k := len(allLetters)

	for i := 0; i < n; i++ {
		c := allLetters[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// random integer, but a in the form of string
func RandomDigit(n int) string {
	var sb strings.Builder
	k := len(allDigits)

	for i := 0; i < n; i++ {
		c := allDigits[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// generates a random (KZ) number
func RandomPhoneNumber() string {
	return "+7" + RandomDigit(10)
}

// generates a random string
func RandomEmail() string {
	return RandomDigit(20) + "@gmail.com"
}

