package model

import (
	"fmt"
	"invest/utils"
	"strings"
)

func (fr *Finresult) Recalculate_sum_values() {
	return
}

func (fr *Finresult) Create_and_store_financial_results_of_project_on_db() (map[string]interface{}, error) {
	fr.Recalculate_sum_values()

	var trans = GetDB().Begin()
	defer func() {
		if trans != nil {
			trans.Rollback()
		}
	}()

	var count int
	if err := GetDB().Table(Finresult{}.TableName()).Where("project_id=?", fr.ProjectId).Count(&count).Error; count != 0 || err != nil {
		//fmt.Println("create finresult err: ", err)
		return utils.ErrorMethodNotAllowed, err
	}

	/*

	 */
	if err := trans.Create(&fr.TotalIncome).Error; err != nil {
		return utils.ErrorInvalidParameters, err
	}
	fr.TotalIncomeId = fr.TotalIncome.Id

	if err := trans.Create(&fr.TotalProduction).Error; err != nil {
		return utils.ErrorInvalidParameters, err
	}
	fr.TotalProductionId = fr.TotalProduction.Id

	/*
		ProductionCost
	 */
	if err := trans.Create(&fr.ProductionCost).Error; err != nil {
		return utils.ErrorInvalidParameters, err
	}
	fr.PrivateCostId = fr.PrivateCost.Id

	if err := trans.Create(&fr.OperationalProfit).Error; err != nil {
		return utils.ErrorInvalidParameters, err
	}
	fr.OperationalProfitId = fr.OperationalProfit.Id

	/*
		Cancellation
	 */
	if err := trans.Create(&fr.Cancellation).Error; err != nil {
		return utils.ErrorInvalidParameters, err
	}
	fr.CancellationId = fr.Cancellation.Id

	if err := trans.Create(&fr.OtherCost).Error; err != nil {
		return utils.ErrorInvalidParameters, err
	}
	fr.OtherCostId = fr.OtherCost.Id

	if err := trans.Create(&fr.PrivateCost).Error; err != nil {
		return utils.ErrorInvalidParameters, err
	}
	fr.PrivateCostId = fr.PrivateCost.Id

	/*
		Create fin table
	 */
	if err := trans.Create(fr).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}

	fmt.Println(trans)
	trans.Commit()
	trans = nil
	return utils.NoErrorFineEverthingOk,nil
}

func (fr *Finresult) Get_finresult_table() (map[string]interface{}, error) {
	err := GetDB().Preload("TotalProduction").Preload("TotalIncome").
		Preload("PrivateCost").Preload("OtherCost").
		Preload("OperationalProfit").Preload("Cancellation").
		Preload("ProductionCost").Table(fr.TableName()).Where("project_id=?", fr.ProjectId).
		First(fr).Error

	if err != nil {
		return utils.ErrorInternalDbError, err
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(*fr)

	return resp, nil
}

/*
	to run using goroutine
 */
func get_finresult__help(result chan bool, id uint64, fc *FinresultCol) {
	if err := GetDB().Table(FinanceCol{}.TableName()).Where("id=?", id).First(fc).Error;
		err != nil {
			fmt.Println("get finresult column: ", err)
	}

	result <- true
}

/*
	delete table and recreate
*/
func (fi *Finresult) Update_finresult_table() (map[string]interface{}, error) {
	var trans = GetDB().Begin()
	defer func() { if trans != nil { trans.Rollback() } }()

	var main_query = `
		delete from finresult_cols where id in (
			select fc.id from finresult_cols fc join finresults fi
				on fc.id = fi.cancellation_id
				or fc.id = fi.operational_profit_id
				or fc.id = fi.other_cost_id
				or fc.id = fi.private_cost_id
				or fc.id = fi.production_cost_id
				or fc.id = fi.total_income_id
				or fc.id = fi.total_production_id
			where project_id = ?
		);
	`
	main_query = strings.Replace(main_query, "\\n", "", -1)
	main_query = strings.Replace(main_query, "\\t", "", -1)

	if err := trans.Exec(main_query, fi.ProjectId).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}

	main_query = `delete from finresults where project_id = ?;`
	if err := trans.Exec(main_query, fi.ProjectId).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}

	_, err := fi.Create_and_store_financial_results_of_project_on_db()
	if err != nil {
		return utils.ErrorInternalDbError, err
	}

	trans.Commit()
	return utils.NoErrorFineEverthingOk, nil
}
