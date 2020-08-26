package model

import (
	"bytes"
	"fmt"
	"invest/utils"
	"net/http"
)

func(r *Role) Validate() bool {
	if r.Name == "" || r.Description == "" {
		return false
	}

	return true
}

/*
	200 - ok
	400 - invalid parameters
	417 - internal db error
 */
func(r *Role) Create_a_role_with_permissions() (*utils.Msg) {
	if ok := r.Validate(); !ok {
		return &utils.Msg{
			Message: utils.ErrorInvalidParameters, Status:  http.StatusBadRequest, ErrMsg: utils.Error_msg_invalid_parameters_passed,
		}
	}

	if err := GetDB().Exec("select * from permissions where id in (?);", r.PermissionsSent).Scan(&r.Permissions).Error;
		err != nil {
			return &utils.Msg{
				Message: utils.ErrorInternalDbError, Status:  http.StatusExpectationFailed, ErrMsg:  err.Error(),
			}
	}

	if err := GetDB().Create(r).Error; err != nil {
		return &utils.Msg{Message: utils.ErrorInternalDbError, Status: http.StatusExpectationFailed, ErrMsg: err.Error()}
	}

	return &utils.Msg{
		Message: utils.NoErrorFineEverthingOk,
		Status:  200,
		ErrMsg:  "",
	}
}

func (r *Role) Update_role_name_description_and_permissions() (*utils.Msg) {
	if ok := r.Validate(); !ok {
		return &utils.Msg{
			Message: utils.ErrorInvalidParameters, Status:  http.StatusBadRequest, ErrMsg: utils.Error_msg_invalid_parameters_passed,
		}
	}

	if err := GetDB().Table(r.TableName()).Where("id=?", r.Id).Updates(map[string]string{
		"name": r.Name,
		"description": r.Description,
	}).Error; err != nil {
		return &utils.Msg{
			Message: utils.ErrorInternalDbError, Status:  http.StatusExpectationFailed, ErrMsg: err.Error(),
		}
	}

	return &utils.Msg{
		Message: utils.NoErrorFineEverthingOk, Status:  200, ErrMsg: "",
	}
}

/*
	this method expects
		1. role_id
		2. ids of permissions
 */
func (r *Role) Set_a_list_of_permissions() (*utils.Msg) {
	var arrper []uint64
	for i, _ := range r.Permissions {
		arrper = append(arrper, r.Permissions[i].Id)
	}

	if err := GetDB().Table(Permission{}.TableName()).Where("id in ?", arrper).Find(&r.Permissions).Error;
		err != nil {
			return &utils.Msg{
				Message: utils.ErrorInternalDbError, Status:  http.StatusExpectationFailed, ErrMsg:  err.Error(),
			}
	}

	var main_query bytes.Buffer
	_, _ = main_query.WriteString(" insert into roles_permissions values ")

	for i, _ := range r.Permissions {
		if i != 0 {
			main_query.WriteString(" , ")
		}
		main_query.WriteString(fmt.Sprintf(" (%d, %d) ", r.Id, r.Permissions[i].Id))
	}

	main_query.WriteString(" ; ")

	if err := GetDB().Exec(main_query.String()).Error; err != nil {
		return &utils.Msg{
			Message: utils.ErrorInternalDbError, Status:  http.StatusExpectationFailed, ErrMsg: err.Error() ,
		}
	}

	return &utils.Msg{
		Message: utils.NoErrorFineEverthingOk, Status:  http.StatusOK, ErrMsg: "",
	}
}

func (r *Role) Remove_a_list_of_permissions() (*utils.Msg) {
	var arrper []uint64
	for i, _ := range r.Permissions {
		arrper = append(arrper, r.Permissions[i].Id)
	}

	if err := GetDB().Table("roles_permissions").
		Delete("role_id=$1 and permission_id in $2", r.Id, arrper).Error; err != nil {
			return &utils.Msg{
				Message: utils.ErrorInternalDbError, Status:  http.StatusExpectationFailed, ErrMsg: err.Error(),
			}
	}

	return &utils.Msg{
		Message: utils.NoErrorFineEverthingOk, Status:  http.StatusOK, ErrMsg: "",
	}
}

/*
	delete the role if there is no user with such a role
 */
func (r *Role) Delete_this_role() (*utils.Msg) {
	var count int
	if err := GetDB().Table(User{}.TableName()).Where("role_id=?", r.Id).Count(&count).Error;
		err != nil {
			return &utils.Msg{
				Message: utils.ErrorInternalDbError, Status:  http.StatusExpectationFailed, ErrMsg: err.Error(),
			}
	}

	if count > 0 {
		return &utils.Msg{
			Message: utils.ErrorMethodNotAllowed, Status:  http.StatusMethodNotAllowed, ErrMsg: utils.Error_msg_method_not_allowed,
		}
	}

	var trans = GetDB().Begin()
	defer func() {if trans != nil {trans.Rollback()} }()

	if err := trans.Exec("delete from roles_permissions where role_id = ? ; ", r.Id).Error;
		err != nil {
			return &utils.Msg{
				Message: utils.ErrorInternalDbError, Status:  http.StatusExpectationFailed, ErrMsg: err.Error(),
			}
	}

	/*
		delete a role
			note: rows with role_id in roles_permissions will be automatically deleted
	 */
	if err := trans.Table(Role{}.TableName()).Where("id=?", r.Id).Delete(&Role{}).Error;
		err != nil {
			return &utils.Msg{
				Message: utils.ErrorInternalDbError, Status:  http.StatusExpectationFailed, ErrMsg: err.Error(),
			}
	}

	trans.Commit()
	trans = nil

	return &utils.Msg{
		Message: utils.NoErrorFineEverthingOk, Status:  http.StatusOK, ErrMsg: "",
	}
}

/*
	get roles
 */
func (r *Role) Get_roles(offset string) (*utils.Msg) {
	var roles = struct {
		Info			[]Role
	}{}
	if err := GetDB().Preload("permissions").Table(Role{}.TableName()).Offset(offset).Limit(GetLimit).Find(&roles.Info).Error; err != nil {
		var resp = utils.ErrorInternalDbError
		resp["info"] = []map[string]interface{}{}
		return &utils.Msg{
			Message: resp, Status:  http.StatusExpectationFailed, ErrMsg: err.Error(),
		}
	}

	var resp = utils.ErrorInternalDbError
	resp["info"] = Struct_to_map(roles)

	return &utils.Msg{
		Message: utils.NoErrorFineEverthingOk, Status:  http.StatusOK, ErrMsg: "",
	}
}

func (r *Role) Get_role_info() (*utils.Msg) {
	if err := GetDB().Preload("Permissions").Table(r.TableName()).Where("id=?", r.Id).Error;
		err != nil {
			return &utils.Msg{Message: utils.ErrorInternalDbError, Status: http.StatusExpectationFailed, ErrMsg: err.Error()}
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(&r)

	return &utils.Msg{
		Message: resp, Status:  http.StatusOK, ErrMsg: "",
	}
}
