package app

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/utils/constants"
	"github.com/yerassyldanay/invest/utils/randomer"
	"testing"
	"time"
)

func HelperTestGenerateUserWithoutAnyInfoStored() model.User {
	return model.User{
		Password:       randomer.RandomString(20),
		Fio:            randomer.RandomString(30),
		Role:           model.Role{
			Name:            constants.RoleInvestor,
		},
		Email:          model.Email{
			Address: 	randomer.RandomEmail(),
		},
		Phone:          model.Phone{
			Ccode:    "+7",
			Number:   randomer.RandomDigit(10),
		},
		Verified:       true,
		Lang:           constants.ContentLanguageRu,
		Organization:   model.Organization{
			Lang:    constants.ContentLanguageRu,
			Bin:     randomer.RandomDigit(10),
			Name:    gofakeit.Company(),
			Fio:     gofakeit.Name(),
			Regdate: time.Now(),
			Address: gofakeit.Address().Address,
		},
		Blocked:        false,
		Created:        time.Now(),
	}
}

// helperTestCreateUser
// password is not hashed
func helperTestGenerateUser(roleArgs string, t *testing.T) model.User {
	// create email address
	var email = model.Email{
		Address: 	randomer.RandomEmail(),
	}
	require.NoError(t, TestGorm.Create(&email).Error)

	// create phone number
	var phone = model.Phone{
		Ccode:    "+7",
		Number:   randomer.RandomDigit(10),
	}
	require.NoError(t, TestGorm.Create(&phone).Error)

	// create role
	var role = model.Role{}
	require.NoError(t, TestGorm.First(&role, "name = ?", roleArgs).Error)

	// create organization
	organization := model.Organization{
		Lang:    constants.ContentLanguageRu,
		Bin:     randomer.RandomDigit(10),
		Name:    gofakeit.Company(),
		Fio:     gofakeit.Name(),
		Regdate: time.Now(),
		Address: gofakeit.Address().Address,
	}
	require.NoError(t, TestGorm.Create(&organization).Error)

	// generate user
	var user = model.User{
		Password:       randomer.RandomString(20),
		Fio:            randomer.RandomString(30),
		RoleId:         role.Id,
		Role:           role,
		EmailId:        email.Id,
		Email:          email,
		PhoneId: 		phone.Id,
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

func helperTestCreateUser(roleArgs string, t *testing.T) model.User {
	user := helperTestGenerateUser(roleArgs, t)
	require.NoError(t, TestGorm.Create(&user).Error)

	return user
}

func TestServer_CreateUser(t *testing.T) {
	_ = helperTestCreateUser(constants.RoleInvestor, t)
}