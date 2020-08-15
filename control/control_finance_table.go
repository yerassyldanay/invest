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
var Finance_table = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Finance_table"
	var finance = model.Finance{}

	var errmsg string
	var resp = make(map[string]interface{})

	if err := json.NewDecoder(r.Body).Decode(&finance); err == nil {
		/*
			create or update finance table in one place
				no need to repeat the same code
		 */
		switch r.Method {
		case http.MethodPost:
			resp, err = finance.Create_and_store_on_db()
		case http.MethodPut:
			resp, err = finance.Update_finance_table()
		}

		if err != nil {
			errmsg = err.Error()
		}
	} else {
		utils.Respond(w, r, &utils.Msg{
			Message: utils.ErrorInvalidParameters,
			Status:  400,
			Fname:   fname + " 1",
			ErrMsg:  err.Error(),
		})
		return
	}

	utils.Respond(w, r, &utils.Msg{
		Message: resp,
		Status:  utils.If_condition_then(errmsg == "", 200, 400).(int),
		Fname:   fname,
		ErrMsg:  errmsg,
	})
}

/*

 */
var Finance_table_get = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Finance_table_get"
	var finance = model.Finance{}

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
	resp, err := finance.Get_table()
	if err != nil {errmsg = err.Error()}

	utils.Respond(w, r, &utils.Msg{
		Message: resp,
		Status:  utils.If_condition_then(errmsg == "", 200, 400).(int),
		Fname:   fname + " 2",
		ErrMsg:  errmsg,
	})
}
