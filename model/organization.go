package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"invest/utils"

	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

/*
	this will provide info by bin number
		to fill fields automatically
*/
func (o *Organization) Make_request_to_get_organization() (*Organization, error) {
	if o.Lang == "" {
		o.Lang = utils.DefaultContentLanguage
	}

	var bin = o.Bin

	var url = fmt.Sprintf("https://stat.gov.kz/api/juridicalusr/counter/gov/?bin=%s&lang=%s", o.Bin, o.Lang)
	resp, err := http.Get(url)
	if err != nil {
		return o, err
	}
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return o, readErr
	}

	var temp = make(map[string]interface{})
	err = json.Unmarshal(body, &temp)

	/*
		marshal then unmarshal in order to convert interface to map
	 */
	b, _ := json.Marshal(temp["obj"])

	temp = map[string]interface{}{}
	_ = json.Unmarshal(b, &temp)

	var registerDate = temp["registerDate"]
	var katoAddress = temp["katoAddress"]
	var name = temp["name"]
	var fio = temp["fio"]

	if utils.Is_any_of_these_nil(registerDate, katoAddress, name, fio) {
		return o, errors.New("there is a nil element inside the map. get info by bin")
	}

	ti, err := time.Parse("2006-01-02T15:04:05", strings.Replace(registerDate.(string), ".000+0000", "", 100))
	if err != nil {
		return o, errors.New("could not parse date. get info by bin")
	}

	o.Name = strings.Replace(name.(string), "\\\"", "", 50)
	o.Fio = strings.Replace(fio.(string), "\\\"", "", 50)
	o.Regdate = ti.UTC()
	o.Address = strings.Replace(katoAddress.(string), "\\\"", "", 50)

	o.Bin = bin

	return o, nil
}

/*
	create a company
		expects that all fields are filled
*/
func (o *Organization) Create_or_get_organization_from_db_by_bin(tx *gorm.DB) (utils.Msg) {
	if o.Lang == "" {
		o.Lang = utils.DefaultContentLanguage
	}

	var bin = o.Bin
	var err error

	/*
		check whether such organization is already on db
	 */
	if err = o.OnlyGetByBinAndLang(tx); err == nil {
		var resp = utils.NoErrorFineEverthingOk
		resp["info"] =  Struct_to_map(*o)
		return ReturnNoErrorWithResponseMessage(resp)
	} else if err != gorm.ErrRecordNotFound {
		return ReturnInternalDbError(err.Error())
	}

	/*
		obtain info on an organization
	*/
	o, err = o.Make_request_to_get_organization()
	if err != nil {
		return utils.Msg{utils.ErrorExternalServiceErrorNoOrganizationInfo, http.StatusServiceUnavailable, "", err.Error(),}
	}

	/*
		store data on db
	*/
	o.Bin = bin
	if err := o.OnlyCreate(tx); err != nil {
		return ReturnInternalDbError(err.Error())
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] =  Struct_to_map(*o)

	return ReturnNoError()
}

/*
	update company info by admin
*/
func (o *Organization) Update_organization_info(tx *gorm.DB) (utils.Msg) {
	if o.Lang == "" {
		o.Lang = "kaz"
	}

	if err := tx.Table(Organization{}.TableName()).Where("bin=? and lang=?", o.Bin, o.Lang).
		Updates(map[string]interface{}{
			"name": o.Name,
			"regdate": o.Regdate,
			"fio": o.Fio,
			"address": o.Address,
		}).Error;
			err != nil {
				return ReturnInvalidParameters(err.Error())
	}

	return ReturnNoError()
}

