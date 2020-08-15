package model

import (
	"errors"
	"fmt"
	"invest/utils"
	"strings"
)

func (fi *Finance) Recalculate_sum_to_avoid_missguidance() {
	return
}

func (fi *Finance) Create_and_store_on_db() (map[string]interface{}, error) {
	fi.Recalculate_sum_to_avoid_missguidance()

	var trans = GetDB().Begin()
	defer func() {
		if trans != nil {
			trans.Rollback()
		}
	}()

	//var count int
	//if GetDB().Table(Project{}.TableName()).Where("id=?", fi.ProjectId).Count(&count); count != 1 {
	//	return utils.ErrorInvalidParameters, errors.New("there is no such project")
	//}

	var count int
	if GetDB().Table(Finance{}.TableName()).Where("project_id=?", fi.ProjectId).Count(&count); count != 0 {
			return utils.ErrorMethodNotAllowed, errors.New("the finance table of the project already exists")
	}

	if err := trans.Create(&fi.Land).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}
	fi.LandId = fi.Land.Id

	if err := trans.Create(&fi.Tech).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}
	fi.TechId = fi.Tech.Id

	if err := trans.Create(&fi.Capital).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}
	fi.CapitalId = fi.Capital.Id

	if err := trans.Create(&fi.Other).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}
	fi.OtherId = fi.Other.Id

	if err := trans.Create(&fi.Sum).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}
	fi.SumId = fi.Sum.Id

	if err := trans.Create(fi).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}

	trans.Commit()
	return utils.NoErrorFineEverthingOk, nil
}

/*
	get a finance table
 */
func (fi *Finance) Get_table() (map[string]interface{}, error) {

	err := GetDB().Preload("Project").Preload("Land").
		Preload("Tech").Preload("Capital").Preload("Other").
		Preload("Sum").Table(fi.TableName()).
		Where("project_id=?", fi.ProjectId).
		First(fi).Error

	if err != nil {
		return utils.ErrorInternalDbError, err
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(*fi)

	return resp, nil
}

func get_table_help(result chan bool, fcid uint64, fc *FinanceCol) {
	if err := GetDB().Table(Ganta{}.TableName()).Where("id=?", fcid).First(fc); err != nil {
		fmt.Println("get fin table: ", err)
	}

	result <- true
}

/*
	delete table and recreate
 */
func (fi *Finance) Update_finance_table() (map[string]interface{}, error) {
	var trans = GetDB().Begin()
	defer func() { if trans != nil { trans.Rollback() } }()

	var main_query = `
		delete from finance_cols where id in (
			select fc.id from finance_cols fc join finances fi
			on fc.id = fi.project_id or fc.id = fi.capital_id
			or fc.id = fi.land_id or fc.id = fi.other_id
			or fc.id = fi.sum_id
			where fi.project_id = ?
		);
	`
	main_query = strings.Replace(main_query, "\\n", "", -1)
	main_query = strings.Replace(main_query, "\\t", "", -1)

	if err := trans.Exec(main_query, fi.ProjectId).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}

	main_query = `delete from finances where project_id = ?;`
	if err := trans.Exec(main_query, fi.ProjectId).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}

	_, err := fi.Create_and_store_on_db()
	if err != nil {
		return utils.ErrorInternalDbError, err
	}

	trans.Commit()
	return utils.NoErrorFineEverthingOk, nil
}

