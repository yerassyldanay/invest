package utils

import (
	"fmt"
	"net/http"
)

const (
	Error_msg_invalid_parameters_passed = "invalid parameters have been passed"
	Error_msg_internal_db_error = "internal db error has occured"
	Error_msg_method_not_allowed = "method is not allowed"
)

var (
	MsgNoErrorMessageOk = Msg{
		NoErrorFineEverthingOk, http.StatusOK, "", "",
	}
	MsgNoErrorMessageCreated = Msg{
		NoErrorFineEverthingOk, http.StatusCreated, "", "",
	}
)

func OnlyPrintQueueIsFull(where string) {
	fmt.Println("somehow the queue is full. where: ", where)
}
