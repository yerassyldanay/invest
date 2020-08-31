package model

import (
	"invest/utils"
)

/*
	create a category
 */
func(ca *Categor) Create_category() (map[string]interface{}, error) {
	if err := GetDB().Create(ca).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}

	return utils.NoErrorFineEverthingOk, nil
}

/*

 */
func (ca *Categor) Update() (map[string]interface{}, error) {
	if err := GetDB().Table(Categor{}.TableName()).Update("name", ca.Name).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}

	return utils.NoErrorFineEverthingOk, nil
}

/*
	delete category + delete from projects
 */
func (ca *Categor) Delete_category_from_tabe_and_projects() (map[string]interface{}, error) {

	if err := GetDB().Delete(ca, "id = ?", ca.Id).Error;
		err != nil {
			return utils.ErrorInternalDbError, err
	}

	return utils.NoErrorFineEverthingOk, nil
}

/*
	get all categories
 */
func (ca *Categor) Get_all_categors(offset string) (map[string]interface{}, error) {
	var cas = []Categor{}
	if err := GetDB().Table(ca.TableName()).Offset(offset).Limit(GetLimit).Find(&cas).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}

	var carr = []map[string]interface{}{}
	for _, each := range cas {
		carr = append(carr, Struct_to_map(each))
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = carr

	return resp, nil
}

