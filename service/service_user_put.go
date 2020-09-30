package service

import (
	"invest/model"
	"invest/utils"
)

func (is *InvestService) Update_user_profile(user *model.User) (utils.Msg) {
	/*
		get the user account, which is being modified
	 */
	var trans = model.GetDB().Begin()
	defer func() { if trans != nil {trans.Rollback()} }()

	var tempuser = model.User{
		Id: user.Id,
	}

	// get user info
	if err := tempuser.OnlyGetByIdPreloaded(trans); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// if this is a new phone number
	if tempuser.Phone.Ccode + tempuser.Phone.Number != user.Phone.Ccode + user.Phone.Number {
		// delete phone number
		if err := tempuser.Phone.OnlyDeleteById(trans); err != nil {
			return model.ReturnInternalDbError(err.Error())
		}

		tempuser.Phone.Ccode = user.Phone.Ccode
		tempuser.Phone.Number = user.Phone.Number

		// validate
		if err := user.Phone.Validate(); err != nil {
			return model.ReturnInternalDbError(err.Error())
		}

		// create a new phone number
		if err := tempuser.Phone.OnlyCreate(trans); err != nil {
			return model.ReturnInternalDbError(err.Error())
		}

		tempuser.PhoneId = tempuser.Phone.Id
	}

	// save changes
	if err := tempuser.OnlySave(trans); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// commit changes
	if err := trans.Commit().Error; err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	trans = nil
	return model.ReturnNoError()
}

