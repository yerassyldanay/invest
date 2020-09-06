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
func (sis *SignIn) Sign_in() (*utils.Msg) {
	var user = User{}
	regexpOnlyLetters, err := regexp.Compile("[a-z]+")
	if err != nil {
		return &utils.Msg{
			Message: utils.ErrorInternalServerError, Status:  http.StatusInternalServerError, ErrMsg:  err.Error(),
		}
	}

	//fmt.Println(strings.ToLower(sis.KeyUsername))
	if ok := regexpOnlyLetters.Match([]byte(strings.ToLower(sis.KeyUsername))); !ok {
			return &utils.Msg{utils.ErrorInvalidParameters, http.StatusBadRequest, "", "regexp does not match with key"}
	}

	user.Username = sis.Username
	msg := user.Get_full_info_of_this_user(sis.KeyUsername)
	if msg.ErrMsg != "" {
		return msg
	} else {
		if err := GetDB().Table(User{}.TableName()).Where(sis.KeyUsername + "=?", sis.Username).First(&user).Error; err != nil {
			return &utils.Msg{utils.ErrorNoSuchUser, http.StatusNotFound, "", err.Error()}
		}
	}

	/*
		here we check whether two passwords (a provided password and password on db_create_fake_data) MATCH
	*/
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(sis.Password)); err == bcrypt.ErrMismatchedHashAndPassword || err != nil {
		return &utils.Msg{utils.ErrorInvalidPassword, http.StatusBadRequest, "", "password either does not match or invalid"}
	}
	
	/*
		Provide a role of a client
	*/
	var token = &Token{
		UserId:         user.Id,
		RoleId:         user.RoleId,
		Exp:            time.Now().Add(time.Hour * 24),
	}

	var token_hashed = jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), token)
	token_string, err := token_hashed.SignedString([]byte(os.Getenv("TOKEN_PASSWORD")))

	if err != nil {
		return &utils.Msg{utils.ErrorInternalIssueOrInvalidPassword, http.StatusInternalServerError, "", err.Error()}
	}

	token_string = "Bearer " + token_string
	sis.TokenCompound = token_string

	user.Password = ""

	resp := utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(user)

	return &utils.Msg{
		resp, http.StatusOK,"", "",
	}
}

/*
	get the full user info
 */
func (c *User) Get_full_info_of_this_user(by string) (*utils.Msg) {
	var d = GetDB().Preload("Role").Preload("Email").Preload("Phone")
	switch by {
	case "username":
		d = d.Table(User{}.TableName()).Where("username=?", c.Username)
	default:
		d = d.Table(User{}.TableName()).Where("id=?", c.Id)
	}

	if err := d.Preload("Phone").Preload("Email").First(c).Error; err != nil {
		return &utils.Msg{utils.ErrorNoSuchUser, http.StatusNotFound, "", err.Error()}
	}
	c.Password = ""

	/*
		permissions will be uploaded to data being sent
	 */
	_ = c.Role.Get_permissions_by_role_id()

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(*c)

	return &utils.Msg{
		resp, http.StatusOK, "", "",
	}
}
