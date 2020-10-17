package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"invest/utils"
	"strconv"
	"time"
)

/*
	Explanation of the most important part.

	This is a Gantt step, which will be shown in the Gantt table.
	There might be child steps, but after some reconsideration, it is decided not to
	use it

	Note:
	Status of the project: The Gantt table will be sorted by start_date and step (for each project).
	The top Gantt step will be considered as the status of the project
	There is a field called 'responsible', which is indicates who is responsible for the current step
		of the project
 */
type Ganta struct {
	Id								uint64					`json:"id" gorm:"primary key"`

	IsAdditional					bool					`json:"is_additional" gorm:"default:false"`
	ProjectId						uint64					`json:"project_id" gorm:"foreignkey:projects.id"`
	//Project							Project					`json:"project" gorm:"foreignkey:ProjectId"`

	Kaz								string					`json:"kaz" gorm:"default:''"`
	Rus								string					`json:"rus" gorm:"default:''"`
	Eng								string					`json:"eng" gorm:"default:''"`

	Start 							int64					`json:"start" gorm:"-"`
	StartDate						time.Time				`json:"start_date" gorm:"default:now()"`
	DurationInDays					time.Duration			`json:"duration_in_days"`
	Deadline						time.Time				`json:"deadline" gorm:"default:null"`

	GantaParentId					uint64					`json:"ganta_parent_id"`
	GantaChildren					[]Ganta					`json:"ganta_children" gorm:"-"`

	//IsHidden						bool					`json:"is_hidden" gorm:"default:true"`
	Step							int						`json:"step" gorm:"default:1"`
	Status							string					`json:"status" gorn:"default:'prending_manager'"`

	IsDone							bool					`json:"is_done" gorm:"default:false"`
	Responsible						string					`json:"responsible" gorm:"default:'spk'"`
	IsDocCheck						bool					`json:"-" gorm:"default:false"`

	NotToShow							bool					`json:"-" gorm:"default:false"`
}

func (Ganta) TableName() string {
	return "gantas"
}

/*
	prettify the description
*/
func (g *Ganta) prepare_name_and_validate() bool {

	var temp string
	for _, lang := range []string{g.Kaz, g.Rus, g.Eng} {
		if lang != "" {
			temp = lang
			break
		}
	}

	if temp == "" {
		return false
	}

	if g.Kaz == "" {
		g.Kaz = temp
	}

	if g.Rus == "" {
		g.Rus = temp
	}

	if g.Eng == "" {
		g.Eng = temp
	}

	return true
}

/*
	one or none of the field out of kaz, rus & eng might be set
		fill them by hand
*/
var errorGantaInvalidName = errors.New("invalid ganta name")
var errorGantaInvalidProjectId = errors.New("invalid project id")
var errorGantaInvalidDuration = errors.New("invalid duration in days")

func (g *Ganta) Validate() error {
	switch {
	case !g.prepare_name_and_validate():
		return errorGantaInvalidName
	case g.ProjectId < 1:
		return errorGantaInvalidProjectId
	case g.DurationInDays < 1:
		return errorGantaInvalidDuration
	}

	return nil
}

func (g *Ganta) OnlyCreate(tx *gorm.DB) (error) {
	return tx.Create(g).Error
}

func (g *Ganta) OnlyGetGantaById(tx *gorm.DB) (error) {
	return tx.First(g, "id = ?", g.Id).Error
}

func (g *Ganta) OnlyGetParentsByProjectId(stage interface{}, tx *gorm.DB) (gantas []Ganta, err error) {
	err = tx.Raw("select * from gantas where not_to_show = false and project_id = ? and ganta_parent_id = 0 and step = ?  order by is_done desc, start_date asc ; ", g.ProjectId, stage).Scan(&gantas).Error
	return gantas, err
}

func (g *Ganta) OnlyGetChildrenByIdAndProjectId(tx *gorm.DB) (error) {
	return tx.Find(g.GantaChildren, "not_to_show = false and ganta_parent_id = ? and project_id = ?", g.Id, g.ProjectId).Error
}

