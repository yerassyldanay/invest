package helper

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	mrand "math/rand"
)

const gen_only_letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const gen_letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const gen_numbers = "0123456789"

func generateSecureKey() string {
	k := make([]byte, 32)
	_, _ = io.ReadFull(rand.Reader, k)
	return fmt.Sprintf("%x", k)
}

/*
	converts provided string to hashed string
 */
func Generate_a_hashed_string(provided string) (string, error) {
	var hashedPassword, err =  bcrypt.GenerateFromPassword([]byte(provided), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

/*
	The following two functions will generate a token
*/
func Generate_Random_Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}
	return b, nil
}

func generate_random_string_or_number(n int, letters string) string {
	bytes, err := Generate_Random_Bytes(n)
	for err != nil {
		bytes, err = Generate_Random_Bytes(n)
	}

	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}

	return string(bytes)
}

func Generate_Random_String_No_Digits(n int) string {
	return generate_random_string_or_number(n, gen_only_letters)
}

func Generate_Random_String(n int) string {
	return generate_random_string_or_number(n, gen_letters)
}

func Generate_Random_Number(n int) string {
	return generate_random_string_or_number(n, gen_numbers)
}

// generate an integer between two values
func OnlyGenerateNumberBetweenTwoNumbers(start, end int) int {
	return mrand.Intn((end - start) + start)
}
