package model

import "invest/utils"

func (p Project) Get_projects_grouped_by_statuses() (*utils.Msg) {
	type TempStat struct {
		Number				int				`json:"number"`
		Status				string			`json:"status"`
	}

	type TempInto struct {
		Info		[]TempStat
	}

	var stats = TempInto{}
	var main_query = `select count(*) as number, status from projects group by status;`

	err := GetDB().Raw(main_query).Scan(&stats.Info).Error
	if err != nil {
		return &utils.Msg{utils.ErrorInternalDbError, 417, "", ""}
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(stats)["info"]

	return &utils.Msg{resp, 200, "", ""}
}
