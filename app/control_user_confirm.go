package app

import (
	"encoding/json"
	"invest/model"
	"invest/service"
	"invest/utils/errormsg"
	"invest/utils/message"
	"net/http"
)

/*
	this is to confirm that a user owns this email
 */
var User_email_confirm = func(w http.ResponseWriter, r *http.Request) {
	var fname = "User_email_confirm"

	// headersOnlyGetByAddress
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// request body
	var email = model.Email{}
	if err := json.NewDecoder(r.Body).Decode(&email); err != nil {
		message.Respond(w, r, message.Msg{errormsg.ErrorInvalidParameters, 400, fname + " 2", err.Error()})
		return
	}

	defer r.Body.Close()

	// logic
	msg := is.EmailConfirm(email)
	msg.SetFname(fname, "c")

	message.Respond(w, r, msg)
}

