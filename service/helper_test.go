package service

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/utils/constants"
	"github.com/yerassyldanay/invest/utils/helper"
	"github.com/yerassyldanay/invest/utils/randomer"
	"testing"
	"time"
)

func HelperTestGenerateUserWithoutAnyInfoStored(t *testing.T) model.User {
	var organization = model.Organization{
		Lang:    constants.ContentLanguageRu,
		Bin:     randomer.RandomDigit(10),
		Name:    gofakeit.Company(),
		Fio:     gofakeit.Name(),
		Regdate: time.Now(),
		Address: gofakeit.Address().Address,
	}
	require.NoError(t, TestGorm.Create(&organization).Error)

	return model.User{
		Password: randomer.RandomString(20),
		Fio:      randomer.RandomString(30),
		Role: model.Role{
			Name: constants.RoleInvestor,
		},
		Email: model.Email{
			Address: randomer.RandomEmail(),
		},
		Phone: model.Phone{
			Ccode:  "+7",
			Number: randomer.RandomDigit(10),
		},
		Verified:       true,
		Lang:           constants.ContentLanguageRu,
		OrganizationId: organization.Id,
		Organization:   organization,
		Blocked:        false,
		Created:        time.Now(),
	}
}

func helperGenerateOrganization() model.Organization {
	return model.Organization{
		Lang:    constants.ContentLanguageRu,
		Bin:     randomer.RandomDigit(10),
		Name:    gofakeit.Company(),
		Fio:     gofakeit.Name(),
		Regdate: time.Now(),
		Address: gofakeit.Address().Address,
	}
}

// helperTestCreateUser
// password is not hashed
func helperTestGenerateUser(roleArgs string, t *testing.T) model.User {
	// create email address
	var email = model.Email{
		Address: randomer.RandomEmail(),
	}
	require.NoError(t, TestGorm.Create(&email).Error)

	// create phone number
	var phone = model.Phone{
		Ccode:  "+7",
		Number: randomer.RandomDigit(10),
	}
	require.NoError(t, TestGorm.Create(&phone).Error)

	// create role
	var role = model.Role{}
	require.NoError(t, TestGorm.First(&role, "name = ?", roleArgs).Error)

	// create organization
	organization := helperGenerateOrganization()
	require.NoError(t, TestGorm.Create(&organization).Error)

	// generate user
	var user = model.User{
		Password:       randomer.RandomString(20),
		Fio:            randomer.RandomString(30),
		RoleId:         role.Id,
		Role:           role,
		EmailId:        email.Id,
		Email:          email,
		PhoneId:        phone.Id,
		Phone:          phone,
		Verified:       true,
		Lang:           constants.ContentLanguageRu,
		OrganizationId: organization.Id,
		Organization:   organization,
		Blocked:        false,
		Created:        time.Time{},
	}

	// ok
	return user
}

type ElementHelperCreateProject struct {
	Project  model.Project
	Investor model.User
}

func helperGenerateProject(t *testing.T) ElementHelperCreateProject {
	// create investor
	investor := helperTestCreateUser(constants.RoleInvestor, t)

	// create org
	var organization = helperGenerateOrganization()
	require.NoError(t, TestGorm.Create(&organization).Error)

	// generate project
	var project = model.Project{
		Name:              randomer.RandomString(20),
		Description:       randomer.RandomString(40),
		Info:              randomer.RandomString(20),
		InfoSent:          nil,
		EmployeeCount:     1000,
		Email:             gofakeit.Email(),
		PhoneNumber:       gofakeit.Phone(),
		OrganizationId:    organization.Id,
		Organization:      organization,
		Users:             nil,
		Categors:          nil,
		OfferedById:       investor.Id,
		OfferedByPosition: "manager",
		Status:            "status",
		Step:              1,
		LandPlotFrom:      "form-land",
		LandArea:          1000,
		LandAddress:       gofakeit.Address().Address,
		IsManagerAssigned: false,
		CurrentStep:       model.Ganta{},
		AddInfo:           model.AddInfo{},
	}

	return ElementHelperCreateProject{
		Project:  project,
		Investor: investor,
	}
}

func helperCreateProject(t *testing.T) ElementHelperCreateProject {
	// generate
	projectElement := helperGenerateProject(t)

	// create project
	require.NoError(t, TestGorm.Create(&projectElement.Project).Error)

	return projectElement
}

func helperTestCreateUser(roleArgs string, t *testing.T) model.User {
	user := helperTestGenerateUser(roleArgs, t)
	require.NoError(t, TestGorm.Create(&user).Error)

	return user
}

func TestServer_CreateUser(t *testing.T) {
	_ = helperTestCreateUser(constants.RoleInvestor, t)
}

func helperRedisCheckExists(key string, t *testing.T) bool {
	cmdStatus := model.GetRedis().Exists(key)
	require.NotZero(t, cmdStatus)

	n, err := cmdStatus.Result()
	require.NoError(t, err)

	return n == 1
}

type ElementHelperCreateManagerWithProject struct {
	Manager  model.User
	Investor model.User
	Project  model.Project
}

func helperCreateManagerWithProject(t *testing.T) ElementHelperCreateManagerWithProject {
	manager := helperTestCreateUser(constants.RoleManager, t)

	// create project
	element := helperCreateProject(t)

	// assign
	var pu = model.ProjectsUsers{
		ProjectId: element.Project.Id,
		UserId:    manager.Id,
		Created:   time.Now(),
	}
	helper.HelperPrint(pu)
	require.NoError(t, TestGorm.Create(&pu).Error)

	// ok
	return ElementHelperCreateManagerWithProject{
		Manager:  manager,
		Investor: element.Investor,
		Project:  element.Project,
	}
}

func helperTestGenerateProjectWithTables(t *testing.T) model.ProjectWithFinanceTables {
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

	return projectsWithTables
}

// helperDeeplyCompareUsers
// check everything except for ids
func helperDeeplyCompareUsers(expected, actual model.User, t *testing.T) {
	require.Equal(t, expected.Email.Address, actual.Email.Address)
	require.Equal(t, expected.Phone.Ccode, actual.Phone.Ccode)
	require.Equal(t, expected.Phone.Number, actual.Phone.Number)
	require.Equal(t, expected.Fio, actual.Fio)
	require.Equal(t, expected.Organization.Name, actual.Organization.Name)
	require.Equal(t, expected.Role.Name, actual.Role.Name)
}

func helperTestUserGetFullInfo(id uint64, t *testing.T) model.User {
	var checkUser = model.User{}
	require.NoError(t, TestGorm.First(&checkUser, "id = ?", id).Error)
	require.NoError(t, TestGorm.First(&checkUser.Phone, "id = ?", checkUser.PhoneId).Error)
	require.NoError(t, TestGorm.First(&checkUser.Email, "id = ?", checkUser.EmailId).Error)
	require.NoError(t, TestGorm.First(&checkUser.Role, "id = ?", checkUser.RoleId).Error)
	require.NoError(t, TestGorm.First(&checkUser.Organization, "id = ?", checkUser.OrganizationId).Error)

	return checkUser
}
