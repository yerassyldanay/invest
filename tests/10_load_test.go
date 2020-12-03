package tests

import (
	"bytes"
	"fmt"
	"invest/app"
	"invest/utils/helper"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestLoadTestOnServiceCreateProject(t *testing.T) {

	NumberOfTimeToTest := 2
	NumberOfConcurrentRequest := 5

	sessionTokenForAdmin := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozLCJyb2xlX2lkIjowLCJyb2xlX25hbWUiOiJpbnZlc3RvciIsImV4cCI6IjIwMjAtMTEtMDhUMTc6MTg6MzQuNTQzMTcxOTE1KzA2OjAwIn0.UxG-6Rww3-en2ou-gNSGrcPZbwhQza_rfi3zCvfg8hM"

	requestBody := `{
		"project": {
			"name": "Тестовый проект - %s",
			"description": "Описание проекта пишете сюда",
			"info_sent": {
				"add-info": "доп инфо"
			},
			"employee_count": 10,
			"email": "invest@spk.com",
			"phone_number": "+77171204560",
			"organization": {
				"bin": "190940011748"
			},
			"categors": [
				{
					"id": 2
				}
			],
			"offered_by_position": "инициатор проекта",
			"land_plot_from": "spk",
			"land_area": 1000,
			"land_address": "город, название улицы, дом"
		},
		"cost": {
		  "building_repair_investor": 1000,
		  "building_repair_involved": 2000,
		  "technology_equipment_investor": 3000,
		  "technology_equipment_involved": 4000,
		  "working_capital_investor": 5000,
		  "working_capital_involved": 6000,
		  "other_cost_investor": 7000,
		  "other_cost_involved": 8000,
		  "share_in_project_investor": 9000,
		  "share_in_project_involved": 10000
		},
		"finance": {
			"total_income": 10000,
			"total_production": 20000,
			"production_cost": 30000,
			"operational_profit": 40000,
			"settlement_obligations": 50000,
			"other_cost": 60000,
			"pure_profit": 70000,
			"taxes": 80000
		}
	}`

	router := app.Create_new_invest_router()
	ts := httptest.NewServer(router)

	for i := 0; i < NumberOfTimeToTest; i++ {
		wg := sync.WaitGroup{}
		for iroutine := 0; iroutine < NumberOfConcurrentRequest; iroutine++ {
			iroutine := iroutine

			wg.Add(1)
			go func(j int, gRequestBody string, gwg *sync.WaitGroup) {
				defer wg.Done()

				gRequestBody = fmt.Sprintf(gRequestBody, helper.Generate_Random_String(20))

				inBytes := []byte(gRequestBody)

				request, err := http.NewRequest(http.MethodPost, ts.URL + "/v1/project", bytes.NewBuffer(inBytes))
				if err != nil {
					t.Log("create new request: ", err)
					return
				}
				request.Header.Set("Authorization", sessionTokenForAdmin)

				response := httptest.NewRecorder()
				router.ServeHTTP(response, request)

				if response == nil {
					//fmt.Println(i, " failed")
					return
				}

				//fmt.Println(i, response.Code)
				////fmt.Println(response.Body.String())
				//fmt.Println()

			}(iroutine + (i * 10) + 1, requestBody, &wg)

		}

		wg.Wait()
	}
}
