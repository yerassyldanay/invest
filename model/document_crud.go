package model

import (
	"errors"
	"invest/utils"
)

/*
	validate
 */
func (d *Document) Validate() bool {
	if d.Name == "" || d.ProjectId == 0 {
		return false
	}
	return true
}

/*
	add docs to the project by project_id
		at this moment document is already stored on db
 */
func (d *Document) Add() (map[string]interface{}, error) {
	if ok := d.Validate(); !ok {
		return utils.ErrorInvalidParameters, errors.New("invalid parameters in docs add")
	}

	if d.Url == "" {
		return utils.ErrorInternalServerError, errors.New("could not store a document on hard disk")
	}

	//d.Info = "'{}'"
	//if d.InfoSent != nil {
	//	if b, err := json.Marshal(d.InfoSent); err == nil {
	//		d.Info = string(b)
	//	}
	//}

	/*
		find the project
	 */
	var project = Project{}
	if err := GetDB().Table(Project{}.TableName()).Where("id=?", d.ProjectId).First(&project).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}

	d.ProjectId = project.Id

	/*
		create a document row on db
	 */
	if err := GetDB().Create(d).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}

	return utils.NoErrorFineEverthingOk, nil
}

/*
	update doc info
 */
func (d *Document) Update_name_or_info() (map[string]interface{}, error) {
	if d.Id == 0 || d.Validate() == false {
		return utils.ErrorInvalidParameters, errors.New("invalid parameters. document update")
	}

	/*
		marshal data & store
	 */
	//b, err := json.Marshal(d.InfoSent)
	//if err == nil {
	//	d.Info = string(b)
	//}

	if err := GetDB().Table(Document{}.TableName()).Where("id=?", d.Id).Updates(map[string]interface{}{
		"name": d.Name,
		//"info": d.Info,
	}).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}

	/*
		send data back
	 */
	var tdoc = Document{}
	if err := GetDB().Table(Document{}.TableName()).Where("id=?", d.Id).First(&tdoc).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(tdoc)

	return resp, nil
}

func (d *Document) Remove() (map[string]interface{}, error) {
	if d.Id == 0 {
		return utils.ErrorInvalidParameters, errors.New("id = 0. remove doc. project")
	}

	if err := GetDB().Where("id=?", d.Id).Delete(&Document{}).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}

	return utils.NoErrorFineEverthingOk, nil
}
