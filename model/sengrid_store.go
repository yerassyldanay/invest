package model

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"invest/utils"
	"os"
	"time"
)

/*
	this struct is needed to send messages using goroutines
*/
type sendMessage struct {
	FromName			string
	From				string
	ToName				string
	To					string
	Status				int
	SendgridMessage		*SendgridMessage
}

/*
	send message
*/
func (sm *SendgridMessageStore) Send_message() (map[string]interface{}, error) {
	from := mail.NewEmail(sm.FromName, sm.From)
	to := mail.NewEmail(sm.ToName, sm.To)

	message := mail.NewSingleEmail(from, sm.SendgridMessage.Subject, to, sm.SendgridMessage.PlainText, sm.SendgridMessage.HTML)
	sendgrid_api_key := os.Getenv("SENDGRID_API_KEY")
	client := sendgrid.NewSendClient(sendgrid_api_key)

	response, err := client.Send(message)

	var msg map[string]interface{}
	if err != nil {
		msg = map[string]interface{}{ "eng": "could not send a message" }
	} else {
		msg = map[string]interface{}{ "eng": string(response.Body) }
	}

	/*
		save the message on db
	*/
	//fmt.Println(sm)
	var trans = GetDB().Begin()

	/*
	* save the message itself
	* if ok then register on db
	 */
	if err := trans.Create(&sm.SendgridMessage).Error; err == nil {
		if err := trans.Create(sm).Error; err == nil {
			trans.Commit()
			return utils.NoErrorFineEverthingOk, err
		}
	}

	trans.Rollback()
	return msg, err
}

//func (sm *SendgridMessageStore) Set_message_fields_to_this_object(c *User) {
//	/*
//		send message to a receiver on behalf of organization (spk)
//	*/
//	var subject, html, page string
//	switch c.Lang {
//	case "kaz":
//		subject = templates.Base_email_subject_kaz
//		html = templates.Base_email_html_kaz
//		page = templates.Base_email_page_kaz
//	case "rus":
//		subject = templates.Base_email_subject_rus
//		html = templates.Base_email_html_rus
//		page = templates.Base_email_page_rus
//	default:
//		subject = templates.Base_email_subject_eng
//		html = templates.Base_email_html_eng
//		page = templates.Base_email_page_rus
//	}
//
//	sm = &SendgridMessageStore{
//		From:              utils.BaseEmailAddress,
//		To:                c.Email.Address,
//		FromName:          utils.BaseEmailName,
//		ToName:            c.Username,
//		SendgridMessage:   SendgridMessage{
//			Subject:   		subject,
//			PlainText: 		page,
//			HTML:      		html,
//			Date:      		time.Now(),
//		},
//		Status: 200,
//	}
//}

/*
	this method sends email to the provided list of email addresses
		this method has been created to optimize the process of sending messages
		using goroutines
*/
func (sm *SendgridMessageStore) SendMessageToList() (map[string]interface{}, error) {

	sm.SendgridMessage.Date = time.Now()

	if err := GetDB().Table(SendgridMessage{}.TableName()).Create(&sm.SendgridMessage).Error; err != nil {
		return utils.ErrorCouldNotSendEmail, err
	}
	sm.SendgridMessageId = sm.SendgridMessage.Id

	var msgchan = make(chan sendMessage, len(sm.ToList) + 1)
	var reschan = make(chan sendMessage, len(sm.ToList) + 1)

	fmt.Println()

	for i := 0; i < len(sm.ToList); i++ {
		go Send_message_to_function(i + 1, msgchan, reschan)
	}

	for i := 0; i < len(sm.ToList); i++ {
		msgchan <- sendMessage{
			FromName:        utils.BaseEmailName,
			From:            utils.BaseEmailAddress,
			ToName:          sm.ToListName[i],
			To:              sm.ToList[i],
			SendgridMessage: &sm.SendgridMessage,
		}
	}
	close(msgchan)

	var sendm sendMessage
	for i := 0; i < len(sm.ToList); i++ {
		select {
		case sendm = <- reschan:
		case <- time.Tick(time.Second * 3):
			break
		}
		var temp = SendgridMessageStore{
			From:              		sendm.From,
			FromName:          		sendm.FromName,
			To:                		sendm.To,
			ToName:            		sendm.ToName,
			SendgridMessageId: 		sm.SendgridMessageId,
			Status:            		sm.Status,
		}

		err := GetDB().Create(&temp).Error
		fmt.Println("sent message to ", sendm.To, err)
	}
	close(reschan)

	return utils.NoErrorFineEverthingOk, nil
}

/*
	this function has been created based on the pttern called
		"worker pool"
*/
func Send_message_to_function(i int, msgs <- chan sendMessage, result chan <- sendMessage) {
	for sm := range msgs {
		from := mail.NewEmail(sm.FromName, sm.From)
		to := mail.NewEmail(sm.ToName, sm.To)

		message := mail.NewSingleEmail(from, sm.SendgridMessage.Subject, to, sm.SendgridMessage.PlainText, sm.SendgridMessage.HTML)
		sendgrid_api_key := os.Getenv("SENDGRID_API_KEY")
		client := sendgrid.NewSendClient(sendgrid_api_key)

		response, err := client.Send(message)

		fmt.Println("message: ", sm.To, response, err)

		if err != nil {
			sm.Status = 400
		} else {
			sm.Status = response.StatusCode
		}

		result <- sm
	}

	fmt.Println("> goroutine #", i, "done")
}
