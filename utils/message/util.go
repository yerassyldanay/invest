package message

import "net/http"

type LogMessage struct {
	Ok					bool
	FuncName			string
	Message 			string
	Req					*http.Request
}

func Is_any_of_these_nil(elems... interface{}) bool {
	for _, elem := range elems {
		if elem == nil {
			return true
		}
	}

	return false
}

