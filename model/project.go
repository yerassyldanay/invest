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
var errorProjectInvalidProjectName = errors.New("invalid project name")
var errorProjectInvalidProjectDescription = errors.New("invalid project description")
var errorProjectInvalidEmployeeCount = errors.New("invalid employee count")

func (p *Project) Validate() (error) {
	switch {
	case p.OfferedById < 1:
		return errorProjectInvalidOfferedById
	case p.OfferedByPosition == "":
		return errorProjectInvalidInitiatorPosition
	case p.Name == "":
		return errorProjectInvalidProjectName
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

// save
func (p *Project) OnlySave(tx *gorm.DB) (err error) {
	err = tx.Save(p).Error
	return err
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

// get projects of manager, lawyer and financier (and other spk user if there is any)
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

	// get status & step of the project by gantt step
	var ganta = Ganta{ProjectId: p.Id}
	err = ganta.OnlyGetCurrentStepByProjectId(tx)

	switch {
	case err == nil && !utils.Does_a_slice_contain_element([]string{utils.ProjectStatusPreliminaryReject,
			utils.ProjectStatusPreliminaryReconsider,
				utils.ProjectStatusReject}, p.Status):
		p.CurrentStep = ganta
		p.Step = ganta.Step
		// status will be changed
		// if the project is not rejected by spk or put into reconsideration
		p.Status = ganta.Status
	case err == nil:
		// in case it is preliminary reject or reconsider (status is set by manager or expert)
		// no need to set step & status to new values
		p.Step = ganta.Step
		p.CurrentStep = ganta
	case err == gorm.ErrRecordNotFound :
		p.Step = 3
		p.Status = utils.ProjectStatusAgreement
		p.CurrentStep = DefaultGantaFinalStep
	default:
		return err
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
	if err == gorm.ErrRecordNotFound {
		// pass
	} else if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(*p)

	return ReturnNoErrorWithResponseMessage(resp)
}

