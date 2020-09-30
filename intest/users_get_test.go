package intest

import (
	"fmt"
	"invest/app"

	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestGetUsersByAdmin(t *testing.T) {
	var fname = "TestGetUsersByAdmin"
	var router = app.Create_new_invest_router()

	var datas = []TestData{
		{
			URL:        "/v1/administrate/user?offset=0&role=manager&role=investor",
			Method:     http.MethodGet,
			RespStatus: 200,
		},
		{
			URL: 		"/v1/administrate/user?offset=0",
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

		r.Header.Add(constants.HeaderContentType, "application/x-www-form-urlencoded")
		r.Header.Set(constants.HeaderAuthorization, constants.AuthorizationAdminToken)

		router.ServeHTTP(resp, r)

		status, _ := strconv.Atoi(resp.Header().Get(constants.HeaderCustomStatus))

		if status != data.RespStatus {
			t.Errorf(fmt.Sprintf("%s 3" + " | %d", fname, resp.Code))
		}
	}
}
