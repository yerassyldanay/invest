package model

import (
	"bytes"
	"fmt"
	"invest/utils"
	"net/http"
	"strconv"
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
func(r *Role) Create_a_role_with_permissions() (utils.Msg) {
	if ok := r.Validate(); !ok {
		return utils.Msg{
			Message: utils.ErrorInvalidParameters, Status:  http.StatusBadRequest, ErrMsg: utils.Error_msg_invalid_parameters_passed,
		}
	}

	var permForQuery bytes.Buffer

	for i, perm := range r.PermissionsSent {
		if i != 0 {
			permForQuery.WriteString(", ")
		}
		permForQuery.WriteString(strconv.Itoa(int(perm)))
	}

	var main_query = fmt.Sprintf(" id in (%s) ", permForQuery.String())
	if err := GetDB().Find(&r.Permissions, main_query).Error;
		err != nil {
			return utils.Msg{
				Message: utils.ErrorInternalDbError, Status:  http.StatusExpectationFailed, ErrMsg:  err.Error(),
			}
	}

	if err := GetDB().Create(r).Error; err != nil {
		return utils.Msg{Message: utils.ErrorInternalDbError, Status: http.StatusExpectationFailed, ErrMsg: err.Error()}
	}

	return utils.Msg{
		Message: utils.NoErrorFineEverthingOk,
		Status:  200,
		ErrMsg:  "",
	}
}

func (r *Role) Update_role_name_description_and_permissions() (utils.Msg) {
	if ok := r.Validate(); !ok {
		return utils.Msg{
			Message: utils.ErrorInvalidParameters, Status:  http.StatusBadRequest, ErrMsg: utils.Error_msg_invalid_parameters_passed,
		}
	}

	if err := GetDB().Table(r.TableName()).Where("id=?", r.Id).Updates(map[string]string{
		"name": r.Name,
		"description": r.Description,
	}).Error; err != nil {
		return utils.Msg{
			Message: utils.ErrorInternalDbError, Status:  http.StatusExpectationFailed, ErrMsg: err.Error(),
		}
	}

	return utils.Msg{
		Message: utils.NoErrorFineEverthingOk, Status:  200, ErrMsg: "",
	}
}

/*
	helper function
 */
func Convert_list_to_string_seperate_by_given_string(list []uint64, sep string) string {
	var query bytes.Buffer
	for i, elem := range list {
		if i != 0 {
			query.WriteString(sep)
		}

		query.WriteString(fmt.Sprintf("%v", elem))
	}

	//fmt.Println("query.String(): ", query.String())
	return query.String()
}

/*
	this method expects
		1. role id
		2. ids of permissions
 */
func (r *Role) Add_a_list_of_permissions() (utils.Msg) {

	if err := GetDB().Preload("Permissions").Find(r, "id = ?", r.Id).Error;
		err != nil {
			return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	var querystr = Convert_list_to_string_seperate_by_given_string(r.PermissionsSent, ", ")
	querystr = fmt.Sprintf(" id in (%s) ", querystr)

	var permissions = []Permission{}
	if err := GetDB().Find(&permissions, querystr).Error;
		err != nil {
			return utils.Msg{
				Message: utils.ErrorInternalDbError, Status:  http.StatusExpectationFailed, ErrMsg:  err.Error(),
			}
	}

	/*
		add permissions to the list
	 */
	for _, permission := range permissions {
		r.Permissions = append(r.Permissions, permission)
	}

	/*
		save the results
	 */
	if err := GetDB().Save(r).Error; err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	return utils.Msg{
		Message: utils.NoErrorFineEverthingOk, Status:  http.StatusOK, ErrMsg: "",
	}
}

func (r *Role) Remove_a_list_of_permissions() (utils.Msg) {

	var querystr = Convert_list_to_string_seperate_by_given_string(r.PermissionsSent, ", ")

	/*
		deleting role & permissions relations
	 */
	querystr = fmt.Sprintf( " delete from roles_permissions where role_id = %d and permission_id in (%s); ", r.Id, querystr )
	if err := GetDB().Exec(querystr).Error; err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	return utils.Msg{
		Message: utils.NoErrorFineEverthingOk, Status:  http.StatusOK, ErrMsg: "",
	}
}

/*
	delete the role if there is no user with such a role
 */
func (r *Role) Delete_this_role() (utils.Msg) {
	var count int
	if err := GetDB().Table(User{}.TableName()).Where("role_id=?", r.Id).Count(&count).Error;
		err != nil {
			return utils.Msg{
				Message: utils.ErrorInternalDbError, Status:  http.StatusExpectationFailed, ErrMsg: err.Error(),
			}
	}

	if count > 0 {
		return utils.Msg{
			Message: utils.ErrorMethodNotAllowed, Status:  http.StatusMethodNotAllowed, ErrMsg: utils.Error_msg_method_not_allowed,
		}
	}

	var trans = GetDB().Begin()
	defer func() {if trans != nil {trans.Rollback()} }()

	if err := trans.Exec("delete from roles_permissions where role_id = ? ; ", r.Id).Error;
		err != nil {
			return utils.Msg{
				Message: utils.ErrorInternalDbError, Status:  http.StatusExpectationFailed, ErrMsg: err.Error(),
			}
	}

	/*
		delete a role
			note: rows with role_id in roles_permissions will be automatically deleted
	 */
	if err := trans.Table(Role{}.TableName()).Where("id=?", r.Id).Delete(&Role{}).Error;
		err != nil {
			return utils.Msg{
				Message: utils.ErrorInternalDbError, Status:  http.StatusExpectationFailed, ErrMsg: err.Error(),
			}
	}

	trans.Commit()
	trans = nil

	return utils.Msg{
		Message: utils.NoErrorFineEverthingOk, Status:  http.StatusOK, ErrMsg: "",
	}
}

/*
	get roles
 */
func (r *Role) Get_roles(offset string) (utils.Msg) {
	var roles = struct {
		Info			[]Role
	}{}
	if err := GetDB().Preload("Permissions").Table(Role{}.TableName()).Offset(offset).Limit(GetLimit).Find(&roles.Info).Error; err != nil {
		var resp = utils.ErrorInternalDbError
		resp["info"] = []map[string]interface{}{}
		return utils.Msg{
			Message: resp, Status:  http.StatusExpectationFailed, ErrMsg: err.Error(),
		}
	}

	var resp = utils.ErrorInternalDbError
	resp["info"] = Struct_to_map(roles)

	return utils.Msg{resp, http.StatusOK, "", ""}
}

func (r *Role) Get_role_info() (utils.Msg) {
	if err := GetDB().Preload("Permissions").Table(r.TableName()).Where("id=?", r.Id).Error;
		err != nil {
			return utils.Msg{Message: utils.ErrorInternalDbError, Status: http.StatusExpectationFailed, ErrMsg: err.Error()}
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(&r)

	return utils.Msg{
		Message: resp, Status:  http.StatusOK, ErrMsg: "",
	}
}
