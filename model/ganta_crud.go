package model

import (
	"invest/utils"
	"time"
)

func (g *Ganta) Add_new_step() (utils.Msg) {

	/*
		validate the ganta step
	 */
	if err := g.Validate(); err != nil {
		return ReturnInternalDbError(err.Error())
	}

	if g.Start != 0 {
		g.StartDate = time.Unix(g.Start, 0)
		if g.StartDate.Before(utils.GetCurrentTime()) {
			g.StartDate = utils.GetCurrentTime()
		}
	}

	/*
		set the start date to max
	 */
	if g.StartDate.IsZero() {
		g.GantaParentId = 0
		g.StartDate = utils.GetCurrentTime()
	}

	if err := GetDB().Create(g).Error; err != nil {
		return ReturnInternalDbError(err.Error())
	}

	var resp = utils.NoErrorFineEverthingOk
	//resp["info"] = Struct_to_map_with_escape(*g, []string{"project"})

	return ReturnNoErrorWithResponseMessage(resp)
}



/*
	get ganta steps by:
		* project_id
*/
func (g *Ganta) Get_ganta_with_documents_by_project_id() (utils.Msg) {
	var gantas = []Ganta{}
	if err := GetDB().Preload("Document").Find(&gantas, "project_id = ?", g.ProjectId).Error;
		err != nil {
		return utils.Msg{ utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	var gantasMap = []map[string]interface{}{}
	for _, ganta := range gantas {
		gantasMap = append(gantasMap, Struct_to_map(ganta))
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = gantasMap

	return utils.Msg{resp, 200, "", ""}
}

/*
	update ganta steps
 */
func (g *Ganta) Update_ganta_step(fields... string) (utils.Msg) {
	if err := g.Validate(); err != nil {
		return utils.Msg{utils.ErrorInvalidParameters, 400, "", err.Error()}
	}

	err := GetDB().Table(g.TableName()).Select(fields).Where("id = ?", g.Id).Updates(map[string]interface{}{
		"kaz": g.Kaz,
		"rus": g.Rus,
		"eng": g.Eng,
		"start_date": g.StartDate,
		"duration_in_days": g.DurationInDays,
	}).Error

	if err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	return MsgNoErrorEverythingIsOk
}

/*
	delete ganta step
		* delete ganta step
		* delete all child steps
		* delete all documents
 */
func (g *Ganta) Delete_ganta_step() (utils.Msg) {
	return MsgNoErrorEverythingIsOk
}

/*
	get single ganta step
 */
func (g *Ganta) Get_only_one_with_docs() (utils.Msg) {
	if err := GetDB().Preload("Document").First(g, "id = ?", g.Id).Error;
		err != nil {
			return ReturnInternalDbError(err.Error())
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(*g)

	return ReturnNoError()
}

func (g *Ganta) Change_time() (utils.Msg) {

	g.StartDate = time.Unix(g.Start, 0)
	if g.StartDate.Before(utils.GetCurrentTime()) {
		g.StartDate = utils.GetCurrentTime()
	}

	return g.Update_ganta_step("start_date", "duration_in_days")
}
