package model

import (
	"bytes"
	"errors"
	"fmt"
	"invest/utils"
	"time"
)

var GantaDefaultStepHours time.Duration = 3 * 24

func (p * Project) Create_ganta_table_for_this_project() (map[string]interface{}, error) {
	if p.Id == 0 {
		return utils.ErrorInvalidParameters, errors.New("invalid project id. id = 0")
	}

	var gantas = []Ganta{
		{
			ProjectId: p.Id,
			Kaz:       "Бірінші қадам",
			Rus:       "Первый шаг",
			Eng:       "First step",
			Start:     p.Created,
		},
		{
			ProjectId: p.Id,
			Kaz:       "Екінші қадам",
			Rus:       "Второй шаг",
			Eng:       "Second step",
			Start:     p.Created.Add(time.Hour * GantaDefaultStepHours),
		},
		{
			ProjectId: p.Id,
			Kaz:       "Үшінші қадам",
			Rus:       "Третий шаг",
			Eng:       "Third step",
			Start:     p.Created.Add(time.Hour * GantaDefaultStepHours * 2),
		},
	}

	var main_query bytes.Buffer
	main_query.WriteString("insert into gantas(project_id, kaz, rus, eng, start) values ")

	for i, ganta := range gantas {
		main_query.WriteString(
			fmt.Sprintf("(%d, '%s', '%s', '%s', '%s')",
				ganta.ProjectId,
				ganta.Kaz,
				ganta.Rus,
				ganta.Eng,
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
		return utils.ErrorInternalDbError, errors.New("no ros are affected. create ganta for the project")
	}

	return utils.NoErrorFineEverthingOk, nil
}


