package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"invest/templates"
	"invest/utils"

	"net/http"
	"net/url"
	"strings"
	"time"
)

/*
	validate sign up user info
 */
func (c *User) Is_user_info_valid() bool {
	if c.Username != "" && c.Password != "" && c.Organization.Bin != "" &&
			c.Email.Address != "" {
		return true
	}

	return false
}

/*
	signup for investors
		201 - created
		400 - bad request
		409 - already in use
		417 - db error
		422 - could not sent message & not stored on db
 */
func (c *User) Sign_Up() (utils.Msg) {
	var user User
	var count int

	if ok := c.Is_user_info_valid(); !ok {
		return utils.Msg{utils.ErrorInvalidParameters, 400, "", "invalid parameters. user sign up"}
	}

	/*
		is username or fio in use
	 */
	if err := GetDB().Preload("Email").Preload("Phone").Preload("Role").Preload("Organization").
		Table(User{}.TableName()).Where("username = ?", c.Username).First(&user).Count(&count).Error;
		err != nil && err != gorm.ErrRecordNotFound {
			return utils.Msg{utils.ErrorInternalDbError, http.StatusExpectationFailed, "", err.Error()}
	} else if count > 0 {
		return utils.Msg{utils.ErrorUsernameOrFioIsAreadyInUse, http.StatusConflict, "", "already in use: " + c.Fio}
	}

	// validate & hash
	err := Validate_password(c.Password)
	hashed, err2 := utils.Convert_string_to_hash(c.Password)

	switch {
	case err != nil:
		return ReturnInternalDbError(err.Error())
	case err2 != nil:
		return ReturnInternalDbError(err2.Error())
	}

	if err := GetDB().Table(Role{}.TableName()).Where("name = ?", utils.RoleInvestor).
		First(&c.Role).Error; err != nil {
			return utils.Msg{utils.ErrorInternalDbError, 417, "", ""}
	}

	c.RoleId = c.Role.Id
	c.Password = hashed

	/*
		create organization or get old one
	 */
	c.Organization.Lang = c.Lang
	c.Organization.Create_or_get_organization_from_db_by_bin(GetDB())
	c.OrganizationId = c.Organization.Id

	/*
		starting a transaction, which is be rolled back in case of an error
	 */
	trans := GetDB().Begin()
	defer func() {
		if trans != nil {
			trans.Rollback()
		}
	}()

	/*
		these code and link will be sent to the user
	 */
	scode := utils.Generate_Random_Number(utils.MaxNumberOfDigitsSentByEmail)
	shash := utils.Generate_Random_String(utils.MaxNumberOfCharactersSentByEmail)

	/*
		NOTE THAT it an email address can be used by somebody already
			case:
				case: verified & being used by someone
				case: not verified & missed deadline & once message was sent
				case: not verified & have time to verify (message has been sent to email)
			case: no message -> no verification
	*/
	if trans.Table(Email{}.TableName()).Where("address = ?", c.Email.Address).First(&c.Email).Count(&count); count > 0 {

		if c.Email.Verified {
			return utils.Msg{utils.ErrorEmailIsAreadyInUse, http.StatusConflict, "", "email address already in use"}

		} else if c.Email.Verified == false && c.Email.Deadline.Before(time.Now()) {
			var tuser = User{}
			if err := trans.Exec("select u.* from users u inner join emails e on u.email_id = e.id where u.email_id = ?;", c.Email.Id).First(&tuser).Error;
				err != nil && err != gorm.ErrRecordNotFound {
					return utils.Msg{ utils.ErrorInternalDbError, http.StatusExpectationFailed, "", err.Error()}
			} else if err != gorm.ErrRecordNotFound {
					trans.Delete(Email{}, "id=?", tuser.EmailId)
					trans.Delete(Phone{}, "id=?", tuser.PhoneId)
					trans.Delete(User{}, "id=?", tuser.Id)
			}

		} else {
			return utils.Msg{utils.ErrorAlreadySentLinkToEmail, http.StatusConflict, "", "a link has already been sent"}
		}
	}

	/*
		store the email, phone & get ids
	 */
	c.Email.SentCode = scode
	c.Email.SentHash = shash
	c.Email.Deadline = utils.GetCurrentTime().Add(time.Hour * 24)

	if err := trans.Create(&c.Email).Error; err != nil {
		//trans.Rollback()
		return utils.Msg{utils.ErrorInternalDbError, http.StatusExpectationFailed, "", err.Error()}
	}
	c.EmailId = c.Email.Id

	if err := trans.Create(&c.Phone).Error; err != nil {
		return utils.Msg{utils.ErrorInternalDbError, http.StatusExpectationFailed, "", err.Error()}
	}
	c.PhoneId = c.Phone.Id

	if err := trans.Table(Role{}.TableName()).Where("name=?", utils.RoleInvestor).First(&c.Role).Error;
		err != nil {
			return utils.Msg{utils.ErrorInternalDbError, http.StatusExpectationFailed, "", err.Error()}
	}
	c.RoleId = c.Role.Id

	if err := trans.Create(c).Error; err != nil {
		return utils.Msg{utils.ErrorFailedToCreateAnAccount, http.StatusExpectationFailed, "", err.Error()}
	}

	/*
		Send message with a code
	*/
	var subject, html, page string
	switch strings.ToLower(c.Lang) {
	case "kaz":
		subject = templates.Base_email_subject_kaz
		html = templates.Base_email_html_kaz
		page = templates.Base_email_page_kaz
	case "rus":
		subject = templates.Base_email_subject_rus
		html = templates.Base_email_html_rus
		page = templates.Base_email_page_rus
	default:
		subject = templates.Base_email_subject_eng
		html = templates.Base_email_html_eng
		page = templates.Base_email_page_eng
	}

	queryParam := url.Values{
		"hashcode": []string{
			shash,
		},
		"key": []string{
			"shash",
		},
	}

	urlPath := url.URL{
		Scheme:     "https",
		Host:       "tsrk.xyz",
		Path:		"/v1/confirmation/email",
		RawQuery: 	queryParam.Encode(),
	}

	html = fmt.Sprintf(html, scode, urlPath.String())
	page = fmt.Sprintf(page, scode, urlPath.String())

	//fmt.Println("urlPath.String(): ", urlPath.String())

	var sm = SendgridMessageStore{
		From:     utils.BaseEmailAddress,
		To:       c.Email.Address,
		FromName: utils.BaseEmailName,
		ToName:   c.Fio,
		SendgridMessage:   SendgridMessage{
			Subject:   subject,
			PlainText: page,
			HTML:      html,
		},
		Status: 200,
		Created: utils.GetCurrentTime(),
	}

	resp, err := sm.Send_message()
	if err != nil {
		return utils.Msg{resp, http.StatusUnprocessableEntity, "", err.Error()}
	}

	trans.Commit()
	return utils.Msg{
		utils.NoErrorFineEverthingOk, http.StatusCreated, "", "",
	}
}
