package model

import (
	"invest/utils"
)

/*
	delete users except for default ones
 */
func (c *User) Delete_user() (utils.Msg) {

	var trans = GetDB().Begin()
	defer func() {if trans != nil {trans.Rollback()}}()

	// first get user info
	err := c.GetByIdOnlyUser(trans)
	if err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
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
	err = c.Email.DeleteById(trans)
	if err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	/*
		delete
	 */
	err = c.Phone.DeleteById(trans)
	if err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	/*
		delete the user
	 */
	err = c.DeleteOnlyUserById(trans)
	if err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	trans.Commit()
	trans = nil

	return utils.Msg{utils.NoErrorFineEverthingOk, 200, "", ""}
}

/*
	blocks or gets rid of a block on user
		note: default users will not be blocked at all
 */
func (c *User) Block_unblock_user() (utils.Msg) {
	//if c.Id <= utils.DefaultNotAllowedUserToDelete {
	//	return utils.Msg{utils.ErrorMethodNotAllowed, 405, "", "trying to delete default users"}
	//}

	if err := GetDB().First(c, "id = ?", c.Id).Error; err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	c.Blocked = !(c.Blocked)

	if err := GetDB().Save(*c).Error; err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	return utils.Msg{utils.NoErrorFineEverthingOk, 200, "", ""}
}
