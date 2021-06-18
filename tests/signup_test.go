package tests

import (
	"github.com/stretchr/testify/require"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/service"
	"github.com/yerassyldanay/invest/utils/constants"
	"testing"
	"time"
)

// generate a new user

/*
	TestSignUp:
		* generate user
		* create user on database
		* check data
		* email not verified
		* verification of email address
		* check - sign in
 */
func TestSignUp(t *testing.T) {
	// generate
	user := HelperGenerateNewUser()
	password := user.Password
	newUser := user

	// sign up
	is := service.InvestService{}
	msg := is.SignUp(newUser)

	// check
	require.Zero(t, msg.ErrMsg)

	// verify
	newUser = model.User{Email: model.Email{
		Address: user.Email.Address,
	}}
	err := newUser.OnlyGetByEmailAddress(model.GetDB())
	// check
	require.NoError(t, err)

	// preload all information
	err = newUser.OnlyGetByIdPreloaded(model.GetDB())

	// check
	require.NoError(t, err)
	require.NotZero(t, newUser.Fio, newUser.Password, newUser.Email.Address, newUser.Phone.Number)
	require.Equal(t, constants.RoleInvestor, newUser.Role.Name)

	// check - email
	require.NotZero(t, newUser.Email.SentCode)
	require.False(t, newUser.Email.Verified)
	require.Condition(t, func() (bool) { return newUser.Email.Deadline.After(time.Now()) })

	// check verification
	up := model.UserPermission{
		UserId:     newUser.Id,
	}
	msg = up.Check_db_whether_this_user_account_is_confirmed()

	// check - must be not zero
	require.NotZero(t, msg.ErrMsg)

	// confirm email address
	is = service.InvestService{}
	email := newUser.Email
	msg = is.EmailConfirm(email)

	// check
	require.Zero(t, msg.ErrMsg)

	// sign in
	sis := model.SignIn{
		KeyUsername:   "email",
		Value:         newUser.Email.Address,
		Password:      password,
	}
	msg = sis.SignIn()

	// check
	require.Zero(t, msg.ErrMsg)
}
