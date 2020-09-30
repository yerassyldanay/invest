package model

import (
	"github.com/jinzhu/gorm"
	"invest/utils"
)

func (e *Email) Confirm(key string) (utils.Msg) {
	var u = User{}
	var trans = GetDB().Begin()

	defer func() {
		if trans != nil {trans.Rollback()}
	}()

	/*
		check & make sure that sql injection is not being made
	 */
	if ok := utils.Is_it_free_from_sql_injection(key); !ok {
		return ReturnInvalidParameters("sql injection pass. user confirmation")
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
				return ReturnEmailAlreadyInUse(err.Error())
	} else if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	//fmt.Println(u, users)

	if err := trans.Table(Email{}.TableName()).Where("id=?", u.EmailId).Updates(map[string]interface{}{
		"sent_code": 	"",
		"sent_hash": 	"",
		"verified": 	true,
		"deadline": 	nil,
	}).Error; err != nil {
		return ReturnInternalDbError(err.Error())
	}

	if err := trans.Table(User{}.TableName()).Where("id=?", u.Id).Update("verified", true).Error; err != nil{
		return ReturnInternalDbError(err.Error())
	}

	trans.Commit()
	trans = nil

	return ReturnNoError()
}