func (g *Ganta) OnlyGetChildrenByIdAndProjectIdStep(project_step interface{}, tx *gorm.DB) (error) {
	return tx.Find(g.GantaChildren, "not_to_show = false and ganta_parent_id = ? and project_id = ? and step = ?", g.Id, g.ProjectId, project_step).Error
}

func (g *Ganta) OnlyUpdateStartDatesOfAllUndoneGantaStepsByProjectId(shiftInHours int, tx *gorm.DB) (err error) {
	var shift = strconv.Itoa(shiftInHours)
	if shiftInHours == -1 || shiftInHours == 1 {
		shift = shift + " hour"
	} else {
		shift = shift + " hours"
	}

	err = tx.Exec("update gantas set start_date = start_date + $1, deadline = deadline + $1 where project_id = $2 and is_done = false;", shift, g.ProjectId).Error
	return err
}

func (g *Ganta) OnlyCountChildrenByIdAndProjectIdStep(project_step interface{}, tx *gorm.DB) (count int, err error) {
	var counter = struct {
		Count				int
	}{}
	err = tx.Raw("select count(*) as count from gantas where ganta_parent_id = ? and project_id = ? and step = ? ; ", g.Id, g.ProjectId, project_step).Scan(&counter).Error
	return counter.Count, err
}

func (g *Ganta) OnlyGetCurrentStepByProjectId(tx *gorm.DB) (err error) {
	err = tx.Raw("select * from gantas where is_done = false and ganta_parent_id = 0 and project_id = ? order by start_date limit 1;", g.ProjectId).Scan(g).Error
	if err == gorm.ErrRecordNotFound {
		var temp = DefaultGantaFinalStep
		g = &temp
		return nil
	}

	return err
}

func (g *Ganta) OnlyGetById(trans *gorm.DB) error {
	return trans.First(g, "id = ?", g.Id).Error
}

// must never be changed
func (g *Ganta) OnlyChangeStatusById(tx *gorm.DB) (err error) {
	return nil
}

// 'is_done' field is set to true
func (g *Ganta) OnlyChangeStatusToDoneAndUpdateDeadlineById(tx *gorm.DB) (err error) {
	var days = int(utils.GetCurrentTruncatedDate().Sub(g.StartDate).Hours() / 24)

	switch {
	case days == 0:
		days = 1
		g.StartDate = utils.GetCurrentTruncatedDate().Add(time.Hour * (-24))
	case days < 0:
		days = 1
		g.StartDate = utils.GetCurrentTruncatedDate().Add(time.Hour * time.Duration(days))
	}

	// deadline ends up where it is
	g.Deadline = utils.GetCurrentTruncatedDate()

	err = tx.Model(&Ganta{Id: g.Id}).Updates(map[string]interface{}{
		"is_done": true,
		"deadline": g.Deadline,
		"start_date": g.StartDate,
		"duration_in_days": days,
	}).Error
	return err
}

// 'is_done' field is set to true
func (g *Ganta) OnlyUpdateStartDateById(tx *gorm.DB) (err error) {
	g.Deadline = g.StartDate.Add(g.DurationInDays * time.Hour * 24)
	err = tx.Model(&Ganta{Id: g.Id}).Updates(map[string]interface{}{
		"deadline": g.Deadline,
		"start_date": g.StartDate,
	}).Error
	return err
}

func (g *Ganta) OnlyUpdateTimeByIdAndProjectId(tx *gorm.DB) (err error) {
	err = tx.Model(&Ganta{}).Where("id = ? and project_id = ?", g.Id, g.ProjectId).
		Updates(map[string]interface{}{
			"start_date": g.StartDate,
			"duration_in_days": g.DurationInDays,
	}).Error
	return err
}

func (g *Ganta) Get_current_ganta_step_and_handle_error_by_project_id(tx *gorm.DB) (err error) {
	if err := g.OnlyGetCurrentStepByProjectId(tx); err == gorm.ErrRecordNotFound {
		// send default final step in case it is the last gantt step
		var project = Project{Id: g.ProjectId}
		_ = project.OnlyGetById(tx)

		// set the current step to the final gantt step
		project.CurrentStep = DefaultGantaFinalStep

	} else if err != nil {
		// an unknown error occurred
		return err
	}

	return nil
}