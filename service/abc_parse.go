package service

import (
	"encoding/json"
	"fmt"
	"github.com/yerassyldanay/invest/utils/constants"
	"net/http"
	"time"
)

func (is *InvestService) OnlyParseRequest(r *http.Request) {
	is.Time = time.Now()
	is.TimeUTC = time.Now()
	is.Offset = GetQueryParameterStr(r, "offset", "0")
	is.BasicInfo.UserId = Get_header_parameter(r, constants.KeyId, uint64(0)).(uint64)
	is.BasicInfo.RoleName = Get_header_parameter(r, constants.KeyRoleName, "").(string)
	is.BasicInfo.Lang = Get_header_parameter(r, constants.HeaderContentLanguage, "").(string)
}

func HelperPrint(any interface{}) {
	b, err := json.MarshalIndent(any, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
