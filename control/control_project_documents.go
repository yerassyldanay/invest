package control

import (
	"encoding/json"
	"fmt"
	"invest/model"
	"invest/utils"
	"net/http"
	"os"
	"strconv"
)

/*
	store docs on disc & info on db

	provide:
		* project_id
		* document_id
		*
*/
var Project_add_document_to_project = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Project_add_document_to_project"
	var document = model.Document{}

	if err := r.ParseMultipartForm(0); err != nil {
		utils.Respond(w, r, utils.Msg{utils.ErrorInvalidParameters, 400, fname + " 1.0", err.Error()})
		return
	}

	/*
		parse data:
			* name
			* project_id
			* ganta_id
	*/
	document.Name = r.FormValue("name")
	i, err := strconv.ParseInt(r.FormValue("project_id"), 10, 64)
	if err != nil {
		i = 0
	}
	j, err := strconv.ParseInt(r.FormValue("ganta_id"), 10, 64)
	if err != nil {
		j = 0
	}

	document.ProjectId = uint64(i)
	document.GantaId = uint64(j)

	id := Get_header_parameter(r, utils.KeyId, uint64(0)).(uint64)

	/*
		CHECK TIME:
		check that a person is an investor, who created this project
	*/
	var project = model.Project{Id: document.ProjectId}
	if err := project.Get_by_id(model.GetDB()); err != nil || project.OfferedById != id {
		utils.Respond(w, r, utils.Msg{utils.ErrorMethodNotAllowed, 405, fname + " project", "this is not an investor of the project"})
		return
	}

	/*
		you cannot add document
			* to parent ganta step
			* if a ganta step already has a document
	*/
	var ganta = model.Ganta{Id: document.GantaId}
	_ = ganta.Get_by_id(model.GetDB())

	if ganta.GantaParentId == 0 || ganta.Does_this_ganta_step_has_document(model.GetDB()) {
		utils.Respond(w, r, utils.Msg{utils.ErrorMethodNotAllowed, 405, fname + " ganta", "the ganta step already possesses a document"})
		return
	}

	var ds = DocStore{}
	resp, err := ds.Download_and_store_file(r)
	if err != nil {
		utils.Respond(w, r, utils.Msg{resp, 400, fname + " 2", err.Error()})
		return
	}

	document.Uri = ds.Directory + ds.Filename + ds.Format
	msg := document.Add()
	msg.Fname = fname + " 3"

	if msg.ErrMsg != "" {
		/*
			remove file in case of an error
		*/
		if err := os.Remove("." + document.Uri); err != nil {
			fmt.Println(err.Error())
		}
	}

	utils.Respond(w, r, msg)
}

var Project_remove_document = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Project_remove_document"
	var document = model.Document{}

	if err := json.NewDecoder(r.Body).Decode(&document); err != nil {
		utils.Respond(w, r, utils.Msg{utils.ErrorInvalidParameters, 400, fname + " 1", err.Error()})
	}

	id := Get_header_parameter(r, utils.KeyId, uint64(0)).(uint64)
	var project = model.Project{Id: document.ProjectId}
	_ = project.Get_by_id(model.GetDB())

	if project.OfferedById != id {
		utils.Respond(w, r, utils.Msg{utils.ErrorMethodNotAllowed, 405, fname + " project", "not investor of the project"})
		return
	}

	document.ChangesMadeById = Get_header_parameter(r, utils.KeyId, uint64(0)).(uint64)

	msg := document.Remove()
	msg.Fname = fname + " 2"

	utils.Respond(w, r, msg)
}

