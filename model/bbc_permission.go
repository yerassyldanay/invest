package model

type Permission struct {
	Id				uint64				`json:"id" gorm:"AUTO_INCREMENT"`
	Name			string				`json:"name"`
	Description		string				`json:"description"`
}

func (Permission) TableName() string {
	return "permissions"
}
