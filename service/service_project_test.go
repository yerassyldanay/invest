package service

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/utils/constants"
	"testing"
)

func helperTestCreateProject(t *testing.T) model.ProjectWithFinanceTables {
	projectElement := helperGenerateProject(t)

	// create project parameters
	var projectsWithTables = model.ProjectWithFinanceTables{
		Project: projectElement.Project,
		Cost: model.Cost{
			Id:                          0,
			ProjectId:                   0,
			BuildingRepairInvestor:      gofakeit.Number(100, 100000),
			BuildingRepairInvolved:      gofakeit.Number(100, 100000),
			TechnologyEquipmentInvestor: gofakeit.Number(100, 100000),
			TechnologyEquipmentInvolved: gofakeit.Number(100, 100000),
			WorkingCapitalInvestor:      gofakeit.Number(100, 100000),
			WorkingCapitalInvolved:      gofakeit.Number(100, 100000),
			OtherCostInvestor:           gofakeit.Number(100, 100000),
			OtherCostInvolved:           gofakeit.Number(100, 100000),
		},
		Finance: model.Finance{
			Id:                    0,
			ProjectId:             0,
			TotalIncome:           gofakeit.Number(100, 100000),
			TotalProduction:       gofakeit.Number(100, 100000),
			ProductionCost:        gofakeit.Number(100, 100000),
			OperationalProfit:     gofakeit.Number(100, 100000),
			SettlementObligations: gofakeit.Number(100, 100000),
			OtherCost:             gofakeit.Number(100, 100000),
			PureProfit:            gofakeit.Number(100, 100000),
			Taxes:                 gofakeit.Number(100, 100000),
		},
	}

	// investor
	investor := helperTestCreateUser(constants.RoleInvestor, t)

	// role
	var role = model.Role{}
	require.NoError(t, TestGorm.First(&role, "name = ?", constants.RoleInvestor).Error)

	// is
	is := InvestService{
		BasicInfo: BasicInfo{
			UserId:   investor.Id,
			RoleId:   role.Id,
			RoleName: constants.RoleInvestor,
			Lang:     "rus",
		},
	}
	msg := is.ServiceCreateProject(&projectsWithTables)
	_ = msg
	//helper.HelperPrint(msg)

	// check
	var project = model.Project{}
	require.NoError(t, TestGorm.First(&project, "offered_by_id = ?", investor.Id).Error)

	// check other tables
	var cost = model.Cost{}
	require.NoError(t, TestGorm.First(&cost, "project_id = ?", project.Id).Error)
	//helper.HelperPrint(cost)

	// check other tables
	var finance = model.Finance{}
	require.NoError(t, TestGorm.First(&finance, "project_id = ?", project.Id).Error)
	//helper.HelperPrint(finance)

	// docs
	var docs = []model.Document{}
	require.NoError(t, TestGorm.Find(&docs, "project_id = ?", project.Id).Error)
	//helper.HelperPrint(docs)

	// gantt steps
	var ganttSteps = []model.Ganta{}
	require.NoError(t, TestGorm.Order("start_date").Find(&ganttSteps, "project_id = ?", project.Id).Error)
	//helper.HelperPrint(ganttSteps)

	return projectsWithTables
}

func TestInvestService_ServiceCreateProject(t *testing.T) {
	helperCreateProject(t)
}
