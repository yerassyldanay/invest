package control

import (
	"invest/model"
	"invest/utils"
	"net/http"
	"strings"
)

var Stats_on_projects_based_on_user_or_status = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Stats_on_projects_based_on_user_or_status"
	var msg utils.Msg

	var pus = model.ProjectUserStat{
		UserId:    Get_query_parameter_uint64(r, "user_id", 0),
		Status:    strings.ToLower(Get_query_parameter_str(r, "status", "")),
	}
	offset := Get_query_parameter_str(r, "offset", "0")

	ok := pus.Status != "" && pus.UserId != 0
	if ok {
		msg = pus.Get_projects_by_status_and_user_id(offset)
	} else if pus.Status != "" {
		msg = pus.Get_projects_by_status(offset)
	} else {
		msg = utils.Msg{utils.ErrorMethodNotAllowed, 405, "", "method is not allowed. status & user_id are not provided"}
	}

	msg.Fname = fname + " 1"

	utils.Respond(w, r, msg)
}
