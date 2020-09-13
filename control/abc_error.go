package control

import (
	"invest/utils"
	"net/http"
)

var Controller_not_found = func(w http.ResponseWriter, r *http.Request) {
	utils.Respond(w, r, utils.Msg{
		Message: map[string]interface{}{
			"eng": "not found | have no clue what you are looking for",
		},
		Status:  http.StatusNotAcceptable,
		Fname:   "Not_found",
		ErrMsg:  "not found | " + r.RemoteAddr,
	})
}
