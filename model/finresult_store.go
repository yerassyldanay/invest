package model

import (
	"github.com/jinzhu/gorm"
	"invest/utils"
)

/*
	calculate values again to make sure they are valid
 */
func (frc *FinresultCol) Recalculate_sum_to_avoid_misguidance() {
	return
}

func (fr *Finresult) Recalculate_sum_to_avoid_misguidance() {
	fr.TotalIncome.Recalculate_sum_to_avoid_misguidance()
	fr.TotalProduction.Recalculate_sum_to_avoid_misguidance()
	fr.ProductionCost.Recalculate_sum_to_avoid_misguidance()
	fr.OperationalProfit.Recalculate_sum_to_avoid_misguidance()
	fr.Cancellation.Recalculate_sum_to_avoid_misguidance()
	fr.OtherCost.Recalculate_sum_to_avoid_misguidance()
	fr.PrivateCost.Recalculate_sum_to_avoid_misguidance()
}

/*
	load table data
 */
func (fr *Finresult) Load_values_to_this_object_by_project_id() error {
	err := GetDB().Preload("TotalProduction").Preload("TotalIncome").
		Preload("PrivateCost").Preload("OtherCost").
		Preload("OperationalProfit").Preload("Cancellation").
		Preload("ProductionCost").Table(fr.TableName()).Where("project_id=?", fr.ProjectId).
		First(fr).Error

	return err
}

/*
	create a table
 */
func (fr *Finresult) Create_this_table() error {

	fr.Recalculate_sum_to_avoid_misguidance()

	return GetDB().Create(fr).Error
}

/*
	this function will provide a fin table
		in case there is no table it will create a new one and return that table
 */
func (fr *Finresult) Get_finresult_table() (utils.Msg) {
	if fr.ProjectId == 0 {
		return utils.Msg{utils.ErrorInvalidParameters, 400, "", "project id is 0"}
	}

	err := fr.Load_values_to_this_object_by_project_id()

	if err == gorm.ErrRecordNotFound {
		/*
			if not found then create a new fin result table
		 */
		if err = fr.Create_this_table(); err != nil {
			/*
				if could not create then return default one
			 */
			fr = &Finresult{}
		}

	} else if err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map_with_escape(*fr, []string{"project"})

	return utils.Msg{resp, 200, "", ""}
}

/*
	update a table
		the AfterCreate hook will delete unnecessary
*/
func (fr *Finresult) Update_this_table() (utils.Msg) {
	var old_finance = Finresult{
		ProjectId: fr.ProjectId,
	}

	/*
		verify that the sum of the values are correct
	*/
	fr.Recalculate_sum_to_avoid_misguidance()

	/*
		this will load table from db
	*/
	err := old_finance.Load_values_to_this_object_by_project_id()
	if err == gorm.ErrRecordNotFound {
		/*
			if the table is not found then store this table on db
		*/
		fr.Id = 0
		err = fr.Create_this_table()

		if err != nil {
			return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
		}

		return utils.Msg{utils.NoErrorFineEverthingOk, 200, "", ""}

	} else if err != nil {
		/*
			did not expect this kind of error
				thus returning msg
		*/
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	/*
		replacing the ids of tables
			to store the new table on db
		thus can store the new table instead of old one
	*/
	fr.Id = old_finance.Id

	fr.TotalIncomeId = old_finance.TotalIncome.Id
	fr.TotalProductionId = old_finance.TotalProduction.Id

	fr.ProductionCostId = old_finance.ProductionCost.Id
	fr.OperationalProfitId = old_finance.OperationalProfit.Id
	fr.CancellationId = old_finance.Cancellation.Id

	fr.OtherCostId = old_finance.OtherCost.Id
	fr.PrivateCostId = old_finance.PrivateCost.Id

	if err := GetDB().Save(fr).Error; err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	return utils.Msg{utils.NoErrorFineEverthingOk, 200, "", ""}
}
