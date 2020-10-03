package intest

import (
	"fmt"
	"invest/app"
	"invest/utils"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

/*
var org = model.Organization{Bin: "190441005334"}
msg := org.Create_or_get_organization_from_db_by_bin(model.GetDB())

fmt.Println(msg)
 */

func TestGetOrganization(t *testing.T) {
	var fname = "TestGetOrganization"
	var router = app.Create_new_invest_router()

	var datas = []TestData{
		{
			URL:        "/v1/organization?bin=190840000779",
			Method:     http.MethodGet,
			RespStatus: 200,
		},
		{
			URL: 		"/v1/organization?bin=190840026603",
			Method: 	http.MethodGet,
			RespStatus: 200,
		},
	}

	var resp = httptest.NewRecorder()

	for _, data := range datas {
		r, err := http.NewRequest(data.Method, data.URL, nil)
		if err != nil {
			t.Errorf(fname + " 2" + " | " + err.Error())
			continue
		}
		//defer func() {if r != nil {r.Body.Close()} }()

		r.Header.Add(utils.HeaderContentType, "application/x-www-form-urlencoded")
		r.Header.Set(utils.HeaderAuthorization, utils.AuthorizationAdminToken)

		router.ServeHTTP(resp, r)

		status, _ := strconv.Atoi(resp.Header().Get(utils.HeaderCustomStatus))

		if status != data.RespStatus {
			t.Errorf(fmt.Sprintf("%s 3" + " | %d", fname, resp.Code))
		}
	}
}