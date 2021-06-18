package service

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"github.com/yerassyldanay/invest/model"
	"testing"
)

type ElementHelperTestInvestService_SignUp struct {
	EmailAddress string
	Code string
}

func helperTestInvestService_SignUp(t *testing.T) ElementHelperTestInvestService_SignUp {
	// generate
	user := HelperTestGenerateUserWithoutAnyInfoStored(t)

	// is
	is := InvestService{}
	msg := is.SignUp(user)

	// code
	randomCodeInterface, ok := msg.Message["code"]
	require.True(t, ok)
	require.NotZero(t, randomCodeInterface)
	randomCode := randomCodeInterface.(string)

	// redis
	helperRedisCheckExists(randomCode, t)
	userString, err := TestRedis.Get(randomCode).Result()
	require.NoError(t, err)

	// unmarshal
	var userFetched = model.User{}
	require.NoError(t, json.Unmarshal([]byte(userString), &userFetched))

	// equal
	require.Equal(t, user.Fio, userFetched.Fio)
	require.Equal(t, user.Email.Address, userFetched.Email.Address)
	require.Equal(t, user.Phone.Number, userFetched.Phone.Number)

	//HelperPrint(userFetched)

	return ElementHelperTestInvestService_SignUp{
		Code: randomCode,
		EmailAddress: randomCode,
	}
}

func TestInvestService_SignUp(t *testing.T) {
	code := helperTestInvestService_SignUp(t)

	// sign up
	is := InvestService{}
	msg := is.EmailConfirm(model.Email{
		Address:  code.EmailAddress,
		SentCode: code.Code,
	})

	// check
	require.Equal(t, msg.Status, 200)

	// get user
	userInt := msg.Message["user"]
	user := userInt.(model.User)

	require.NotZero(t, user.Id)

	HelperPrint(user)

	require.NotZero(t, user.Fio)
	require.NotZero(t, user.Email.Address)
	require.NotZero(t, user.Phone.Number)
	require.NotZero(t, user.Phone.Ccode)
}

