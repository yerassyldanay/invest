package intest

import (
	"invest/model"
	"invest/service"
	"testing"
)

func TestEmailConfirmation(t *testing.T) {
	// get email
	email := model.Email{Address: "yerassyl.danay@mail.ru"}
	if err := email.OnlyGetByAddress(model.GetDB()); err != nil {
		t.Error(err)
	}

	// logic
	is := service.InvestService{}
	msg := is.EmailConfirm(email)
	if msg.IsThereAnError() {
		t.Error(msg.ErrMsg)
	}
}

func TestResetPassword(t *testing.T) {
	// get email
	fp := model.ForgetPassword{
		NewPassword:  "6z24HXMd7nLeZAE",
		EmailAddress: "invest.dept.spk@inbox.ru",
		Code:         "7777",
	}

	// logic
	is := service.InvestService{}
	msg := is.Password_reset_change_password(fp)
	if msg.IsThereAnError() {
		t.Error(msg.ErrMsg)
	}

	TestSignIn(t)
}
