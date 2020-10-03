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

// get by id (+ add info about organization)
func (p *Project) OnlyGetByIdPreloaded(tx *gorm.DB) (err error) {
	err = tx.Preload("Organization").First(p, "id = ?", p.Id).Error
	return err
}

func (p *Project) OnlyGetCategorsByProjectId(trans *gorm.DB) (err error) {
	err = trans.Raw("select distinct c.* from projects_categors pc join categors c on pc.categor_id = c.id where pc.project_id = ? ;", p.Id).Scan(&p.Categors).Error
	return err
}

func (p *Project) OnlyGetAssignedUsersByProjectId(trans *gorm.DB) (err error) {
	err = trans.Preload("Email").Preload("Role").Preload("Phone").
		Find(&p.Users, "id in (select user_id from projects_users where project_id = ?)", p.Id).
		Error
	return err
}

func (p *Project) OnlyPreloadOrganizationByOrganizationId(tx *gorm.DB) (err error) {
	err = tx.First(&p.Organization, "id = ?", p.OrganizationId).Error
	return err
}

// get projects of manager, lawyer and financier (and other spk users if they are)
// except for an admin, who has access to all projects
func (p *Project) OnlyGetProjectsOfSpkUsers(user_id uint64, statuses []string, offset interface{}, tx *gorm.DB) (projects []Project, err error) {
	err = tx.Preload("Organization").Table("projects as p").
		Joins("join projects_users pu on p.id = pu.project_id").Select("p.*").
		Find(&projects, "pu.user_id = ? and p.status in (?)", user_id, statuses).
		Offset(offset).Limit(GetLimit).Error
	return projects, err
}

// get projects of a particular investor
func (p *Project) OnlyGetProjectsOfInvestor(user_id uint64, statuses []string, offset interface{}, tx *gorm.DB) (projects []Project, err error) {
	err = tx.Preload("Organization").Find(&projects, "offered_by_id = ? and status in (?)", user_id, statuses).
		Offset(offset).Limit(GetLimit).Error
	return projects, err
}

// get all projects, but based on statuses
func (p *Project) OnlyGetProjectsByStatuses(offset interface{}, statuses []string, tx *gorm.DB) (projects []Project, err error) {
	err = tx.Preload("Organization").Find(&projects, "status in (?)", statuses).Offset(offset).Limit(GetLimit).Error
	return projects, err
}

func (p *Project) GetAndUpdateStatusOfProject(tx *gorm.DB) (err error) {
	err = p.OnlyGetByIdPreloaded(tx)
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

	err = tx.Save(p).Error
	return err
}

func (p *Project) Get_project_with_current_status() (utils.Msg) {
	// update & set status
	// set the current ganta step
	err := p.GetAndUpdateStatusOfProject(GetDB())
	if err != nil {
		return ReuturnInternalServerError(err.Error())
	}

	// no need
	err = p.CurrentStep.OnlyGetCurrentStepByProjectId(GetDB())
	if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(*p)

	return ReturnNoErrorWithResponseMessage(resp)
}

