package middleware

import (
	"invest/model"
	"invest/service"
	"invest/utils/constants"
	"invest/utils/errormsg"
	"invest/utils/message"

	"net/http"
	"strings"
)

/*
	This wrapper will check whether a request valid or invalid
	The pather is composed of several parts:
		/v1/permission_type/other/pather/part
	This handler will:
		1. parse id of the user
		2. get permission_type
		3. find whether given user has such a permission
			if yes: a request will be forwarded
			else: 'the method is not allowed' message will be sent
 */
var HasPermissionWrapper = func(w http.ResponseWriter, r *http.Request)  (message.Msg) {
		var fname = "check_whether_user_has_such_permission"
		var up = model.UserPermission{}

		up.UserId = service.Get_header_parameter(r, constants.KeyId, uint64(0)).(uint64)

		/*
			/v1/permission/1/2 -> [v1, permission, 1, 2]
		*/
		paths := strings.Split(r.URL.Path, "/")
		if len(paths) < 2 {
			return message.Msg{
				Message: errormsg.ErrorInternalServerError,
				Status:  http.StatusMisdirectedRequest,
				Fname:   fname + " 1",
				ErrMsg:  "the pather is invalid",
			}
		}

		up.Permission = paths[2]

		var msg = up.Check_db_whether_this_user_has_such_a_permission()
		msg.Fname = fname + " 2"

		if msg.ErrMsg != "" {
			return msg
		}

		/*
			this means user has such a permission
		 */
		return Parse_prefered_language_of_user(w, r)
}

