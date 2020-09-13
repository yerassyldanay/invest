package control

import (
	"encoding/json"
	"invest/model"
	"invest/utils"
	"net/http"
)

var Finresult_table_update = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Finresult_table_update"
	var finance = model.Finresult{}

	if err := json.NewDecoder(r.Body).Decode(&finance); err != nil {
		utils.Respond(w, r, utils.Msg{
			Message: utils.ErrorInvalidParameters,
			Status:  400,
			Fname:   fname + " 1",
			ErrMsg:  err.Error(),
		})
		return
	}
	defer r.Body.Close()

	var msg = utils.Msg{
		Fname: fname + " 1",
	}

	switch r.Method {
	//case http.MethodPost:
	//	resp, err = finance.Create_and_store_financial_results_of_project_on_db()
	case http.MethodPut:
		msg = finance.Update_this_table()
	default:
		msg = utils.Msg{utils.ErrorMethodNotAllowed, 405, "", "not allowed. fin result put/post"}
	}

	utils.Respond(w, r, msg)
}

var Finresult_table_get = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Finresult_table_get"
	var finres = model.Finresult{
		ProjectId: Get_query_parameter_uint64(r, "project_id", 0),
	}

	var msg = finres.Get_finresult_table()
	msg.Fname = fname + " 1"

	utils.Respond(w, r, msg)
}
