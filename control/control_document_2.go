package control

import (
	"encoding/json"
	"fmt"
	"invest/model"
	"invest/service"
	"invest/utils"
	"os"
	"path/filepath"

	"net/http"
)

/*
	store docs on disc & info on db

	provide:
		* project_id
		* document_id
		*
*/
var Document_upload_document = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Project_add_document_to_project"

	// get header values (id & role)
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// parse form-data
	if err := r.ParseMultipartForm(0); err != nil {
		utils.Respond(w, r, utils.Msg{utils.ErrorInvalidParameters, 400, fname + " 1.0", err.Error()})
		return
	}

	// project_id & user_id are ready
	project_id := service.OnlyConvertString(r.FormValue("project_id"), uint64(0)).(uint64)
	document_id := service.OnlyConvertString(r.FormValue("document_id"), uint64(0)).(uint64)

	// security - user must have an access to the project
	msg := is.Check_whether_this_user_can_get_access_to_project_info(project_id)
	if msg.IsThereAnError() {
		msg.SetFname(fname, "access project")
		utils.Respond(w, r, msg)
		return
	}

	// security #2 - is this user allowed to upload a document - each document has a responsible role
	msg = is.Check_whether_this_user_is_responsible_for_document(document_id, project_id)
	if msg.IsThereAnError() {
		msg.SetFname(fname, "access doc")
		utils.Respond(w, r, msg)
		return
	}

	// get values of fields of struct 'document'
	var document = model.Document{
		Id:        document_id,
		ProjectId: project_id,
	}

	// download file
	var ds = service.DocStore{}
	_, err := ds.Download_and_store_file(r)
	if err != nil {
		utils.Respond(w, r, utils.Msg{utils.ErrorInvalidParameters, 400, fname + " doc", err.Error()})
		return
	}

	// clear the field
	ds.ContentBytes = nil

	// set document uri
	path := ds.Directory + ds.Filename + ds.Format
	document.Uri = path
	document.Modified = utils.GetCurrentTime()

	// store on db
	msg = is.Upload_documents_to_project(&document)
	msg.Fname = fname + " add"

	// remove file in case of an error
	if msg.IsThereAnError() {
		path, err = filepath.Abs("./" + path)
		err = os.Remove(path)
		fmt.Println(err)
	}

	utils.Respond(w, r, msg)
}

// remove document & delete a file
var Document_remove_file = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Project_remove_document"
	
	/*
		parse request body:
			* document_id
			* project_id
	 */
	var document = model.Document{}
	if err := json.NewDecoder(r.Body).Decode(&document); err != nil {
		utils.Respond(w, r, utils.Msg{utils.ErrorInvalidParameters, 400, fname + " 1", err.Error()})
	}
	defer r.Body.Close()

	// parse header
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	/*
		Security:
			* check that exactly this user can access project
			* is this user responsible for current step
	 */
	msg := is.Check_whether_this_user_can_get_access_to_project_info(document.ProjectId)
	if msg.IsThereAnError() {
		msg.SetFname(fname, "project")
		utils.Respond(w ,r, msg)
		return
	}

	msg = is.Check_whether_this_user_responsible_for_current_step(document.ProjectId)
	if msg.IsThereAnError() {
		msg.SetFname(fname, "step")
		utils.Respond(w, r, msg)
		return
	}

	// delete a document by doc & project id
	msg = is.Document_remove_document_from_project(document.Id)
	msg.SetFname(fname, "remove")

	utils.Respond(w, r, msg)
}

