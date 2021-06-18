package service

import (
	"encoding/json"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/utils/constants"
	"github.com/yerassyldanay/invest/utils/helper"
	"github.com/yerassyldanay/invest/utils/message"
	"time"
)

/*
	signup for investors
		201 - created
		400 - bad request
		409 - already in use
		417 - database error
		422 - could not sent message & not stored on database
*/
func (is *InvestService) SignUp(c model.User) (message.Msg) {

	//// validate
	//if err := c.Validate(); err != nil {
	//	return model.ReturnInvalidParameters(err.Error())
	//}

	//convert password to hash
	//hashed, err := helper.Convert_string_to_hash(c.Password)
	//if err != nil {
	//	return model.ReturnInternalDbError(err.Error())
	//}
	//c.Password = hashed

	// these code and link will be sent to the user
	randomCode := helper.Generate_Random_Number(constants.MaxNumberOfDigitsSentByEmail)

	// marshal to store data
	userInBytes, err := json.Marshal(c)
	if err != nil {
		return model.ReturnFailedToCreateAnAccount(err.Error())
	}

	// store it in redis
	cmd := model.GetRedis().Set(randomCode, string(userInBytes), time.Second * 120)
	if cmd.Err() != nil {
		return model.ReturnFailedToCreateAnAccount(cmd.Err().Error())
	}

	// send notification
	nc := model.NotifyCode{
		Code:    randomCode,
		Address: c.Email.Address,
	}

	select {
	case model.GetMailerQueue().NotificationChannel <- &nc:
	default:
	}

	return model.ReturnNoErrorWithResponseMessage(map[string]interface{}{
		"code": randomCode,
	})
}
