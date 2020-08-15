package model

import (
	"errors"
	"invest/utils"
)

/*
	delete users except for default ones
 */
func (c *User) Delete_user() (map[string]interface{}, error) {
	if c.Id <= utils.DefaultNotAllowedUserToDelete {
		return utils.ErrorMethodNotAllowed, errors.New("trying to delete a default user")
	}

	var trans = GetDB().Begin()
	defer func() { if trans != nil { trans.Rollback() }}()

	if err := trans.Exec("delete from projects_users where user_id = ?", c.Id).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}

	if err := trans.Exec("delete from users where id = ?", c.Id).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}

	return utils.NoErrorFineEverthingOk, nil
}

/*
	blocks or gets rid of a block on user
		note: default users will not be blocked at all
 */
func (c *User) Block_unblock_user() (map[string]interface{}, error) {
	if c.Id <= utils.DefaultNotAllowedUserToDelete {
		return utils.ErrorMethodNotAllowed, errors.New("trying to delete deafult users")
	}

	c.Blocked = !(c.Blocked)

	if err := GetDB().Save(*c).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}

	return utils.NoErrorFineEverthingOk, nil
}
