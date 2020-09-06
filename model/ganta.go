package model

import (
	"bytes"
	"fmt"
	"invest/utils"
	"time"
)

var GantaDefaultStepHours time.Duration = 3 * 24

func (p * Project) Create_ganta_table_for_this_project() (*utils.Msg) {
	if p.Id == 0 {
		return &utils.Msg{utils.ErrorInvalidParameters, 400, "", "invalid project id. id = 0"}
	}

	/*
		create ganta steps & set project id to them
	 */
	var gantas = GantaDefaultSteps
	for i, _ := range gantas {
		gantas[i].ProjectId = p.Id
	}

	var main_query bytes.Buffer
	main_query.WriteString("insert into gantas(project_id, kaz, rus, eng, start) values ")

	for i, ganta := range gantas {
		main_query.WriteString(
			fmt.Sprintf("(%d, '%s', '%s', '%s', '%s')",
				ganta.ProjectId,
				ganta.Start.Format("2006-01-02 15:04:05")),
		)

		if i != len(gantas) - 1 {
			main_query.WriteString(" , ")
		}
	}

	main_query.WriteString(" ; ")

	var a = main_query.String()
	fmt.Println(a)

	if n := GetDB().Exec(main_query.String()).RowsAffected; n == 0 {
		return &utils.Msg{utils.ErrorInternalDbError, 417, "", "no ros are affected. create ganta for the project"}
	}

	return &utils.Msg{utils.NoErrorFineEverthingOk, 200, "", ""}
}


