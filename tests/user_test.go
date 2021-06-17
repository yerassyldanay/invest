package tests

import (
	"github.com/stretchr/testify/require"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/service"
	"github.com/yerassyldanay/invest/utils/constants"
	"github.com/yerassyldanay/invest/utils/helper"
	"testing"
)

/*
	* create user by admin
	* test profile update
		own
		by admin
	* update password
		own
		by admin
	* get users by role
 */
func createUserProfile(user model.User, t *testing.T) {

	is := service.InvestService{
		BasicInfo: service.BasicInfo{
			UserId:   1,
			RoleName: constants.RoleInvestor,
		},
	}

	savedUser := user

	// create a new user by admin
	msg := is.Create_user_based_on_role(&user)

	// check
	require.Zero(t, msg.ErrMsg)

	// get user
	gotUser := model.User{
		Email: model.Email{
			Address: savedUser.Email.Address,
		},
	}
	err := gotUser.OnlyGetByEmailAddress(model.GetDB())
	require.NoError(t, err)

	err = gotUser.OnlyGetByIdPreloaded(model.GetDB())

	// check
	require.NoError(t, err)
	HelperValidateUser(gotUser, t)
}

func TestCreateUserProfile(t *testing.T) {
	// new user
	user := HelperGenerateNewUser()
	user.Role.Name = constants.RoleManager

	// this function does all other work
	createUserProfile(user, t)
}

func TestServiceUpdateUserInfo(t *testing.T) {
	// new user
	user := HelperGenerateNewUser()
	user.Role.Name = constants.RoleManager

	// create new user
	createUserProfile(user, t)

	// get user
	_ = user.OnlyGetByEmailAddress(model.GetDB())
	err := user.OnlyGetByIdPreloaded(model.GetDB())

	// check
	require.NoError(t, err)

	// update
	user.Fio = helper.Generate_Random_String(30)
	user.Phone.Number = helper.Generate_Random_Number(10)

	is := service.InvestService {
		BasicInfo: service.BasicInfo {
			UserId: 1,
			RoleName: constants.RoleInvestor,
		},
	}
	tempUser := user
	msg := is.Update_user_profile(&tempUser)

	// check
	require.Zero(t, msg.ErrMsg)

	// compare
	err = tempUser.OnlyGetByIdPreloaded(model.GetDB())
	require.NoError(t, err)
	require.Equal(t, tempUser.Fio, user.Fio)
	require.Equal(t, tempUser.Phone.Number, user.Phone.Number)
	require.Equal(t, tempUser.Phone.Ccode, user.Phone.Ccode)
}

func TestServiceUpdateUserPassword(t *testing.T) {

	// copy user info
	user := HelperGenerateNewUser()
	user.Role.Name = constants.RoleManager
	createUserProfile(user, t)

	oldPassword := user.Password
	newPassword := helper.Generate_Random_String(20)

	// need user id
	_ = user.OnlyGetByEmailAddress(model.GetDB())
	err := user.OnlyGetByIdPreloaded(model.GetDB())

	// check
	require.NoError(t, err)

	// this is admin
	is := service.InvestService{
		BasicInfo: service.BasicInfo{
			UserId:   user.Id,
			RoleName: user.Role.Name,
		},
	}
	msg := is.Update_user_password(oldPassword, newPassword)

	// check
	require.Zero(t, msg.ErrMsg)

	// sign in - check
	sis := model.SignIn{
		KeyUsername:   "email",
		Value:         user.Email.Address,
		Password:      newPassword,
	}
	msg = sis.Sign_in()

	// check
	require.Zero(t, msg.ErrMsg)
}

func TestModelGetUsersByRole(t *testing.T) {
	// create manager
	user := HelperGenerateNewUser()
	user.Role.Name = constants.RoleManager
	createUserProfile(user, t)

	// get managers - it must be more than 1
	users, err := user.OnlyGetUsersByRolePreloaded([]string{constants.RoleManager}, "0", model.GetDB())

	// check
	require.NoError(t, err)
	require.Condition(t, func() (bool) { return len(users) >= 1 })
}

func TestServiceGetUsersByRole(t *testing.T) {
	// create manager
	user := HelperGenerateNewUser()
	user.Role.Name = constants.RoleManager
	createUserProfile(user, t)

	// admin
	is := service.InvestService{
		BasicInfo: service.BasicInfo{
			UserId: 1,
			RoleName: constants.RoleInvestor,
		},
	}

	// get managers
	msg := is.Get_users_by_roles([]string{constants.RoleManager})

	// check
	require.Zero(t, msg.ErrMsg)
}
