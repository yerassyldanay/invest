package model

import (
	"invest/utils"
)

/*
	prepare sendgrid message store object
 */
func (sm *SendgridMessageStore) Prepare_message_this_object(e *Email, lang string, message_map map[string]map[string]string) (SendgridMessageStore, error) {
	var newsm = SendgridMessageStore{
		From:              utils.BaseEmailAddress,
		FromName:          utils.BaseEmailName,
		To:                e.Address,
		ToName:            "Қолданушы. Пользователь. User",
		SendgridMessage:   SendgridMessage{
			Subject:   	message_map[utils.KeyEmailSubject][lang],
			PlainText: 	message_map[utils.KeyEmailPlainText][lang],
			HTML:      	message_map[utils.KeyEmailHtml][lang],
		},
		Created:      		utils.GetCurrentTime(),
	}

	return newsm, nil
}

/*
	prepare a sendgrid message
 */
func (sm *SendgridMessageStore) Set_default_values() {
	sm.From = utils.BaseEmailAddress
	sm.FromName = utils.BaseEmailName
}


