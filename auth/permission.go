package auth

import (
	"invest/control"
	"invest/model"
	"invest/utils"
	"net/http"
	"strings"
)

/*
	This wrapper will check whether a request valid or invalid
	The path is composed of several parts:
		/v1/permission_type/other/path/part
	This handler will:
		1. parse id of the user
		2. get permission_type
		3. find whether given user has such a permission
			if yes: a request will be forwarded
			else: 'the method is not allowed' message will be sent
 */
var HasPermissionWrapper = func(next http.Handler, w http.ResponseWriter, r *http.Request) {
		var fname = "check_whether_user_has_such_permission"
		var up = model.UserPermission{}

		up.UserId = control.Get_header_parameter(r, utils.KeyId, uint64(0)).(uint64)

		/*
			/v1/permission/1/2 -> [v1, permission, 1, 2]
		*/
		paths := strings.Split(r.URL.Path, "/")
		if len(paths) < 2 {
			utils.Respond(w, r, utils.Msg{
				Message: utils.ErrorInternalServerError,
				Status:  http.StatusMisdirectedRequest,
				Fname:   fname + " 1",
				ErrMsg:  "the path is invalid",
			})
			return
		}

		up.Permission = paths[2]

		var msg = up.Check_db_whether_this_user_has_such_a_permission()
		msg.Fname = fname + " 2"

		if msg.ErrMsg != "" {
			utils.Respond(w, r, msg)
			return
		}

		/*
			this means user has such a permission
		 */
		Parse_prefered_language_of_user(next, w, r)
}

