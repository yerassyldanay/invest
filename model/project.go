package model

import (
	"encoding/json"
	"errors"
	"invest/utils"
	"time"
)

func (p *Project) Validate() (bool) {
	if p.CreatedBy == 0 || p.Email == "" || p.Phone == "" || p.Name == "" || p.Description == "" || p.EmployeeCount == 0 {
		return false
	}
	return true
}

func (p *Project) Create_project(bin string, lang string) (map[string]interface{}, error){

	if ok := p.Validate(); !ok {
		return utils.ErrorInvalidParameters, errors.New("invalid parameters have been passed")
	}
	
	var trans = GetDB().Begin()
	defer Rollback(trans)
	
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
	var org = Organization{
		Bin:	bin,
		Lang: 	utils.If_condition_then(lang == "kaz", lang, "rus").(string),
	}
	if resp, err := org.Create_or_get_organization_from_db_by_bin(); err != nil {
		return resp, err
	}

	p.Organization = org
	p.OrganizationId = org.Id
	p.Created = time.Now().UTC()

	if err := trans.Table(Project{}.TableName()).Create(p).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}
	
	var pu = ProjectsUsers{
		ProjectId: 	p.Id,
		UserId:    	p.CreatedBy,
	}
	
	if err := trans.Table("projects_users").Create(pu).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}

	trans.Commit()
	return utils.NoErrorFineEverthingOk, nil
}

func (p *Project) Update(bin string, lang string) (map[string]interface{}, error) {
	if p.Id == 0 {
		return utils.ErrorInvalidParameters, errors.New("invalid id. project update")
	}

	if lang == "" {
		lang = utils.DefaultLContentanguage
	}
	
	/*
		get org id
	 */
	var org = &Organization{
		Bin: bin,
		Lang: lang,
	}
	org, err := org.Get_and_assign_info_on_organization_by_bin()
	if err == nil && bin != "" {
		p.OrganizationId = org.Id
	} else {
		return utils.ErrorInvalidParameters, err
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
		Ccode:          	p.Ccode,
		Phone:          	p.Phone,
		OrganizationId: 	p.OrganizationId,
	}).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}

	return utils.NoErrorFineEverthingOk, nil
}
