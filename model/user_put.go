package model

import "invest/utils"

func (c *User) Update_user_profile_including_password(user *User) (utils.Msg) {
	// load user by id
	if err := c.OnlyGetByIdPreloaded(GetDB()); err != nil {
		return ReturnInternalDbError(err.Error())
	}

	// validate new info
	if err := user.ValidateForUpdate(); err != nil {
		return ReturnInvalidParameters(err.Error())
	}

	// set fields
	c.Email.Address = user.Email.Address
	c.Phone.Number = user.Phone.Number
	c.Phone.Ccode = user.Phone.Ccode
	c.Fio = user.Fio

	ok := Validate_password(user.Password, "", "")
	new_password, err := utils.Convert_string_to_hash(user.Password)
	if err == nil || ok {
		c.Password = new_password
	}

	var trans = GetDB().Begin()
	defer func() { if trans != nil {trans.Rollback()} }()

	// save email & phone
	err = c.Email.OnlySave(trans)
	err2 := c.Phone.OnlySave(trans)

	if err != nil || err2 != nil {
		return ReturnInternalDbError("could not update email or phone")
	}

	// save the profile of a user
	if err := c.OnlySave(trans); err != nil {
		return ReturnInternalDbError(err.Error())
	}

	if err := trans.Commit().Error; err != nil {
		return ReturnInternalDbError(err.Error())
	}

	trans = nil
	return ReturnNoError()
}

