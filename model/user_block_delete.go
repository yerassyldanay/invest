package model

import (
	"invest/utils"
)

/*
	delete users except for default ones
 */
func (c *User) Delete_user() (*utils.Msg) {
	//if c.Id <= utils.DefaultNotAllowedUserToDelete {
	//	return &utils.Msg{utils.ErrorMethodNotAllowed, 405, "", "trying to delete a default user"}
	//}

	//var trans = GetDB().Begin()
	//defer func() { if trans != nil { trans.Rollback() }}()

	//if err := trans.Exec("delete from projects_users where user_id = ?", c.Id).Error; err != nil {
	//	return &utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	//}

	if err := GetDB().Delete(&User{}, "id = ?", c.Id).Error; err != nil {
		return &utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	return &utils.Msg{utils.NoErrorFineEverthingOk, 200, "", ""}
}

/*
	blocks or gets rid of a block on user
		note: default users will not be blocked at all
 */
func (c *User) Block_unblock_user() (*utils.Msg) {
	//if c.Id <= utils.DefaultNotAllowedUserToDelete {
	//	return &utils.Msg{utils.ErrorMethodNotAllowed, 405, "", "trying to delete default users"}
	//}

	if err := GetDB().First(c, "id = ?", c.Id).Error; err != nil {
		return &utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	c.Blocked = !(c.Blocked)

	if err := GetDB().Save(*c).Error; err != nil {
		return &utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	return &utils.Msg{utils.NoErrorFineEverthingOk, 200, "", ""}
}
