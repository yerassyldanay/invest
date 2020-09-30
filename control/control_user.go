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

	var is = service.InvestService{}
	is.OnlyParseRequest(r)

	if msg := is.Check_is_it_admin(); msg.IsThereAnError() {
		utils.Respond(w, r, msg);
		return
	}

	var user = model.User{}
	var roles = r.URL.Query()["role"]

	if len(roles) < 1 {
		roles = []string{utils.RoleManager, utils.RoleExpert}
	}

	msg := user.Get_users_by_roles(roles, is.Offset)
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
var Update_password_help = func(whose string, w http.ResponseWriter, r* http.Request) {
	utils.Respond(w, r, utils.Msg{utils.NoErrorFineEverthingOk, 200, "", ""})
}

var Update_own_password = func(w http.ResponseWriter, r *http.Request) {
	Update_password_help("own", w, r)
}

var Update_other_password = func(w http.ResponseWriter, r *http.Request) {
	Update_password_help("other", w, r)
}


