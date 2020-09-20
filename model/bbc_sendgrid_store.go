package model

import "time"

type SendgridMessageStore struct {
	Id					uint64			`json:"id" gorm:"AUTO_INCREMENT; primary_key"`
	From				string			`json:"from" gorm:"size: 155" validate:"required, email"`
	FromName			string			`json:"from_name" validate:"required"`

	To					string			`json:"to" gorm:"size: 155" validate:"required, email"`
	ToName				string			`json:"to_name" validate:"required"`

	ToAddresser			[]EmailAddresser			`json:"to_addresser"`

	SendgridMessageId		uint64					`json:"sendgrid_message_id"`
	SendgridMessage			SendgridMessage			`json:"sendgrid_message" gorm:"foreignkey:SendgridMessageId"`
	
	ProjectId				uint64					`json:"project_id" gorm:"foreignkey:projects.id"`
	
	Status					int						`json:"status" gorm:"default:200"`
	Opened					bool					`json:"opened" gorm:"default:false"`

	Created					time.Time				`json:"created_date" gorm:"now()"`
}

type EmailAddresser struct {
	Name			string				`json:"name"`
	Address			string				`json:"address"`
}

func (SendgridMessageStore) TableName() string {
	return "sendgrid_message_stores"
}


