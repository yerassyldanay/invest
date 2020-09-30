package model

type Comment struct {
	Id					uint64					`json:"id" gorm:"primary key"`
	Body				string					`json:"body" gorm:"not null"`

	UserId				uint64					`json:"user_id" gorm:"foreignkey:users.id"`
	ProjectId			uint64					`json:"project_id" gorm:"foreignkey:projects.id"`

	//GantaId				uint64					`json:"ganta_id"`

	//DocumentUrl			string					`json:"document_url"`

	Status				string					`json:"status" gorm:"-"`
	DocStatuses			[]Document				`json:"doc_statuses" gorm:"-"`
}

func (Comment) TableName() string {
	return "comments"
}

