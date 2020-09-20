package model

import (
	"gopkg.in/validator.v2"
)

/*
	a message structure, which will be used to send a email message
 */
type SendgridMessage struct {
	Id					uint64			`json:"id" gorm:"AUTO_INCREMENT"`

	Subject				string			`json:"subject" validate:"required"`
	PlainText 			string			`json:"plain_text" validate:"required"`
	HTML      			string			`json:"html" validate:"required"`
	//Created				time.Time		`json:"date" gorm:"default: now()"`
}

func (SendgridMessage) TableName() string {
	return "sendgrid_messages"
}

func (sm SendgridMessage) Is_valid() bool {
	if err := validator.Validate(sm); err != nil {
		return false
	}
	return true
}

