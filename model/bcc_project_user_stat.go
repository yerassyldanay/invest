package model

type ProjectUserStat struct {
	ProjectId				uint64				`json:"project_id;omitempty"`
	UserId					uint64				`json:"user_id;omitempty"`
	Status					string				`json:"status"`
}

