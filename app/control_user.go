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

/* */
var UsersGetByRole = func(w http.ResponseWriter, r *http.Request) {
	var fname = "User_get_by_role"

	// headers
	var is = service.InvestService{}
	is.OnlyParseRequest(r)

	// only admins can get users by role
	if msg := is.CheckIsItAdmin(); msg.IsThereAnError() {
		message.Respond(w, r, msg)
		return
	}

	// if nothing is provided then all manager & expert will be given
	var roles = r.URL.Query()["role"]
	if len(roles) < 1 {
		roles = []string{constants.RoleManager, constants.RoleExpert}
	}

	// logic
	msg := is.GetUsersByRoles(roles)
	msg.Fname = fname + " get"

	message.Respond(w, r, msg)
}

var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Create_user_based_on_role"
	var user = model.User{}

	// headers
	var is = service.InvestService{}
	is.OnlyParseRequest(r)

	// decode user info
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		message.Respond(w, r, message.Msg{errormsg.ErrorInvalidParameters, 400, fname + " json", err.Error()})
		return
	}
	defer r.Body.Close()

	// check
	if msg := is.CheckIsItAdmin(); msg.IsThereAnError() {
		message.Respond(w, r, msg)
		return
	}

	// create
	msg := is.CreateUserBasedOnRole(&user)
	if msg.IsThereAnError() != true {
		msg = model.ReturnNoError()
	}
	msg.Fname = fname + " 1"

	message.Respond(w, r, msg)
}

var UpdateUserProfileHelp = func(whose string, w http.ResponseWriter, r *http.Request) {
	var fname = "Update_user_profile"
	var user = model.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		message.Respond(w, r, message.Msg{errormsg.ErrorInvalidParameters, 400, fname + " json", err.Error()})
		return
	}
	defer r.Body.Close()

	//headers
	var is = service.InvestService{}
	is.OnlyParseRequest(r)

	// check & logic
	var msg message.Msg
	switch whose {
	case "other":
		// pass
		if msg := is.CheckIsItAdmin(); msg.IsThereAnError() {
			message.Respond(w, r, msg)
			return
		}
		msg = is.UpdateUserProfile(&user)
	case "own":
		user.Id = is.UserId
		msg = is.UpdateUserProfile(&user)
	default:
		msg = model.ReturnMethodNotAllowed("this is not supported | which profile are updating")
	}

	msg.SetFname(fname, " update")
	message.Respond(w, r, msg)
}

var UpdateOwnProfile = func(w http.ResponseWriter, r *http.Request) {
	UpdateUserProfileHelp("own", w, r)
}

var UpdateOtherProfile = func(w http.ResponseWriter, r *http.Request) {
	UpdateUserProfileHelp("other", w, r)
}

// UpdatePasswordHelp Password
var UpdatePasswordHelp = func(whose string, w http.ResponseWriter, r *http.Request) {
	var fname = "Update_password_help"
	var requestBody = struct {
		Id          uint64 `json:"id"`
		Password    string `json:"password"`
		OldPassword string `json:"old_password"`
	}{}

	var msg message.Msg

	// get request body
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		msg = model.ReturnInvalidParameters(err.Error())
		msg.Fname = fname + " json"
		message.Respond(w, r, msg)
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
		msg = is.UpdateUserPassword(requestBody.OldPassword, requestBody.Password)
	case whose == "other" && is.RoleName == constants.RoleAdmin:

		// security check
		if is.RoleName != constants.RoleAdmin {
			OnlyReturnMethodNotAllowed(w, r, "only admin can access, your role "+is.RoleName, fname, "sec")
			return
		}

		// if this is a profile of another user
		// then set is.UserId to the id of that user
		is.UserId = requestBody.Id
		msg = is.UpdateUserPassword(requestBody.OldPassword, requestBody.Password)

	default:
		// this is not allowed
		msg = model.ReturnMethodNotAllowed("requesting " + whose + " | role is " + is.RoleName)
	}

	msg.Fname = fname + " update"
	message.Respond(w, r, msg)
}

var UpdateOwnPassword = func(w http.ResponseWriter, r *http.Request) {
	UpdatePasswordHelp("own", w, r)
}

var UpdateOtherPassword = func(w http.ResponseWriter, r *http.Request) {
	UpdatePasswordHelp("other", w, r)
}
