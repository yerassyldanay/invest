package model

import "C"
import (
	"invest/utils/constants"
	"invest/utils/helper"
	"invest/utils/message"
	"strings"
	"time"

	//"gorm.io/gorm/clause"
)

const GetLimit = 20

func (c *User) Remove_all_users_with_not_confirmed_email() (map[string]interface{}, error) {
	return nil, nil
}

/*
	create
*/
func (c *User) Create_user_without_check() (message.Msg) {
	if c.Lang == "" {
		c.Lang = constants.DefaultContentLanguage
	}

	/*
		update seq ids
			refer to function description for more information
	 */
	//_ = Update_sequence_id_thus_avoid_duplicate_primary_key_error(GetDB(), "users")
	//_ = Update_sequence_id_thus_avoid_duplicate_primary_key_error(GetDB(), "phones")
	//_ = Update_sequence_id_thus_avoid_duplicate_primary_key_error(GetDB(), "emails")

	// if email address is not confirmed, then delete user
	email := c.Email
	if err := email.OnlyGetByAddress(GetDB()); err == nil && !email.Verified {
		user := User{
			Email: c.Email,
		}
		// get user preloaded
		if err := user.OnlyGetByEmailAddress(GetDB()); err != nil {
			return ReturnInternalDbError(err.Error())
		}

		// delete all user info by email address
		if err := user.DeleteUserByEmail(c.Email, GetDB()); err != nil {
			return ReturnInternalDbError(err.Error())
		}
	}

	trans := GetDB().Begin()
	defer func(){
		if trans != nil {
			trans.Rollback()
		}
	}()

	// validate password
	if err := OnlyValidatePassword(c.Password); err != nil {
		return ReturnInvalidParameters(err.Error())
	}

	// hash the password
	hashed, err := helper.Convert_string_to_hash(c.Password)
	if err != nil {
		return ReturnInvalidParameters("could not hash password. user. hash")
	}
	c.Password = string(hashed)
	
	// get role of the user
	if err := c.Role.OnlyGetByName(trans); err != nil {
		return ReturnInternalDbError(err.Error())
	}
	c.RoleId = c.Role.Id
	c.Email.Verified = true
	c.Email.SentCode = ""
	c.Email.Deadline = time.Time{}

	// store email
	if err := c.Email.OnlyCreate(trans); err != nil && strings.Contains(err.Error(), "duplicate key") {
		return ReturnEmailAlreadyInUse(err.Error())
	} else if err != nil {
		return ReturnInternalDbError(err.Error())
	}
	c.EmailId = c.Email.Id

	// store phone
	c.Phone.Verified = true
	if err := c.Phone.OnlyCreate(trans); err != nil {
		return ReturnInternalDbError(err.Error())
	}
	c.PhoneId = c.Phone.Id

	// create a user with provided info
	c.Verified = true
	if err := c.OnlyCreate(trans); err != nil {
		return ReturnInternalDbError(err.Error())
	}

	if err := trans.Commit().Error; err != nil {
		return ReturnInternalDbError(err.Error())
	}

	trans = nil
	msg := ReturnSuccessfullyCreated()

	return msg
}


