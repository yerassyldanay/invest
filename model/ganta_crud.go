package model

import (
	"invest/utils"
	"time"
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

	for _, lang := range []*string{&g.Kaz, &g.Rus, &g.Eng} {
		if *(lang) != "" {
			*lang = temp
		}
	}

	return true
}

/*
	one or none of the field out of kaz, rus & eng might be set
		fill them by hand
 */
func (g *Ganta) IsValid() bool {
	return g.prepare_name_and_validate() && g.ProjectId != 0 && g.DurationInDays > 0
}

func (g *Ganta) Add_new_step() (utils.Msg) {
	/*
		validate the ganta step
	 */
	if ok := g.IsValid(); !ok {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", "invalid parameters (lang) have been provided"}
	}

	/*
		set the start date to max
	 */
	if g.StartDate.IsZero() {
		g.GantaParentId = 0
		var tganta = Ganta{}
		if err := GetDB().First(&tganta, "project_id = $1 and start = (select max(start) from gantas where project_id = $1)", g.ProjectId).Error
		//	err == gorm.ErrRecordNotFound {
		//	g.StartDate = utils.GetCurrentTime()
		//} else if
		err != nil {
			return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
		} else {
			g.StartDate = tganta.StartDate.Add(time.Hour * time.Duration(tganta.DurationInDays) * 24)
		}
	}

	if err := GetDB().Create(g).Error; err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	var resp = utils.NoErrorFineEverthingOk
	//resp["info"] = Struct_to_map_with_escape(*g, []string{"project"})

	return utils.Msg{resp, 200, "", ""}
}

/*
	get ganta steps by:
		* project_id
 */
func (g *Ganta) Get_only_ganta_by_project_id() (utils.Msg) {
	var gantas = []Ganta{}
	if err := GetDB().Find(&gantas, "project_id = ?", g.ProjectId).Error;
		err != nil {
			return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
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
func (g *Ganta) Update_ganta_step() (utils.Msg) {
	if ok := g.IsValid(); !ok {
		return utils.Msg{utils.ErrorInvalidParameters, 400, "", "empty lang fields. ganta up date"}
	}

	err := GetDB().Table(g.TableName()).Where("id = ?", g.Id).Updates(map[string]interface{}{
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
			return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(*g)

	return utils.Msg{resp, 200, "", ""}
}
