package model

import "invest/utils"

func Prepare_project_statuses(status string) (statuses []string) {
	switch status {
	case "new_one":
		statuses = []string{
			utils.ProjectStatusPendingInvestor,
			utils.ProjectStatusReconsider,
		}
	case "pending":
		statuses = []string{
			utils.ProjectStatusPendingManager,
			utils.ProjectStatusPendingExpert,
			utils.ProjectStatusPendingInvCommittee,
			utils.ProjectStatusRegistrationOfLandPlot,
			utils.ProjectStatusPendingBoard,
		}
	case "waiting":
		statuses = []string{
			utils.ProjectStatusPendingAdmin,
		}
	case "rejected":
		statuses = []string{
			utils.ProjectStatusReject,
		}
	case "completed":
		statuses = []string{
			utils.ProjectStatusAgreement,
		}
	default:
		statuses = []string{
			utils.ProjectStatusPendingInvestor,
			utils.ProjectStatusReconsider,
			utils.ProjectStatusPendingManager,
			utils.ProjectStatusPendingExpert,
			utils.ProjectStatusPendingInvCommittee,
			utils.ProjectStatusRegistrationOfLandPlot,
			utils.ProjectStatusPendingBoard,
		}
	}

	return statuses
}
