package test

import (
	//"encoding/json"
	"fmt"
	"invest/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProjectGetAll(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Add(utils.KeyRoleId, "3")
		r.Header.Add(utils.KeyId, "3")
	}))
	defer ts.Close()

	req := httptest.NewRequest("GET", ts.URL + "/v1/projects_see_all/project?offset=0", nil)
	cln := http.Client{}

	resp, err := cln.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}

	//var temp = make(map[string]interface{})
	//if err := json.NewDecoder(resp.Body).Decode(&temp); err != nil {
	//	t.Error(err)
	//}
	//defer resp.Body.Close()
	//t.Log()

	fmt.Println(resp.Status, resp.Header)
}

