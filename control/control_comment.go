package control

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"invest/model"
	"invest/utils"
	"net/http"
	"os"
	"strconv"
	"strings"
)

/*
	is_docked - true if a document (e.g. pdf) is attached
	project_id
	subject
	body - the comment itself

	id - from session token
 */
var Add_comment_to_project = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Add_comment_to_project"

	if err := r.ParseMultipartForm(0); err != nil {
		utils.Respond(w, r, &utils.Msg{utils.ErrorInvalidParameters, 400, fname + " 1.0", err.Error()})
		return
	}

	docked := r.FormValue("is_docked")
	is_docked := utils.If_condition_then(strings.ToLower(docked) == "true", true, false).(bool)

	var ds = DocStore{}
	if is_docked {
		resp, err := ds.Download_and_store_file(r)
		if err != nil {
			utils.Respond(w, r, &utils.Msg{resp, 400, fname + " 2", err.Error()})
			return
		}
	}

	path := ds.Directory + ds.Filename + ds.Format

	var id, pid int64
	id, err := strconv.ParseInt(r.Header.Get(utils.KeyId), 0, 64)
	pid, err2 := strconv.ParseInt(r.FormValue("project_id"), 0, 64)

	if err != nil || err2 != nil {
		if path != "" {
			fmt.Println(os.Remove("." + path))
		}
		utils.Respond(w, r, &utils.Msg{
			Message: utils.ErrorInvalidParameters,
			Status:  400,
			Fname:   fname + " 1",
			ErrMsg:  "invalid id or pid number",
		})
		return
	}

	var comment = model.Comment{
		Subject:    	r.FormValue("subject"),
		Body:       	r.FormValue("body"),
		UserId:     	uint64(id),
		ProjectId:  	uint64(pid),
		DocumentUrl: 	ds.Directory + ds.Filename + ds.Format,
	}

	var errmsg string
	resp, err := comment.Create_comment_after_saving_its_document()
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

var Get_comments_of_the_project = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Get_comments_of_the_project"
	var vars, ok = mux.Vars(r)["project_id"]

	var errmsg string
	var err error
	var resp map[string]interface{}

	if ok && len(vars) > 0 {
		var c = model.Comment{}
		resp, err = c.Get_all_comments_to_the_project()
	} else {
		resp, err = utils.ErrorInvalidParameters, errors.New("parameters are not valid. get comment")
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