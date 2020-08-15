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

var Create_project = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Create_project"
	var errmsg string

	var project = struct{
		model.Project
		Bin					string 					`json:"bin"`
	}{}

	var id = utils.GetHeader(r, utils.KeyId)
	fmt.Println(id)

	project.CreatedBy = id

	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		utils.Respond(w, r, &utils.Msg{
			Message: 	utils.ErrorInvalidParameters,
			Status:  	400,
			Fname:   	fname + " 1",
			ErrMsg:  	err.Error(),
		})
		return
	}
	defer r.Body.Close()

	lang := r.Header.Get(utils.HeaderContentLanguage)

	var resp = make(map[string]interface{})
	var err error

	if resp, err = project.Project.Create_project(project.Bin, lang); err == nil {
		/*
			create a table of ganta for this project
		 */
		if _, err = project.Create_ganta_table_for_this_project(); err != nil {
			errmsg = err.Error()
		} 
		
		/*
			create finance table for this project
		 */
		var finance = model.Finance{
			ProjectId: project.Id,
		}
		if _, err = finance.Create_and_store_on_db(); err != nil {
			errmsg = errmsg + err.Error()
		}
		
		/*
			create finresult table for this project
		 */
		var finresult = model.Finresult{
			ProjectId: project.Id,
		}
		if _, err = finresult.Create_and_store_financial_results_of_project_on_db();
			err != nil {
				errmsg = errmsg + err.Error()
		}
	} else {
		errmsg = err.Error()
	}

	utils.Respond(w, r, &utils.Msg{
		Message: 	resp,
		Status:  	utils.If_condition_then(errmsg == "", 200, 400).(int),
		Fname:   	fname + " 2",
		ErrMsg:  	errmsg,
	})
}

/*
	store docs on disc & info on db
 */
var Project_add_document_to_project = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Project_add_document_to_project"
	var rbody = model.Document{}

	var a, b = r.Header.Get(utils.KeyId), r.Header.Get(utils.KeyRoleId)
	fmt.Println(a, b)

	if err := r.ParseMultipartForm(0); err != nil {
		utils.Respond(w, r, &utils.Msg{utils.ErrorInvalidParameters, 400, fname + " 1.0", err.Error()})
		return
	}

	rbody.Name = r.FormValue("name")
	i, err := strconv.ParseInt(r.FormValue("project_id"), 0, 10)
	rbody.ProjectId = uint64(i)

	/*
		unmarshal string to json
	 */
	_ = json.Unmarshal([]byte(r.FormValue("info_sent")), &rbody.InfoSent)

	//if err := json.NewDecoder(r.Body).Decode(&rbody); err != nil {
	//	utils.Respond(w, r, &utils.Msg{utils.ErrorInvalidParameters, 400, fname + " 1.5", err.Error()})
	//	return
	//}
	//defer r.Body.Close()

	var ds = DocStore{}
	resp, err := ds.Download_and_store_file(r)
	if err != nil {
		utils.Respond(w, r, &utils.Msg{resp, 400, fname + " 2", err.Error()})
		return
	}

	rbody.Url = ds.Directory + ds.Filename + ds.Format
	resp, err = rbody.Add()

	var errmsg string
	if err != nil {
		errmsg = err.Error()
		/*
			remove file in case of an error
		*/
		if err := os.Remove("." + rbody.Url); err != nil {
			fmt.Println(err.Error())
		}
	}

	utils.Respond(w, r, &utils.Msg{
		Message: 	resp,
		Status:  	utils.If_condition_then(err == nil, 200, 400).(int),
		Fname:   	fname,
		ErrMsg:  	errmsg,
	})
}

var Update_project_by_investor = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Update_project_by_investor"
	var project = struct {
		model.Project
		Bin				string				`json:"lang"`
	}{}

	var errmsg string
	var resp = utils.NoErrorFineEverthingOk

	var lang = r.Header.Get(utils.HeaderContentLanguage)

	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		utils.Respond(w, r, &utils.Msg{
			Message: utils.ErrorInvalidParameters,
			Status:  400,
			Fname:   fname + " 1",
			ErrMsg:  err.Error(),
		})
		return
	}

	resp, err := project.Project.Update(project.Bin, lang)
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