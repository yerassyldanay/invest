package control

import (
	"encoding/json"
	"invest/model"
	"invest/utils"
	"net/http"
)

/*
	Land					uint64								`json:"land"`
	Tech					uint64								`json:"tech"`
	Capital					uint64								`json:"capital"`
	Other					uint64								`json:"other"`
	Sum						uint64								`json:"sum"`
 */
/*
	update - delete & recreate
 */
var Finance_table_update = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Finance_table"
	var finance = model.Finance{}

	var msg = utils.Msg{}

	if err := json.NewDecoder(r.Body).Decode(&finance); err != nil {
		msg = utils.Msg{utils.ErrorInvalidParameters, 400, fname + " 1", err.Error() }
		utils.Respond(w, r, msg)
		return
	}

	/*
		create or update finance table in one place
			no need to repeat the same code
	 */
	switch r.Method {
	case http.MethodPut:
		msg = finance.Update_finance_table_with_this_table_by_project_id()
	default:
		msg = utils.Msg{utils.ErrorMethodNotAllowed, 405, fname + " 2", "method not allowed. crud finance table"}
	}
	defer r.Body.Close()

	utils.Respond(w, r, msg)
}

/*
	provide:
		* project_id
 */
var Finance_table_get = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Finance_table_get"
	var finance = model.Finance{
		ProjectId: Get_query_parameter_uint64(r, "project_id", 0),
	}

	var msg = finance.Get_table()
	msg.Fname = fname + " 1"

	utils.Respond(w, r, msg)
}
