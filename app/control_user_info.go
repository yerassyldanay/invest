package app

import (
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/service"
	"github.com/yerassyldanay/invest/utils/constants"
	"github.com/yerassyldanay/invest/utils/errormsg"
	"github.com/yerassyldanay/invest/utils/message"
	"net/http"
)

var Get_full_user_info = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Get_full_user_info"
	var user = model.User{
		Id: service.Get_query_parameter_uint64(r, "user_id", 0),
	}

	// admins only
	roleName := service.Get_header_parameter(r, constants.KeyRoleName, "").(string)
	if roleName != constants.RoleAdmin {
		message.Respond(w, r, message.Msg{errormsg.ErrorMethodNotAllowed, 405, "", "only admins can access. your role is: " + roleName})
		return
	}

	msg := user.GetFullInfoOfThisUser(model.ElementGetFullInfoOfThisUser{})
	msg.Fname = fname + " 1"

	message.Respond(w, r, msg)
}
