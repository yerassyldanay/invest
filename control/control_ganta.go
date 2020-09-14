package control

import (
	"encoding/json"
	"invest/model"
	"invest/utils"
	"net/http"
)

var Ganta_create_update_delete = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Ganta_create_update_delete"
	var ganta = model.Ganta{}

	if err := json.NewDecoder(r.Body).Decode(&ganta); err != nil {
		utils.Respond(w, r, utils.Msg{
			Message: utils.ErrorInvalidParameters,
			Status:  400,
			Fname:   fname + " 1",
			ErrMsg:  err.Error(),
		})
		return
	}
	defer r.Body.Close()

	var msg = utils.Msg{}
	switch r.Method {
	case http.MethodPost:
		msg = ganta.Add_new_step()
		msg.Fname = fname + " post"
	case http.MethodPut:
		msg = ganta.Update_ganta_step()
		msg.Fname = fname + " put"
	case http.MethodDelete:
		msg.Fname = fname + " delete"
	default:
		msg = utils.Msg{utils.ErrorMethodNotAllowed, 405, "", "making a not supported request: " + r.Method}
	}

	utils.Respond(w, r, msg)
}

/*
	get only ganta steps
 */
var Ganta_only_ganta_steps_by_project_id = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Ganta_get_all_steps_by_project_id"
	var choice = Get_query_parameter_str(r, "choice", "")
	var ganta = model.Ganta{
		ProjectId: uint64(Get_query_parameter_int(r, "project_id", 0)),
	}

	var msg = utils.Msg{}
	switch choice {
	case "withdoc":
		msg = ganta.Get_ganta_with_documents_by_project_id()
		msg.Fname = fname + " docs"
	case "one":
		msg = ganta.Get_only_one_with_docs()
		msg.Fname = fname + " one"
	default:
		msg = ganta.Get_only_ganta_by_project_id()
		msg.Fname = fname + " default"
	}

	utils.Respond(w, r, msg)
}

