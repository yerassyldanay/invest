package control

import (
	"encoding/json"
	"invest/model"
	"invest/service"
	"invest/utils/message"
	"net/http"
)

/*
	201 - created
	400 - bad request
	409 - already in use
	417 - db error
	422 - could not sent message & not stored on db
	500 - internal server error
 */
var Sign_up = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Sign_up"
	var user = model.User{}

	// headers
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// parse request body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		OnlyReturnInvalidParametersError(w, r, err.Error(), fname, "json")
		return
	}
	defer r.Body.Close()

	// logic
	msg := is.SignUp(user)
	msg.SetFname(fname, "up")

	message.Respond(w, r, msg)
}
