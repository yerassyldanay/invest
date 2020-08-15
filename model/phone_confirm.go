package model

import "invest/utils"

/*
	confirm that you own phone number
 */
func (p *Phone) Confirm() (map[string]interface{}, error) {
	if err := GetDB().Table(Phone{}.TableName()).Where("number=? and sent_code=?", p.Number, p.SentCode).
		Update("verified", true).Error;
		err != nil {
			return utils.ErrorInvalidParameters, err
	}

	return utils.NoErrorFineEverthingOk, nil
}