package app

import (
	"encoding/json"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/service"
	"github.com/yerassyldanay/invest/utils/constants"
	"github.com/yerassyldanay/invest/utils/errormsg"
	"github.com/yerassyldanay/invest/utils/message"

	"net/http"
)

var Update_organization_data = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Update_organization_data"
	var org = model.Organization{}

	if err := json.NewDecoder(r.Body).Decode(&org); err != nil {
		message.Respond(w, r, message.Msg{
			Message: errormsg.ErrorInvalidParameters,
			Status:  400,
			Fname:   fname + " 1",
			ErrMsg:  err.Error(),
		})
		return
	}
	defer r.Body.Close()

	org.Lang = r.Header.Get(constants.HeaderContentLanguage)

	msg := org.Update_organization_info(model.GetDB())
	message.Respond(w, r, msg)
}

var Get_organization_info_by_bin = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Get_organization_info_by_bin"
	var bin = service.Get_query_parameter_str(r, "bin", "")
	var msg = message.Msg{
		errormsg.ErrorInvalidParameters, http.StatusBadRequest, fname + " 1", "invalid parameters. invalid bin number",
	}

	var org = &model.Organization{
		Lang: r.Header.Get(constants.HeaderContentLanguage),
		Bin:  bin,
	}
	
	if bin != "" {
		msg = org.Create_or_get_organization_from_db_by_bin(model.GetDB())
	}

	message.Respond(w, r, msg)
}
