package model

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"invest/utils"
	"os"
	"sync"
	"time"
)

/*
	this method sends email to the provided list of email addresses
		this method has been created to optimize the process of sending messages
		using goroutines
*/
func (sm *SendgridMessageStore) SendMessageToList() (map[string]interface{}, error) {
	var wg = sync.WaitGroup{}

	/*
		create sendgrid message on db
			this will save the body of the message on db
		after this step, based on id of the sendgrid message
	 */
	sm.SendgridMessage.Created = utils.GetCurrentTime()
	_ = sm.SendgridMessage.Create_on_db()
	sm.SendgridMessageId = sm.SendgridMessage.Id

	/*
		create channels
	 */
	var msgchan = make(chan *sendMessage, len(sm.ToAddresser) + 1)
	var reschan = make(chan *sendMessage, len(sm.ToAddresser) + 1)

	/*
		run goroutines & pass channels
	*/
	for i := 0; i < len(sm.ToAddresser); i++ {
		go Send_message_to_function(i + 1, msgchan, reschan)
	}

	/*
		pass messages through channels
	*/
	for i := 0; i < len(sm.ToAddresser); i++ {
		msgchan <- &sendMessage{
			FromName:        utils.BaseEmailName,
			From:            utils.BaseEmailAddress,
			ToName:          sm.ToAddresser[i].Name,
			To:              sm.ToAddresser[i].Address,
			SendgridMessage: &sm.SendgridMessage,
		}
	}

	/*
		receive messages from channels
	*/
	var sendm *sendMessage
	for i := 0; i < len(sm.ToAddresser); i++ {
		select {
		case sendm = <- reschan:
		case <- time.Tick(time.Second * 3):
			break
		}

		if sendm == nil {
			continue
		}

		var temp = SendgridMessageStore{
			From:              		sendm.From,
			FromName:          		sendm.FromName,
			To:                		sendm.To,
			ToName:            		sendm.ToName,
			SendgridMessageId: 		sm.SendgridMessageId,
			Status:            		sendm.Status,
		}

		/*
			save messages
		 */
		wg.Add(1)
		Store_sendgrid_message_on_db(&wg, &temp)
	}

	select {
	case <- time.Tick(time.Second * 3):
		break
	default:
		wg.Wait()
	}

	close(msgchan)
	close(reschan)

	return utils.NoErrorFineEverthingOk, nil
}

/*
	this function has been created based on the pttern called
		"worker pool"
*/
func Send_message_to_function(i int, msgchan chan *sendMessage, reschan chan *sendMessage) {
	for sm := range msgchan {
		go Send_message_helper(sm, reschan)
	}

	//fmt.Println("> goroutine #", i, "done")
}

/*
	get sendgrid message & send it using external service
*/
func Send_message_helper(sm *sendMessage, result chan *sendMessage) {
	from := mail.NewEmail(sm.FromName, sm.From)
	to := mail.NewEmail(sm.ToName, sm.To)

	message := mail.NewSingleEmail(from, sm.SendgridMessage.Subject, to, sm.SendgridMessage.PlainText, sm.SendgridMessage.HTML)
	sendgrid_api_key := os.Getenv("SENDGRID_API_KEY")
	client := sendgrid.NewSendClient(sendgrid_api_key)

	response, err := client.Send(message)

	if err != nil {
		sm.Status = 400
	} else {
		sm.Status = response.StatusCode
	}

	result <- sm
}

/*
	create sendgrid messages
*/
func Store_sendgrid_message_on_db(wg *sync.WaitGroup, sm *SendgridMessageStore) {
	defer wg.Done()
	err := GetDB().Create(&sm).Error
	if err != nil {
		fmt.Println("sent message to ", sm, err)
	}
}
