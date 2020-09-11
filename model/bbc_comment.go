package model

type Comment struct {
	Id					uint64					`json:"id" gorm:"primary key"`

	Subject				string					`json:"subject"`
	Body				string					`json:"body" gorm:"not null"`

	UserId				uint64					`json:"user_id"`
	ProjectId			uint64					`json:"project_id" foreignkey:"projects.id"`

	GantaId				uint64					`json:"ganta_id"`

	DocumentUrl			string					`json:"document_url"`
}

func (Comment) TableName() string {
	return "comments"
}

