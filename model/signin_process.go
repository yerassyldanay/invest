package model

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"github.com/yerassyldanay/invest/utils/errormsg"
	"github.com/yerassyldanay/invest/utils/helper"
	"github.com/yerassyldanay/invest/utils/message"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

/*
	return token, response, err

	status:
		200 - ok
		400 - bad request | request parameters are not correct
		404 - info not found on database
		500 - internal server error (code error)
 */
func (sis *SignIn) Sign_in() (message.Msg) {
	var user = User{}
	regexpOnlyLetters, err := regexp.Compile("[a-z]+")
	if err != nil {
		return ReturnInternalServerError(err.Error())
	}

	//fmt.Println(strings.ToLower(sis.KeyUsername))
	if ok := regexpOnlyLetters.Match([]byte(strings.ToLower(sis.KeyUsername))); !ok {
			return ReturnInvalidParameters("regexp does not match with key")
	}

	switch sis.KeyUsername {
	case "email":
		user.Email.Address = sis.Value
	case "phone":
		user.Phone.Ccode = ""
		user.Phone.Number = sis.Value
	default:
		return ReturnInvalidParameters(sis.KeyUsername + " is not supported")
	}

	// get user
	msg := user.Get_full_info_of_this_user(sis.KeyUsername)
	//fmt.Printf("\n user %#v \n", user)
	if msg.ErrMsg != "" {
		return msg
	}

	// check whether email address is confirmed
	// confirmation check
	if ok := user.Email.Verified; !ok {
		return ReturnEmailIsNotVerified("email address has not been verified yet")
	}

	/*
		here we check whether two passwords (a provided password and password on db_create_fake_data) MATCH
	*/
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(sis.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return ReturnWrongPassword("password is wrong")
	} else if err != nil {
		return ReturnInvalidPassword("password either does not match or invalid")
	}
	
	/*
		Provide a role of a client
	*/
	var token = &Token {
		UserId:         user.Id,
		RoleName:		user.Role.Name,
		Exp:            time.Now().Add(time.Hour * 24),
	}

	// create a session token
	var token_hashed = jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), token)
	token_string, err := token_hashed.SignedString([]byte(os.Getenv("TOKEN_PASSWORD")))

	if err != nil {
		return message.Msg{errormsg.ErrorInternalIssueOrInvalidPassword, http.StatusInternalServerError, "", err.Error()}
	}

	sis.TokenCompound = "Bearer " + token_string
	user.Password = ""

	resp := errormsg.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(user)
	resp["token"] = token_string
	resp["role"] = user.Role.Name

	return ReturnNoError()
}


func (si *SignIn) Is_on_db() bool {
	if ok := helper.Is_it_free_from_sql_injection(si.KeyUsername); !ok {
		return false
	}

	var count int
	if err := GetDB().Table(User{}.TableName()).Where(si.KeyUsername + "=?", si.Value).Count(&count).Error;
		err != nil {
		return false
	}

	return count == 0
}

