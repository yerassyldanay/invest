package control

import (
	"encoding/json"
	"fmt"
	"invest/model"
	"invest/service"
	"invest/utils"

	"net/http"
)

/*
	is_docked - true if a document (e.g. pdf) is attached
	project_id
	subject
	body - the comment itself

	id - from session token
 */
//var Add_comment_to_project = func(w http.ResponseWriter, r *http.Request) {
//	var fname = "Add_comment_to_project"
//
//	if err := r.ParseMultipartForm(0); err != nil {
//		utils.Respond(w, r, utils.Msg{utils.ErrorInvalidParameters, 400, fname + " 1.0", err.Error()})
//		return
//	}
//
//	docked := r.FormValue("is_docked")
//	is_docked := utils.If_condition_then(strings.ToLower(docked) == "true", true, false).(bool)
//
//	var ds = DocStore{}
//	if is_docked {
//		resp, err := ds.Download_and_store_file(r)
//		if err != nil {
//			utils.Respond(w, r, utils.Msg{resp, 400, fname + " 2", err.Error()})
//			return
//		}
//	}
//
//	path := ds.Directory + ds.Filename + ds.Format
//
//	id := Get_header_parameter(r, utils.KeyId, uint64(0)).(uint64)
//	project_id, err := strconv.ParseInt(r.FormValue("project_id"), 10, 64)
//	if err != nil {
//		msg := utils.Msg{utils.ErrorInvalidParameters, 400, "", err.Error()}
//		utils.Respond(w, r, msg)
//		return
//	}
//
//	/*
//		parse statuses of documents
//	 */
//	var documents = []model.Document{}
//	var docsStr = r.FormValue("documents")
//
//	if err := json.Unmarshal([]byte(docsStr), &documents); err != nil {
//		if path != "" {
//			fmt.Println(os.Remove("." + path))
//		}
//		utils.Respond(w, r, utils.Msg{utils.ErrorInvalidParameters, 400, "", err.Error()})
//		return
//	}
//
//	var comment = model.Comment{
//		Subject:    	r.FormValue("subject"),
//		Body:       	r.FormValue("body"),
//		UserId:     	uint64(id),
//		ProjectId:  	uint64(project_id),
//		DocumentUrl: 	path,
//		DocStatuses: 	documents,
//	}
//
//	var is = service.InvestService{
//		service.BasicInfo{
//			UserId: id,
//		},
//	}
//	msg := is.Comment_on_project_documents(comment)
//
//	if msg.ErrMsg != "" {
//		if path != "" {
//			fmt.Println(os.Remove("." + path))
//		}
//	}
//
//	utils.Respond(w, r, msg)
//}

var Project_comment_on_documents = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Add_comment_to_project"
	var spkComment = model.SpkComment{}

	// parse request body
	if err := json.NewDecoder(r.Body).Decode(&spkComment); err != nil {
		utils.Respond(w, r, utils.Msg{utils.ErrorInvalidParameters, 400, fname + " 1", err.Error()})
		return
	}
	defer r.Body.Close()

	// parse headers
	var is = service.InvestService{}
	is.OnlyParseRequest(r)

	/*
		Security:
			* only users, who are assigned to the project, can comment
	 */
	var project = model.Project{
		Id: spkComment.Comment.ProjectId,
	}

	if is.RoleName == utils.RoleAdmin {
		// if this is an admin then pass this point
	} else if err := project.OnlyCheckUserByProjectAndUserId(spkComment.Comment.ProjectId, is.UserId, model.GetDB()); err != nil {
		utils.Respond(w, r, utils.Msg{utils.ErrorMethodNotAllowed, 405, fname + " security", err.Error()})
		return
	}

	/*
		Security:
			* is this user responsible?
	 */
	err := project.GetAndUpdateStatusOfProject(model.GetDB())
	if err != nil {
		fmt.Println(err)
	}

	if project.CurrentStep.Responsible != is.RoleName {
		utils.Respond(w, r, utils.Msg{utils.ErrorMethodNotAllowed, 405, fname + " step", "this user is not responsible for the current step"})
		return
	}

	// the logic
	msg := is.Comment_on_project_documents(spkComment)
	msg.SetFname(fname, "2")

	utils.Respond(w, r, msg)
}

/*
	provide
		* project_id
 */
var Project_get_comments = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Get_comments_of_the_project"

	// parse
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// get query parameters
	var project_id = service.OnlyGetQueryParameter(r, "project_id", uint64(0)).(uint64)

	// security
	msg := is.Check_whether_this_user_can_get_access_to_project_info(project_id)
	if msg.ErrMsg != "" {
		msg.SetFname(fname, "acc")
		utils.Respond(w, r, msg)
		return
	}

	// get comments
	msg = is.Get_comments_of_project(project_id)
	msg.SetFname(fname, "get")

	utils.Respond(w, r, msg)
}
