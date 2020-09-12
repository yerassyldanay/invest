package model

import "invest/utils"

func (p *ProjectStatsOnStatuses) Put_status_on_this_object(raw_stats []ProjectStatsRaw) {
	for _, stat := range raw_stats {
		switch stat.Status {
		case utils.ProjectStatusInprogress:
			p.Inprogress += stat.Number

		case utils.ProjectStatusNewone:
			p.Newone += stat.Number

		case utils.ProjectStatusRejected:
			p.Rejected += stat.Number

		case utils.ProjectStatusDone:
			p.Done += stat.Number

		default:

		}
	}
}
