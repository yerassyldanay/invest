package model

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"invest/utils"
	"os"
	"regexp"
	"strings"
	"time"
)

/*
	return token, response, err
 */
func (sis *SignIn) Sign_in() (string, map[string]interface{}, error) {
	var user = User{}
	regexpOnlyLetters, err := regexp.Compile("[a-z]+")

	if err != nil {
		return "", utils.ErrorInternalServerError, err
	}

	//fmt.Println(strings.ToLower(sis.KeyUsername))
	if ok := regexpOnlyLetters.Match([]byte(strings.ToLower(sis.KeyUsername))); !ok {
			return "", utils.ErrorInvalidParameters, errors.New("regexp does not match with key")
	}

	user.Username = sis.Username
	resp, err := user.Get_full_info_of_this_user("username")
	if err != nil {
		return "", resp, err
	} else {
		if err := GetDB().Table(User{}.TableName()).Where(sis.KeyUsername + "=?", sis.Username).First(&user).Error; err != nil {
			return "", utils.ErrorNoSuchUser, err
		}
	}

	/*
		here we check whether two passwords (a provided password and password on db_create_fake_data) MATCH
	*/
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(sis.Password)); err == bcrypt.ErrMismatchedHashAndPassword || err != nil {
		return "", utils.ErrorInvalidPassword, err
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
		return "", utils.ErrorInternalIssueOrInvalidPassword, err
	}

	token_string = "Bearer " + token_string
	user.Password = ""

	resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(user)

	return token_string, resp, nil
}

/*
	get the full user info
 */
func (c *User) Get_full_info_of_this_user(by string) (map[string]interface{}, error) {
	var d *gorm.DB
	if by == "username" {
		d = GetDB().Table(User{}.TableName()).Where("username=?", c.Username)
	} else {
		d = GetDB().Table(User{}.TableName()).Where("id=?", c.Id)
	}

	if err := d.First(c).Error; err != nil {
		return utils.ErrorNoSuchUser, err
	}
	c.Password = ""

	if err := GetDB().Table(Email{}.TableName()).Where("id=?", c.EmailId ).First(&c.Email).Error;
		err != nil {
			return utils.ErrorNoSuchUser, err
	}

	if err := GetDB().Model(&Phone{}).Where("id=?", c.PhoneId).First(&c.Phone).Error;
		err != nil {
			return utils.ErrorNoSuchUser, err
	}

	if err := GetDB().Model(&Role{}).Where("id=?", c.RoleId).First(&c.Role).Error;
		err != nil {
			return utils.ErrorNoSuchUser, err
	}

	/*
		permissions will be uploaded to data being sent
	 */
	_ = c.Role.Get_permissions_by_role_id()

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(*c)

	return resp, nil
}

/*
	get user info by other keywords
 */
//func (c *User) Get_full_info_by_other_keywords(key string) (map[string]interface{}, error) {
//
//}


