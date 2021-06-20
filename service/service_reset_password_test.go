package service

import (
	"github.com/stretchr/testify/require"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/utils/constants"
	"github.com/yerassyldanay/invest/utils/randomer"
	"testing"
	"time"
)

type caseTestInvestService_PasswordResetSendMessage struct {
	EmailAddress string
	Ok           bool
	Check        func()
}

func TestInvestService_PasswordResetSendMessage(t *testing.T) {
	// user
	user := helperTestCreateUser(constants.RoleInvestor, t)

	// test cases
	testCases := []caseTestInvestService_PasswordResetSendMessage{
		{
			EmailAddress: "notFound" + user.Email.Address,
			Ok:           false,
			Check: func() {

			},
		},
		{
			EmailAddress: user.Email.Address,
			Ok:           true,
			Check: func() {

			},
		},
	}

	// go through
	for _, testCase := range testCases {
		is := InvestService{}
		msg := is.PasswordResetSendMessage(model.ForgetPassword{
			EmailAddress: testCase.EmailAddress,
			Lang:         "rus",
		})
		if testCase.Ok {
			require.Zero(t, msg.ErrMsg)
			require.Less(t, msg.Status, 300)
		} else {
			require.NotZero(t, msg.ErrMsg)
			require.GreaterOrEqual(t, msg.Status, 300)
		}

		//helper.HelperPrint(msg)
	}
}

type caseTestInvestService_PasswordResetChangePassword struct {
	Pass  model.ForgetPassword
	Ok    bool
	Check func()
}

func TestInvestService_PasswordResetChangePassword(t *testing.T) {
	// user
	user := helperTestCreateUser(constants.RoleInvestor, t)

	// get code
	is := InvestService{}
	msg := is.PasswordResetSendMessage(model.ForgetPassword{
		EmailAddress: user.Email.Address,
		Lang:         "rus",
	})

	// get code
	respMap := msg.Message
	//helper.HelperPrint(respMap)
	codeInterface, ok := respMap["code"]
	require.True(t, ok)
	code := codeInterface.(string)

	// test cases
	testCases := []caseTestInvestService_PasswordResetChangePassword{
		{
			Pass: model.ForgetPassword{
				NewPassword:  randomer.RandomString(20),
				EmailAddress: user.Email.Address,
				Code:         "12" + code,
				Deadline:     time.Time{},
				Lang:         "",
			},
			Ok:    false,
			Check: func() {},
		},
		{
			Pass: model.ForgetPassword{
				NewPassword:  randomer.RandomString(20),
				EmailAddress: user.Email.Address,
				Code:         code,
			},
			Ok: true,
			Check: func() {
				var newUser = model.User{}
				require.NoError(t, TestGorm.First(&newUser, "email_id = ?", user.EmailId).
					Error)
				require.Equal(t, user.Id, newUser.Id)
				require.NotEqual(t, user.Password, newUser.Password)
			},
		},
	}

	// go through
	for _, testCase := range testCases {
		is := InvestService{}
		msg = is.PasswordResetChangePassword(testCase.Pass)
		if testCase.Ok {
			require.Zero(t, msg.ErrMsg)
		} else {
			require.NotZero(t, msg.ErrMsg)
		}

		//helper.HelperPrint(msg)
	}
}
