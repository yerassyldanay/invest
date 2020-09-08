package model

import (
	"errors"
	"invest/utils"
	"time"
)

/*
	one or none of the field out of kaz, rus & eng might be set
		fill them by hand
 */
func (g *Ganta) Validate() bool {
	return true
}

func (g *Ganta) Add_new_step() (map[string]interface{}, error) {

	if ok := g.Validate(); !ok {
		return utils.ErrorInvalidParameters, errors.New("invalid parameters (lang) have been provided")
	}

	if g.Start.IsZero() {
		var tganta = Ganta{}
		if err := GetDB().Table(Ganta{}.TableName()).Where("project_id = $1 and start = (select max(start) from gantas where project_id = $1)", g.ProjectId).First(&tganta).Error
			err != nil {
			return utils.ErrorInternalDbError, err
		}
		g.Start = tganta.Start.Add(time.Hour * GantaDefaultStepHours)
	}

	if err := GetDB().Create(g).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map_with_escape(*g, []string{"project"})

	return resp, nil
}

func (g *Ganta) Get_ganta_by_project_id() (map[string]interface{}, error) {
	var gantas = struct {
		Ganta				[]Ganta					`json:"ganta"`
	}{}
	if err := GetDB().Table(Ganta{}.TableName()).Where("project_id=?", g.ProjectId).Find(&gantas.Ganta).Error;
		err != nil {
			return utils.ErrorInternalDbError, err
	}


	var ganmap = []map[string]interface{}{}
	for _, ganta := range gantas.Ganta {
		ganmap = append(ganmap, Struct_to_map_with_escape(ganta, []string{"project", "organization"}))
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = ganmap

	return resp, nil
}

func (gu *GantaUpDate) Update_step_start_thus_others() (map[string]interface{}, error) {
	if ok := gu.Ganta.Validate(); !ok {
		return utils.ErrorInvalidParameters, errors.New("empty lang fields. ganta up date")
	}

	rows, err := GetDB().Table(Ganta{}.TableName()).Where("project_id=1 and start >= (select start from gantas where id = ?)", gu.ProjectId, gu.Id).Rows()
	if err != nil {
		return utils.ErrorInternalDbError, err
	}
	defer rows.Close()

	var ganta = Ganta{}
	for rows.Next() {
		if err := GetDB().ScanRows(rows, &ganta); err != nil {
			continue
		}

		ganta.Start.Add(time.Hour * (24 * gu.Day + gu.Hour))
		if err := GetDB().Save(&ganta).Error; err != nil {
			break
		}
	}

	return utils.NoErrorFineEverthingOk, nil
}

/*
	delete ganta step
 */
func (g *Ganta) Delete_ganta_step() (map[string]interface{}, error) {
	if err := GetDB().Table(Ganta{}.TableName()).Where("project_id=? and id=?", g.ProjectId, g.Id).Delete(&Ganta{}).Error;
		err != nil {
			return utils.ErrorInternalDbError, err
	}

	return utils.NoErrorFineEverthingOk, nil
}