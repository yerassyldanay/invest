package model

import "C"
import (
	"bytes"
	"github.com/jinzhu/gorm"
	"invest/utils"
	"sync"

	//"gorm.io/gorm/clause"
)

const GetLimit = 20

/*
	validate user information
 */
func (c *User) Is_spk_user_valid() bool {
	if ok := Validate_password(c.Password, "", ""); !ok {
		return false
	}

	if c.Username == "" || c.Fio == "" || c.Email.Address == "" || c.Phone.Ccode == "" || c.Phone.Number == "" {
		return false
	}

	return true
}

func (c *User) Remove_all_users_with_not_confirmed_email() (map[string]interface{}, error) {
	return nil, nil
}

/*
	create
*/
func (c *User) Create_user() (utils.Msg) {
	if c.Lang == "" {
		c.Lang = utils.DefaultContentLanguage
	}

	/*
		update seq ids
			refer to function description for more information
	 */
	_ = Update_sequence_id_thus_avoid_duplicate_primary_key_error(GetDB(), "default")

	trans := GetDB().Begin()
	defer func(){
		if trans != nil {
			trans.Rollback()
		}
	}()

	/*
		remove if the user account is not confirmed
	*/
	_, _ = c.Remove_all_users_with_not_confirmed_email()

	if ok := Validate_password(c.Password, nil, ""); !ok {
		return utils.Msg{
			Message: utils.ErrorInvalidPassword,
			Status:  400,
			ErrMsg:  "invalid password",
		}
	}

	if ok := c.Is_spk_user_valid(); !ok || !c.Phone.Is_valid() {
		return utils.Msg{
			Message: utils.ErrorInvalidParameters,
			Status:  400,
			ErrMsg:  "did not pass validation. crud user",
		}
	}

	/*
		hash the password
	*/
	hashed, err := utils.Convert_string_to_hash(c.Password)
	if err != nil {
		return utils.Msg{
			Message: utils.ErrorInternalServerError,
			Status:  500,
			ErrMsg:  err.Error(),
		}
	}
	c.Password = string(hashed)
	
	/*
		get role of the user
	 */
	if err := trans.Table(Role{}.TableName()).Where("name = ?", c.Role.Name).First(&c.Role).Error; err != nil {
		return utils.Msg{
			Message: utils.ErrorInternalDbError,
			Status:  417,
			ErrMsg:  err.Error(),
		}
	}
	c.RoleId = c.Role.Id

	/*
		create email & phone on db
	 */
	c.Email.Verified = true

	if err := trans.Create(&c.Email).Error; err != nil {
		return utils.Msg{
			Message: utils.ErrorInternalDbError,
			Status:  417,
			ErrMsg:  err.Error(),
		}
	}
	c.EmailId = c.Email.Id

	c.Phone.Verified = true
	if err := trans.Create(&c.Phone).Error; err != nil {
		return utils.Msg{
			Message: utils.ErrorInternalDbError,
			Status:  417,
			ErrMsg:  err.Error(),
		}
	}
	c.PhoneId = c.Phone.Id

	/*
		create a user with provided info
	*/
	c.Verified = true
	if err := trans.Create(&c).Error; err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	///*
	//	generate message
	// */
	//var sm = SendgridMessageStore{}
	//sm, _ = sm.Prepare_message_this_object(c, templates.Base_message_map_1_welcome)
	//
	//_, _ = sm.Send_message()

	if err := trans.Commit().Error; err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	trans = nil
	return utils.Msg{utils.NoErrorFineEverthingOk, 201, "", ""}
}

/*
	update password
*/
func (c *User) Update_own_user_password_by_user_id() (utils.Msg) {
	if ok := Validate_password(c.Password, nil, ""); !ok {
		return utils.Msg{utils.ErrorInvalidPassword, 400, "", "invalid password"}
	}

	b, err := utils.Convert_string_to_hash(c.Password)
	if err != nil {
		return utils.Msg{utils.ErrorInternalServerError, 500, "", err.Error()}
	}
	c.Password = string(b)

	if err := GetDB().Model(&User{Id: c.Id}).Update("password", c.Password).Error; err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	return utils.Msg{utils.NoErrorFineEverthingOk, 200, "", ""}
}

/*
	who can update info:
		admin other's info & user their own info
 */
