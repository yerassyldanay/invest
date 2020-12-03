package tests

import (
	"github.com/stretchr/testify/require"
	"invest/model"
	"invest/service"
	"invest/utils/helper"
	"testing"
	"time"
)

func TestEmailConfirmation(t *testing.T) {
	// generate
	address := helper.Generate_Random_String(10) + "_test@gmail.com"
	code := helper.Generate_Random_Number(5)

	// email
	email := model.Email{
		Address:  address,
		Verified: false,
		SentCode: code,
		Deadline: time.Now().Add(time.Hour * 24),
	}

	err := email.OnlyCreate(model.GetDB())

	// check
	require.NoError(t, err)
	require.NotZero(t, email.Id)
	require.Equal(t, address, email.Address)
	require.Equal(t, code, email.SentCode)

	// get email
	newEmail := model.Email{
		Address: email.Address,
		SentCode: code,
	}

	// logic
	is := service.InvestService{}
	msg := is.EmailConfirm(newEmail)

	// check]
	require.Zero(t, msg.ErrMsg)
	require.Equal(t, msg.Status, 200)
}

/*
	steps:
	* get user
	* reset password
	* sign in
 */
func TestResetPassword(t *testing.T) {
	// get email address
	user := model.User{Id: 1}
	err := user.OnlyGetByIdPreloaded(model.GetDB())
	email := user.Email

	// check
	require.NoError(t, err)

	// reset password
	fp := model.ForgetPassword{EmailAddress: user.Email.Address}

	is := service.InvestService{}
	msg := is.Password_reset_send_message(fp)

	// check
	require.Zero(t, msg.ErrMsg)
	require.Equal(t, email.Address, user.Email.Address)

	// get code
	fp = model.ForgetPassword{EmailAddress: email.Address}
	err = fp.OnlyGetByAddress(model.GetDB())

	// check
	require.NoError(t, err)
	require.NotZero(t, fp.Code)
	require.NotZero(t, fp.EmailAddress)

	//change password
	newPassword := helper.Generate_Random_String(20)
	newFp := model.ForgetPassword {
		NewPassword:  newPassword,
		EmailAddress: fp.EmailAddress,
		Code:         fp.Code,
	}

	// check
	msg = is.Password_reset_change_password(newFp)

	// check
	require.Zero(t, msg.ErrMsg)

	// sign in
	sis := model.SignIn{
		KeyUsername:   "email",
		Value:         email.Address,
		Password:      newPassword,
	}
	msg = sis.Sign_in()

	// check
	require.Zero(t, msg.ErrMsg)

	// get user & make sure that password has been changed
	newUser := model.User{Id: user.Id}
	err = newUser.OnlyGetByIdPreloaded(model.GetDB())

	// check
	require.NoError(t, err)
	require.Equal(t, user.Fio, newUser.Fio)
	require.NotEqual(t, user.Password, newUser.Password)
	require.Equal(t, user.Email.Address, newUser.Email.Address)
}
