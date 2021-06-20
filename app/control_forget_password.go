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
var ForgetPasswordSendMessage = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Forget_password_send_message"

	// headers
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	var msg message.Msg
	var fp = model.ForgetPassword{}

	// request methods
	switch r.Method {
	case http.MethodGet:
		fp.EmailAddress = service.OnlyGetQueryParameter(r, "email", "").(string)
		msg = is.PasswordResetSendMessage(fp)

	case http.MethodPost:
		if err := json.NewDecoder(r.Body).Decode(&fp); err != nil {
			message.Respond(w, r, message.Msg{errormsg.ErrorInvalidParameters, 400, fname + " 2", err.Error()})
			return
		}
		defer r.Body.Close()

		msg = is.PasswordResetChangePassword(fp)

	default:
		msg = model.ReturnMethodNotAllowed("only post & get requests are supported")
	}

	msg.SetFname(fname, " 2")
	message.Respond(w, r, msg)
}
