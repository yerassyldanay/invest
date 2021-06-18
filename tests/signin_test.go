package tests

import (
	"github.com/yerassyldanay/invest/model"
	"reflect"
	"testing"
)

type signInDataStruct struct {
	Sis						model.SignIn
	ResultError				error
	ResultErrorString		string
}

func TestSignInData(t *testing.T) {
	var siCases = []signInDataStruct{
		{
			Sis:	model.SignIn{
				KeyUsername:   "",
				Value:         "investor",
				Password:      "KeRXaTaq5Ce8ULO",
			},
			ResultError: model.ErrorInvalidSignInKey,
		},
		{
			Sis: model.SignIn{
				KeyUsername:   "username",
				Value:         "investor",
				Password:      "",
				Id:            0,
				TokenCompound: "",
			},
			ResultError: model.ErrorInvalidSignInPassword,
		},
		{
			Sis: model.SignIn{
				KeyUsername:   "username",
				Value:         "investor",
				Password:      "somepassword",
			},
			ResultError: nil,
		},
	}

	for _, siCase := range siCases {
		err := siCase.Sis.Validate()
		if !reflect.DeepEqual(err, siCase.ResultError) {
			t.Error("expected: ", siCase.ResultError, "  but got: ", err)
		}
	}
}

func TestSignInByPhone(t *testing.T) {
	sis := model.SignIn{
		KeyUsername:   "email",
		Value:         "investor.spk@inbox.ru",
		Password:      "KeRXaTaq5Ce8ULO",
	}

	msg := sis.SignIn()
	if msg.IsThereAnError() {
		t.Error(msg.ErrMsg)
	}

	sis = model.SignIn{
		KeyUsername:   "email",
		Value:         "investor.spk@inbox.ru",
		Password:      "invalidpassword",
	}

	msg = sis.SignIn()
	if !msg.IsThereAnError() {
		t.Error("expected an error, but got nil")
	}
}


