package control

import (
	"encoding/json"
	"invest/model"
	"invest/utils/message"
	"net/http"
)

// handle parsing of a request body in one place
func OnlyParseRequestBody(r *http.Request, to interface{}) (interface{}, message.Msg) {
	if err := json.NewDecoder(r.Body).Decode(&to); err != nil {
		return to, model.ReturnInvalidParameters(err.Error())
	}
	defer r.Body.Close()

	return to, model.ReturnNoError()
}

