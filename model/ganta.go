package model

import (
	"github.com/jinzhu/gorm"
	"github.com/yerassyldanay/invest/utils/errormsg"
	"github.com/yerassyldanay/invest/utils/helper"
	"github.com/yerassyldanay/invest/utils/message"
	"time"
)

func (g *Ganta) Add_new_step() (message.Msg) {

	/*
		validate the ganta step
	 */
	if err := g.Validate(); err != nil {
		return ReturnInternalDbError(err.Error())
	}

	if g.Start != 0 {
		g.StartDate = time.Unix(g.Start, 0)
		if g.StartDate.Before(helper.GetCurrentTime()) {
			g.StartDate = helper.GetCurrentTime()
		}
	}

	/*
		set the start date to max
	 */
	if g.StartDate.IsZero() {
		g.GantaParentId = 0
		g.StartDate = helper.GetCurrentTime()
	}

	if err := GetDB().Create(g).Error; err != nil {
		return ReturnInternalDbError(err.Error())
	}

	var resp = errormsg.NoErrorFineEverthingOk
	//resp["info"] = Struct_to_map_with_escape(*g, []string{"project"})

	return ReturnNoErrorWithResponseMessage(resp)
}



/*
	get ganta steps by:
		* project_id
*/
func (g *Ganta) Get_ganta_with_documents_by_project_id() (message.Msg) {
	var gantas = []Ganta{}
	if err := GetDB().Preload("Document").Find(&gantas, "project_id = ?", g.ProjectId).Error;
		err != nil {
		return message.Msg{errormsg.ErrorInternalDbError, 417, "", err.Error()}
	}

	var gantasMap = []map[string]interface{}{}
	for _, ganta := range gantas {
		gantasMap = append(gantasMap, Struct_to_map(ganta))
	}

	var resp = errormsg.NoErrorFineEverthingOk
	resp["info"] = gantasMap

	return message.Msg{resp, 200, "", ""}
}

/*
	update ganta steps
 */
func (g *Ganta) Update_ganta_step(fields... string) (message.Msg) {
	if err := g.Validate(); err != nil {
		return message.Msg{errormsg.ErrorInvalidParameters, 400, "", err.Error()}
	}

	err := GetDB().Table(g.TableName()).Select(fields).Where("id = ?", g.Id).Updates(map[string]interface{}{
		"kaz": g.Kaz,
		"rus": g.Rus,
		"eng": g.Eng,
		"start_date": g.StartDate,
		"duration_in_days": g.DurationInDays,
	}).Error

	if err != nil {
		return message.Msg{errormsg.ErrorInternalDbError, 417, "", err.Error()}
	}

	return ReturnNoError()
}

/*
	delete ganta step
		* delete ganta step
		* delete all child steps
		* delete all documents
 */
func (g *Ganta) Delete_ganta_step() (message.Msg) {
	return ReturnNoError()
}

/*
	get single ganta step
 */
func (g *Ganta) Get_only_one_with_docs() (message.Msg) {
	if err := GetDB().Preload("Document").First(g, "id = ?", g.Id).Error;
		err != nil {
			return ReturnInternalDbError(err.Error())
	}

	var resp = errormsg.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(*g)

	return ReturnNoError()
}

func (g *Ganta) Change_time() (message.Msg) {

	g.StartDate = time.Unix(g.Start, 0)
	if g.StartDate.Before(helper.GetCurrentTime()) {
		g.StartDate = helper.GetCurrentTime()
	}

	msg := g.Update_ganta_step("start_date", "duration_in_days")
	return msg
}

func (g *Ganta) Add_ganta_step_to_the_top(tx *gorm.DB) (error) {
	var project_id = g.ProjectId

	var tempGanta = Ganta{Id: g.Id}
	err := tempGanta.OnlyGetCurrentStepByProjectId(tx);

	if err == gorm.ErrRecordNotFound {
		// create a new one
		g = &Ganta{
			IsAdditional:   true,
			ProjectId:      project_id,
			Kaz:            g.Kaz,
			Rus:            g.Rus,
			Eng:            g.Eng,
			StartDate:      helper.GetCurrentTime(),
			DurationInDays: g.DurationInDays,
			Step:           g.Step,
			Status:         g.Status,
			Responsible:    g.Responsible,
			IsDocCheck:     g.IsDocCheck,
		}
	} else if err != nil {
		return err
	} else {
		// this puts this gantt step on the top
		g.StartDate = tempGanta.StartDate.Add(time.Hour * (-1))
	}

	// make sure id is 0
	g.Id = 0

	// prettify name of the gantt step
	if err = g.Validate(); err != nil {
		return err
	}

	// create a new gantt step
	if err := g.OnlyCreate(tx); err != nil {
		return err
	}

	return nil
}

/*
	check whether this ganta step already possesses a document
*/
func (g *Ganta) Does_this_ganta_step_has_document(trans *gorm.DB) (bool) {
	var document = Document{}
	_ = trans.First(document, "ganta_id = ?", g.Id).Error
	return document.Id != 0
}
