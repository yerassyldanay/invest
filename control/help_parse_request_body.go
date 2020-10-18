package control

import (
	"encoding/json"
	"invest/model"
	"invest/utils"
	"net/http"
)

// handle parsing of a request body in one place
func OnlyParseRequestBody(r *http.Request, to interface{}) (interface{}, utils.Msg) {
	if err := json.NewDecoder(r.Body).Decode(&to); err != nil {
		return to, model.ReturnInvalidParameters(err.Error())
	}

	return to, model.ReturnNoError()
}

