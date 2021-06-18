package service

import (
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/utils/constants"
	"testing"
)

func TestModule_UserGetFullInfo(t *testing.T) {
	// user
	user := helperTestCreateUser(constants.RoleInvestor, t)
	//HelperPrint(user)

	// get info
	var newUser = model.User{}
	newUser.GetFullInfoOfThisUser(model.ElementGetFullInfoOfThisUser{
		Key:   "email",
		Value: user.Email.Address,
	})
	helperDeeplyCompareUsers(user, newUser, t)

	// get info
	newUser = model.User{}
	newUser.GetFullInfoOfThisUser(model.ElementGetFullInfoOfThisUser{
		Key:   "phone",
		Value: user.Phone.Ccode + user.Phone.Number,
	})
	helperDeeplyCompareUsers(user, newUser, t)

	//HelperPrint(newUser)
}

func Test_SignIn(t *testing.T) {
	// user
	user := helperTestCreateUser(constants.RoleInvestor, t)

	// sign in
	sis := model.SignIn{
		KeyUsername:   "email",
		Value:         user.Email.Address,
		Password:      user.Password,
		Id:            0,
		TokenCompound: "",
	}
	msg := sis.SignIn()

	// get user
	userInterface := msg.Message["user"]
	HelperPrint(userInterface)

	HelperPrint(msg)
}