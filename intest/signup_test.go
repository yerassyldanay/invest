package intest

import (
	"fmt"
	"invest/model"
	"invest/service"
	"invest/utils/constants"
	"testing"
	"time"
)

/*
	TestSignUp:
		* create user
		* send message to an email address
 */
func TestSignUp(t *testing.T) {
	var user = model.User{
		Username:       "yerassyl_danay",
		Password:       "YerassylDanay1234",
		Fio:            "Yerassyl Danay",
		Email:          model.Email{
			Address:  "yerassyl.danay@mail.ru",
		},
		Phone:          model.Phone{
			Ccode:    "+7",
			Number:   "7058686509",
		},
		Organization:   model.Organization{
			Bin: "190940011748",
		},
	}

	is := service.InvestService{
		BasicInfo: service.BasicInfo{
			Lang: constants.DefaultContentLanguage,
		},
	}

	//// create or use a mailer queue
	//mq := model.GetMailerQueue()
	//
	//// handle messages (a handler with context)
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//go mq.Handle(ctx)

	msg := is.SignUp(user)

	if msg.IsThereAnError() {
		t.Error(msg, msg.ErrMsg)
	}
}

func TestEmailVerification(t *testing.T) {
	// get the one, which not confirmed
	var email = model.Email{}
	err := email.OnlyGetNotConfirmedOne(model.GetDB())
	if err != nil {
		t.Error("expected nil, but got ", err)
		return
	}

	fmt.Println("not confirmed email", email)

	// confirm the email
	email.Id = 0
	email.Deadline = time.Time{}
	email.Verified = false

	// logic
	is := service.InvestService{}

	msg := is.EmailConfirm(email)
	if msg.IsThereAnError() {
		t.Error("expected not err, but got ", msg.ErrMsg)
		return
	}
}

func TestEmailUpdate(t *testing.T) {
	var email = model.Email{
		Id:       7,
		Address:  "yerassyl.danay@mail.ru",
		Verified: false,
		SentCode: "7223",
	}

	if ok := email.OnlyUpdateAfterConfirmation(model.GetDB()); !ok {
		t.Error("expected to update, but could not")
		return
	}
}
