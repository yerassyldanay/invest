package app

import (
	"github.com/gorilla/mux"
	"github.com/yerassyldanay/invest/middleware"
	"net/http"
)

func NewRouter() *mux.Router {
	// new router
	var generalRouter = mux.NewRouter().StrictSlash(true)

	// prefixes
	var v1 = generalRouter.PathPrefix("/v1").Subrouter()
	var v1Free = generalRouter.PathPrefix("/v1").Subrouter()

	var docRouter = generalRouter.PathPrefix("/documents").Subrouter()
	var download = generalRouter.PathPrefix("/download").Subrouter()

	/*
		check for session token
			also will go through: middleware.EmailVerifiedWrapper
	*/
	v1.Use(middleware.JwtAuthentication)

	//v1.Use(middleware.CORSMethodMiddleware)
	//v1Free.Use(middleware.CORSMethodMiddleware)

	// log every request
	v1.Use(middleware.LoggerMiddleware)
	v1Free.Use(middleware.LoggerMiddleware)

	var STATIC_DIR = "/documents/docs"
	docRouter.Handle("/docs/{file}", http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("."+STATIC_DIR))))

	var STATIC_DIR_ANALYSIS = "/documents/analysis"
	docRouter.Handle("/analysis/{file}", http.StripPrefix(STATIC_DIR_ANALYSIS, http.FileServer(http.Dir("."+STATIC_DIR_ANALYSIS))))

	// download a binary file
	download.HandleFunc("/documents/docs/{file}", Document_download).Methods("GET")

	/*
		Registration
	*/
	v1Free.HandleFunc("/signup", SignUp).Methods("POST")
	v1Free.HandleFunc("/signin", SignIn).Methods("POST")

	/*
		Confirm
	*/
	v1Free.HandleFunc("/confirmation/email", UserProfileConfirmEmail).Methods("POST")

	/*
		Profile
	*/
	v1.HandleFunc("/profile", GetFullUserInfo).Methods("GET")
	v1.HandleFunc("/profile/own", UserGetOwnInfo).Methods("GET")
	v1.HandleFunc("/profile/other", GetFullUserInfo).Methods("GET")

	v1.HandleFunc("/profile/own", UpdateOwnProfile).Methods("PUT")
	v1.HandleFunc("/profile/other", UpdateOtherProfile).Methods("PUT")

	// 6z24HXMd7nLeZAE
	v1.HandleFunc("/profile/password/own", UpdateOwnPassword).Methods("PUT")
	v1.HandleFunc("/profile/password/other", UpdateOtherPassword).Methods("PUT")

	/*
		CRUD user by admin
	*/
	v1.HandleFunc("/user", UsersGetByRole).Methods("GET")
	v1.HandleFunc("/user", CreateUser).Methods("POST")

	/*
		Assign & remove user from project
	*/
	v1.HandleFunc("/assign", RemoveUserFromProject).Methods("DELETE")
	v1.HandleFunc("/assign", AssignUserToProject).Methods("POST")

	/*
		Gantt
	*/
	v1.HandleFunc("/ganta/restricted/parents", GantaRestrictedGetParentGantaSteps).Methods("GET")
	v1.HandleFunc("/ganta/restricted/children", GantaRestrictedGetChildGantaSteps).Methods("GET")

	v1.HandleFunc("/ganta/change/check_permission", Ganta_can_user_change_current_status).Methods("GET")
	v1.HandleFunc("/ganta/change/status", Ganta_confirm_the_ganta_step).Methods("POST")

	/*
		Documents
	*/
	v1.HandleFunc("/project/docs/box", DocumentAddBoxToUploadDocument).Methods("POST")
	v1.HandleFunc("/project/docs/file", Document_upload_document).Methods("POST")
	v1.HandleFunc("/project/docs/file/delete", DocumentRemoveFile).Methods("GET")
	v1.HandleFunc("/project/docs", DocumentGet).Methods("GET")

	/*
		Project
	*/
	v1.HandleFunc("/project", GetProjectByProjectId).Methods("GET")
	v1.HandleFunc("/project", CreateProject).Methods("POST")

	/*
		Status
	*/
	v1.HandleFunc("/project/status", Project_get_status_of_project).Methods("GET")

	/*
		Analysis
	*/
	v1.HandleFunc("/analysis", Analysis_get).Methods("POST")
	v1.HandleFunc("/analysis/file", Analysis_get_file).Methods("POST")

	/*
		Role & Permissions
	*/
	v1.HandleFunc("/role", Role_create_update_add_and_remove_permissions).Methods("GET")
	//v1.HandleFunc("/role/{role_id}", app.Role_delete_or_get_with_role_id).Methods("GET", "DELETE")
	//v1.HandleFunc("/permissions", app.Role_add_and_remove_permissions).Methods( "POST", "DELETE")

	/*
		Categories
	*/
	v1.HandleFunc("/categor", Categors_create_read_update_delete).Methods("GET", "POST", "DELETE")

	/*
		List of something or someone
	*/
	v1.HandleFunc("/list/users/by_project", Get_all_assigned_users_to_project).Methods("GET")
	v1.HandleFunc("/list/projects/by_user", Get_all_projects_by_user_and_status).Methods("GET")
	v1.HandleFunc("/list/projects/all", Get_all_projects_by_statuses).Methods("GET")
	v1.HandleFunc("/list/projects/own", Get_own_projects).Methods("GET")
	//v1.HandleFunc("/list/projects/by_user_and_status", app.Get_projects_based_on_user_or_status).Methods("GET")

	/*
		/ad../stat/project?status=? - provides all projects by status
		/ad../stat/project?user_id=?&&status=? - provides projects by user_id & status
	*/
	//v1.HandleFunc("/stats/projects/grouped_by_status", nil).Methods("GET")
	//v1.HandleFunc("/stats/docs/by_project", nil).Methods("GET")

	/*
		Leave a COMMENT on the project
	*/
	v1.HandleFunc("/spk_comment", Project_get_comments).Methods("GET")
	v1.HandleFunc("/spk_comment", Project_comment_on_documents).Methods("POST")

	/*
		Organization
	*/
	v1.HandleFunc("/organization", Get_organization_info_by_bin).Methods("GET")
	v1.HandleFunc("/organization", Update_organization_data).Methods("PUT")

	/*
		Reset password
	*/
	v1.HandleFunc("/reset_password", ForgetPasswordSendMessage).Methods("GET", "POST")

	/*
		Notifications
	*/
	v1.HandleFunc("/notifications", Notification_get).Methods("GET")

	/*
		SMTP
	*/
	v1.HandleFunc("/smtp", Smtp_create_update_put).Methods("POST", "PUT", "DELETE")
	v1.HandleFunc("/smtp", Smtp_get).Methods("GET")

	/*
		Test API
	*/
	v1Free.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Tested!!!"))
	})

	return generalRouter
}

// Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozLCJyb2xlX2lkIjozLCJleHAiOiIyMDIwLTA4LTA4VDIzOjE5OjIwLjI1ODQ2NTQyMSswNjowMCJ9.Ffqpg5W0VK-1sxGZdXsX6tEzSgN4Jv19WFGmdGBBeUs
//
