package control

import (
	"encoding/json"
	"errors"
	"invest/model"
	"invest/utils"
	"net/http"
)

var Update_organization_data = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Update_organization_data"
	var org = model.Organization{}

	if err := json.NewDecoder(r.Body).Decode(&org); err != nil {
		utils.Respond(w, r, &utils.Msg{
			Message: 	utils.ErrorInvalidParameters,
			Status:  	400,
			Fname:   	fname + " 1",
			ErrMsg:  	err.Error(),
		})
		return
	}

	org.Lang = r.Header.Get(utils.HeaderContentLanguage)

	var errmsg string
	resp, err := org.Update_organization_info()
	if err != nil {
		errmsg = err.Error()
	}

	utils.Respond(w, r, &utils.Msg{
		Message: resp,
		Status:  utils.If_condition_then(errmsg == "", 200, 400).(int),
		Fname:   fname + " 2",
		ErrMsg:  errmsg,
	})
}

var Get_organization_info_by_bin = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Get_organization_info_by_bin"
	var bin = Get_query_parameter_str(r, "bin", "")

	var err = errors.New("invalid parameters. invalid bin number")
	var resp = utils.ErrorInvalidParameters
	var errmsg string

	var org = &model.Organization{
		Lang: r.Header.Get(utils.HeaderContentLanguage),
		Bin:  bin,
	}
	
	if bin != "" {
		resp, err = org.Create_or_get_organization_from_db_by_bin()
	}

	if err != nil {
		errmsg = err.Error()
	}

	utils.Respond(w, r, &utils.Msg{
		Message: resp,
		Status:  utils.If_condition_then(errmsg == "", 200, 400).(int),
		Fname:   fname + " 1",
		ErrMsg:  errmsg,
	})
}
