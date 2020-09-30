package model

import (
	"encoding/json"
	"invest/utils"

	"strings"
)

func (p *Project) Create_project() (utils.Msg){
	if p.Lang == "" {
		p.Lang = utils.DefaultContentLanguage
	}

	var trans = GetDB().Begin()
	defer func() { if trans != nil {trans.Rollback()} }()

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
	p.Organization.Lang = p.Lang
	if msg := p.Organization.Create_or_get_organization_from_db_by_bin(trans); msg.ErrMsg != "" {
		return msg
	}

	p.OrganizationId = p.Organization.Id
	p.Created = utils.GetCurrentTime()

	var categors = p.Categors
	p.Categors = []Categor{}

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

	err = trans.Commit().Error
	if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	trans = nil
	return ReturnNoError()
}

func (p *Project) Update() (utils.Msg) {
	if p.Lang == "" {
		p.Lang = utils.DefaultContentLanguage
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

