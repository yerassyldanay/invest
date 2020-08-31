package model

type Phone struct {
	Id				uint64				`json:"id" gorm:"AUTO_INCREMENT;primary key"`

	Ccode			string				`json:"ccode"`
	Number			string				`json:"number"`

	SentCode		string				`json:"sent_code"`
	Verified		bool				`json:"verified"`
}

func (Phone) TableName() string {
	return "phones"
}

