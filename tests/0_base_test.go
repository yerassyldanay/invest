package tests

import (
	"github.com/stretchr/testify/require"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/utils/helper"
	"testing"
)

// this message is repeated again & again
func HelperExpectedNoErrorButGot(errmsg string) string {
	return "expected no error, but got " + errmsg
}

// >>> WARN
func HelperWarn(msg string) string {
	return ">>> WARN  " + msg
}

// get any project
func HelperGetAnyProject(t *testing.T) model.Project {
	project := model.Project{}
	db := model.GetDB()
	err := project.OnlyGetAny(db)
	require.NoError(t, err)
	return project
}

func HelperGenerateNewUser() model.User {
	return model.User{
		Password:       "P" + helper.Generate_Random_String(15),
		Fio:            helper.Generate_Random_String_No_Digits(10) + " " + helper.Generate_Random_String(10),
		Email:          model.Email{
			Address:  	helper.Generate_Random_String_No_Digits(15) + "_test@mail.ru",
		},
		Phone:          model.Phone{
			Ccode:    "+7",
			Number:   helper.Generate_Random_Number(10),
		},
		Organization:   model.Organization{
			Bin: "190940011748",
		},
	}
}

func HelperGenerateSMTP() model.SmtpServer {
	return model.SmtpServer{
		Host:     helper.Generate_Random_String(10) + ".test.com",
		Port:     876,
		Username: helper.Generate_Random_String(10) + "_test",
		Password: helper.Generate_Random_String(50),
		LastUsed: helper.GetCurrentTime(),
		Headers:  []model.SmtpHeaders{
			{
				Key:          "X-First-Header",
				Value:        "First-Header-Value",
			},
			{
				Key:          "X-Second-Header",
				Value:        "Second-Header-Value",
			},
		},
	}
}

// validate smtp
func HelperValidateSmtpCredentials(smtp model.SmtpServer, t *testing.T) {
	require.NotZero(t, smtp.Id)
	require.NotZero(t, smtp.Username)
	require.NotZero(t, smtp.Password)
	require.NotZero(t, smtp.Host)
	require.NotZero(t, smtp.Port)
}

// validate user
func HelperValidateUser(user model.User, t *testing.T) {
	require.NotZero(t, user.Id)
	require.NotZero(t, user.Fio)
	require.NotZero(t, user.Email.Address)
	require.NotZero(t, user.Phone.Number)
	require.NotZero(t, user.Role.Name)
	require.NotZero(t, user.Password)
}
