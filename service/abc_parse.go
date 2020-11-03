package service

import (
	"invest/utils/constants"
	"net/http"
	"time"
)

func (is *InvestService) OnlyParseRequest(r *http.Request) {
	is.Time = time.Now()
	is.TimeUTC = time.Now()
	is.Offset = Get_query_parameter_str(r, "offset", "0")
	is.BasicInfo.UserId = Get_header_parameter(r, constants.KeyId, uint64(0)).(uint64)
	is.BasicInfo.RoleName = Get_header_parameter(r, constants.KeyRoleName, "").(string)
	is.BasicInfo.Lang = Get_header_parameter(r, constants.HeaderContentLanguage, "").(string)
}

