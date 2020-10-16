package intest

import (
	"fmt"
	"invest/model"
	"invest/service"
	"invest/utils"
	"testing"
	"time"
)

/*

 */
func TestProjectCreate(t *testing.T) {

	project := model.ProjectWithFinanceTables{
		Project: model.Project{
			Name:              "Тестовый проект - спк",
			Description:       "Описание проекта пишете сюда",
			InfoSent:          map[string]interface{}{
				"add-info": "доп инфо",
			},
			EmployeeCount:     100,
			Email:             "any@gmail.com",
			PhoneNumber:       "+77781254856",
			Organization:      model.Organization{
				Bin:     "190940011748",
			},
			OfferedByPosition: "инициатор проекта",
			LandPlotFrom:      "что то нужно здесь написать",
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
			ShareInProjectInvestor:      11000,
			ShareInProjectInvolved:      5000,
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

	// headers
	is := service.InvestService{
		TimeUTC:   time.Now().UTC(),
		Time:      time.Now(),
		Offset:    "0",
		BasicInfo: service.BasicInfo{
			UserId:   3,
			RoleName: utils.RoleInvestor,
			Lang:     utils.DefaultContentLanguage,
		},
	}

	// logic
	msg := is.Service_create_project(&project)
	if msg.IsThereAnError() {
		t.Error("expected no error, but got " + msg.ErrMsg)
	}

	// create thee same project twice
	msg = is.Service_create_project(&project)
	if !msg.IsThereAnError() {
		t.Error("expected kind of 'duplicate error' error, but got ''")
	}

}

func getAnyProject(t *testing.T) model.Project {
	project := model.Project{}
	if err := project.OnlyGetAny(model.GetDB()); err != nil {
		t.Error("expected to get at least one project, but got ", err)
	}

	return project
}

func TestProjectGet(t *testing.T) {
	project := getAnyProject(t)
	fmt.Println(project.Id, project.Name)
}

func TestProjectDocuments(t *testing.T) {
	project := getAnyProject(t)

	document := model.Document{}
	documents, err := document.OnlyGetDocumentsByStepsAndProjectId(project.Id, []interface{}{interface{}(1)}, model.GetDB())
	switch {
	case err != nil:
		t.Error("expected no error, but got ", err)
	case len(documents) < 1:
		t.Error("expected at least one document, but got 0")
	default:
		fmt.Println("len of documents: ", len(documents))
	}
}

func TestProjectFinance(t *testing.T) {
	project := getAnyProject(t)

	finance := model.Finance{ProjectId: project.Id}
	err := finance.OnlyGetByProjectId(model.GetDB())

	switch {
	case err != nil :
		t.Error("expected no err, but got ", err)
	case finance.Id < 1:
		t.Error("expected id to be > 1, but got ", finance.Id)
	default:
		fmt.Println("finance: ", finance.Id, " | project: ", project.Id)
	}
}

func TestProjectCost(t *testing.T) {
	project := getAnyProject(t)

	cost := model.Cost{ProjectId: project.Id}
	err := cost.OnlyGetByProjectId(model.GetDB())

	switch {
	case err != nil :
		t.Error("expected no err, but got ", err)
	case cost.Id < 1:
		t.Error("expected id to be > 1, but got ", cost.Id)
	default:
		fmt.Println("cost: ", cost.Id, " | project: ", project.Id)
	}
}

func TestProjectGanta(t *testing.T) {
	project := getAnyProject(t)

	// check first step
	ganta := model.Ganta{ProjectId: project.Id, Step: project.Step}
	gantas, err := ganta.OnlyGetParentsByProjectId(1, model.GetDB())

	switch {
	case err != nil:
		t.Error("expected no error, but got ", err)
	case len(gantas) != len(model.DefaultGantaParentsOfStep1):
		t.Error("expected ", len(gantas), ", but got ", len(model.DefaultGantaParentsOfStep1))
	default:
		fmt.Println("gantt 1 step: ", len(gantas))
	}

	// check for second step
	gantas, err = ganta.OnlyGetParentsByProjectId(2, model.GetDB())

	switch {
	case err != nil:
		t.Error("expected no err, but got ", err)
	case len(gantas) != len(model.DefaultGantaParentsOfStep2) - 1:
		t.Error("expected ", len(gantas), ", but got ", len(model.DefaultGantaParentsOfStep2) - 1)
	default:
		fmt.Println("gantt 2 step: ", len(gantas))
	}

	if len(gantas) < 1 {
	} else if gantas[len(gantas) - 1].Rus == model.DefaultGantaFinalStep.Rus {
		t.Error("expected something else, but got ", model.DefaultGantaFinalStep.Rus)
	}
}



