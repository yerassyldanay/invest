package model

import (
	"github.com/yerassyldanay/invest/utils/constants"
)

func Prepare_project_statuses(status string) (statuses []string) {
	switch status {
	case "investor":
		statuses = []string{
			constants.ProjectStatusPendingInvestor,
			constants.ProjectStatusReconsider,
		}
	case "spk":
		statuses = []string{
			constants.ProjectStatusPendingManager,
			constants.ProjectStatusPendingExpert,
			constants.ProjectStatusPendingInvCommittee,
			constants.ProjectStatusRegistrationOfLandPlot,
			constants.ProjectStatusPendingBoard,
		}
	case "admin":
		statuses = []string{
			constants.ProjectStatusPendingAdmin,
		}
	case "reject":
		statuses = []string{
			constants.ProjectStatusReject,
		}
	case "agreement":
		statuses = []string{
			constants.ProjectStatusAgreement,
		}
	default:
		statuses = []string{
			constants.ProjectStatusPendingInvestor,
			constants.ProjectStatusPendingManager,
			constants.ProjectStatusPendingAdmin,
			constants.ProjectStatusPendingExpert,
			constants.ProjectStatusPendingInvCommittee,
			constants.ProjectStatusRegistrationOfLandPlot,
			constants.ProjectStatusPendingBoard,
			constants.ProjectStatusAgreement,
			constants.ProjectStatusReject,
			constants.ProjectStatusReconsider,
			constants.ProjectStatusAccept,
		}
	}

	return statuses
}