func (c *User) Get_permission_for_update(whois string) utils.Msg {
	//
	//if err := GetDB().Table(c.TableName()).Where("id = ?", c.Id).Error; err != {
	//
	//}
	return utils.Msg{}
}

/*
	update info
*/
func (c *User) Update_user_info_except_for_email_address_and_password_by_user_id() (utils.Msg) {

	var trans = GetDB().Begin()
	defer func() { if trans != nil { trans.Rollback() } }()

	/*
		validate info that will be stored
			note: omit password, email info as they won't be not stored / updated
	 */
	var ok =  c.Phone.Is_valid() && c.Fio != "" && c.Role.Name != "" && c.Id != 0
	if !ok {
			return utils.Msg{utils.ErrorInvalidParameters, 400, "", "invalid parameters. failed user info update validation"}
	}

	/*
		get role id as a new role can be set
	 */
	if err := trans.Table(Role{}.TableName()).Where("name = ?", c.Role.Name).First(&c.Role).Error;
		err != nil {
			return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}
	c.RoleId = c.Role.Id

	/*
		check db and find the phone id
			or create such phone
	 */
	var phone = Phone{}
	if err := trans.Table(phone.TableName()).
		Where("ccode=$1 and number=$2", c.Phone.Ccode, c.Phone.Number).First(&phone).Error;
	err != nil && err != gorm.ErrRecordNotFound {
		/*
			db error occurred
		 */
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	} else if err == gorm.ErrRecordNotFound {
		/*
			record not found -> create a phone
		 */
		c.Phone.Verified = true
		if err := trans.Create(&c.Phone).Error; err != nil {
			return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
		}
	} else {
		/*
			record found -> delete old phone -> assign the number id
		 */
		//if err = trans.Exec(" delete from phones p where p.id = (select p.id from users u where u.id = ? ) ;", c.Id).Error; err != nil {
		//	return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
		//}

		c.Phone.Id = phone.Id
	}
	c.PhoneId = c.Phone.Id

	/*
		at this moment:
			phone
			role
			other info: fio & position
		are valid & ready
	 */
	if err := trans.Model(&User{Id: c.Id}).Where("id = ?", c.Id).Updates(User{
		Fio:      		c.Fio,
		RoleId:  	 	c.RoleId,
		PhoneId:  		c.PhoneId,
		Verified: 		true,
	}).Error;
		err != nil {
			return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	trans.Commit()
	trans = nil

	return utils.Msg{utils.NoErrorFineEverthingOk, 200, "", ""}
}

/*
	[]{"manager", "lawyer", "financier"}, 10 - starting from this point
		limit = 20
*/
func (a *User) Get_users_by_roles(roles []string, offset string) (utils.Msg) {
	type Info struct {
		Users 		[]User
	}

	var users = Info{}
	var bquery = bytes.Buffer{}

	bquery.WriteString("select * from users u join roles r on u.role_id = r.id where r.name in (")
	for i, role := range roles {
		if ok := utils.Is_it_free_from_sql_injection(role); !ok {
			return utils.Msg{utils.ErrorInternalDbError, 417, "", ""}
		}

		bquery.WriteString("'" + role + "'")
		if i != len(roles) - 1 {
			bquery.WriteString(", ")
		}
	}
	bquery.WriteString(");");

	err := GetDB().Raw(bquery.String()).Scan(&users.Users).Error
	if err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", ""}
	}

	var wg = sync.WaitGroup{}
	for i, _ := range users.Users {
		users.Users[i].Password = ""

		wg.Add(1)
		go users.Users[i].Add_statistics_to_this_user_on_project_statuses(&wg)
	}
	wg.Wait()

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(users)["users"]

	return utils.Msg{resp, 200, "", ""}
}

func (a *User) Get_all_users_except_admins(offset string) (utils.Msg) {

	var users []User
	
	if err := GetDB().Preload("Role").Preload("Email").Preload("Phone").
		Offset(offset).Limit(GetLimit).Where("id != (select id from roles where name = 'admin' limit 1)").Find(&users).Error; err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	var wg = sync.WaitGroup{}
	var infos = []map[string]interface{}{}
	for i, _ := range users {
		wg.Add(1)
		go users[i].Add_statistics_to_this_user_on_project_statuses(&wg)
	}
	wg.Wait()

	for _, user := range users {
		user.Password = ""
		infos = append(infos, Struct_to_map(user))
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = infos

	return utils.Msg{resp, 200, "", ""}
}

