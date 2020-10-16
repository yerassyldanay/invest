package service

import (
	"invest/utils"
	"net/http"
	"time"
)

func (is *InvestService) OnlyParseRequest(r *http.Request) {
	is.Time = time.Now()
	is.TimeUTC = time.Now()
	is.Offset = Get_query_parameter_str(r, "offset", "0")
	is.BasicInfo.UserId = Get_header_parameter(r, utils.KeyId, uint64(0)).(uint64)
	is.BasicInfo.RoleName = Get_header_parameter(r, utils.KeyRoleName, "").(string)
	is.BasicInfo.Lang = Get_header_parameter(r, utils.HeaderContentLanguage, "").(string)
}

