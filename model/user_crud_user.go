package model

import (
	"encoding/json"
	"errors"
	"gopkg.in/validator.v2"
	"invest/templates"
	"invest/utils"
	//"gorm.io/gorm/clause"
)

const GetLimit = 20

func (c *User) Remove_all_users_with_not_confirmed_email() (map[string]interface{}, error) {
	return nil, nil
}

/*
	create
*/
func (c *User) Create_user() (map[string]interface{}, error) {
	if c.Lang == "" {
		c.Lang = utils.DefaultLContentanguage
	}

	trans := GetDB().Begin()
	defer func(){
		if trans != nil {
			trans.Rollback()
		}
	}()

	/*
		remove if the user is not confirmed
	*/
	if resp, err := c.Remove_all_users_with_not_confirmed_email(); err != nil {
		return resp, err
	}

	if ok := Validate_password(c.Password); !ok {
		return utils.ErrorInvalidPassword, errors.New("invalid password")
	}

	if err := validator.Validate(c); err != nil {
		return utils.ErrorInvalidParameters, err
	}

	/*
		hash the password
	*/
	hashed, err := utils.Convert_string_to_hash(c.Password)
	if err != nil {
		return utils.ErrorInternalServerError, err
	}
	c.Password = string(hashed)
	
	/*
		get role of the user
	 */
	if err := trans.Model(&Role{}).Where("name", c.Role.Name).First(&c.Role).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}

	/*
		create email & phone on db
	 */
	if err := GetDB().Create(&c.Email).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}
	c.EmailId = c.Email.Id

	if err := GetDB().Create(&c.Phone).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}
	c.PhoneId = c.Phone.Id

	/*
		create a user with provided info
	*/
	if err := trans.Create(&c).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}

	/*
		generate message
	 */
	var sm = SendgridMessageStore{}
	sm, _ = sm.Prepare_message_this_object(c, templates.Base_message_map_1_welcome)

	if resp, err := sm.Send_message(); err != nil {
		return resp, err
	}

	return utils.NoErrorFineEverthingOk, trans.Commit().Error
}

/*
	update password
*/
func (c *User) Update_user_password() (map[string]interface{}, error) {
	if ok := Validate_password(c.Password); !ok {
		return utils.ErrorInvalidPassword, errors.New("invalid password")
	}

	if err := GetDB().Table(User{}.TableName()).Where("id=?", c.Id).First(c).Error; err != nil {
		return utils.ErrorInvalidParameters, errors.New("invalid parameters")
	}

	b, err := utils.Convert_string_to_hash(c.Password)
	if err != nil {
		return utils.ErrorInternalDbError, errors.New("invalid parameters")
	}

	c.Password = string(b)

	if err := GetDB().Model(&User{}).Where("id=?", c.Id).Update("password", c.Password).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}

	return utils.NoErrorFineEverthingOk, nil
}

/*
	update info
*/
func (c *User) Update_user_info() (map[string]interface{}, error) {

	if err := GetDB().Table(User{}.TableName()).Where("id=?", c.Id).First(&User{}).Error; err != nil {
		return utils.ErrorInvalidParameters, err
	}

	if err := GetDB().Table(User{}.TableName()).Where("id=?", c.Id).Select("fio", "position").Error;
		err != nil {
			return utils.ErrorInternalDbError, err
	}

	return utils.NoErrorFineEverthingOk, nil
}

/*
	[]{"manager", "lawyer", "financier"}, 10 - starting from this point
		limit = 20
*/
func (a *User) Get_users_by_roles(roles []string, offset string) (map[string]interface{}, error) {
	var users = []User{}
	if err := GetDB().Table(User{}.TableName()).Omit("password").Where("role in (?)", roles).Offset(offset).Limit(GetLimit).Find(&users).Error; err != nil {
		return utils.ErrorNoSuchUser, err
	}

	for i, user := range users {
		user.Password = ""
		user.Phone.SentCode = ""
		user.Email.SentCode = ""
		user.Email.SentHash = ""

		users[i] = user
	}

	var t = utils.NoErrorFineEverthingOk
	if b, err := json.Marshal(users); err != nil {
		return utils.ErrorInternalServerError, nil
	} else {
		t["info"] = string(b)
	}

	return t, nil
}

func (a *User) Get_all_users(offset string) (map[string]interface{}, error) {
	type Temp struct{
		User
		Address			string				`json:"address"`
		Ccode			string				`json:"ccode"`
		Number			string				`json:"number"`
	}
	var users []Temp

	var main_query = `select u.*, e.address as address, p.ccode as ccode, p.number as number from users u
			join emails e on u.email_id = e.id
			join phones p on u.phone_id = p.id 
			offset ? limit ? ; `
	
	if err := GetDB().Raw(main_query, offset, GetLimit).Scan(&users).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}
	
	var infos = []map[string]interface{}{}
	for _, user := range users {
		infos = append(infos, Struct_to_map(user))
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = infos

	return resp, nil
}

