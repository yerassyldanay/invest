package control

import (
	"invest/model"
	"invest/utils"
	"net/http"
)

var Get_all_assigned_users_to_project = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Project_get_all_assigned_users"
	var project = model.Project{
		Id: Get_query_parameter_uint64(r, "project_id", 0),
	}

	msg := project.Get_this_project_with_its_users()
	msg.Fname = fname + " 1"

	utils.Respond(w, r, msg)
}
