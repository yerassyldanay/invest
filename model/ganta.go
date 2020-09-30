package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"invest/utils"
)

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
	err = tx.Raw("select * from gantas where is_done = false and ganta_parent_id = 0 order by start_date limit 1;").Scan(g).Error
	return err
}

func (g *Ganta) OnlyGetPreloadedChildStepsByProjectIdAndStep(project_step interface{}, tx *gorm.DB) (err error) {
	err = tx.Preload("Document").Find(&g.GantaChildren, "project_id = ? and step = ? and ganta_parent_id != 0", g.ProjectId, project_step).Error
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

/*
	Project
 */
func (g *Ganta) OnlyUpdateReconsiderStatusByProjectId(status bool, tx *gorm.DB) (err error) {
	err = tx.Model(&Project{Id: g.ProjectId}).Update("reconsider", status).Error
	return err
}

func (g *Ganta) OnlyUpdateRejectStatusByProjectId(status bool, tx *gorm.DB) (err error) {
	err = tx.Model(&Project{Id: g.ProjectId}).Update("reject", status).Error
	return err
}

func (g *Ganta) OnlySetReconsiderStatusForProjectByProjectId(tx *gorm.DB) (err error) {
	err = tx.Model(&Project{Id: g.ProjectId}).Updates(map[string]interface{}{
		"status": utils.ProjectStatusReconsider,
	}).Error
	return err
}

func (g *Ganta) OnlySetRejectStatusForProjectByProjectId(tx *gorm.DB) (err error) {
	err = tx.Model(&Project{Id: g.ProjectId}).Updates(map[string]interface{}{
		"status": utils.ProjectStatusReject,
		"responsible": utils.RoleInvestor,
	}).Error
	return err
}


