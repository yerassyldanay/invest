package control

import (
	"invest/model"
	"invest/utils"
	"net/http"
)

var Get_projects_by_user_id = func (w http.ResponseWriter, r *http.Request) {
	var fname = "Get_projects_by_user_id"
	var project = model.Project{}

	user_id := Get_query_parameter_uint64(r, "user_id", 0)
	msg := project.Get_projects_by_user_id(user_id)

	msg.Fname = fname + " 1"

	utils.Respond(w, r, msg)
}
