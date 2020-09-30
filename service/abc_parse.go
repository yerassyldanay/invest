package service

import (
	"invest/utils"
	"net/http"
)

func (is *InvestService) OnlyParseRequest(r *http.Request) {
	is.Offset = Get_query_parameter_str(r, "offset", "0")
	is.BasicInfo.UserId = Get_header_parameter(r, utils.KeyId, uint64(0)).(uint64)
	is.BasicInfo.RoleName = Get_header_parameter(r, utils.KeyRoleName, "").(string)
	is.BasicInfo.Lang = Get_header_parameter(r, utils.HeaderContentLanguage, "").(string)
}
