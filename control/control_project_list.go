package control

import (
	"invest/model"
	"invest/service"
	"invest/utils"
	"net/http"
)

var Get_projects_by_user_id = func (w http.ResponseWriter, r *http.Request) {
	var fname = "Get_projects_by_user_id"
	var project = model.Project{}

	roleName := service.Get_header_parameter(r, utils.KeyRoleName, "").(string)
	if roleName != utils.RoleAdmin {
		utils.Respond(w, r, utils.Msg{utils.ErrorMethodNotAllowed, 405, "", "your role is " + roleName})
		return
	}

	user_id := service.Get_query_parameter_uint64(r, "user_id", 0)
	msg := project.Get_projects_by_user_id(user_id)

	msg.Fname = fname + " 1"

	utils.Respond(w, r, msg)
}
