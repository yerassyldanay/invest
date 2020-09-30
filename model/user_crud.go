package model

import "C"
import (
	"bytes"
	"invest/utils"
	"time"

	"sync"

	//"gorm.io/gorm/clause"
)

const GetLimit = 20

func (c *User) Remove_all_users_with_not_confirmed_email() (map[string]interface{}, error) {
	return nil, nil
}

/*
	create
*/
func (c *User) Create_user_without_check() (utils.Msg) {
	if c.Lang == "" {
		c.Lang = utils.DefaultContentLanguage
	}

	/*
		update seq ids
			refer to function description for more information
	 */
	_ = Update_sequence_id_thus_avoid_duplicate_primary_key_error(GetDB(), "users")
	_ = Update_sequence_id_thus_avoid_duplicate_primary_key_error(GetDB(), "phones")
	_ = Update_sequence_id_thus_avoid_duplicate_primary_key_error(GetDB(), "emails")

	trans := GetDB().Begin()
	defer func(){
		if trans != nil {
			trans.Rollback()
		}
	}()

	// remove if the user account is not confirmed
	_, _ = c.Remove_all_users_with_not_confirmed_email()

	// validate password
	if ok := Validate_password(c.Password, nil, ""); !ok {
		return ReturnInvalidParameters("password is invalid")
	}

	// hash the password
	hashed, err := utils.Convert_string_to_hash(c.Password)
	if err != nil {
		return ReturnInvalidParameters("could not hash password. user. hash")
	}
	c.Password = string(hashed)
	
	// get role of the user
	if err := c.Role.OnlyGetByName(trans); err != nil {
		return ReturnInternalDbError(err.Error())
	}
	c.RoleId = c.Role.Id

	c.Email.Verified = true
	c.Email.SentCode = ""
	c.Email.SentHash = ""
	c.Email.Deadline = time.Time{}

	// store email
	if err := c.Email.OnlyCreate(trans); err != nil {
		return ReturnInternalDbError(err.Error())
	}
	c.EmailId = c.Email.Id

	// store phone
	c.Phone.Verified = true
	if err := c.Phone.OnlyCreate(trans); err != nil {
		return ReturnInternalDbError(err.Error())
	}
	c.PhoneId = c.Phone.Id

	// create a user with provided info
	c.Verified = true
	if err := c.OnlyCreate(trans); err != nil {
		return ReturnInternalDbError(err.Error())
	}

	if err := trans.Commit().Error; err != nil {
		return ReturnInternalDbError(err.Error())
	}

	trans = nil
	return ReturnSuccessfullyCreated()
}

/*
	[]{"manager", "lawyer", "financier"}, 10 - starting from this point
		limit = 20
*/
func (a *User) Get_users_by_roles(roles []string, offset interface{}) (utils.Msg) {
	type Info struct {
		Users 		[]User
	}

	var users = Info{}
	var bquery = bytes.Buffer{}

	bquery.WriteString("select * from users u join roles r on u.role_id = r.id where r.name in (")
	for i, role := range roles {
		if ok := utils.Is_it_free_from_sql_injection(role); !ok {
			return ReturnInvalidParameters("sql injection check. did not pass")
		}

		bquery.WriteString("'" + role + "'")
		if i != len(roles) - 1 {
			bquery.WriteString(", ")
		}
	}
	bquery.WriteString(");");

	var main_query = bquery.String()
	err := GetDB().Raw(main_query).Scan(&users.Users).Error
	if err != nil {
		return ReturnInternalDbError(err.Error())
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

	return ReturnNoErrorWithResponseMessage(resp)
}

func (a *User) Get_all_users_except_admins(offset string) (utils.Msg) {

	var users []User

	if err := GetDB().Preload("Role").Preload("Email").Preload("Phone").
		Offset(offset).Limit(GetLimit).Where("id != (select id from roles where name = 'admin' limit 1)").Find(&users).Error; err != nil {
		return ReturnInternalDbError(err.Error())
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

	return ReturnNoErrorWithResponseMessage(resp)
}

