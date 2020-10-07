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

	// headers
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// security check
	if is.RoleName != utils.RoleAdmin {
		utils.Respond(w, r, utils.Msg{utils.ErrorMethodNotAllowed, 405, fname + " role", "not admin"})
		return
	}

	// check project
	if msg := is.Does_project_exist(pu.ProjectId); msg.IsThereAnError() {
		utils.Respond(w, r, msg)
		return
	}

	// check user
	if msg := is.Does_user_has_given_role(pu.UserId, []string{utils.RoleExpert, utils.RoleManager}); msg.IsThereAnError() {
		utils.Respond(w, r, msg)
		return
	}

	// logic
	msg := is.Assign_user_to_project(pu)
	msg.Fname = fname + " asg"

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

	// headers
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// logic
	msg := is.Assign_remove_relation(pu)
	utils.Respond(w, r, msg)
}

