package model

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"invest/utils"
	"os"
	"path/filepath"
)

/*
	validate
 */
func (d *Document) Validate() bool {
	if d.Name == "" || d.ProjectId == 0 || d.GantaId == 0 {
		return false
	}
	return true
}

/*
	create a document
 */
func (d *Document) Only_create(trans *gorm.DB) error {
	return  trans.Create(d).Error
}

/*
	add docs to the project by project_id
		at this moment document is already stored on db
 */
func (d *Document) Add() (utils.Msg) {
	if ok := d.Validate(); !ok {
		return utils.Msg{utils.ErrorInvalidParameters, 400, "", "invalid parameters in docs add"}
	}

	if d.Uri == "" {
		return utils.Msg{utils.ErrorInvalidParameters, 400, "", "could not store a document on hard disk"}
	}

	var trans = GetDB().Begin()
	defer func() {if trans != nil {trans.Rollback()}}()

	/*
		create a document row on db
	 */
	if err := d.Only_create(trans); err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	/*
		set document id to ganta
	 */
	var ganta = Ganta{ Id: d.GantaId }
	err := ganta.Only_add_document_by_ganta_id(d.Id, trans)
	if err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	trans.Commit()
	trans = nil

	return utils.Msg{utils.NoErrorFineEverthingOk, 200, "", ""}
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

func (d *Document) Remove() (utils.Msg) {

	if !d.Is_it_investor() {
		return utils.Msg{utils.ErrorMethodNotAllowed, 405, "", "not investor or did not create this project"}
	}

	if err := GetDB().Raw("delete from documents where project_id = ? and id = ? returning *;", d.ProjectId, d.Id).
		Scan(d).Error; err != nil {
			return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	var docpath, err = filepath.Abs(".." + "/invest" + d.Uri)
	if err == nil {
		err = os.Remove(docpath)
	}

	fmt.Println("removed file: ", err)

	return utils.Msg{utils.NoErrorFineEverthingOk, 200, "", ""}
}
