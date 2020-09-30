package model

import (
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"invest/utils"
)

/*
	errors to validate
 */
var errorProjectInvalidOfferedById = errors.New("id of user (who offered) is not indicated")
var errorProjectInvalidInitiatorPosition = errors.New("invalid position of an initiator")
var errorProjectInvlaidProjectName = errors.New("invalid project name")
var errorProjectInvalidProjectDescription = errors.New("invalid project description")
var errorProjectInvalidEmployeeCount = errors.New("invalid employee count")

func (p *Project) Validate() (error) {
	switch {
	case p.OfferedById < 1:
		return errorProjectInvalidOfferedById
	case p.OfferedByPosition == "":
		return errorProjectInvalidInitiatorPosition
	case p.Name == "":
		return errorProjectInvlaidProjectName
	case p.Description == "":
		return errorProjectInvalidProjectDescription
	case p.EmployeeCount < 1:
		return errorProjectInvalidEmployeeCount
	}

	return nil
}

func (p *Project) OnlyUpdateStatusById (tx *gorm.DB) (error) {
	return tx.Model(&Project{Id: p.Id}).Update("status", p.Status).Error
}

// check whether this is an investor of the project
func (p *Project) OnlyCheckInvestorByProjectAndInvestorId(tx *gorm.DB) (err error) {
	err = tx.First(p, "id = ? and offered_by_id = ?", p.Id, p.OfferedById).Error
	return err
}

// check whether this is a user, who is assign to the project
func (p *Project) OnlyCheckUserByProjectAndUserId(project_id uint64, user_id uint64, tx *gorm.DB) (err error) {
	err = tx.Raw("select p.* from projects_users pu join projects p on pu.project_id = p.id where pu.project_id = ? and pu.user_id = ? limit 1;", project_id, user_id).Scan(p).Error
	return err
}

// create only
func (p *Project) OnlyCreate(tx *gorm.DB) error {
	return tx.Create(p).Error
}

func (p *Project) OnlyUnmarshalInfo() (err error) {
	err = json.Unmarshal([]byte(p.Info), &p.InfoSent)
	p.Info = ""
	return err
}

// get project by id
func (p *Project) OnlyGetById(trans *gorm.DB) error {
	return trans.First(p, "id=?", p.Id).Error
}

func (p *Project) OnlyGetCategorsByProjectId(trans *gorm.DB) (err error) {
	err = trans.Raw("select distinct c.* from projects_categors pc join categors c on pc.categor_id = c.id where pc.project_id = ? ;", p.Id).Scan(&p.Categors).Error
	return err
}

func (p *Project) OnlyGetAssignedUsers(trans *gorm.DB) (err error) {
	return trans.Preload("Email").Preload("Role").Omit("password, created").Find(&p.Users, "id in (select user_id from projects_users where project_id = ?)", p.Id).Error
}

func (p *Project) GetAndUpdateStatusOfProject(tx *gorm.DB) (err error) {
	err = tx.First(p, "id = ?", p.Id).Error
	if err != nil {
		return err
	}

	// get status & step of the project by ganta step
	var ganta = Ganta{}
	err = ganta.OnlyGetCurrentStepByProjectId(tx)

	if err == nil {
		p.Status = ganta.Status
		p.Step = ganta.Step
		p.CurrentStep = ganta
	} else if err == gorm.ErrRecordNotFound {
		// means the project is finished
		p.Completed = true
	} else {
		return err
	}

	// if the status is such then ganta step will not be considered
	if p.Reject || p.Reconsider {
		return nil
	}

	return tx.Save(p).Error
}

func (p *Project) Get_project_with_current_status() (utils.Msg) {
	err := p.GetAndUpdateStatusOfProject(GetDB())
	if err != nil {
		return ReuturnInternalServerError(err.Error())
	}

	err = p.CurrentStep.OnlyGetCurrentStepByProjectId(GetDB())
	if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(*p)

	return ReturnNoErrorWithResponseMessage(resp)
}

