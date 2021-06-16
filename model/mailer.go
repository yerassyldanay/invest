package model

import (
	"errors"
	"fmt"
	"gopkg.in/gomail.v2"
	"invest/utils/helper"
)

// all notifications messages must have following methods,
// which prepare message to be sent
type InterMessage interface {
	// this indicates the sender address
	GetFrom() string
	// indicates email addresses of receivers
	GetToList() []string
	// the subject of the email message
	GetSubject() string
	// body of the message in html format
	GetHtml() string
	// body of the message in plain text format
	GetPlainText() string
	// get project id
	GetProjectId() uint64
}

type InterDialer interface {
	DialAndSend(m... *gomail.Message) error
}

// errors
var errorMessageInvalidFrom = errors.New("invalid sender address")
var errorMessageInvalidToList = errors.New("invalid receiver address")
var errorMessageInvalidSubject = errors.New("invalid subject")
var errorMessageInvalidBody = errors.New("invalid body of a message")

// validate
func MessageOnlyValidate(message InterMessage) error {
	switch {
	case message.GetFrom() == "":
		return errorMessageInvalidFrom
	case len(message.GetToList()) == 0:
		return errorMessageInvalidToList
	case message.GetSubject() == "":
		return errorMessageInvalidSubject
	case len(message.GetHtml()) + len(message.GetPlainText()) == 0:
		return errorMessageInvalidBody
	}

	return nil
}

// prepare message to send
func MessageOnlyPrepareMail(message InterMessage) (*gomail.Message, error) {
	// validate message
	if err := MessageOnlyValidate(message); err != nil {
		return &gomail.Message{}, err
	}

	// prepare mail message to send
	msgToSend := gomail.NewMessage()

	// set headers here
	msgToSend.SetHeader("From", message.GetFrom())
	msgToSend.SetHeader("To", message.GetToList()...)
	msgToSend.SetHeader("Subject", message.GetSubject())

	var temp = message.GetHtml()
	_ = temp

	// set body
	if len(message.GetHtml()) > 0 {
		msgToSend.SetBody("text/html", message.GetHtml())
	}

	if len(message.GetPlainText()) > 0 {
		msgToSend.SetBody("text/plain", message.GetPlainText())
	}
	//else {
	//	return msgToSend, errors.New("content body is not specified")
	//}

	return msgToSend, nil
}

// send message
func MessageOnlySend(dialer InterDialer, message *gomail.Message) error {
	startTime := helper.GetCurrentTime()
	err := dialer.DialAndSend(message)
	fmt.Println("Dial & Send Message Took: ", helper.GetCurrentTime().Sub(startTime).Seconds(), " s")
	return err
}

// get connection & send message
func MessageDialAndSend(message *gomail.Message) error {
	// get smtp credential from database
	var smtpServer = SmtpServer{}
	if err := smtpServer.OnlyGetOne(GetDB()); err != nil {
		return err
	}

	// change last_used time
	// this helps to send messages from different smtp servers
	smtpServer.LastUsed = helper.GetCurrentTime()
	if err := smtpServer.OnlySaveById(GetDB()); err != nil {
		return err
	}

	// create dialer (establish connection)
	var d = &gomail.Dialer{
		Host:      smtpServer.Host,
		Port:      smtpServer.Port,
		Username:  smtpServer.Username,
		Password:  smtpServer.Password,
	}

	err := MessageOnlySend(d, message)
	return err
}

func MessageStoreNotificationOnDb(n InterMessage) error {
	// create notification
	notification := Notification{
		FromAddress: n.GetFrom(),
		ProjectId:   n.GetProjectId(),
		Html:        n.GetHtml(),
		Plain:       n.GetPlainText(),
		Created:     helper.GetCurrentTime(),
	}

	// store notification body on database
	if err := notification.OnlyCreate(GetDB()); err != nil {
		return err
	}

	// get a list of receivers
	emails := n.GetToList()

	//var wg = sync.WaitGroup{}

	for _, email := range emails {
		// not for fun
		email := email

		notInstance := &NotificationInstance{
			ToAddress:      email,
			NotificationId: notification.Id,
		}

		//wg.Add(1)
		func(email string, ni *NotificationInstance) {
			//defer wg.Done()

			// get notification instance
			if err := ni.OnlyCreate(GetDB()); err != nil {
				fmt.Println("could not store notification on database. err: ", err)
			}

		}(email, notInstance)
	}

	// wait for goroutines
	//wg.Wait()

	return nil
}
