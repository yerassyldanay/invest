package control

import (
	"encoding/json"
	"invest/model"
	"invest/service"
	"invest/utils/constants"
	"invest/utils/message"
	"net/http"
)

var Smtp_create_update_put = func(w http.ResponseWriter, r *http.Request) {
	var fname = "smtp_create"
	var smtp = model.SmtpServer{}

	// parse values
	if err := json.NewDecoder(r.Body).Decode(&smtp); err != nil {
		OnlyReturnInvalidParametersError(w, r, err.Error(), fname, "json")
		return
	}
	defer r.Body.Close()

	// headers
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// security check
	if is.RoleName != constants.RoleAdmin {
		OnlyReturnMethodNotAllowed(w, r, "only admins can access this", fname, "2")
		return
	}

	// methods
	var msg message.Msg
	switch r.Method {
	case http.MethodPost:
		msg = is.SmtpCreate(&smtp)
	case http.MethodPut:
		msg = is.SmtpUpdate(&smtp)
	case http.MethodDelete:
		msg = is.SmtpDelete(&smtp)
	default:
		msg = model.ReturnMethodNotAllowed("this method is not allowed")
	}

	msg.SetFname(fname, "resp")
	message.Respond(w, r, msg)
}

var Smtp_get = func (w http.ResponseWriter, r *http.Request) {
	fname := "Smtp_get"

	// headers
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// get all
	msg := is.SmtpGet()

	msg.SetFname(fname, "1")
	message.Respond(w, r, msg)
}
