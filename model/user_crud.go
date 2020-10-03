package model

import "C"
import (
	"invest/utils"
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
func (c *User) Create_user_without_check() (utils.Msg) {
	if c.Lang == "" {
		c.Lang = utils.DefaultContentLanguage
	}

	/*
		update seq ids
			refer to function description for more information
	 */
	//_ = Update_sequence_id_thus_avoid_duplicate_primary_key_error(GetDB(), "users")
	//_ = Update_sequence_id_thus_avoid_duplicate_primary_key_error(GetDB(), "phones")
	//_ = Update_sequence_id_thus_avoid_duplicate_primary_key_error(GetDB(), "emails")

	trans := GetDB().Begin()
	defer func(){
		if trans != nil {
			trans.Rollback()
		}
	}()

	// remove if the user account is not confirmed
	_, _ = c.Remove_all_users_with_not_confirmed_email()

	// validate password
	if err := Validate_password(c.Password); err != nil {
		return ReturnInvalidParameters(err.Error())
	}

	// hash the password
	hashed, err := utils.Convert_string_to_hash(c.Password)
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
	c.Email.SentHash = ""
	c.Email.Deadline = time.Time{}

	// store email
	if err := c.Email.OnlyCreate(trans); err != nil {
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


