package control

import (
	"invest/model"
	"invest/service"
	"invest/utils"
	"net/http"
)

var Get_full_user_info = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Get_full_user_info"
	var user = model.User{
		Id: service.Get_query_parameter_uint64(r, "user_id", 0),
	}

	// admins only
	roleName := service.Get_header_parameter(r, utils.KeyRoleName, "").(string)
	if roleName != utils.RoleAdmin {
		utils.Respond(w, r, utils.Msg{utils.ErrorMethodNotAllowed, 405, "", "only admins can access. your role is: " + roleName})
		return
	}

	msg := user.Get_full_info_of_this_user("id")
	msg.Fname = fname + " 1"

	utils.Respond(w, r, msg)
}
