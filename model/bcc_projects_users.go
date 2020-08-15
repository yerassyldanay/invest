package model

type ProjectsUsers struct {
	ProjectId				uint64		`json:"project_id"`
	UserId					uint64			`json:"user_id"`
}

func (pu *ProjectsUsers) TableName() string {
	return "projects_users"
}
