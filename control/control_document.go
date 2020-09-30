package control

import (
	"invest/model"
	"invest/service"
	"invest/utils"
	"net/http"
)

var Get_stat_on_documents_of_project = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Get_number_of_documents"
	var document = model.Document{
		ProjectId: service.Get_query_parameter_uint64(r, "project_id", 0),
	}

	msg := document.Get_stat_on_docs_by_project_id()
	msg.Fname = fname + " 1"

	utils.Respond(w, r, msg)
}
