package model

type SendgridMessageStore struct {
	From				string			`json:"from" gorm:"size: 155" validate:"required, email"`
	FromName			string			`json:"from_name" validate:"required"`

	To					string			`json:"to" gorm:"size: 155" validate:"required, email"`
	ToName				string			`json:"to_name" validate:"required"`

	ToAddresser			[]EmailAddresser			`json:"to_addresser"`

	SendgridMessageId		uint64					`json:"sendgrid_message_id"`
	SendgridMessage			SendgridMessage			`json:"sendgrid_message" gorm:"foreignkey:SendgridMessageId"`

	Status					int						`json:"status" gorm:"default:200"`
	Opened					bool					`json:"opened" gorm:"default:false"`
}

type EmailAddresser struct {
	Name			string				`json:"name"`
	Address			string				`json:"address"`
}

func (SendgridMessageStore) TableName() string {
	return "sendgrid_message_stores"
}


