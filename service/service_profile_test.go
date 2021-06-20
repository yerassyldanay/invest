package service

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/utils/constants"
	"github.com/yerassyldanay/invest/utils/helper"
	"github.com/yerassyldanay/invest/utils/randomer"
	"testing"
)

type caseTestInvestService_UserGetProfile struct {
	User model.User
	Ok   bool
}

func TestInvestService_UserGetProfile(t *testing.T) {
	userElement := helperTestCreateUser(constants.RoleAdmin, t)
	userElement.GetFullInfoOfThisUser(model.ElementGetFullInfoOfThisUser{})

	//
	testCases := []caseTestInvestService_UserGetProfile{
		{
			User: model.User{},
			Ok:   false,
		},
		{
			User: userElement,
			Ok:   true,
		},
	}

	// go through
	for _, testCase := range testCases {
		var user = model.User{
			Id: testCase.User.Id,
		}
		msg := user.GetFullInfoOfThisUser(model.ElementGetFullInfoOfThisUser{})
		if testCase.Ok {
			require.Zero(t, msg.ErrMsg)
			require.NotZero(t, user.Id)
			require.NotZero(t, user.Fio)
			require.NotZero(t, user.Email.Id)
			require.NotZero(t, user.Phone.Id)
			require.NotZero(t, user.Organization.Id)
			require.NotZero(t, user.Role.Id)

			helper.HelperPrint(user)
		} else {
			require.NotZero(t, msg.ErrMsg)
		}
	}
}

func Test_GetFullInfoOfThisUserWithoutPasswordById(t *testing.T) {
	user := helperTestCreateUser(constants.RoleInvestor, t)

	// get
	var newUser = model.User{
		Id: user.Id,
	}
	newUser.GetFullInfoOfThisUserWithoutPasswordById()

	// equal
	helperDeeplyCompareUsers(user, newUser, t)
}

type caseTestInvestService_UpdateUserProfile struct {
	NewUser model.User
	Ok      bool
	Check   func()
}

func TestInvestService_UpdateUserProfile(t *testing.T) {
	// data
	user := helperTestCreateUser(constants.RoleInvestor, t)

	// cases
	testCases := []caseTestInvestService_UpdateUserProfile{
		//{
		//	NewUser: model.User{
		//		Id:    user.Id,
		//		Fio:   "",
		//		Phone: model.Phone{},
		//	},
		//	Ok:      true,
		//	Check: func() {
		//		var checkUser = helperTestUserGetFullInfo(user.Id, t)
		//		helperDeeplyCompareUsers(user, checkUser, t)
		//	},
		//},
		//{
		//	NewUser: model.User{
		//		Id:    user.Id,
		//		Fio:   gofakeit.FirstName(),
		//		Phone: model.Phone{},
		//	},
		//	Ok:      true,
		//	Check: func() {
		//		var checkUser = helperTestUserGetFullInfo(user.Id, t)
		//		require.NotEqual(t, user.Fio, checkUser.Fio)
		//		require.Equal(t, user.Phone.Number, checkUser.Phone.Number)
		//	},
		//},
		{
			NewUser: model.User{
				Id:  user.Id,
				Fio: gofakeit.FirstName(),
				Phone: model.Phone{
					Ccode:  "+7",
					Number: randomer.RandomDigit(10),
				},
			},
			Ok: true,
			Check: func() {
				var checkUser = helperTestUserGetFullInfo(user.Id, t)
				require.NotEqual(t, user.Fio, checkUser.Fio)
				require.NotEqual(t, user.Phone.Number, checkUser.Phone.Number)

				//helper.HelperPrint(user)
				//helper.HelperPrint(checkUser)
			},
		},
	}

	// go through test cases
	for _, testCase := range testCases {
		var is = InvestService{}
		msg := is.UpdateUserProfile(&testCase.NewUser)

		if testCase.Ok {
			require.Zero(t, msg.ErrMsg)
		} else {
			require.NotZero(t, msg.ErrMsg)
		}

		testCase.Check()

		//helper.HelperPrint(user)
		//helper.HelperPrint(helperTestUserGetFullInfo(user.Id, t))
	}
}
