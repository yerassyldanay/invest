package service

import (
	"github.com/stretchr/testify/require"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/utils/constants"
	"github.com/yerassyldanay/invest/utils/helper"
	"github.com/yerassyldanay/invest/utils/randomer"
	"testing"
	"time"
)

func TestInvestService_AddBoxToUploadDocument(t *testing.T) {
	// create admin
	admin := helperTestCreateUser(constants.RoleAdmin, t)

	// create data
	project := helperCreateProject(t)
	_ = project

	// document
	var document = model.Document{
		Kaz:          "kaz" + randomer.RandomString(20),
		Rus:          "rus" + randomer.RandomString(20),
		Eng:          "eng" + randomer.RandomString(20),
		Uri:          "",
		Modified:     time.Now(),
		Created:      time.Now(),
		Status:       constants.ProjectStatusReconsider,
		Step:         0,
		IsAdditional: true,
		ProjectId:    project.Project.Id,
		Responsible:  constants.RoleInvestor,
		Count:        0,
	}
	_ = document

	// add box
	is := InvestService{
		BasicInfo: BasicInfo{
			UserId:   admin.Id,
			RoleId:   admin.RoleId,
			RoleName: constants.RoleAdmin,
			Lang:     "rus",
		},
	}
	msg := is.AddBoxToUploadDocument(document)
	helper.HelperPrint(msg)

	// check
	var newDocument = model.Document{}
	require.NoError(t, TestGorm.First(&newDocument, "kaz = ? and rus = ?",
		document.Kaz,
		document.Rus).
		Error)

	helper.HelperPrint(document)
	helper.HelperPrint(newDocument)
}
