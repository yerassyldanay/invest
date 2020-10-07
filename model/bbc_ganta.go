package model

import (
	"errors"
	"github.com/jinzhu/gorm"
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

	GantaParentId					uint64					`json:"ganta_parent_id"`
	GantaChildren					[]Ganta					`json:"ganta_children" gorm:"-"`

	//IsHidden						bool					`json:"is_hidden" gorm:"default:true"`
	Step							int						`json:"step" gorm:"default:1"`
	Status							string					`json:"status" gorn:"default:'prending_manager'"`

	IsDone							bool					`json:"is_done" gorm:"default:false"`
	Responsible						string					`json:"responsible" gorm:"default:'spk'"`
	IsDocCheck						bool					`json:"-" gorm:"default:false"`
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
	err = tx.Raw("select * from gantas where project_id = ? and ganta_parent_id = 0 and step = ?  order by is_done desc, start_date asc ; ", g.ProjectId, stage).Scan(&gantas).Error
	return gantas, err
}

func (g *Ganta) OnlyGetChildrenByIdAndProjectId(tx *gorm.DB) (error) {
	return tx.Find(g.GantaChildren, "ganta_parent_id = ? and project_id = ?", g.Id, g.ProjectId).Error
}

func (g *Ganta) OnlyGetChildrenByIdAndProjectIdStep(project_step interface{}, tx *gorm.DB) (error) {
	return tx.Find(g.GantaChildren, "ganta_parent_id = ? and project_id = ? and step = ?", g.Id, g.ProjectId, project_step).Error
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
func (g *Ganta) OnlyChangeStatusToDoneById(tx *gorm.DB) (err error) {
	err = tx.Model(&Ganta{Id: g.Id}).Update("is_done", true).Error
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
