package service

import (
	"github.com/jinzhu/gorm"
	"invest/model"
	"invest/utils"
	"time"
)

func (is *InvestService) Password_reset_send_message(fp model.ForgetPassword) (utils.Msg) {

	// get email
	var email = model.Email{Address: fp.EmailAddress}
	if err := email.OnlyGetByAddress(model.GetDB()); err != nil {
		return model.ReturnInvalidParameters(err.Error())
	}

	// validate email
	if err := fp.Validate(); err != nil {
		return model.ReturnInvalidParameters(err.Error())
	}

	// create a hash that will be sent to email address
	var hashCodeToSend = utils.Generate_Random_String(utils.MaxNumberOfCharactersSentByEmail)

	// check whether once message was sent
	err := fp.OnlyGet(model.GetDB())

	// set hash
	fp.Code = hashCodeToSend
	fp.Deadline = utils.GetCurrentTime().Add(time.Hour * 24)

	switch {
	case err == gorm.ErrRecordNotFound:
		// it is possible that an email has been sent to this email address
		// this indicates it is not
		// create
		if err = fp.OnlyCreate(model.GetDB()); err != nil {
			return model.ReturnInternalDbError(err.Error())
		}

	case err != nil:
		// an unknown / unexpected error has occurred
		return model.ReturnInternalDbError(err.Error())

	default:
		// there is one
		if err = fp.OnlyUpdateByEmailAddress(model.GetDB(), "code", "deadline"); err != nil {
			return model.ReturnInternalDbError(err.Error())
		}
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

	var resp = utils.NoErrorFineEverthingOk
	//resp["info"] = model.Struct_to_map(fp)

	return model.ReturnNoErrorWithResponseMessage(resp)
}

// change the actual password
// hash & password
func (is *InvestService) Password_reset_change_password(fp model.ForgetPassword) (utils.Msg) {

	var trans = model.GetDB().Begin()
	defer func() { if trans != nil {trans.Rollback()} }()

	// get email from forget password
	if err := fp.OnlyGetByCode(trans); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// get user by email
	var user = model.User{Email: model.Email{Address: fp.EmailAddress}}
	if err := user.OnlyGetByEmailAddress(trans); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// validate password
	if err := model.OnlyValidatePassword(fp.NewPassword); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// convert password to hash
	hashedPassword, err := utils.Convert_string_to_hash(fp.NewPassword)
	if err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// update password
	if err := user.OnlyUpdatePasswordById(hashedPassword, model.GetDB()); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// delete - otherwise it can be used to reset password
	if err = fp.OnlyDelete(trans); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// commit changes
	if err := trans.Commit().Error; err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	return model.ReturnNoError()
}