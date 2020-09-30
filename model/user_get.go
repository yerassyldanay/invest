package model

import (
	"github.com/jinzhu/gorm"
	"invest/utils"
	"net/http"
	"sync"
)

/*
	get the full user info
*/
func (c *User) Get_full_info_of_this_user(by string) (utils.Msg) {
	var err error
	switch by {
	case "username":
		err = GetDB().First(c, "username=?", c.Username).Error
	case "email":
		err = GetDB().
			Raw("select u.* from users u join emails e on u.email_id = e.id where e.address = ?;", c.Email.Address).
				Scan(c).Error
	case "phone":
		phoneNumber := c.Phone.Ccode + c.Phone.Number
		err = GetDB().
			Raw("select u.* from users u join phones p on p.id = u.phone_id where p.ccode || p.number = ? ;", phoneNumber).
				Scan(c).Error
	default:
		err = GetDB().First(c, "id = ?", c.Id).Error
	}

	if err == gorm.ErrRecordNotFound {
		return utils.Msg{utils.ErrorNoSuchUser, 404, "", err.Error()}
	}

	var wg = sync.WaitGroup{}
	wg.Add(3)

	/*
		Phone
	 */
	go func(wgi *sync.WaitGroup) {
		defer wg.Done()
		_ = GetDB().First(&c.Phone, "id = ?", c.PhoneId)
	}(&wg)

	/*
		Email
	 */
	go func(wgi *sync.WaitGroup) {
		defer wg.Done()
		_ = GetDB().First(&c.Email, "id = ?", c.EmailId)
	}(&wg)

	/*
		Role
	 */
	go func(wgi *sync.WaitGroup) {
		defer wg.Done()
		_ = GetDB().First(&c.Role, "id = ?", c.RoleId)
	}(&wg)

	if c.OrganizationId > 0 {
		wg.Add(1)
		go func(wgi *sync.WaitGroup) {
			defer wg.Done()
			_ = GetDB().First(&c.Organization, "id = ?", c.OrganizationId)
		}(&wg)
	}

	wg.Wait()

	/*
		add statistics
	*/
	wg.Add(1)
	c.Add_statistics_to_this_user_on_project_statuses(&wg)
	wg.Wait()

	var password = c.Password
	c.Password = ""

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(*c)

	c.Password = password

	return utils.Msg{
		resp, http.StatusOK, "", "",
	}
}

/*
	Get Admins
 */
func (c *User) Get_admins() (utils.Msg) {
	users, err := c.OnlyGetPreloadedUsersByRole(utils.RoleAdmin, GetDB())
	if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	var usersMap = make([]map[string]interface{}, len(users))
	for _, user := range users {
		usersMap = append(usersMap, Struct_to_map(user))
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = usersMap

	return ReturnNoErrorWithResponseMessage(resp)
}

// get users by role (email, phone & role preloaded)
func (c *User) OnlyGetPreloadedUsersByRole(role string, tx *gorm.DB) (users []User, err error) {
	err = tx.Preload("Phone").Preload("Email").Preload("Role").
		Find(&users, "role_id = (select id from roles where name = ? limit 1)", role).Error
	return users, err
}

