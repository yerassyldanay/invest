package model

import "invest/utils"

func (p Project) Get_projects_grouped_by_statuses() (*utils.Msg) {
	var stats = []ProjectStatsRaw{}
	var main_query = `select count(*) as number, status from projects group by status;`

	err := GetDB().Raw(main_query).Scan(&stats).Error
	if err != nil {
		return &utils.Msg{utils.ErrorInternalDbError, 417, "", ""}
	}

	var prettifiedStats = ProjectStatsOnStatuses{}
	prettifiedStats.Put_status_on_this_object(stats)

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(prettifiedStats)

	return &utils.Msg{resp, 200, "", ""}
}
