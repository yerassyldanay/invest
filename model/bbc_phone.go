package model

import "invest/utils"

type Phone struct {
	Id				uint64				`json:"id" gorm:"AUTO_INCREMENT;primary key"`

	Ccode			string				`json:"ccode" validate:"required"`
	Number			string				`json:"number" validate:"required"`

	SentCode		string				`json:"sent_code"`
	Verified		bool				`json:"verified"`
}

func (Phone) TableName() string {
	return "phones"
}

func (p *Phone) Is_verified() (map[string]interface{}, error) {
	var err error
	if err = GetDB().Model(&Phone{}).Where("ccode=? and number=?", p.Ccode, p.Number).First(p).Error; err == nil {
		if p.Verified {
			return nil, nil
		}
	}

	return utils.ErrorPhoneNumberIsNotVerified, err
}

