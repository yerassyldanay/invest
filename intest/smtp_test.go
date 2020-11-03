package intest

import (
	"fmt"
	"invest/model"
	"invest/service"
	"invest/utils/helper"
	"testing"
)

func TestServiceSmtpCreate(t *testing.T) {
	smtp := model.SmtpServer{
		Host:     "service.test.com",
		Port:     876,
		Username: "test.user",
		Password: "YA6654qwd",
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

	// headers
	is := service.InvestService{}
	msg := is.SmtpCreate(&smtp)

	// check
	if msg.IsThereAnError() {
		t.Error(msg.ErrMsg)
	}
}

func TestServiceSmtpUpdate(t *testing.T) {
	// we need id to update
	smtp := model.SmtpServer{}

	// get test
	err := model.GetDB().First(&smtp, "host like '%.test.%'").Error
	if err != nil {
		t.Error(err)
	}

	smtp.Headers = []model.SmtpHeaders{
		{
			Key:          "X-First-Header",
			Value:        "First-Header-Value-Updated",
		},
		{
			Key:          "X-Second-Header",
			Value:        "Second-Header-Value-Updated",
		},
	}

	smtp.Host = "service.test.com.updated"
	smtp.Username = "test.user.updated"

	// headers
	is := service.InvestService{}
	msg := is.SmtpUpdate(&smtp)

	// check
	if msg.IsThereAnError() {
		t.Error(msg.ErrMsg)
	}
}

func TestServiceSmtpGet(t *testing.T) {

	is := service.InvestService{}
	msg := is.SmtpGet()

	if msg.IsThereAnError() {
		t.Error(msg.ErrMsg)
	}

	smtps, ok := msg.Message["info"]
	if !ok {
		t.Error("could not find info inside the map")
	}

	if smtps == nil {
		t.Error("there is no any smtp credential")
	}
}

func TestModelSmtpGet(t *testing.T) {
	smtp := model.SmtpServer{}

	smtps, err := smtp.OnlyGetAll(model.GetDB())
	if err != nil {
		t.Error(err)
	}

	if len(smtps) <= 0 {
		t.Error("could not find any smtp credential")
	} else {
		fmt.Println("smtps found: ", len(smtps))
		for _, smtp := range smtps {
			fmt.Println(smtp)
		}
	}
}

func TestServiceSmtpDelete(t *testing.T) {
	smtp := model.SmtpServer{}

	// get test
	err := model.GetDB().First(&smtp, "host like '%.test.%'").Error
	if err != nil {
		t.Error(err)
	}

	// delete
	is := service.InvestService{}
	msg := is.SmtpDelete(&smtp)

	// check
	if msg.IsThereAnError() {
		t.Error(msg.ErrMsg)
	}
}

func TestLogicSmtpRotate(t *testing.T) {
	// get smtp server
	firstSmtp := model.SmtpServer{}
	if err := firstSmtp.OnlyGetOne(model.GetDB()); err != nil {
		t.Error(err)
	}

	// change time
	firstSmtp.LastUsed = helper.GetCurrentTime()
	if err := firstSmtp.OnlySaveById(model.GetDB()); err != nil {
		t.Error(err)
	}

	// check is it rotated
	secondSmtp := model.SmtpServer{}
	if err := secondSmtp.OnlyGetOne(model.GetDB()); err != nil {
		t.Error(err)
	}

	if firstSmtp.Host == secondSmtp.Host && firstSmtp.Username == secondSmtp.Username {
		t.Error("smtp credentials have not been rotated")
	}
}


