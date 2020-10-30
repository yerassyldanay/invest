package model

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"invest/utils"
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
func (sis *SignIn) Sign_in() (utils.Msg) {
	var user = User{}
	regexpOnlyLetters, err := regexp.Compile("[a-z]+")
	if err != nil {
		return utils.Msg{
			Message: utils.ErrorInternalServerError, Status:  http.StatusInternalServerError, ErrMsg:  err.Error(),
		}
	}

	//fmt.Println(strings.ToLower(sis.KeyUsername))
	if ok := regexpOnlyLetters.Match([]byte(strings.ToLower(sis.KeyUsername))); !ok {
			return utils.Msg{utils.ErrorInvalidParameters, http.StatusBadRequest, "", "regexp does not match with key"}
	}

	switch sis.KeyUsername {
	case "email":
		user.Email.Address = sis.Value
	case "phone":
		user.Phone.Ccode = ""
		user.Phone.Number = sis.Value
	case "username":
		user.Username = sis.Value
	}

	msg := user.Get_full_info_of_this_user(sis.KeyUsername)
	//fmt.Printf("\nuser %#v \n", user)
	if msg.ErrMsg != "" {
		return msg
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
	var token = &Token{
		UserId:         user.Id,
		RoleName:		user.Role.Name,
		Exp:            time.Now().Add(time.Hour * 24),
	}

	// create a session token
	var token_hashed = jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), token)
	token_string, err := token_hashed.SignedString([]byte(os.Getenv("TOKEN_PASSWORD")))

	if err != nil {
		return utils.Msg{utils.ErrorInternalIssueOrInvalidPassword, http.StatusInternalServerError, "", err.Error()}
	}

	sis.TokenCompound = "Bearer " + token_string
	user.Password = ""

	resp := utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(user)
	resp["token"] = token_string
	resp["role"] = user.Role.Name

	return utils.Msg{
		resp, http.StatusOK,"", "",
	}
}


func (si *SignIn) Is_on_db() bool {
	if ok := utils.Is_it_free_from_sql_injection(si.KeyUsername); !ok {
		return false
	}

	var count int
	if err := GetDB().Table(User{}.TableName()).Where(si.KeyUsername + "=?", si.Value).Count(&count).Error;
		err != nil {
		return false
	}

	return count == 0
}

