package control

import (
	"encoding/json"
	"errors"
	"invest/model"
	"invest/utils"
	"net/http"
)

var Categors_create_read_update_delete = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Categors_create_read_update_delete"

	var errmsg string
	var resp = utils.ErrorInternalServerError
	var err = errors.New("no of the switch cases. categor crud")

	var c = model.Categor{}

	if r.Method != http.MethodGet {
		if err = json.NewDecoder(r.Body).Decode(&c); err != nil {
			utils.Respond(w, r, &utils.Msg{
				Message: resp,
				Status:  400,
				Fname:   fname + " 1",
				ErrMsg:  err.Error(),
			})
			return
		}
		defer r.Body.Close()
	}

	switch r.Method {
	case http.MethodGet:
		var offset = Get_query_parameter_str(r, "offset", "0")
		resp, err = c.Get_all_categors(offset)

	case http.MethodPost:
		resp, err = c.Create_category()

	//case http.MethodPut:
	//	resp, err = c.Update()

	case http.MethodDelete:
		resp, err = c.Delete_category_from_tabe_and_projects()

	default:
		resp, err = utils.ErrorMethodNotAllowed, errors.New("not supported")
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
