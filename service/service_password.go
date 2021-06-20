package service

import (
	"fmt"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/utils/helper"
	"github.com/yerassyldanay/invest/utils/message"
)

func (is *InvestService) PasswordResetSendMessage(fp model.ForgetPassword) message.Msg {
	// get email
	var email = model.Email{}
	if err := model.GetDB().First(&email, "address = ?", fp.EmailAddress).Error; err != nil {
		return model.ReturnInvalidParameters(fmt.Sprintf("failed to get email. err: %v", err))
	}

	// validate email
	if err := fp.Validate(); err != nil {
		return model.ReturnInvalidParameters(fmt.Sprintf("failed to validate. err: %v", err))
	}

	// create a hash that will be sent to email address
	var codeToSend = helper.Generate_Random_Number(6)

	// store on redis
	cmdStatus := model.GetRedis().Set("password_reset"+codeToSend, fp.EmailAddress, 0)
	if cmdStatus.Err() != nil {
		return model.ReturnInternalDbError(fmt.Sprintf("failed to fetch data from redis. err: %v", cmdStatus.Err()))
	}

	// send notification
	nc := model.NotifyCode{
		Code:    fp.Code,
		Address: fp.EmailAddress,
	}

	// this handle all further operation
	select {
	case model.GetMailerQueue().NotificationChannel <- &nc:
	default:
	}

	return model.ReturnNoErrorWithResponseMessage(map[string]interface{}{
		"code":    codeToSend,
		"address": fp.EmailAddress,
	})
}

// PasswordResetChangePassword
func (is *InvestService) PasswordResetChangePassword(fp model.ForgetPassword) message.Msg {
	// get from redis
	emailAddress, err := model.GetRedis().Get("password_reset" + fp.Code).Result()
	if err != nil {
		return model.ReturnInvalidParameters(err.Error())
	}

	// get user by email address
	var email = model.Email{}
	if err := model.GetDB().First(&email, "address = ?", emailAddress).Error; err != nil {
		return model.ReturnInvalidParameters(err.Error())
	}

	// update password
	if err := model.GetDB().Model(&model.User{}).
		Where("email_id = ?", email.Id).
		Updates(map[string]interface{}{
			"password": fp.NewPassword,
		}).Error; err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// ok
	return model.ReturnNoError()
}
