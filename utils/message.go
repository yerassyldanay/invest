package utils

import (
	"encoding/json"
	logr "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Msg struct {
	Message				map[string]interface{}
	Status				int
	Fname				string
	ErrMsg				string
}

/*
	this function creates for us a message struct
 */
func Message(fname string, status_to_send int, message_to_send map[string]interface{}) *Msg {
	return &Msg{
		Message: 		message_to_send,
		Status:  		status_to_send,
		Fname:			fname,
	}
}

func (msg *Msg) Log(r *http.Request) {
	logr.WithFields(map[string]interface{}{
		"path":				r.URL.Path,
		"host":				r.RemoteAddr,
		"error":			msg.ErrMsg,
		"status":			msg.Status,
	}).Info(msg.Fname)
}

/*
	note: the order, how headers are set, matters
		1. headers
		2. status
		3. encoder - body of response
 */
func Respond(w http.ResponseWriter, r *http.Request, msg *Msg) {
	var fname = "RESPOND"

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Status", strconv.Itoa(msg.Status))

	/*
		this header will bear a auth token
	 */
	w.Header().Add("Authentication", r.Header.Get("Authentication"))

	w.WriteHeader(msg.Status)

	if err := json.NewEncoder(w).Encode(msg.Message); err != nil {
		SysMessage {
			FuncName: fname,
			Message:  err.Error(),
			Ok:       false,
			Lev:      WarnLevel,
		}.Log_system_message()
		return
	}

	msg.Log(r)
}

