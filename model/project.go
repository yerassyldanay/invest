package model

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"invest/utils/constants"
	"invest/utils/helper"
	"invest/utils/message"

	"strings"
)

func (p *Project) Create_project(trans *gorm.DB) (message.Msg){
	if p.Lang == "" {
		p.Lang = constants.DefaultContentLanguage
	}

	if err := p.Validate(); err != nil {
		return ReturnInvalidParameters(err.Error())
	}

	/*
		inside the 'info' field, we will store any info connected with a project
	*/
	b, err := json.Marshal(p.InfoSent)
	if err == nil {
		p.Info = string(b)
	}

	/*
		get org. id by bin or create one
	*/
	//p.Organization.Lang = p.Lang
	//msg := p.Organization.Create_or_get_organization_from_db_by_bin(trans)
	//if msg.IsThereAnError() {
	//	fmt.Println(msg)
	//}

	p.OrganizationId = p.Organization.Id
	p.Created = helper.GetCurrentTime()

	var categors = p.Categors
	p.Categors = []Categor{}
	p.Organization = Organization{}

	if err := p.OnlyCreate(trans); err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return ReturnDuplicateKeyError(err.Error())
		}
		return ReturnInternalDbError(err.Error())
	}

	var categor = Categor{}
	err = categor.Create_project_and_categor_relationships(categors, p.Id, trans)
	if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	return ReturnNoError()
}

func (p *Project) Update() (message.Msg) {
	if p.Lang == "" {
		p.Lang = constants.DefaultContentLanguage
	}

	if p.Id == 0 {
		return ReturnInvalidParameters("invalid id. project update")
	}

	/*
		convert info_sent to string
	*/
	b, err := json.Marshal(p.InfoSent)
	if err != nil {
		p.Info = string(b)
	}

	if err := GetDB().Model(&Project{}).Where("id=?", p.Id).Updates(Project{
		Name:           	p.Name,
		Description:    	p.Description,
		Info:           	p.Info,
		EmployeeCount:  	p.EmployeeCount,
		Email:          	p.Email,
		PhoneNumber: 		p.PhoneNumber,
		OrganizationId: 	p.OrganizationId,
	}).Error; err != nil {
		return ReturnInternalDbError(err.Error())
	}

	return ReturnNoError()
}

func (p *Project) Get_this_project_by_project_id() (error) {
	// update & get status
	err := p.GetAndUpdateStatusOfProject(GetDB())
	if err != nil {
		return err
	}

	// preload all other info
	err = p.OnlyGetByIdPreloaded(GetDB())
	if err != nil {
		return err
	}

	// categories
	err = p.OnlyGetCategorsByProjectId(GetDB())
	if err != nil {
		return err
	}

	// get info
	err = p.OnlyUnmarshalInfo()
	if err != nil {
		return err
	}

	// get assigned users
	if err := p.OnlyGetAssignedUsersByProjectId(GetDB()); err != nil {
		return err
	}

	// get rid of password
	for i, _ := range p.Users {
		p.Users[i].Password = ""
	}

	return err
}
