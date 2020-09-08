package model

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"invest/utils"
	"os"
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

