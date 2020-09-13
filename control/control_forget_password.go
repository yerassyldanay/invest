package control

import (
	"encoding/json"
	"invest/model"
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
	var msg = utils.Msg{
		Fname: fname + " 1",
	}

	var fp = model.ForgetPassword{
		Lang: Get_header_parameter(r, utils.HeaderAcceptLanguage, "").(string),
	}
	
	switch r.Method {
	case http.MethodGet:
		fp.EmailAddress = Get_query_parameter_str(r, "email", "")
		msg = fp.SendMessage()

	case http.MethodPost:
		if err := json.NewDecoder(r.Body).Decode(&fp); err != nil {
			utils.Respond(w, r, utils.Msg{utils.ErrorInvalidParameters, 400, fname + " 2", err.Error()})
			return
		}

		msg = fp.Change_password_of_user_by_hash()
	}
	
	utils.Respond(w, r, msg)
}