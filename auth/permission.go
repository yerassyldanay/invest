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
var HasPermissionAndEmailVerifiedWrapper = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var fname = "check_whether_user_has_such_permission"
		var up = model.UserPermission{}
		up.UserId = control.Get_query_parameter_uint64(r, utils.KeyId, 0)

		/*
			/v1/permission/1/2 -> [v1, permission, 1, 2]
		*/
		paths := strings.Split(r.URL.Path, "/")
		if len(paths) < 2 {
			utils.Respond(w, r, &utils.Msg{
				Message: utils.ErrorInternalServerError,
				Status:  http.StatusInternalServerError,
				Fname:   fname + " 1",
				ErrMsg:  "the path is invalid",
			})
			return
		}

		up.Permission = paths[1]

		var msg = up.Check_on_db_whether_this_user_has_such_a_permission_and_user_account_is_confirmed()
		msg.Fname = fname + " 2"

		if msg.ErrMsg != "" {
			utils.Respond(w, r, msg)
		}

		/*
			this means user has such a permission
		 */
		next.ServeHTTP(w, r)
	})
}
