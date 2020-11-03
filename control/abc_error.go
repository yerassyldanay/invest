package control

import (
	"invest/utils/message"
	"net/http"
)

var Controller_not_found = func(w http.ResponseWriter, r *http.Request) {
	message.Respond(w, r, message.Msg{
		Message: map[string]interface{}{
			"eng": "not found | have no clue what you are looking for",
		},
		Status:  http.StatusNotAcceptable,
		Fname:   "Not_found",
		ErrMsg:  "not found | " + r.RemoteAddr,
	})
}
