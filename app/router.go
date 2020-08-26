package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"invest/auth"
	"invest/control"
	"invest/model"
	"invest/utils"
	"net/http"
	"regexp"
)

func Create_new_invest_router() (*mux.Router) {
	/*
		new router
	*/
	var router = mux.NewRouter().StrictSlash(true)

	/*
		When POST request is going to be made, client agent (browser) sends first OPTIONS requests
			to check whether CORS is enables or not
		For this reason, method that handles OPTIONS requests based on the url pattern is
			provided below
	*/
	router.Methods("OPTIONS").MatcherFunc(func(r *http.Request, match *mux.RouteMatch) bool {
		matchCase, err := regexp.MatchString("/*", r.URL.Path)
		if err != nil {
			return false
		}
		return matchCase
	}).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Origin")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Add("Content-Type", "application/json")
	})

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("id: ", utils.GetContext(r, utils.KeyId))
		//fmt.Println("role: ", utils.GetContext(r, utils.KeyRole))
		utils.Respond(w, r, &utils.Msg{
			Message: 	map[string]interface{}{
				"eng": 		"Welcome Home",
				"rus":		"",
				"kaz":		"",
			},
			Status:  	http.StatusBadRequest,
			Fname:   	"MAIN",
		})
	}).Methods("GET", "POST")

	router.HandleFunc("/v1/all/signup", control.Investor_sign_up).Methods("POST")
	router.HandleFunc("/v1/all/signin", control.Sign_in).Methods("POST")

	router.HandleFunc("/v1/all/confirmation/email", control.User_email_confirm).Methods("GET")
	router.HandleFunc("/v1/all/confirmation/phone", control.User_phone_confirm).Methods("GET")

	router.HandleFunc("/v1/all/info", control.User_get_own_info).Methods("GET", "POST")

	/*
		check
	 */
	router.HandleFunc("/v1/all/user/{which}", control.User_create_read_update_delete).Methods("PUT")
	router.HandleFunc("/v1/administrate/user", control.User_create_read_update_delete).Methods("GET")
	router.HandleFunc("/v1/administrate/user", control.User_create_read_update_delete).Methods("POST")
	router.HandleFunc("/v1/administrate/user/{which}", control.User_create_read_update_delete).Methods("DELETE")

	/*
		check
	 */
	router.HandleFunc("/v1/administrate/role", control.Role_create_update_add_and_remove_permissions).Methods("GET", "POST", "PUT")
	router.HandleFunc("/v1/administrate/role/{role_id}", control.Role_delete_or_get_with_role_id).Methods("GET", "DELETE")

	/*
		check
	 */
	router.HandleFunc("/v1/administrate/categor", control.Categors_create_read_update_delete).Methods("GET")
	router.HandleFunc("/v1/administrate/categor", control.Categors_create_read_update_delete).Methods("POST")
	router.HandleFunc("/v1/administrate/categor", control.Categors_create_read_update_delete).Methods("PUT")
	router.HandleFunc("/v1/administrate/categor", control.Categors_create_read_update_delete).Methods("DELETE")

	router.HandleFunc("/v1/projects_make_changes/project", control.Update_project_by_investor).Methods("PUT")

	router.HandleFunc("/v1/projects_make_changes/project", control.Create_project).Methods("POST")
	router.HandleFunc("/v1/projects_make_changes/project/docs", control.Project_add_document_to_project).Methods("POST")

	/*
		check
	 */
	router.HandleFunc("/v1/projects_comment/project", control.Get_comments_of_the_project).Methods("GET")

	router.HandleFunc("/v1/projects_comment/project", control.Add_comment_to_project).Methods("POST")

	/*
		check
	 */
	router.HandleFunc("/v1/projects_make_changes/finance", control.Finance_table_get).Methods("GET")
	router.HandleFunc("/v1/projects_make_changes/finresult", control.Finresult_table_get).Methods("GET")
	router.HandleFunc("/v1/projects_make_changes/finance", control.Finance_table).Methods("PUT")
	router.HandleFunc("/v1/projects_make_changes/finresult", control.Finresult_and_project_evaluation_table).Methods("PUT")


	router.HandleFunc("/v1/projects_make_changes/finance", control.Finance_table).Methods("POST")
	router.HandleFunc("/v1/projects_make_changes/finresult", control.Finresult_and_project_evaluation_table).Methods("POST")

	router.HandleFunc("/v1/projects_see_all/project", control.User_project_get_all).Methods("GET")
	router.HandleFunc("/v1/projects_see_own/project", control.User_project_get_own).Methods("GET")

	/*
		check
	 */
	router.HandleFunc("/v1/administrate/project", control.Remove_user_from_project).Methods("DELETE")

	router.HandleFunc("/v1/administrate/project", control.Admin_assign_user_to_project).Methods("POST")
	router.HandleFunc("/v1/all/organization", control.Get_organization_info_by_bin).Methods("GET")

	router.HandleFunc("/v1/administrate/ganta", control.Ganta_get_all_steps_by_project_id).Methods("GET")
	router.HandleFunc("/v1/administrate/ganta", control.Ganta_add_new_step).Methods("POST")
	router.HandleFunc("/v1/administrate/ganta", control.Update_ganta_step).Methods("PUT")

	router.HandleFunc("/v1/administrate/ganta", control.Update_ganta_step).Methods("DELETE")

	/*
		check
	 */
	router.HandleFunc("/v1/administrate/organization", control.Update_organization_data).Methods("PUT")

	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		var fr = model.Project{
			Id: 1,
		}
		err := model.GetDB().Set("gorm:auto_preload", true).Table(fr.TableName()).Where("id=?", fr.Id).
			First(&fr).Error

		fmt.Println(err)
		fmt.Println(fr)
	})

	/*
		admin's functionality
		GET: /admin/civil?role=&offset=
	 */

	/*
		check for session token
	 */
	router.Use(auth.HasPermissionAndEmailVerifiedWrapper, auth.JwtAuthentication)

	return router
}

// Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozLCJyb2xlX2lkIjozLCJleHAiOiIyMDIwLTA4LTA4VDIzOjE5OjIwLjI1ODQ2NTQyMSswNjowMCJ9.Ffqpg5W0VK-1sxGZdXsX6tEzSgN4Jv19WFGmdGBBeUs
//
