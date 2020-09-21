package app

/*

*/
const (
	Perm_1_crud_user = "administrate"
	Perm_2_projects_see_all = "projects_see_all"
	Perm_3_projects_see_own = "projects_see_own"
	Perm_4_projects_make_changes = "projects_make_changes"
	Perm_5_projects_comment = "projects_comment"
	Perm_6_projects_accept = "projects_accept"
	Perm_7_analysis_see = "analysis_see"
	Perm_8_projects_submit = "projects_submit"

	Perm_9_comment_add = "comment_add"
	Perm_10_comment_get = "comment_get"
)

var PMapPermissionMap = map[string]string{
	Perm_1_crud_user: `{"eng": ["create, remove & block the user"]}`,
	Perm_2_projects_see_all: `{"eng": ["see all details of the project", "stages & comments", "documents"]}`,
	Perm_3_projects_see_own: `{"eng": ["see only own projects"]}`,
	Perm_4_projects_make_changes: `{"eng": ["create a project", "can make changes"]}`,
	Perm_5_projects_comment: `{"eng": ["comment on projects"]}`,
	Perm_6_projects_accept: `{"eng": ["accept the project", "deny", "send to resubmission"]}`,
	Perm_7_analysis_see: `{"eng": ["see analysis"]}`,
	Perm_8_projects_submit: `{"eng": ["submit project"]}`,

	Perm_9_comment_add: `{"eng":["add comment with a document or without a document", "delete a comment"]}`,
	Perm_10_comment_get: `{"eng": ["get comments of the project"]}`,
}
