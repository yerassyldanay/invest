package randomer

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
)

func TestRandomDigit(t *testing.T) {
	number := RandomDigit(6)
	r, err := regexp.Compile("[0-9]{6}")
	require.NoError(t, err)

	fmt.Println("number: ", number)

	ok := r.Match([]byte(number))
	require.True(t, ok)

	ok = r.Match([]byte("xa xa"))
	require.False(t, ok)
}
