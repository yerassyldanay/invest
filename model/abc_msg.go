package model

import (
	"invest/utils"
	"net/http"
)

func ReturnInternalDbError(errmsg string) utils.Msg {
	return utils.Msg{utils.ErrorInternalDbError, 417, "", errmsg}
}

func ReturnNoError() (utils.Msg) {
	return utils.Msg{utils.NoErrorFineEverthingOk, 200, "", ""}
}

func ReturnSuccessfullyCreated() (utils.Msg) {
	return utils.Msg{utils.NoErrorFineEverthingOk, 201, "", ""}
}

func ReturnInvalidParameters(errmsg string) (utils.Msg) {
	return utils.Msg{utils.ErrorInvalidParameters, 400, "", errmsg}
}

func ReturnEmailAlreadyInUse(errmsg string) (utils.Msg) {
	return utils.Msg{utils.ErrorEmailIsAreadyInUse, 409, "", errmsg}
}

func ReturnDuplicateKeyError(errmsg string) (utils.Msg) {
	return utils.Msg{utils.ErrorDupicateKeyOnDb, http.StatusConflict, "", errmsg}
}

func ReturnNoErrorWithResponseMessage(resp map[string]interface{}) utils.Msg {
	return utils.Msg{resp, http.StatusOK, "", ""}
}

func ReuturnInternalServerError(errmsg string) utils.Msg {
	return utils.Msg{utils.ErrorInternalServerError, http.StatusInternalServerError, "", errmsg}
}

func ReturnNotFoundError(errmsg string) utils.Msg {
	return utils.Msg{utils.ErrorNotFound, 400, "", errmsg}
}

func ReturnMethodNotAllowed(errmsg string) utils.Msg {
	return utils.Msg{utils.ErrorMethodNotAllowed, 405, "", errmsg}
}

func ReturnNoSuchUser(errmsg string) utils.Msg {
	return utils.Msg{utils.ErrorNoSuchUser, 404, "", errmsg}
}
