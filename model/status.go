package model

import "invest/utils"

func Prepare_project_statuses(status string) (statuses []string) {
	switch status {
	case "investor":
		statuses = []string{
			utils.ProjectStatusPendingInvestor,
			utils.ProjectStatusReconsider,
		}
	case "spk":
		statuses = []string{
			utils.ProjectStatusPendingManager,
			utils.ProjectStatusPendingExpert,
			utils.ProjectStatusPendingInvCommittee,
			utils.ProjectStatusRegistrationOfLandPlot,
			utils.ProjectStatusPendingBoard,
		}
	case "admin":
		statuses = []string{
			utils.ProjectStatusPendingAdmin,
			utils.ProjectStatusPreliminaryReject,
			utils.ProjectStatusPreliminaryReconsider,
			utils.ProjectStatusPreliminaryAccept,
		}
	case "reject":
		statuses = []string{
			utils.ProjectStatusReject,
		}
	case "agreement":
		statuses = []string{
			utils.ProjectStatusAgreement,
		}
	default:
		statuses = []string{
			utils.ProjectStatusPendingInvestor,
			utils.ProjectStatusPendingManager,
			utils.ProjectStatusPendingAdmin,
			utils.ProjectStatusPendingExpert,
			utils.ProjectStatusPendingInvCommittee,
			utils.ProjectStatusRegistrationOfLandPlot,
			utils.ProjectStatusPendingBoard,
			utils.ProjectStatusAgreement,
			utils.ProjectStatusReject,
			utils.ProjectStatusReconsider,
			utils.ProjectStatusAccept,
			utils.ProjectStatusPreliminaryReject,
			utils.ProjectStatusPreliminaryReconsider,
			utils.ProjectStatusPreliminaryAccept,
		}
	}

	return statuses
}
