
package tests

import (
	"github.com/stretchr/testify/require"
	"invest/model"
	"invest/service"
	"invest/utils/constants"
	"invest/utils/helper"
	"testing"
	"time"
)

/*

 */
const bin = "190940011748"
func HelperGetNewProject() model.ProjectWithFinanceTables {
	project := model.ProjectWithFinanceTables {
		Project: model.Project{
			Name:              "Тестовый проект - " + helper.Generate_Random_String(20),
			Description:       helper.Generate_Random_String(30),
			InfoSent:          map[string]interface{}{
				"add-info": helper.Generate_Random_String(20),
			},
			EmployeeCount:     100,
			Email:             "any@gmail.com",
			PhoneNumber:       "+77781254856",
			Organization:      model.Organization{
				Bin:    bin,
			},
			Categors: []model.Categor{
				{
					Id: 1,
				},
			},
			OfferedByPosition: "инициатор проекта",
			LandPlotFrom:      "что то нужно здесь написать - " + helper.Generate_Random_String(20),
			LandArea:          10000,
			LandAddress:       "город, название улицы, дом",
		},
		Cost:    model.Cost{
			BuildingRepairInvestor:      2000,
			BuildingRepairInvolved:      3000,
			TechnologyEquipmentInvestor: 5000,
			TechnologyEquipmentInvolved: 6000,
			WorkingCapitalInvestor:      7000,
			WorkingCapitalInvolved:      8000,
			OtherCostInvestor:           9000,
			OtherCostInvolved:           10000,
		},
		Finance: model.Finance{
			TotalIncome:           1000,
			TotalProduction:       2000,
			ProductionCost:        3000,
			OperationalProfit:     4000,
			SettlementObligations: 5000,
			OtherCost:             6000,
			PureProfit:            7000,
			Taxes:                 8000,
		},
	}

	return project
}

func TestProjectCreate(t *testing.T) {
	// get organization
	org := model.Organization{Bin: bin}
	msg := org.Create_or_get_organization_from_db_by_bin(model.GetDB())

	// check
	require.Zero(t, msg.ErrMsg)
	require.NotZero(t, org.Id)
	require.NotZero(t, org.Name)
	require.NotZero(t, org.Fio)
	require.NotZero(t, org.Address)

	// get new project
	project := HelperGetNewProject()

	// headers
	is := service.InvestService{
		TimeUTC:   time.Now().UTC(),
		Time:      time.Now(),
		Offset:    "0",
		BasicInfo: service.BasicInfo{
			UserId:   3,
			RoleName: constants.RoleInvestor,
			Lang:     constants.DefaultContentLanguage,
		},
	}

	// logic
	msg = is.Service_create_project(&project)
	if msg.IsThereAnError() {
		t.Error("expected no error, but got " + msg.ErrMsg)
	}
}

// check documents
func TestProjectDocuments(t *testing.T) {
	// get any project from database
	project := HelperGetAnyProject(t)

	document := model.Document{}
	steps := []interface{}{1}
	documents, err := document.OnlyGetDocumentsByStepsAndProjectId(project.Id, steps, model.GetDB())

	require.NoError(t, err)
	require.Condition(t, func() (success bool) { return len(documents) > 1 }, nil)
}

func TestProjectFinance(t *testing.T) {
	project := HelperGetAnyProject(t)

	finance := model.Finance{ProjectId: project.Id}
	err := finance.OnlyGetByProjectId(model.GetDB())

	require.NoError(t, err)
	require.NotZero(t, finance.Id)
}

func TestProjectCost(t *testing.T) {
	project := HelperGetAnyProject(t)

	cost := model.Cost{ProjectId: project.Id}
	err := cost.OnlyGetByProjectId(model.GetDB())

	// check
	require.NoError(t, err)
	require.NotZero(t, cost.Id)
}

func TestProjectGanta(t *testing.T) {
	project := HelperGetAnyProject(t)

	// check first step
	ganta := model.Ganta{ProjectId: project.Id, Step: project.Step}
	gantas, err := ganta.OnlyGetParentsByProjectId(1, model.GetDB())

	// check
	require.NoError(t, err)
	require.True(t, len(gantas) == len(model.DefaultGantaParentsOfStep1))

	// check for second step
	gantas, err = ganta.OnlyGetParentsByProjectId(2, model.GetDB())

	// check
	require.NoError(t, err)
	require.True(t, len(gantas) == len(model.DefaultGantaParentsOfStep2) - 1)
	require.True(t, len(gantas) > 0)

	length := len(model.DefaultGantaParentsOfStep2)
	require.Equal(t, gantas[len(gantas) - 1].Rus, model.DefaultGantaParentsOfStep2[length - 2].Rus)
}



