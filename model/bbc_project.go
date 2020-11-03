package model

import (
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"invest/utils/constants"
	"invest/utils/errormsg"
	"invest/utils/message"
)

/*
	A project might have different statuses, which are mainly set by the current
		(top after sorting) Gantt step

	However, there are some nuances:
		* pre-reject, pre-reconsider are the statuses,
		which are set to this project apart from the Gantt step

		These statuses indicates admin that a project was not accepted
		by manager or expert
 */
type Project struct {
	Id					uint64					`json:"id" gorm:"primary key"`

	Name				string						`json:"name" validate:"required"`
	Description			string						`json:"description" gorm:"default:''" validate:"required"`

	Info				string						`json:"-" default:"''"`
	InfoSent			map[string]interface{}		`json:"info_sent" gorm:"-"`
	
	EmployeeCount			uint					`json:"employee_count" validate:"required"`

	Email					string						`json:"email" gorm:"default:''"`
	PhoneNumber				string						`json:"phone_number" gorm:"default:''"`

	OrganizationId			uint64					`json:"organization_id"`
	Organization			Organization			`json:"organization" gorm:"foreignkey:OrganizationId"`

	Users					[]User					`json:"user" gorm:"many2many:projects_users;"`
	//Documents				[]Document				`json:"documents" gorm:"-"`
	Categors				[]Categor				`json:"categors" gorm:"many2many:projects_categors"`

	OfferedById					uint64					`json:"offered_by_id" gorm:"not null"`
	OfferedByPosition			string					`json:"offered_by_position" gorm:"not null"`

	Status						string					`json:"status" gorm:"default:'pending_admin'"`
	Step						int						`json:"step"`
	
	LandPlotFrom				string					`json:"land_plot_from" gorm:"default:'investor'"`
	LandArea 					int						`json:"land_area"`
	LandAddress					string					`json:"land_address"`
	
	IsManagerAssigned			bool					`json:"is_manager_assigned" gorm:"default:false"`
	CurrentStep						Ganta					`json:"current_ganta_step" gorm:"-"`
	
	AddInfo
}

type ProjectWithFinanceTables struct {
	Project						Project					`json:"project"`
	Cost						Cost					`json:"cost"`
	Finance						Finance					`json:"finance"`
}

type ProjectExtended struct {
	Project
	Cost						Cost					`json:"cost"`
	Finance						Finance					`json:"finance"`
}

func (Project) TableName() string {
	return "projects"
}

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

// update fields
func (p *Project) OnlyUpdateById(tx *gorm.DB, fields ... interface{}) error {
	err := tx.Model(Project{Id: p.Id}).Select(fields).Updates(p).Error
	return err
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

func (p *Project) OnlyGetAny(tx *gorm.DB) error {
	err := tx.First(p).Error
	return err
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
func (p *Project) OnlyGetProjectsOfSpkUsers(user_id uint64, statuses []string, steps []int, offset interface{}, tx *gorm.DB) (projects []Project, err error) {
	err = tx.Preload("Organization").Table("projects as p").
		Joins("join projects_users pu on p.id = pu.project_id").Select("p.*").
		Order("p.created desc").Offset(offset).Limit(GetLimit).
		Find(&projects, "pu.user_id = ? and p.status in (?) and step in (?)", user_id, statuses, steps).Error
	return projects, err
}

// get projects of a particular investor
func (p *Project) OnlyGetProjectsOfInvestor(user_id uint64, statuses []string, steps []int, offset interface{}, tx *gorm.DB) (projects []Project, err error) {
	err = tx.Preload("Organization").Order("created desc").Offset(offset).Limit(GetLimit).
		Find(&projects, "offered_by_id = ? and status in (?) and step in (?)", user_id, statuses, steps).Error
	return projects, err
}

// get all projects, but based on statuses
func (p *Project) OnlyGetProjectsByStatusesAndSteps(offset interface{}, statuses []string, steps []int, tx *gorm.DB) (projects []Project, err error) {
	err = tx.Preload("Organization").Order("created desc").
		Find(&projects, "status in (?) and step in (?)", statuses, steps).Error
	return projects, err
}

func (p *Project) GetAndUpdateStatusOfProject(tx *gorm.DB) (err error) {
	// get project
	err = p.OnlyGetByIdPreloaded(tx)
	if err != nil {
		return err
	}

	// get status & step of the project by gantt step
	var ganta = Ganta{ProjectId: p.Id}
	err = ganta.OnlyGetCurrentStepByProjectId(tx)

	switch {
	case err == nil && constants.ProjectStatusReject == p.Status:
		p.Step = ganta.Step
		p.CurrentStep = ganta
	case err == nil:
		p.CurrentStep = ganta
		p.Step = ganta.Step
		// status will be changed
		// if the project is not rejected by spk or put into reconsideration
		p.Status = ganta.Status
	case err == gorm.ErrRecordNotFound:
		p.Step = 3
		p.Status = constants.ProjectStatusAgreement
		p.CurrentStep = DefaultGantaFinalStep
	default:
		return err
	}

	err = tx.Save(p).Error
	return err
}

func (p *Project) Get_project_with_current_status() (message.Msg) {
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

	var resp = errormsg.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(*p)

	return ReturnNoErrorWithResponseMessage(resp)
}

