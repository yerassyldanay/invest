package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"invest/utils"
)

func (e *Email) Confirm(key string) (map[string]interface{}, error) {
	var u = User{}
	var trans = GetDB().Begin()

	defer func() {
		if trans != nil {trans.Rollback()}
	}()

	/*
		check & make sure that sql injection is not being made
	 */
	if ok := utils.Is_it_free_from_sql_injection(key); !ok {
		return utils.ErrorInvalidParameters, errors.New("sql injection pass. user confirmation")
	}

	var query string
	var param string

	switch key {
	case "shash":
		query = " emails.sent_hash = ? "
		param = e.SentHash
	default:
		query = " emails.sent_code = "
		param = e.SentCode
	}

	if err := trans.Model(&User{}).Joins(" join emails on users.email_id = emails.id ").
		Where(query, param).Limit("1").Scan(&u).Error; err == gorm.ErrRecordNotFound {
				return utils.ErrorEmailIsAreadyInUse, err
	} else if err != nil {
		return utils.ErrorInternalDbError, err
	}

	//fmt.Println(u, users)

	if err := trans.Table(Email{}.TableName()).Where("id=?", u.EmailId).Updates(map[string]interface{}{
		"sent_code": 	"",
		"sent_hash": 	"",
		"verified": 	true,
		"deadline": 	nil,
	}).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}

	if err := trans.Table(User{}.TableName()).Where("id=?", u.Id).Update("verified", true).Error; err != nil{
		return utils.ErrorInternalDbError, err
	}

	trans.Commit()
	trans = nil

	return utils.NoErrorFineEverthingOk, nil
}
