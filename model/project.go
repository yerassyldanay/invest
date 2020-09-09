package model

import (
	"encoding/json"
	"invest/utils"
	"net/http"
	"strings"
)

func (p *Project) IsValid() (bool) {
	if p.OfferedById == 0 || p.OfferedByPosition == "" || p.Email == "" || p.PhoneNumber == "" || p.Name == "" || p.Description == "" || p.EmployeeCount <= 0 {
		return false
	}
	return true
}

func (p *Project) Create_project() (*utils.Msg){
	if p.Lang == "" {
		p.Lang = utils.DefaultContentLanguage
	}

	if ok := p.IsValid(); !ok {
		return &utils.Msg{
			utils.ErrorInvalidParameters, http.StatusBadRequest, "", "invalid parameters have been passed",
		}
	}
	
	/*
		convert map to string
	 */
	b, err := json.Marshal(p.InfoSent)
	if err == nil {
		p.Info = string(b)
	}
	
	/*
		get org. id by bin or create one
	 */
	p.Organization.Lang = p.Lang
	if msg := p.Organization.Create_or_get_organization_from_db_by_bin(); msg.ErrMsg != "" {
		return msg
	}

	p.OrganizationId = p.Organization.Id
	p.Created = utils.GetCurrentTime()

	var trans = GetDB().Begin()
	defer func() { if trans != nil {trans.Rollback()} }()
	
	if err := trans.Create(p).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return &utils.Msg{
				utils.ErrorDupicateKeyOnDb, http.StatusConflict, "", err.Error(),
			}
		}
		return &utils.Msg{
			utils.ErrorInternalDbError, http.StatusExpectationFailed, "", err.Error(),
		}
	}

	trans.Commit()
	return &utils.Msg{
		utils.NoErrorFineEverthingOk, http.StatusOK, "", "",
	}
}

func (p *Project) Update() (*utils.Msg) {
	if p.Lang == "" {
		p.Lang = utils.DefaultContentLanguage
	}

	if p.Id == 0 {
		return &utils.Msg{
			utils.ErrorInvalidParameters, http.StatusBadRequest, "", "invalid id. project update",
		}
	}
	
	/*
		get org id
	 */
	org, err := p.Organization.Get_and_assign_info_on_organization_by_bin()
	if err == nil {
		p.OrganizationId = org.Id
	} else {
		return &utils.Msg{
			utils.ErrorInvalidParameters, http.StatusBadRequest, "", "could not get info on organization",
		}
	}

	/*
		convert info_sent to string
	 */
	b, err := json.Marshal(p.InfoSent)
	if err == nil {
		_ = GetDB().Table(Project{}.TableName()).Select("info").Where("id=?", p.Id).First(&p.Info).Error
	} else { p.Info = string(b) }

	if err := GetDB().Model(&Project{}).Where("id=?", p.Id).Updates(Project{
		Name:           	p.Name,
		Description:    	p.Description,
		Info:           	p.Info,
		EmployeeCount:  	p.EmployeeCount,
		Email:          	p.Email,
		PhoneNumber: 		p.PhoneNumber,
		OrganizationId: 	p.OrganizationId,
	}).Error; err != nil {
		return &utils.Msg{
			utils.ErrorInternalDbError, http.StatusExpectationFailed, "", err.Error(),
		}
	}

	return &utils.Msg{
		utils.NoErrorFineEverthingOk, http.StatusOK, "", "",
	}
}
