package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/yerassyldanay/invest/utils/errormsg"
	"github.com/yerassyldanay/invest/utils/message"
	"net/http"
	"os"
	"time"
)

/*
	return token, response, err
 */
func (sis *SignIn) SignIn() (message.Msg) {
	var err error
	var user = User{}

	// get user
	msg := user.GetFullInfoOfThisUser(ElementGetFullInfoOfThisUser{
		Key:   sis.KeyUsername,
		Value: sis.Value,
	})
	if msg.ErrMsg != "" {
		return msg
	}

	// provide a role of a client
	var token = &Token {
		UserId:         user.Id,
		RoleName:		user.Role.Name,
		Exp:            time.Now().Add(time.Hour * 24),
	}

	// create a session token
	var tokenHashed = jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), token)
	tokenString, err := tokenHashed.SignedString([]byte(os.Getenv("TOKEN_PASSWORD")))

	if err != nil {
		return message.Msg{errormsg.ErrorInternalIssueOrInvalidPassword, http.StatusInternalServerError, "", err.Error()}
	}

	if user.Password != sis.Password {
		return ReturnWrongPassword("password is not correct")
	}

	sis.TokenCompound = "Bearer " + tokenString
	user.Password = ""

	resp := errormsg.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(user)
	resp["token"] = tokenString
	resp["role"] = user.Role.Name

	return ReturnNoErrorWithResponseMessage(resp)
}

//func (si *SignIn) IsOnDb() bool {
//	if ok := helper.Is_it_free_from_sql_injection(si.KeyUsername); !ok {
//		return false
//	}
//
//	var count int
//	if err := GetDB().Table(User{}.TableName()).Where(si.KeyUsername + "=?", si.Value).Count(&count).Error;
//		err != nil {
//		return false
//	}
//
//	return count == 0
//}

