package model

import (
	"invest/utils"
)

/*
	create & store a comment on db
*/
func (c *Comment) Create_comment_after_saving_its_document() (utils.Msg) {
	if err := c.Validate(); err != nil {
		return utils.Msg{utils.ErrorInvalidParameters, 400, "", err.Error()}
	}

	if err := c.Only_create(GetDB()); err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	var resp = utils.NoErrorFineEverthingOk
	//resp["info"] = Struct_to_map(*c)

	return utils.Msg{resp, 200, "", ""}
}

/*
	get comments of the project
*/
func (c *Comment) Get_all_comments_of_the_project_by_project_id(offset interface{}) (utils.Msg) {

	var commentsMap = []map[string]interface{}{}
	comments, err := c.Only_get_comments_by_project_id(offset, GetDB())

	for _, comment := range comments {
		commentsMap = append(commentsMap, Struct_to_map(comment))
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = commentsMap

	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}

	return utils.Msg{resp, 200, "", errMsg}
}

func (c *Comment) Get_comment_by_comment_id() (utils.Msg) {
	/*
		get only one comment
	 */
	err := c.Only_get_comment_by_comment_id(GetDB())
	if err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(*c)

	return utils.Msg{resp, 200, "", ""}
}
