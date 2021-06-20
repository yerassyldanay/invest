package app

import (
	"encoding/json"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/service"
	"github.com/yerassyldanay/invest/utils/constants"
	"github.com/yerassyldanay/invest/utils/errormsg"
	"github.com/yerassyldanay/invest/utils/message"
	"net/http"
)

/*
	each ganta step has its 'is_done' field, which is an indication of whether
		this step is passed
	this function helps users pass the step manually
*/
var Ganta_confirm_the_ganta_step = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Ganta_shift_to_the_next_ganta_step"

	var reqBody = struct {
		ProjectId uint64 `json:"project_id"`
		Status    string `json:"status"`
	}{}

	//var project_id = service.OnlyGetQueryParameter(r, "project_id", uint64(0)).(uint64)
	//var status = service.OnlyGetQueryParameter(r, "status", "").(string)
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		message.Respond(w, r, message.Msg{errormsg.ErrorInvalidParameters, 400, fname + " json", err.Error()})
		return
	}
	defer r.Body.Close()

	// headers
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// security check - access to project
	msg := is.CheckWhetherThisUserCanGetAccessToProjectInfo(reqBody.ProjectId)
	if msg.IsThereAnError() {
		message.Respond(w, r, msg)
		return
	}

	// security check - can change status
	msg = is.Check_whether_this_user_responsible_for_current_step(reqBody.ProjectId)
	if msg.IsThereAnError() {
		message.Respond(w, r, msg)
		return
	}

	// security check
	_, msg = is.Ganta_can_user_change_current_status(reqBody.ProjectId)
	if msg.IsThereAnError() {
		message.Respond(w, r, msg)
		return
	}

	msg = is.GantaChangeTheStatusOfProject(reqBody.ProjectId, reqBody.Status)
	msg.Fname = fname + " status"

	message.Respond(w, r, msg)
}

var Ganta_change_ganta_time = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Ganta_change_ganta_time"

	// parsing the request body
	var ganta = model.Ganta{}
	if err := json.NewDecoder(r.Body).Decode(&ganta); err != nil {
		message.Respond(w, r, message.Msg{errormsg.ErrorInvalidParameters, 400, fname + " json", err.Error()})
		return
	}
	defer r.Body.Close()

	// headers
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// security check
	if is.RoleName != constants.RoleAdmin {
		errmsg := "only admin can access. your role is " + is.RoleName
		OnlyReturnMethodNotAllowed(w, r, errmsg, fname, "role")
		return
	}

	// logic
	msg := is.Ganta_change_time(ganta)
	msg.SetFname(fname, "time")

	message.Respond(w, r, msg)
}

// change time of gantt step
var Ganta_can_user_change_current_status = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Ganta_can_user_change_current_status"

	// headers
	var is = service.InvestService{}
	is.OnlyParseRequest(r)

	// parameters
	var project_id = service.OnlyGetQueryParameter(r, "project_id", uint64(0)).(uint64)

	// logic
	_, msg := is.Ganta_can_user_change_current_status(project_id)
	msg.SetFname(fname, "ganta")

	message.Respond(w, r, msg)
}
