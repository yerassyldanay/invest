package app

import (
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/utils/constants"
	"github.com/yerassyldanay/invest/utils/message"

	"net/http"
	"strconv"
)

/*
	get info for personal info
 */
var User_get_own_info = func(w http.ResponseWriter, r *http.Request) {
	var fname = "User_get_own_info"
	var id, _ = strconv.ParseInt(r.Header.Get(constants.KeyId), 0, 16)

	var user = model.User{
		Id: uint64(id),
	}

	msg := user.GetFullInfoOfThisUser(model.ElementGetFullInfoOfThisUser{})
	msg.Fname = fname + " 1"

	message.Respond(w, r, msg)
}