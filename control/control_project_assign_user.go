package control

import (
	"encoding/json"
	"invest/model"
	"invest/service"
	"invest/utils"

	"net/http"
)

var Assign_user_to_project = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Admin_assign_user_to_project"
	var pu = model.ProjectsUsers{}

	if err := json.NewDecoder(r.Body).Decode(&pu); err  != nil {
		utils.Respond(w, r, utils.Msg{utils.ErrorInvalidParameters, 400, fname + " 1", err.Error()})
		return
	}
	defer r.Body.Close()

	is := service.InvestService{}
	is.OnlyParseRequest(r)

	if is.RoleName != utils.RoleAdmin {
		utils.Respond(w, r, utils.Msg{utils.ErrorMethodNotAllowed, 405, fname + " role", "not admin"})
		return
	}

	msg := pu.Assign_user_after_check()
	msg.Fname = fname + " assign"

	utils.Respond(w, r, msg)
}

/*
	remove user & project relation
 */
var Remove_user_from_project = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Remove_user_from_project"
	var pu = model.ProjectsUsers{}

	if err := json.NewDecoder(r.Body).Decode(&pu); err != nil {
		utils.Respond(w, r, utils.Msg{
			Message: utils.ErrorInvalidParameters,
			Status:  400,
			Fname:   fname + " 1",
			ErrMsg:  err.Error(),
		})
		return
	}
	defer r.Body.Close()

	msg := pu.Remove_relation()
	utils.Respond(w, r, msg)
}

