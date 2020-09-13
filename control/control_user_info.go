package control

import (
	"invest/model"
	"invest/utils"
	"net/http"
)

var Get_full_user_info = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Get_full_user_info"
	var user = model.User{
		Id: Get_query_parameter_uint64(r, "user_id", 0),
	}

	msg := user.Get_full_info_of_this_user("id")
	msg.Fname = fname + " 1"

	utils.Respond(w, r, msg)
}