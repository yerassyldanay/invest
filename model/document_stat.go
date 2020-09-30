package model

import (
	"invest/utils"

)

func (d *Document) Get_stat_on_docs_by_project_id() (utils.Msg) {
	var stats = []struct {
		Status					string				`json:"name"`
		Num						int					`json:"num"`
	}{}

	var main_query = `select status, count(*) as num from documents where project_id = ? group by status;`
	if err := GetDB().Raw(main_query, d.ProjectId).Scan(&stats).Error; err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	var docStat = DocumentStat{}
	//for _, stat := range stats {
	//	switch stat.Status {
	//	case utils.ProjectStatusDone:
	//		docStat.Done += stat.Num
	//
	//	case utils.ProjectStatusPendingAdmin:
	//		docStat.Inprogress += stat.Num
	//
	//	case utils.ProjectStatusRejected:
	//		docStat.Rejected += stat.Num
	//
	//	default:
	//		docStat.Sum -= stat.Num
	//	}
	//
	//	docStat.Sum += stat.Num
	//}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(docStat)

	return utils.Msg{resp, 200, "", ""}
}

