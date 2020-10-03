package control

import (
	"encoding/json"
	"invest/model"
	"invest/service"
	"invest/utils"

	"net/http"
)

/* */
var Users_get_by_role = func(w http.ResponseWriter, r *http.Request) {
	var fname = "User_get_by_role"

	// headers
	var is = service.InvestService{}
	is.OnlyParseRequest(r)

	// only admins can get users by role
	if msg := is.Check_is_it_admin(); msg.IsThereAnError() {
		utils.Respond(w, r, msg);
		return
	}

	// if nothing is provided then all manager & expert will be given
	var roles = r.URL.Query()["role"]
	if len(roles) < 1 {
		roles = []string{utils.RoleManager, utils.RoleExpert}
	}

	// logic
	msg := is.Get_users_by_roles(roles)
	msg.Fname = fname + " get"

	utils.Respond(w, r, msg)
}

var Create_user = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Create_user_based_on_role"
	var user = model.User{}

	var is = service.InvestService{}
	is.OnlyParseRequest(r)

	// decode user info
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.Respond(w, r, utils.Msg{utils.ErrorInvalidParameters, 400, fname + " json", err.Error()})
		return
	}
	defer r.Body.Close()

	// check
	if msg := is.Check_is_it_admin(); msg.IsThereAnError() {
		utils.Respond(w, r, msg); return
	}

	// create
	msg := is.Create_user_based_on_role(&user)
	if msg.IsThereAnError() != true {
		msg = model.ReturnNoError()
	}
	msg.Fname = fname + " 1"

	utils.Respond(w, r, msg)
}

var Update_user_profile_help = func(whose string, w http.ResponseWriter, r *http.Request) {
	var fname = "Update_user_profile"
	var user = model.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.Respond(w, r, utils.Msg{utils.ErrorInvalidParameters, 400, fname + " json", err.Error()})
		return
	}
	defer r.Body.Close()

	var is = service.InvestService{}
	is.OnlyParseRequest(r)

	// check
	if msg := is.Check_is_it_admin(); msg.IsThereAnError() {
		utils.Respond(w, r, msg); return
	}

	switch whose {
	case "other":
		// pass
	default:
		user.Id = is.UserId
	}

	msg := is.Update_user_profile(&user)
	msg.Fname = fname + " update"

	utils.Respond(w, r, msg)
}

var Update_own_profile = func(w http.ResponseWriter, r *http.Request) {
	Update_user_profile_help("own", w, r)
}

var Update_other_profile = func(w http.ResponseWriter, r *http.Request) {
	Update_user_profile_help("other", w, r)
}

/*
	Password
 */
var Update_password_help = func(whose string, w http.ResponseWriter, r *http.Request) {
	var fname = "Update_password_help"
	var user = model.User{}

	var msg utils.Msg

	// get request body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		msg = model.ReturnInvalidParameters(err.Error())
		msg.Fname = fname + " json"
		utils.Respond(w, r, msg)
		return
	}
	defer r.Body.Close()

	// headers
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// security check
	switch {
	case whose == "own":
		// if this is own then is.UserId will be used
		msg = is.Update_user_password(user.Password)
	case whose == "other" && is.RoleName == utils.RoleAdmin:
		// if this is a profile of another user
		// then set is.UserId to the id of that user
		is.UserId = user.Id
		msg = is.Update_user_password(user.Password)
	default:
		// this is not allowed
		msg = model.ReturnMethodNotAllowed("requesting " + whose + " | role is " + is.RoleName)
	}

	msg.Fname = fname + " update"
	utils.Respond(w, r, msg)
}

var Update_own_password = func(w http.ResponseWriter, r *http.Request) {
	Update_password_help("own", w, r)
}

var Update_other_password = func(w http.ResponseWriter, r *http.Request) {
	Update_password_help("other", w, r)
}


