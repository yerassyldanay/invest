package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"invest/auth"
	"invest/control"
	"invest/model"
	"invest/utils/message"
	"net/http"
)

func Create_new_invest_router() (*mux.Router) {

	/*
		new router
	*/
	var generalRouter = mux.NewRouter().StrictSlash(true)

	var v1 = generalRouter.PathPrefix("/v1").Subrouter()
	var docRouter = generalRouter.PathPrefix("/documents").Subrouter()
	var download = generalRouter.PathPrefix("/download").Subrouter()

	v1.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		msg := model.ReturnNoError()
		message.Respond(w, r, msg)
	}).Methods("GET", "POST")

	var STATIC_DIR = "/documents/docs"
	docRouter.Handle("/docs/{file}", http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("." + STATIC_DIR))))

	var STATIC_DIR_ANALYSIS = "/documents/analysis"
	docRouter.Handle("/analysis/{file}", http.StripPrefix(STATIC_DIR_ANALYSIS, http.FileServer(http.Dir("." + STATIC_DIR_ANALYSIS))))

	// download a binary file
	download.HandleFunc("/documents/docs/{file}", control.Document_download).Methods("GET")

	/*
		Registration
	 */
	v1.HandleFunc("/signup", control.Sign_up).Methods("POST")
	v1.HandleFunc("/signin", control.Sign_in).Methods("POST")

	/*
		Confirm
	 */
	v1.HandleFunc("/confirmation/email", control.User_email_confirm).Methods("POST")
	//v1.HandleFunc("/confirmation/phone", control.User_phone_confirm).Methods("GET")

	/*
		Profile
	 */
	v1.HandleFunc("/profile", control.Get_full_user_info).Methods("GET")
	v1.HandleFunc("/profile/own", control.User_get_own_info).Methods("GET")
	v1.HandleFunc("/profile/other", control.Get_full_user_info).Methods("GET")

	v1.HandleFunc("/profile/own", control.Update_own_profile).Methods("PUT")
	v1.HandleFunc("/profile/other", control.Update_other_profile).Methods("PUT")

	// 6z24HXMd7nLeZAE
	v1.HandleFunc("/profile/password/own", control.Update_own_password).Methods("PUT")
	v1.HandleFunc("/profile/password/other", control.Update_other_password).Methods("PUT")

	/*
		CRUD user by admin
	*/
	v1.HandleFunc("/user", control.Users_get_by_role).Methods("GET")
	v1.HandleFunc("/user", control.Create_user).Methods("POST")

	/*
		Assign & remove user from project
	*/
	v1.HandleFunc("/assign", control.Remove_user_from_project).Methods("DELETE")
	v1.HandleFunc("/assign", control.Assign_user_to_project).Methods("POST")

	/*
		Gantt
	 */
	v1.HandleFunc("/ganta/restricted/parents", control.Ganta_restricted_get_parent_ganta_steps).Methods("GET")
	v1.HandleFunc("/ganta/restricted/children", control.Ganta_restricted_get_child_ganta_steps).Methods("GET")

	v1.HandleFunc("/ganta/change/check_permission", control.Ganta_can_user_change_current_status).Methods("GET")
	v1.HandleFunc("/ganta/change/status", control.Ganta_confirm_the_ganta_step).Methods("POST")
	//v1.HandleFunc("/ganta/change/time", control.Ganta_change_ganta_time).Methods("POST")

	/*
		Documents
	 */
	v1.HandleFunc("/project/docs/box", control.Document_add_box_to_upload_document).Methods("POST")
	v1.HandleFunc("/project/docs/file", control.Document_upload_document).Methods("POST")
	v1.HandleFunc("/project/docs/file/delete", control.Document_remove_file).Methods("GET")
	v1.HandleFunc("/project/docs", control.Document_get).Methods("GET")

	/*
		Project
	*/
	v1.HandleFunc("/project", control.Get_project_by_project_id).Methods("GET")
	v1.HandleFunc("/project", control.Create_project).Methods("POST")

	/*
		Status
	 */
	v1.HandleFunc("/project/status", control.Project_get_status_of_project).Methods("GET")

	/*
		Analysis
	 */
	v1.HandleFunc("/analysis", control.Analysis_get).Methods("POST")
	v1.HandleFunc("/analysis/file", control.Analysis_get_file).Methods("POST")

	/*
		Role & Permissions
	 */
	v1.HandleFunc("/role", control.Role_create_update_add_and_remove_permissions).Methods("GET")
	//v1.HandleFunc("/role/{role_id}", control.Role_delete_or_get_with_role_id).Methods("GET", "DELETE")
	//v1.HandleFunc("/permissions", control.Role_add_and_remove_permissions).Methods( "POST", "DELETE")

	/*
		Categories
	 */
	v1.HandleFunc("/categor", control.Categors_create_read_update_delete).Methods("GET", "POST", "DELETE")

	/*
		List of something or someone
	*/
	v1.HandleFunc("/list/users/by_project", control.Get_all_assigned_users_to_project).Methods("GET")
	v1.HandleFunc("/list/projects/by_user", control.Get_all_projects_by_user_and_status).Methods("GET")
	v1.HandleFunc("/list/projects/all", control.Get_all_projects_by_statuses).Methods("GET")
	v1.HandleFunc("/list/projects/own", control.Get_own_projects).Methods("GET")
	//v1.HandleFunc("/list/projects/by_user_and_status", control.Get_projects_based_on_user_or_status).Methods("GET")

	/*
		/ad../stat/project?status=? - provides all projects by status
		/ad../stat/project?user_id=?&&status=? - provides projects by user_id & status
	*/
	//v1.HandleFunc("/stats/projects/grouped_by_status", nil).Methods("GET")
	//v1.HandleFunc("/stats/docs/by_project", nil).Methods("GET")

	/*
		Leave a COMMENT on the project
	 */
	v1.HandleFunc("/spk_comment", control.Project_get_comments).Methods("GET")
	v1.HandleFunc("/spk_comment", control.Project_comment_on_documents).Methods("POST")

	/*
		Organization
	 */
	v1.HandleFunc("/organization", control.Get_organization_info_by_bin).Methods("GET")
	v1.HandleFunc("/organization", control.Update_organization_data).Methods("PUT")

	/*
		Reset password
	 */
	v1.HandleFunc("/reset_password", control.Forget_password_send_message).Methods("GET", "POST")

	/*
		Notifications
	 */
	v1.HandleFunc("/notifications", control.Notification_get).Methods("GET")

	/*
		SMTP
	 */
	v1.HandleFunc("/smtp", control.Smtp_create_update_put).Methods("POST", "PUT", "DELETE")
	v1.HandleFunc("/smtp", control.Smtp_get).Methods("GET")

	/*
		Test API
	 */
	v1.HandleFunc("/tests", func(w http.ResponseWriter, r *http.Request) {

	})

	v1.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header)
	})

	/*
		check for session token
			also will go through: auth.EmailVerifiedWrapper
	 */
	v1.Use(auth.JwtAuthentication)
	v1.Use(mux.CORSMethodMiddleware(v1))

	return generalRouter
}

// Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozLCJyb2xlX2lkIjozLCJleHAiOiIyMDIwLTA4LTA4VDIzOjE5OjIwLjI1ODQ2NTQyMSswNjowMCJ9.Ffqpg5W0VK-1sxGZdXsX6tEzSgN4Jv19WFGmdGBBeUs
//
