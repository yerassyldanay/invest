package control

import (
	"invest/service"
	"invest/utils"

	"net/http"
)

var Get_own_emails_by_project_id = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Get_emails_of_any_project"
	var user_id = service.Get_header_parameter(r, utils.KeyId, uint64(0)).(uint64)
	var project_id = service.Get_query_parameter_uint64(r, "project_id", 0)
	var offset = service.Get_query_parameter_str(r, "offset", "0")

	var is = service.InvestService{
		BasicInfo: service.BasicInfo{
			UserId: 	user_id,
			RoleId: 	0,
			Lang: 		"",
		},
	}

	/*
		this service function:
			* checks the project
			* get emails by project id & email address (which will be retrieved using user id)
	 */
	msg := is.Get_own_emails_by_project_id_after_check(project_id, offset)
	msg.Fname = fname + " 1"

	utils.Respond(w, r, msg)
}
