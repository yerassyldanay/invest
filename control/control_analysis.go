package control

import (
	"encoding/json"
	"invest/model"
	"invest/service"
	"invest/utils"
	"net/http"
)

var Analysis_get = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Analysis_get"

	var analysis = model.Analysis{}
	if err := json.NewDecoder(r.Body).Decode(&analysis); err != nil {
		OnlyReturnInvalidParametersError(w, r, err.Error(), fname, "json")
		return
	}

	// headers
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// security check
	if is.RoleName == utils.RoleInvestor {
		OnlyReturnMethodNotAllowed(w, r,"your role is investor", fname, "role")
		return
	}

	// logic
	msg := is.Analysis_get_on_projects(analysis)
	msg.SetFname(fname, "analysis")

	utils.Respond(w, r, msg)
}
