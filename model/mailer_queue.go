package model

import (
	"context"
	"fmt"
)

/*
	MailerQueue:
		this is to handle all notifications in one place
 */
type MailerQueue struct {
	NotificationChannel				chan InterMessage			`json:"notification_channel"`
}

// to initiate mail queue
var mq *MailerQueue
func InitiateNewMailerQueue() (*MailerQueue) {
	mq = &MailerQueue{}
	mq.NotificationChannel = make(chan InterMessage, 100)
	return mq
}

// getter for a mailer queue
func GetMailerQueue() (*MailerQueue) {
	if mq == nil {
		fmt.Println("create a new mailer queue...")
		InitiateNewMailerQueue()
	}
	return mq
}

/*
	MailerQueue this function handles all notifications
		* stores notification on database
		* creates a message (with a list of receivers)
		* sends notification
 */
func (mqi *MailerQueue) Handle(ctx context.Context) {
	for {
		select {
		case <- ctx.Done():
			fmt.Println("stop handling messages...")
			return
		case notifyMessage := <- mqi.NotificationChannel:

			//fmt.Println("Sending one more message: ", notifyMessage.GetSubject())

			if len(notifyMessage.GetToList()) < 1 {
				continue
			}

			// prepare notification message
			// store notification body on db
			// store notification instances on db
			err := MessageStoreNotificationOnDb(notifyMessage)
			if err != nil {
				mqi.printCouldNotSendNotification(err)
				return
			}

			// prepare message (message for smtp server)
			message, err := MessageOnlyPrepareMail(notifyMessage)
			if err != nil {
				mqi.printCouldNotSendNotification(err)
				return
			}

			// this function gets smtp server credentials from db
			// sets connection & sends message
			if err = MessageDialAndSend(message); err != nil {
				mqi.printCouldNotSendNotification(err)
			}
		}
	}
}

func (mqi *MailerQueue) printCouldNotSendNotification(err error) {
	fmt.Println("Could not send a notification. Error: ", err)
}
