package message

import (
	"encoding/json"
	logr "github.com/sirupsen/logrus"
	"invest/utils/constants"
	"invest/utils/logist"

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

func (msg *Msg) IsThereAnError() bool {
	return msg.ErrMsg != ""
}

func (msg *Msg) SetFname(fname string, appendix string) {
	msg.Fname = fname + " " + appendix
}

/*
	note: the order, how headers are set, matters
		1. headers
		2. status
		3. encoder - body of response
 */
func Respond(w http.ResponseWriter, r *http.Request, msg Msg) {
	var fname = "RESPOND"

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Origin")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, HEAD, OPTIONS, PATCH, POST, PUT")
	w.Header().Set("Content-Type", "application/json")

	if msg.Status == 0 {
		msg.Status = 200
	}

	w.Header().Set(constants.HeaderCustomStatus, strconv.Itoa(msg.Status))

	/*HEAD
		this header will bear a auth token
	 */
	w.Header().Add(constants.HeaderAuthorization, r.Header.Get(constants.HeaderAuthorization))

	w.WriteHeader(msg.Status)

	//fmt.Println("w.WriteHeader: ", w.Header(), msg.Status)

	//msg.Message["error"] = msg.ErrMsg

	if err := json.NewEncoder(w).Encode(msg.Message); err != nil {
		logist.SysMessage{
			FuncName: fname,
			Message:  err.Error(),
			Ok:       false,
			Lev:      constants.WarnLevel,
		}.Log_system_message()
		return
	}

	//fmt.Println("msg: ", msg)
	msg.Message = nil
	msg.Log(r)
}
