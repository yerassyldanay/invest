package intest

import (
	"invest/app"
	_ "invest/model"
)

var InvestTestRouter = app.Create_new_invest_router()

//func TestSignIn(t *testing.T) {
//	var fname = "TestSignIn"
//	router := InvestTestRouter
//
//	var datas = []TestData{
//		{
//			Body: map[string]interface{}{
//				"username": "investor",
//				"password": "KeRXaTaq5Ce8ULO",
//				"role": "investor",
//				"key": "username",
//			},
//			RespStatus: 200,
//		},
//		{
//			Body: map[string]interface{}{
//				"username": "investor",
//				"password": "invalidpassword",
//				"role": "investor",
//				"key": "username",
//			},
//			RespStatus: 400,
//		},
//	}
//
//	//cln := http.Client{}
//
//	resp := httptest.NewRecorder()
//	for _, data := range datas {
//
//		var datastr = data.Convert_body_to_string()
//		//fmt.Println(datastr)
//
//		r, err := http.NewRequest(http.MethodPost, "/v1/all/signin", bytes.NewBuffer([]byte(datastr)))
//		if err != nil {
//			t.Errorf(fname + " 2" + " | " + err.Error())
//			continue
//		}
//		defer r.Body.Close()
//
//		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
//		router.ServeHTTP(resp, r)
//
//		status, _ := strconv.Atoi(resp.Header().Get(utils.HeaderCustomStatus))
//
//		if status != data.RespStatus {
//			t.Errorf(fmt.Sprintf("%s 3" + " | %d", fname, resp.Code))
//		}
//	}
//}

