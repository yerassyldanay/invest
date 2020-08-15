package model

import (
	"errors"
	"invest/utils"
)

/*
	comment documents must be stored on disk beforehand
		comment can have no docs attached
 */
func (c *Comment) Validate() bool {
	if c.Subject == "" || c.Body == "" || c.ProjectId == 0 {
		return false
	}

	return true
}

/*
	create & store a comment on db
 */
func (c *Comment) Create_comment_after_saving_its_document() (map[string]interface{}, error) {
	if ok := c.Validate(); !ok {
		return utils.ErrorInvalidParameters, errors.New("invalid parameters have been provided")
	}

	if err := GetDB().Create(c).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(*c)

	return resp, nil
}

/*
	get comments of the project
 */
func (c *Comment) Get_all_comments_to_the_project() (map[string]interface{}, error) {
	if err := GetDB().Set("gorm:auto_preload", false).Table(c.TableName()).
			Where("project_id=?", c.ProjectId).First(c).Error; err != nil {
				return utils.ErrorInternalDbError, err
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(*c)

	return resp, nil
}
