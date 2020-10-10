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
	//fmt.Println(user)
	if msg.ErrMsg != "" {
		return msg
	}

	/*
		here we check whether two passwords (a provided password and password on db_create_fake_data) MATCH
	*/
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(sis.Password)); err == bcrypt.ErrMismatchedHashAndPassword || err != nil {
		return utils.Msg{utils.ErrorInvalidPassword, http.StatusBadRequest, "", "password either does not match or invalid"}
	}
	
	/*
		Provide a role of a client
	*/
	var token = &Token{
		UserId:         user.Id,
		RoleId:         user.RoleId,
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

