package control

import (
	"encoding/json"
	"errors"
	"invest/model"
	"invest/service"
	"invest/utils"

	"net/http"
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

	// get header values (id & role)
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// parse form-data
	if err := r.ParseMultipartForm(0); err != nil {
		utils.Respond(w, r, utils.Msg{utils.ErrorInvalidParameters, 400, fname + " 1.0", err.Error()})
		return
	}

	// project_id & user_id are ready
	project_id, err := strconv.ParseInt(r.FormValue("project_id"), 10, 64)
	if err != nil { project_id = 0 }

	// check permission - whether this user is allowed to upload a document
	var project = model.Project{
		Id: uint64(project_id),
		OfferedById: is.UserId,
	}
	if err = project.OnlyCheckInvestorByProjectAndInvestorId(model.GetDB()); err != nil {
		utils.Respond(w, r, utils.Msg{utils.ErrorMethodNotAllowed, 405, fname + " role", err.Error()})
		return
	}

	/* get values of fields of struct 'document' */
	var document = model.Document{}
	ganta_id, err := strconv.ParseInt(r.FormValue("ganta_id"), 10, 64)
	if err != nil { ganta_id = 0 }

	document.ProjectId = uint64(project_id)
	document.GantaId = uint64(ganta_id)

	var ds = service.DocStore{}
	_, err = ds.Download_and_store_file(r)
	if err != nil { utils.Respond(w, r, utils.Msg{utils.ErrorInvalidParameters, 400, fname + " doc", err.Error()}); return }

	// clear the field
	ds.ContentBytes = make([]byte, 1)

	// set document uri
	document.Uri = ds.Directory + ds.Filename + ds.Format
	document.Created = utils.GetCurrentTime()

	// store on db
	msg := is.Add_documents_to_project(&document)
	msg.Fname = fname + " add"

	utils.Respond(w, r, msg)
}

var Project_remove_document = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Project_remove_document"
	
	// parse request body
	var document = model.Document{}
	if err := json.NewDecoder(r.Body).Decode(&document); err != nil {
		utils.Respond(w, r, utils.Msg{utils.ErrorInvalidParameters, 400, fname + " 1", err.Error()})
	}
	defer r.Body.Close()

	// parse header
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// check for permission
	var err error
	var project = model.Project{Id: document.ProjectId, OfferedById: is.UserId}

	switch {
	case is.RoleName == utils.RoleInvestor:
		err = project.OnlyCheckInvestorByProjectAndInvestorId(model.GetDB())
	case is.RoleName == utils.RoleAdmin:
		err = errors.New("admins cannot delete a document")
	default:
		err = project.OnlyCheckUserByProjectAndUserId(document.ProjectId, is.UserId, model.GetDB())
	}

	// delete a document by doc & project id
	var msg = utils.Msg{}
	switch {
	case err != nil:
		msg := model.ReturnMethodNotAllowed(err.Error())
		msg.Fname = fname + " err"
	default:
		var document = model.Document{Id: document.Id, ProjectId: document.ProjectId}
		msg := document.Remove_document_based_on_responsibility(model.GetDB())
		msg.Fname = fname + " 2"
	}

	utils.Respond(w, r, msg)
}

