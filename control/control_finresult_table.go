package control

import (
	"encoding/json"
	"invest/model"
	"invest/utils"
	"net/http"
)

var Finresult_and_project_evaluation_table = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Finresult_and_project_evaluation_table"
	var finance = model.Finresult{}

	if err := json.NewDecoder(r.Body).Decode(&finance); err != nil {
		utils.Respond(w, r, &utils.Msg{
			Message: utils.ErrorInvalidParameters,
			Status:  400,
			Fname:   fname + " 1",
			ErrMsg:  err.Error(),
		})
		return
	}

	var errmsg string
	var resp = make(map[string]interface{})
	var err error

	switch r.Method {
	case http.MethodPost:
		resp, err = finance.Create_and_store_financial_results_of_project_on_db()
	case http.MethodPut:
		resp, err = finance.Update_finresult_table()
	}

	if err != nil {
		errmsg = err.Error()
	}

	utils.Respond(w, r, &utils.Msg{
		Message: resp,
		Status:  utils.If_condition_then(errmsg == "", 200, 400).(int),
		Fname:   fname,
		ErrMsg:  errmsg,
	})
}

var Finresult_table_get = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Finresult_table_get"
	var finres = model.Finresult{}

	if err := json.NewDecoder(r.Body).Decode(&finres); err != nil {
		utils.Respond(w, r, &utils.Msg{
			Message: utils.ErrorInvalidParameters,
			Status:  400,
			Fname:   fname + " 1",
			ErrMsg:  err.Error(),
		})
		return
	}

	var errmsg string
	resp, err := finres.Get_finresult_table()
	if err != nil {
		errmsg = err.Error()
	}

	utils.Respond(w, r, &utils.Msg{
		Message: resp,
		Status:  utils.If_condition_then(errmsg == "", 200, 400).(int),
		Fname:   fname,
		ErrMsg:  errmsg,
	})
}
