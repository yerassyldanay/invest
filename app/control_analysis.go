package app

import (
	"encoding/json"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/service"
	"github.com/yerassyldanay/invest/utils/constants"
	"github.com/yerassyldanay/invest/utils/message"
	"net/http"
)

var Analysis_get_help = func(which string, w http.ResponseWriter, r *http.Request) {
	var fname = "Analysis_get"

	var analysis = model.Analysis{}
	if err := json.NewDecoder(r.Body).Decode(&analysis); err != nil {
		OnlyReturnInvalidParametersError(w, r, err.Error(), fname, "json")
		return
	}
	defer r.Body.Close()

	// headers
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// security check
	if is.RoleName == constants.RoleInvestor {
		OnlyReturnMethodNotAllowed(w, r,"your role is investor", fname, "role")
		return
	}

	// should write to file?
	switch which {
	case "file":
		analysis.WriteToFile = true
	}

	// logic
	msg := is.Analysis_get_on_projects(analysis)
	msg.SetFname(fname, "analysis")

	message.Respond(w, r, msg)
}

// get xls file
var Analysis_get_file = func(w http.ResponseWriter, r *http.Request) {
	Analysis_get_help("file", w, r)
}

var Analysis_get = func(w http.ResponseWriter, r *http.Request) {
	Analysis_get_help("stat", w, r)
}
