package model

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/validator.v2"
	"invest/templates"
	"invest/utils"

	"time"
)

type ForgetPassword struct {
	NewPassword			string				`json:"new_password"`
	EmailAddress				string				`json:"email_address" validate:"email"`

	Code				string				`json:"code"`
	Deadline			time.Time			`json:"deadline"`
	
	Lang 				string				`json:"lang"`
	UserId				uint64				`json:"user_id"`
}

func (fp *ForgetPassword) SendMessage() (utils.Msg) {
	if err := validator.Validate(*fp); err != nil {
		return utils.Msg{utils.ErrorInvalidParameters, 400, "", err.Error()}
	}

	var codeChan = make(chan string, 1)
	go func(codeCh chan<- string) {
		codeCh <- utils.Generate_Random_String(utils.MaxNumberOfCharactersSentByEmail)
	}(codeChan)

	var user = User{}
	if err := GetDB().First(&user.Email, "address = ?", fp.EmailAddress).Error;
		err == gorm.ErrRecordNotFound {
		return utils.Msg{utils.ErrorInvalidParameters, 400, "", err.Error()}
	} else if err != nil || !user.Email.Verified {
		return utils.Msg{utils.ErrorInvalidParameters, 400, "", "email is not confirmed or internal data base error occurred"}
	}

	var email = user.Email
	err := GetDB().First(&user, "email_id = ?", user.Email.Id).Error
	if err != nil {
		return utils.Msg{utils.ErrorInvalidParameters, 400, "", err.Error()}
	}

	user.Email = email

	/*
		create a hash that will be sent to email address
	 */
	var code string
	select {
	case code = <- codeChan:
		fp.Code = code
	case  <- time.Tick(time.Second * 10):
		return utils.Msg{utils.ErrorInternalServerError, 500, "", "timeout. ForgetPassword Link"}
	}

	/*
		this will create a message to send
	 */
	var sms = SendgridMessageStore{}
	sms, err = sms.Prepare_message_this_object(&user, templates.Base_message_map_2_forget_password)
	if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	/*
		store this code in redis
			{
				"somehash": 1, (this is valid for one day)
			}
	 */
	GetRedis().Set(fp.Code, fp.UserId, time.Hour * 24)

	return utils.Msg{utils.NoErrorFineEverthingOk, 200, "", ""}
}

func (fp *ForgetPassword) Change_password_of_user_by_hash() (utils.Msg) {
	// check hash is in redis
	user_id, err := GetRedis().Get(fp.Code).Uint64()
	if err != nil || user_id != fp.UserId || user_id == 0 {
		return ReturnInvalidParameters("invalid code has been provided")
	}

	if err = Validate_password(fp.NewPassword); err != nil {
		return ReturnInternalDbError(err.Error())
	}

	hashed_password, err := utils.Convert_string_to_hash(fp.NewPassword)
	if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	if err := GetDB().Table(User{}.TableName()).Where("id = ?", fp.UserId).Update("password", hashed_password).Error;
		err != nil {
			return ReturnInternalDbError(err.Error())
	}

	return ReturnNoError()
}



