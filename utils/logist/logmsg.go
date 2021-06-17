package logist

import (
	logr "github.com/sirupsen/logrus"
	"github.com/yerassyldanay/invest/utils/constants"
)

/*
	this struct will be used further to log
		system messages into log files
	it makes the logs strict & controllable
 */
type SysMessage struct {
	FuncName					string
	Message						string
	Ok							bool
	Lev							int
}

func (messageToSend SysMessage) Log_system_message() {

	f := logr.WithFields(map[string]interface{}{
		"message": messageToSend.Message,
		"ok":      messageToSend.Ok,
	})

	switch messageToSend.Lev {
	case constants.InfoLevel:
		f.Info(messageToSend.FuncName)
	case constants.WarnLevel:
		f.Warn(messageToSend.FuncName)
	case constants.FatalLevel:
		f.Fatal(messageToSend.FuncName)
	case constants.DebugLevel:
		f.Debug(messageToSend.FuncName)
	default:
		f.Log(logr.Level(messageToSend.Lev), messageToSend)
	}
}
