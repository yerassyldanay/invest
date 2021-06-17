package app

import (
	"encoding/json"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/service"
	"github.com/yerassyldanay/invest/utils/message"
	"net/http"
)

/*
	201 - created
	400 - bad request
	409 - already in use
	417 - database error
	422 - could not sent message & not stored on database
	500 - internal server error
 */
func Sign_up(w http.ResponseWriter, r *http.Request) {
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
