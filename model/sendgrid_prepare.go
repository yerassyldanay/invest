package model

import (
	"invest/utils"
	"strings"
	"time"
)

/*
	prepare sendgrid message store object
 */
func (sm *SendgridMessageStore) Prepare_message_this_object(c *User, message_map map[string]map[string]string) (SendgridMessageStore, error) {
	var lang = strings.ToLower(c.Lang)
	var newsm = SendgridMessageStore{
		From:              utils.BaseEmailAddress,
		FromName:          utils.BaseEmailName,
		To:                c.Email.Address,
		ToName:            c.Fio,
		SendgridMessageId: 0,
		SendgridMessage:   SendgridMessage{
			Subject:   	message_map[utils.KeyEmailSubject][lang],
			PlainText: 	message_map[utils.KeyEmailPlainText][lang],
			HTML:      	message_map[utils.KeyEmailHtml][lang],
			Date:      	time.Now(),
		},
	}

	return newsm, nil
}
