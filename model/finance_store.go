package model

import (
	"github.com/jinzhu/gorm"
	"invest/utils"
)

func (fic *FinanceCol) Recalculate_sum_to_avoid_misguidance() {
	fic.Sum = fic.Initiator + fic.SPK + fic.Funded
}

func (fi *Finance) Recalculate_sum_to_avoid_misguidance() {
	fi.Land.Recalculate_sum_to_avoid_misguidance()
	fi.Capital.Recalculate_sum_to_avoid_misguidance()
	fi.Tech.Recalculate_sum_to_avoid_misguidance()
	fi.Other.Recalculate_sum_to_avoid_misguidance()

	fi.Sum.SPK = fi.Land.SPK + fi.Capital.SPK + fi.Tech.SPK + fi.Other.SPK
	fi.Sum.Funded = fi.Land.Funded + fi.Capital.Funded + fi.Tech.Funded + fi.Other.Funded
	fi.Sum.Initiator = fi.Land.Initiator + fi.Capital.Initiator + fi.Tech.Initiator + fi.Other.Initiator
	fi.Sum.Sum = fi.Land.Sum + fi.Capital.Sum + fi.Tech.Sum + fi.Other.Sum

	//fmt.Println(*fi)
}

/*
	create a project in one place
		expects that values are valid
 */
func (fi *Finance) Create_this_table() error {

	fi.Recalculate_sum_to_avoid_misguidance()

	//_ = fi.

	err := GetDB().Create(fi).Error
	return err
}

/*
	load data in one place
 */
func (fi *Finance) Load_values_to_this_object_by_project_id() error {
	err := GetDB().Preload("Project").Preload("Land").
		Preload("Tech").Preload("Capital").Preload("Other").
		Preload("Sum").Table(fi.TableName()).
		Where("project_id=?", fi.ProjectId).
		First(fi).Error

	return err
}

/*
	updating table values
 */
func (fi *Finance) Update_finance_table_with_this_table_by_project_id() (utils.Msg) {
	var old_finance = Finance{
		ProjectId: fi.ProjectId,
	}

	/*
		verify that the sum of the values are correct
	 */
	fi.Recalculate_sum_to_avoid_misguidance()

	/*
		this will load table from db
	 */
	err := old_finance.Load_values_to_this_object_by_project_id()
	if err == gorm.ErrRecordNotFound {
		/*
			if the table is not found then store this table on db
		 */
		fi.Id = 0
		err = fi.Create_this_table()

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
	fi.Id = old_finance.Id

	fi.LandId = old_finance.Land.Id
	fi.TechId = old_finance.Tech.Id
	fi.CapitalId = old_finance.Capital.Id
	fi.OtherId = old_finance.Other.Id
	fi.SumId = old_finance.Sum.Id

	if err := GetDB().Save(fi).Error; err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	return utils.Msg{utils.NoErrorFineEverthingOk, 200, "", ""}
}

/*
	get a finance table
 */
func (fi *Finance) Get_table() (utils.Msg) {
	if fi.ProjectId == 0 {
		return utils.Msg{utils.ErrorInvalidParameters, 400, "", "project id is 0"}
	}

	err := fi.Load_values_to_this_object_by_project_id()

	if err == gorm.ErrRecordNotFound {
		err = fi.Create_this_table()
		if err != nil {
			fi = &Finance{}
		}
	} else if err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 500, "", err.Error()}
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map_with_escape(*fi, []string{"project"})

	return utils.Msg{resp, 200, "", ""}
}

