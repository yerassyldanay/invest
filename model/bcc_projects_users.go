package model

import (
	"time"
)

type ProjectUserStat struct {
	ProjectId uint64 `json:"project_id;omitempty"`
	UserId    uint64 `json:"user_id;omitempty"`
	Status    string `json:"status"`
}

type ProjectsUsers struct {
	ProjectId uint64    `json:"project_id"`
	UserId    uint64    `json:"user_id"`
	Created   time.Time `json:"created"`
}

func (p *ProjectsUsers) TableName() string {
	return "projects_users"
}
