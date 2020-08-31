package auth

import (
	"invest/control"
	"invest/model"
	"invest/utils"
	"net/http"
	"strings"
)

/*
	statuses: 200, 406, 417
 */
var EmailVerifiedWrapper = func(next http.Handler, w http.ResponseWriter, r *http.Request) {
		var fname = "EmailVerifiedWrapper"
		var up = model.UserPermission{}

		up.UserId = control.Get_header_parameter(r, utils.KeyId, uint64(0)).(uint64)

		for _, path := range utils.NoNeedToConfirmEmail {
			if strings.Contains(r.URL.Path, path){
				next.ServeHTTP(w, r)
			}
		}

		/*
			check whether user has confirmed an ownership on email address
		*/
		msg := up.Check_db_whether_this_user_account_is_confirmed()
		if msg.ErrMsg != "" && msg.Status != 200 {
			msg.Fname = fname
			utils.Respond(w, r, msg)
			return
		}

		/*
			this means user has confirmed an email address
		*/
		HasPermissionWrapper(next, w, r)
}
