package app

import (
	"encoding/json"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/service"
	"github.com/yerassyldanay/invest/utils/errormsg"
	"github.com/yerassyldanay/invest/utils/message"
	"net/http"
)

/*
	this is to confirm that a user owns this email
 */
func UserProfileConfirmEmail(w http.ResponseWriter, r *http.Request) {
	var fname = "UserProfileConfirmEmail"

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

	// ok
	message.Respond(w, r, msg)
}

