package model

import (
	"github.com/yerassyldanay/invest/utils/errormsg"
	"github.com/yerassyldanay/invest/utils/message"
	"net/http"
)

func ReturnInternalDbError(errmsg string) message.Msg {
	return message.Msg{errormsg.ErrorInternalDbError, 417, "", errmsg}
}

func ReturnInternalServerError(errmsg string) message.Msg {
	return message.Msg {
		Message: errormsg.ErrorInternalServerError, Status:  http.StatusInternalServerError, ErrMsg:  errmsg,
	}
}

func ReturnWrongPassword(errmsg string) message.Msg {
	return message.Msg{errormsg.ErrorWrongPassword, 417, "", errmsg}
}

func ReturnInvalidPassword(errmsg string) message.Msg {
	return message.Msg{errormsg.ErrorInvalidPassword, 417, "", errmsg}
}

func ReturnNoError() (message.Msg) {
	return message.Msg{errormsg.NoErrorFineEverthingOk, 200, "", ""}
}

func ReturnSuccessfullyCreated() (message.Msg) {
	return message.Msg{errormsg.NoErrorFineEverthingOk, 201, "", ""}
}

func ReturnInvalidParameters(errmsg string) (message.Msg) {
	return message.Msg{errormsg.ErrorInvalidParameters, 400, "", errmsg}
}

func ReturnEmailAlreadyInUse(errmsg string) (message.Msg) {
	return message.Msg{errormsg.ErrorEmailIsAreadyInUse, 409, "", errmsg}
}

func ReturnEmailAlreadyInUseOrCodeExpired(errmsg string) (message.Msg) {
	return message.Msg{errormsg.ErrorEmailIsAreadyInUseOrCodeExpired, 409, "", errmsg}
}

func ReturnEmailIsNotVerified(errmsg string) (message.Msg) {
	return message.Msg{errormsg.ErrorEmailIsNotVerified, 406, "", errmsg}
}

func ReturnDuplicateKeyError(errmsg string) (message.Msg) {
	return message.Msg{errormsg.ErrorDupicateKeyOnDb, 409, "", errmsg}
}

func ReturnNoErrorWithResponseMessage(resp map[string]interface{}) message.Msg {
	return message.Msg{resp, 200, "", ""}
}

func ReuturnInternalServerError(errmsg string) message.Msg {
	return message.Msg{errormsg.ErrorInternalServerError, 500, "", errmsg}
}

func ReturnNotFoundError(errmsg string) message.Msg {
	return message.Msg{errormsg.ErrorNotFound, 400, "", errmsg}
}

func ReturnMethodNotAllowed(errmsg string) message.Msg {
	return message.Msg{errormsg.ErrorMethodNotAllowed, 405, "", errmsg}
}

func ReturnNoSuchUser(errmsg string) message.Msg {
	return message.Msg{errormsg.ErrorNoSuchUser, 404, "", errmsg}
}

func ReturnCouldNotSendEmailError(errmsg string) (message.Msg) {
	return message.Msg{errormsg.ErrorCouldNotSendEmail, 422, "", errmsg}
}

func ReturnFailedToCreateAnAccount(errmsg string) (message.Msg) {
	return message.Msg{errormsg.ErrorFailedToCreateAnAccount, http.StatusExpectationFailed, "", errmsg}
}
