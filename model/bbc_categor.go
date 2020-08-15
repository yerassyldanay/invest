package model

type Categor struct {
	Id				uint64				`json:"id" gorm:"primary key"`
	Name			string				`json:"name" gorm:"unique;not null"`
}

func (Categor) TableName() string {
	return "categors"
}
