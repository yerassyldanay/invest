package control

import (
	"invest/model"
	"invest/utils"
	"net/http"
)

/*
	this is to confirm that a user owns this email
 */
var User_email_confirm = func(w http.ResponseWriter, r *http.Request) {
	var fname = "User_email_confirm"

	key := model.Get_value_from_query(r, "key")
	hashcode := model.Get_value_from_query(r, "hashcode")

	var email = model.Email{
		SentHash: hashcode,
	}

	msg := email.Confirm(key)
	msg.Fname = fname + " confirm"

	if key == "shash" {
		http.Redirect(w, r, "http://www.spk-saryarka.kz/", 301)
		return
	}

	utils.Respond(w, r, msg)
}



