package model

import "time"

/*
	указывает БИН – для роли Инвестора , при этом БИН может являться логином Инвестора
		(по согласованию с Заказчиком):
		* наименование организации;
		* ФИО руководителя;
		* адрес регистрации;
		* дата регистрации юридического лица.
		Допускается автоматическое заполнение данных сведений по БИН.
*/
type Organization struct {
	Id					uint64				`json:"id" gorm:"primary key"`

	Lang				string				`json:"lang" gorm:"unique_index:make_org_unique"`
	Bin					string				`json:"bin" gorm:"unique_index:make_org_unique" validate:"regexp=[0-9]+"`

	Name				string				`json:"name" gorm:"not null" validate:"required"`
	Fio					string				`json:"fio"`

	Regdate				time.Time			`json:"regdate" validate:"required"`
	Address				string				`json:"address" gorm:"not null" validate:"required"`

	//Info				string				`json:"info" gorm:"default:'{}'"`
	//InfoSent			map[string]interface{}		`json:"info_sent" gorm:"-"`

	//Deleted				time.Time			`json:"deleted" gorm:"default:null"`
}

func (Organization) TableName() string {
	return "organizations"
}

