package model

import (
	"bytes"
	"fmt"
	"github.com/jinzhu/gorm"
	"invest/utils"
)

type Categor struct {
	Id				uint64				`json:"id" gorm:"primary key"`
	Name			string				`json:"name" gorm:"unique;not null"`
}

func (Categor) TableName() string {
	return "categors"
}

/*
	create a category
*/
func(ca *Categor) Create_category() (utils.Msg) {
	if err := GetDB().Create(ca).Error; err != nil {
		return ReturnInternalDbError(err.Error())
	}

	return ReturnNoError()
}

/*

 */
func (ca *Categor) Update() (utils.Msg) {
	if err := GetDB().Table(Categor{}.TableName()).Update("name", ca.Name).Error; err != nil {
		return ReturnInternalDbError(err.Error())
	}

	return ReturnNoError()
}

/*
	delete category + delete from projects
*/
func (ca *Categor) Delete_category_from_tabe_and_projects() (utils.Msg) {

	if err := GetDB().Delete(ca, "id = ?", ca.Id).Error;
		err != nil {
		return ReturnInternalDbError(err.Error())
	}

	return ReturnNoError()
}

/*
	get all categories
*/
func (ca *Categor) Get_all_categors(offset string) (utils.Msg) {
	var cas = []Categor{}
	if err := GetDB().Table(ca.TableName()).Offset(offset).Limit(GetLimit).Find(&cas).Error; err != nil {
		return ReturnInternalDbError(err.Error())
	}

	var carr = []map[string]interface{}{}
	for _, each := range cas {
		carr = append(carr, Struct_to_map(each))
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = carr

	return ReturnNoError()
}

/*
	stores a category or categories of a project
*/
func (ca *Categor) Create_project_and_categor_relationships(categors []Categor, project_id uint64, trans *gorm.DB) error {
	if len(categors) > 0 {
		var main_query = bytes.Buffer{}
		main_query.WriteString(" insert into projects_categors (project_id, categor_id) values ")
		for i, categor := range categors {
			if i != 0 {
				main_query.WriteString(", ")
			}
			main_query.WriteString(fmt.Sprintf("(%d, %d)", project_id, categor.Id))
		}

		main_query.WriteString(";")

		var so = main_query.String()
		return trans.Exec(so).Error
	}

	return nil
}

