package app

import (
	"github.com/gorilla/mux"
	"invest/auth"
	"invest/control"
	"invest/model"
	"invest/utils"
	"net/http"
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
		//matchCase, err := regexp.MatchString("/.*", r.URL.Path)
		//if err != nil {
		//	return false
		//}
		//return matchCase
		return true
	}).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/*
			Explicitly informs the referer how many seconds it should store the preflight
			result. Within this time, it can just send the request,
			and doesn't need to bother sending the preflight request again.
		*/
		w.Header().Set("Access-Control-Max-Age", "86400")

		w.Header().Set("Access-Control-Allow-Credentials", "")

		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:63342")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Origin")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS, EAT")
		w.Header().Add("Content-Type", "application/json")
	})

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("id: ", utils.GetContext(r, utils.KeyId))
		//fmt.Println("role: ", utils.GetContext(r, utils.KeyRole))
		utils.Respond(w, r, utils.Msg{
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

	/*
		create
	 */
	//router.HandleFunc("/v1/all/check/fio", control.Investor_sign_up).Methods("GET")
	//router.HandleFunc("/v1/all/check/email", control.Investor_sign_up)

	router.HandleFunc("/v1/all/confirmation/email", control.User_email_confirm).Methods("GET")
	router.HandleFunc("/v1/all/confirmation/phone", control.User_phone_confirm).Methods("GET")

	router.HandleFunc("/v1/all/info", control.User_get_own_info).Methods("GET", "POST")

	/*
		CRUD user by admin
	 */
	router.HandleFunc("/v1/administrate/user/{which}", control.User_create_read_update_delete).Methods("PUT")
	router.HandleFunc("/v1/administrate/user", control.User_create_read_update_delete).Methods("GET")
	router.HandleFunc("/v1/administrate/user", control.User_create_read_update_delete).Methods("POST")

	router.HandleFunc("/v1/administrate/profile", control.Get_full_user_info)

	/*
		CRUD role & assign permissions
	 */
	router.HandleFunc("/v1/administrate/role", control.Role_create_update_add_and_remove_permissions).Methods("GET", "POST", "PUT")
	router.HandleFunc("/v1/administrate/role/{role_id}", control.Role_delete_or_get_with_role_id).Methods("GET", "DELETE")
	router.HandleFunc("/v1/administrate/permissions", control.Role_add_and_remove_permissions).Methods( "POST", "DELETE")

	/*
		list users & their roles, who are assigned to the project
	 */
	router.HandleFunc("/v1/administrate/", nil).Methods("GET")

	/*
		Categories
	 */
	router.HandleFunc("/v1/all/categor", control.Categors_create_read_update_delete).Methods("GET")
	router.HandleFunc("/v1/administrate/categor", control.Categors_create_read_update_delete).Methods("POST")
	router.HandleFunc("/v1/administrate/categor", control.Categors_create_read_update_delete).Methods("DELETE")

	/*
		Project & Document
	 */
	router.HandleFunc("/v1/projects_make_changes/project", control.Update_project).Methods("PUT")
	router.HandleFunc("/v1/projects_make_changes/project", control.Create_project).Methods("POST")

	router.HandleFunc("/v1/projects_make_changes/project/docs", control.Project_add_document_to_project).Methods("POST")
	router.HandleFunc("/v1/projects_make_changes/project/docs", control.Project_remove_document).Methods("DELETE")

	/*
		/ad../stat/project?status=?
		/ad../stat/project?user_id=?&&status=?
		/ad../stat/project?user_id
	*/
	router.HandleFunc("/v1/administrate/stat/project", control.Stats_on_projects_based_on_user_or_status).Methods("GET")
	router.HandleFunc("/v1/all/project/docs/stat", control.Get_stat_on_documents_of_project).Methods("GET")

	/*
		Leave a COMMENT on the project
	 */
	router.HandleFunc("/v1/projects_comment/project", control.Get_comments_of_the_project).Methods("GET")
	router.HandleFunc("/v1/projects_comment/project", control.Add_comment_to_project).Methods("POST")

	/*
		Read & Update financial tables
			they are automatically created
	 */
	router.HandleFunc("/v1/projects_make_changes/finance", control.Finance_table_get).Methods("GET")
	router.HandleFunc("/v1/projects_make_changes/finance", control.Finance_table_update).Methods("PUT")

	router.HandleFunc("/v1/projects_make_changes/finresult", control.Finresult_table_get).Methods("GET")
	router.HandleFunc("/v1/projects_make_changes/finresult", control.Finresult_table_update).Methods("PUT")

	router.HandleFunc("/v1/projects_see_all/project/analysis", control.User_get_projects_info_grouped_by_statuses).Methods("GET")
	router.HandleFunc("/v1/projects_see_all/project", control.User_project_get_all).Methods("GET")
	router.HandleFunc("/v1/projects_see_own/project", control.User_project_get_own).Methods("GET")

	/*
		Assign & remove user from project
	 */
	router.HandleFunc("/v1/administrate/project", control.Remove_user_from_project).Methods("DELETE")
	router.HandleFunc("/v1/administrate/project", control.Assign_user_to_project).Methods("POST")
	router.HandleFunc("/v1/administrate/project/stat", control.Get_all_assigned_users_to_project).Methods("GET")

	router.HandleFunc("/v1/projects_comment/ganta/{choice}", control.Ganta_only_ganta_steps_by_project_id).Methods("GET")
	router.HandleFunc("/v1/projects_comment/ganta", control.Ganta_create_update_delete).Methods("POST")
	router.HandleFunc("/v1/projects_comment/ganta", control.Ganta_create_update_delete).Methods("PUT")
	router.HandleFunc("/v1/projects_comment/ganta", control.Ganta_create_update_delete).Methods("DELETE")

	/*
		check
	 */
	router.HandleFunc("/v1/all/organization", control.Get_organization_info_by_bin).Methods("GET")
	router.HandleFunc("/v1/administrate/organization", control.Update_organization_data).Methods("PUT")

	/*
		check
	 */
	router.HandleFunc("/v1/all/password", control.Forget_password_send_message).Methods("GET", "POST")

	router.HandleFunc("/droptables", func(w http.ResponseWriter, r *http.Request) {
		model.GetDB().Debug().DropTableIfExists(&model.Categor{}, &model.Comment{}, &model.Document{}, &model.Email{}, &model.Finance{}, &model.FinanceCol{},
			&model.Finresult{}, &model.FinresultCol{}, &model.Ganta{}, &model.Organization{}, &model.Permission{},
			&model.Phone{}, &model.Project{}, &model.ProjectStatus{}, &model.Role{}, &model.SendgridMessage{}, &model.SendgridMessageStore{},
			&model.User{})

		model.GetDB().Debug().AutoMigrate(&model.ProjectsUsers{})

		model.GetDB().DropTableIfExists("goose_db_version")
	})

	router.HandleFunc("/intest", func(w http.ResponseWriter, r *http.Request) {
		var project = model.Project{Id: 2}
		msg := project.Get_this_project_with_its_users()
		utils.Respond(w, r, msg)
	})

	/*
		admin's functionality
		GET: /admin/civil?role=&offset=
	 */

	/*
		check for session token
			also will go through: auth.EmailVerifiedWrapper, auth.HasPermissionWrapper
	 */
	router.Use(auth.JwtAuthentication)
	router.Use(mux.CORSMethodMiddleware(router))

	return router
}


// Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozLCJyb2xlX2lkIjozLCJleHAiOiIyMDIwLTA4LTA4VDIzOjE5OjIwLjI1ODQ2NTQyMSswNjowMCJ9.Ffqpg5W0VK-1sxGZdXsX6tEzSgN4Jv19WFGmdGBBeUs
//
