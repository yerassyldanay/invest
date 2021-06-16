package app

import (
	"invest/model"
	"invest/utils/message"
	"net/http"
)

func OnlyReturnInvalidParametersError(w http.ResponseWriter, r *http.Request, errmsg string, fname string, appendix string) {
	msg := model.ReturnInvalidParameters(errmsg)
	msg.SetFname(fname, appendix)
	message.Respond(w, r, msg)
}

func OnlyReturnMethodNotAllowed(w http.ResponseWriter, r *http.Request, errmsg string, fname string, appendix string) {
	msg := model.ReturnMethodNotAllowed(errmsg)
	msg.SetFname(fname, appendix)
	message.Respond(w, r, msg)
}

func OnlyReturnInternalDbError(w http.ResponseWriter, r *http.Request, errmsg string, fname string, appendix string) {
	msg := model.ReturnInternalDbError(errmsg)
	msg.SetFname(fname, appendix)
	message.Respond(w, r, msg)
}

