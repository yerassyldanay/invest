package model

import (
	"github.com/yerassyldanay/invest/utils/errormsg"
	"github.com/yerassyldanay/invest/utils/message"
)

/*
	create & store a comment on database
*/
func (c *Comment) Create_comment_after_saving_its_document() (message.Msg) {
	if err := c.Validate(); err != nil {
		return message.Msg{errormsg.ErrorInvalidParameters, 400, "", err.Error()}
	}

	if err := c.OnlyCreate(GetDB()); err != nil {
		return message.Msg{errormsg.ErrorInternalDbError, 417, "", err.Error()}
	}

	var resp = errormsg.NoErrorFineEverthingOk

	return message.Msg{resp, 200, "", ""}
}

/*
	get comments of the project
*/
func (c *Comment) Get_all_comments_of_the_project_by_project_id(offset interface{}) (message.Msg) {

	var commentsMap = []map[string]interface{}{}
	comments, err := c.OnlyGetCommentsByProjectId(offset, GetDB())

	for _, comment := range comments {
		commentsMap = append(commentsMap, Struct_to_map(comment))
	}

	var resp = errormsg.NoErrorFineEverthingOk
	resp["info"] = commentsMap

	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}

	return message.Msg{resp, 200, "", errMsg}
}

func (c *Comment) Get_comment_by_comment_id() (message.Msg) {
	/*
		get only one comment
	 */
	err := c.OnlyGetById(GetDB())
	if err != nil {
		return message.Msg{errormsg.ErrorInternalDbError, 417, "", err.Error()}
	}

	var resp = errormsg.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(*c)

	return message.Msg{resp, 200, "", ""}
}
