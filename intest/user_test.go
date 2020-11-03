package intest

import (
	"fmt"
	"invest/model"
	"invest/service"
	"invest/utils/constants"
	"strings"
	"testing"
)

var user = model.User{
	Password:       "6z24HXMd7nLeZAE",
	Fio:            "Тестовый Сотрудник СПК",
	Role:           model.Role{
		Name: constants.RoleManager,
	},
	Email:          model.Email{
		Address: 		"yerassyl.danay.nu@gmail.com",
	},
	Phone:          model.Phone{
		Ccode: 			"+7",
		Number: 		"7051234567",
	},
}

func TestGetUser(t *testing.T) {
	var user = model.User{
		Id: 1,
	}
	err := user.OnlyGetByIdPreloaded(model.GetDB())
	if err != nil {
		return
	}
}

func TestUpdateUserProfile(t *testing.T) {
	var is = service.InvestService{
		BasicInfo: service.BasicInfo{
			UserId:   1,
		},
	}

	user.Username = strings.ToLower(user.Fio)

	msg := is.Create_user_based_on_role(&user)
	if msg.IsThereAnError() {
		t.Error(msg.ErrMsg)
	}


}

func TestNewlyCreatedUser(t *testing.T) {

	// new user
	newUser := user

	// check whether everything is ok
	if err := newUser.Email.OnlyGetByAddress(model.GetDB()); err != nil {
		t.Error(err)
	}

	// check phone
	if err := newUser.Phone.OnlyGetByCcodeAndNumber(model.GetDB()); err != nil {
		t.Error(err)
	}

	// check user
	if err := newUser.OnlyGetByEmailAddress(model.GetDB()); err != nil {
		t.Error(err)
	}

	fmt.Printf("user: %#v \n", newUser)
}

func TestServiceUpdateUserInfo(t *testing.T) {

	// this is admin
	is := service.InvestService{
		BasicInfo: service.BasicInfo{
			UserId:   1,
		},
	}

	// copy user info
	newUser := user

	// get user
	if err := newUser.OnlyGetByEmailAddress(model.GetDB()); err != nil {
		t.Error(err)
	}

	fmt.Printf("found one: %d %#v\n", newUser.Id, newUser)

	// new user info
	newUser.Fio = "Тестовый Сотрудник Новый"
	newUser.Email.Address = "yerassyl.danay@mail.ru"
	newUser.Phone.Number = "7759876543"

	// logic
	msg := is.Update_user_profile(&newUser)
	if msg.IsThereAnError() {
		t.Error(msg.ErrMsg)
	}

}

func TestServiceUpdatePassword(t *testing.T) {
	// copy user info
	newUser := user

	// logic
	// 6z24HXMd7nLeZAE
	if err := newUser.OnlyGetByEmailAddress(model.GetDB()); err != nil {
		t.Error(err)
	}

	// this is a user
	// this is an admin
	is := service.InvestService{
		BasicInfo: service.BasicInfo{
			UserId: newUser.Id,
		},
	}

	// logic
	newPassword := "newUserPassword6sqw"
	msg := is.Update_user_password(user.Password, newPassword)
	if msg.IsThereAnError() {
		t.Error(msg.ErrMsg)
	}

	// check login
	sis := model.SignIn{
		KeyUsername:   "email",
		Value:         newUser.Email.Address,
		Password:      newPassword,
	}

	// sign in
	msg = sis.Sign_in()
	if msg.IsThereAnError() {
		t.Error(msg.ErrMsg)
	}
	fmt.Println("msg: ", msg)

	// set old password
	//newPassword = user.Password
	
	// change own password
	is = service.InvestService{
		BasicInfo: service.BasicInfo{
			UserId: user.Id,
		},
	}

	// logic
	msg = is.Update_user_password(newPassword, user.Password)
	if msg.IsThereAnError() {
		t.Error(msg.ErrMsg)
	}

	// check sign in
	sis = model.SignIn{
		KeyUsername:   "email",
		Value:         user.Email.Address,
		Password:      newPassword,
	}

	msg = sis.Sign_in()
	if msg.IsThereAnError() {
		t.Error(msg.ErrMsg)
	}
}
