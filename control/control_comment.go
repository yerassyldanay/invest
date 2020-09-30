package control

import (
	"encoding/json"
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

var Add_comment_to_project = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Add_comment_to_project"
	var comment = model.Comment{}

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		msg := utils.Msg{utils.ErrorInvalidParameters, 400, fname + " 1", err.Error()}
		utils.Respond(w, r, msg)
		return
	}
	defer r.Body.Close()

	var is = service.InvestService{
		Offset: "",
		BasicInfo: service.BasicInfo {
			UserId: service.Get_header_parameter(r, utils.KeyId, uint64(0)).(uint64),
		},
	}

	comment.UserId = is.UserId

	msg := is.Comment_on_project_documents(comment)
	msg.Fname = fname + " 2"

	utils.Respond(w, r, msg)
}

/*
	provide
		* project_id
 */
var Get_comments_of_the_project_or_comment_by_comment_id = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Get_comments_of_the_project"

	comment_id := service.Get_query_parameter_uint64(r, "comment_id", 0)
	project_id := service.Get_query_parameter_uint64(r, "project_id", 0)
	offset := service.Get_query_parameter_str(r, "offset", "0")

	var c = model.Comment{
		ProjectId: 		project_id,
		Id: 			comment_id,
	}

	var msg = utils.Msg{}
	if c.Id != 0 {
		msg = c.Get_comment_by_comment_id()
	} else {
		msg = c.Get_all_comments_of_the_project_by_project_id(offset)
	}

	msg.Fname = fname + " 1"
	utils.Respond(w, r, msg)
}
