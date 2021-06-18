package service

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/utils/constants"
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
		OrganizationId: organization.Id,
		Organization:   organization,
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

func helperRedisCheckExists(key string, t *testing.T) bool {
	cmdStatus := model.GetRedis().Exists(key)
	require.NotZero(t, cmdStatus)

	n, err := cmdStatus.Result()
	require.NoError(t, err)

	return n == 1
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