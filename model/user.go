package model

import (
	"invest/utils/message"
)

/*
	delete users except for default ones
 */
func (c *User) Delete_user() (message.Msg) {

	var trans = GetDB().Begin()
	defer func() {if trans != nil {trans.Rollback()}}()

	// first get user info
	err := c.OnlyGetUserById(trans)
	if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	/*
		set ids to delete
	 */
	c.Email.Id = c.EmailId
	c.Phone.Id = c.PhoneId
	c.Role.Id = c.RoleId

	/*
		delete email
	 */
	err = c.Email.OnlyDeleteById(trans)
	if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	/*
		delete
	 */
	err = c.Phone.OnlyDeleteById(trans)
	if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	/*
		delete the user
	 */
	err = c.OnlyDeleteUserById(trans)
	if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	trans.Commit()
	trans = nil

	return ReturnNoError()
}

/*
	blocks or gets rid of a block on user
		note: default users will not be blocked at all
 */
func (c *User) Block_unblock_user() (message.Msg) {
	//if c.Id <= utils.DefaultNotAllowedUserToDelete {
	//	return utils.Msg{utils.ErrorMethodNotAllowed, 405, "", "trying to delete default users"}
	//}

	if err := GetDB().First(c, "id = ?", c.Id).Error; err != nil {
		return ReturnInternalDbError(err.Error())
	}

	c.Blocked = !(c.Blocked)

	if err := GetDB().Save(*c).Error; err != nil {
		return ReturnInternalDbError(err.Error())
	}

	return ReturnNoError()
}

/*
	check whether a user has a following role
*/
func (u *User) Is_this_is_role_of_user(role string) (bool) {
	if err := GetDB().Preload("Role").Where("id = ?", u.Id).First(u).Error;
		err != nil {
		return false
	}

	return u.Role.Name == role
}