package utils

import (
	logr "github.com/sirupsen/logrus"
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
	case InfoLevel:
		f.Info(messageToSend.FuncName)
	case WarnLevel:
		f.Warn(messageToSend.FuncName)
	case FatalLevel:
		f.Fatal(messageToSend.FuncName)
	case DebugLevel:
		f.Debug(messageToSend.FuncName)
	default:
		f.Log(logr.Level(messageToSend.Lev), messageToSend)
	}
}
