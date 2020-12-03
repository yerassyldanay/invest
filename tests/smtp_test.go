package tests

import (
	"github.com/stretchr/testify/require"
	"invest/model"
	"invest/service"
	"invest/utils/helper"
	"testing"
)

func TestServiceSmtpCreate(t *testing.T) {
	// generate one
	smtp := HelperGenerateSMTP()

	// headers
	is := service.InvestService{}
	msg := is.SmtpCreate(&smtp)

	// check
	require.Zero(t, msg.ErrMsg)
}

func TestServiceSmtpUpdate(t *testing.T) {
	// we need id to update
	smtp := model.SmtpServer{}

	// get any smtp
	err := model.GetDB().First(&smtp).Error

	// check
	require.NoError(t, err)

	newHeaders := []model.SmtpHeaders{
		{
			Key:          "X-First-Header",
			Value:        "First-Header-Value-Updated",
		},
		{
			Key:          "X-Second-Header",
			Value:        "Second-Header-Value-Updated",
		},
	}

	newHost := helper.Generate_Random_String(20) + ".updated.host"
	newUsername := helper.Generate_Random_String(20) + ".updated.username"

	// headers
	newSmtp := model.SmtpServer{
		Id:       smtp.Id,
		Host:     newHost,
		Port:     smtp.Port,
		Username: newUsername,
		Password: smtp.Password,
		Headers:  newHeaders,
	}
	is := service.InvestService{}
	msg := is.SmtpUpdate(&newSmtp)

	// check
	require.Zero(t, msg.ErrMsg)

	// check that updated
	newSmtp.Id = smtp.Id
	err = newSmtp.OnlyGetById(model.GetDB())

	// check
	//fmt.Println(newSmtp)
	require.NoError(t, err)
	require.Equal(t, newSmtp.Id, smtp.Id)
	require.Equal(t, newSmtp.Password, smtp.Password)
	require.Equal(t, newSmtp.Username, newUsername)
	require.Equal(t, newSmtp.Host, newHost)
}

func TestModelSmtpGet(t *testing.T) {

	is := service.InvestService{}
	msg := is.SmtpGet()

	// check
	require.Zero(t, msg.ErrMsg)

	// create one & get list of smtp servers
	TestServiceSmtpCreate(t)
	smtp := model.SmtpServer{}
	smtps, err := smtp.OnlyGetAll(model.GetDB())

	// check
	require.NoError(t, err)
	require.Condition(t, func() (bool) { return len(smtps) > 0 })

	for _, smtp := range smtps {
		HelperValidateSmtpCredentials(smtp, t)
	}
}

func TestServiceSmtpDelete(t *testing.T) {

	// create one initially
	TestServiceSmtpCreate(t)

	// get test
	smtp := model.SmtpServer{}
	err := model.GetDB().First(&smtp).Error

	// check
	require.NoError(t, err)
	HelperValidateSmtpCredentials(smtp, t)

	// delete
	saveSmtp := model.SmtpServer{Id: smtp.Id}
	is := service.InvestService{}
	msg := is.SmtpDelete(&smtp)

	// check
	require.Zero(t, msg.ErrMsg)

	// make sure
	err = saveSmtp.OnlyGetById(model.GetDB())

	// check
	require.Error(t, err)
}

func TestLogicSmtpRotate(t *testing.T) {
	// get smtp server
	firstSmtp := model.SmtpServer{}
	err := firstSmtp.OnlyGetOne(model.GetDB())

	// check
	require.NoError(t, err)

	// change time
	firstSmtp.LastUsed = helper.GetCurrentTime()
	err = firstSmtp.OnlySaveById(model.GetDB())

	// check
	require.NoError(t, err)

	// -- second --

	// check is it rotated
	secondSmtp := model.SmtpServer{}
	err = secondSmtp.OnlyGetOne(model.GetDB())

	// check
	require.NoError(t, err)
	require.NotEqual(t, firstSmtp.Id, secondSmtp.Id)
}


