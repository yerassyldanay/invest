package control

import (
	"encoding/json"
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
	defer r.Body.Close()

	org.Lang = r.Header.Get(utils.HeaderContentLanguage)

	msg := org.Update_organization_info()
	utils.Respond(w, r, msg)
}

var Get_organization_info_by_bin = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Get_organization_info_by_bin"
	var bin = Get_query_parameter_str(r, "bin", "")
	var msg = &utils.Msg{
		utils.ErrorInvalidParameters, http.StatusBadRequest, fname + " 1", "invalid parameters. invalid bin number",
	}

	var org = &model.Organization{
		Lang: r.Header.Get(utils.HeaderContentLanguage),
		Bin:  bin,
	}
	
	if bin != "" {
		msg = org.Create_or_get_organization_from_db_by_bin()
	}

	utils.Respond(w, r, msg)
}
