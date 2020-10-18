package model

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"invest/utils"
)

type Categor struct {
	Id				uint64				`json:"id" gorm:"primary key"`
	Kaz				string				`json:"kaz" gorm:"unique;not null"`
	Rus				string				`json:"rus" gorm:"unique;not null"`
	Eng				string				`json:"eng" gorm:"unique;not null"`
}

func (Categor) TableName() string {
	return "categors"
}

func (c *Categor) BeforeDelete(tx *gorm.DB) error {
	var count int
	tx.Table("projects_categories").Where("categor_id = ?", c.Id).Count(&count)

	if count != 0 {
		return errors.New("categor is being used")
	}

	return nil
}

// create
func (c *Categor) OnlyCreate(tx *gorm.DB) error {
	err := tx.Create(c).Error
	return err
}

// get one
func (c *Categor) OnlyGetById(tx *gorm.DB) error {
	err := tx.First(c, "id = ?", c.Id).Error
	return err
}

// get all
func (c *Categor) OnlyGetAll(tx *gorm.DB) ([]Categor, error) {
	categors := []Categor{}
	err := tx.Find(&categors).Error
	return categors, err
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
func (ca *Categor) UpdateById() (utils.Msg) {
	if err := GetDB().Model(Categor{Id: ca.Id}).Updates(map[string]interface{}{
		"kaz": ca.Kaz,
		"rus": ca.Rus,
		"eng": ca.Eng,
	}).Error; err != nil {
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

