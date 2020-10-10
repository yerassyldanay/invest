package control

import (
	"encoding/json"
	"invest/model"
	"invest/service"
	"invest/utils"

	"net/http"
)

/*
	GET:
		* ?email=yerassyl.danay@nu.edu.kz
		* will send message to the email address writen here
	POST:
		* using the link in the email you can reset your password
		* data:{
			"code": "",
			"new_password": "",
			""
		}
 */
var Forget_password_send_message = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Forget_password_send_message"

	// headers
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	var msg utils.Msg
	var fp = model.ForgetPassword{}

	// request methods
	switch r.Method {
	case http.MethodGet:
		fp.EmailAddress = service.OnlyGetQueryParameter(r, "email", "").(string)
		msg = is.Password_reset_send_message(fp)

	case http.MethodPost:
		if err := json.NewDecoder(r.Body).Decode(&fp); err != nil {
			utils.Respond(w, r, utils.Msg{utils.ErrorInvalidParameters, 400, fname + " 2", err.Error()})
			return
		}

		defer r.Body.Close()

		msg = is.Password_reset_change_password(fp)

	default:
		msg = model.ReturnMethodNotAllowed("only post & get requests are supported")
	}

	msg.SetFname(fname, " 2")
	utils.Respond(w, r, msg)
}