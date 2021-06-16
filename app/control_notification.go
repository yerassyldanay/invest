package app

import (
	"invest/service"
	"invest/utils/message"
	"net/http"
)

var Notification_get = func(w http.ResponseWriter, r *http.Request) {
	fname := "Notification_get"

	// headers
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// parameters
	project_id := service.OnlyGetQueryParameter(r, "project_id", uint64(0)).(uint64)

	// logic
	msg := is.Notification_get_by_project_id(project_id)
	msg.SetFname(fname, "n")

	message.Respond(w, r, msg)
}