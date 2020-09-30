package auth

import (
	"invest/model"
	"invest/service"
	"invest/utils"

	"net/http"
	"strings"
)

/*
	statuses: 200, 406, 417
 */
var EmailVerifiedWrapper = func(w http.ResponseWriter, r *http.Request) (utils.Msg) {
	var fname = "EmailVerifiedWrapper"
	var up = model.UserPermission{}

	up.UserId = service.Get_header_parameter(r, utils.KeyId, uint64(0)).(uint64)

	for _, path := range utils.NoNeedToConfirmEmail {
		if strings.Contains(r.URL.Path, path){
			return utils.Msg{}
		}
	}

	/*
		check whether user has confirmed an ownership on email address
	*/
	msg := up.Check_db_whether_this_user_account_is_confirmed()
	if msg.ErrMsg != "" && msg.Status != 200 {
		msg.Fname = fname
		return msg
	}

	return Parse_prefered_language_of_user(w, r)
}
