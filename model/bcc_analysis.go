package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/tealeg/xlsx"
	"github.com/yerassyldanay/invest/utils/constants"
	"os"
	"strconv"
	"time"
)

type Analysis struct {
	Steps						[]int				`json:"steps"`
	Statuses					[]string			`json:"statuses"`
	Categors					[]uint64			`json:"categors"`
	Start						int64				`json:"start"`
	StartDate					time.Time			`json:"start_date"`
	End							int64				`json:"end"`
	EndDate						time.Time			`json:"end_date"`
	
	WriteToFile					bool				`json:"-" gorm:"-"`
	ProjectExtendedList					[]ProjectExtended		`json:"-" gorm:"-"`
}

// errors
var errorAnalysisFileExists = errors.New("this file already exists. cannot create one")

func (a *Analysis) OnlyGetProjectsByExtendedFields(tx *gorm.DB) ([]Project, error) {
	projects := []Project{}
	main_query := `select distinct p.* from projects p
	join projects_categors pc on p.id = pc.project_id
	where step in (?) and status in (?) and created > ? and created < ? and pc.categor_id in (?);`

	err := tx.Raw(main_query, a.Steps, a.Statuses, a.StartDate, a.EndDate, a.Categors).Scan(&projects).Error
	return projects, err
}

func (a *Analysis) OnlyWriteDataToFile(absFilepath string, lang string) (error) {
	// check whether this file exists
	// error is not nil = exists
	_, err := os.Stat(absFilepath)
	if err == nil {
		return errorAnalysisFileExists
	}

	// create a file
	excel := xlsx.NewFile()
	sheet, err := excel.AddSheet("sheet-1")
	if err != nil {
		return err
	}

	// prepare names of rows
	var headers = []string{
		"Название",
		"Организация",
		"Создан",
		"Категория",
		"Кол. сотрудников",
		"Налоги",
		"Площадь",
		"Инвестиция",
		"Стадия",
		"Статус",
	}

	// set headers
	// first row
	for i, header := range headers {
		cell := sheet.Cell(0, i)
		cell.Value = header
	}

	// write down info on projects
	for rowIndex, projectEx := range a.ProjectExtendedList {

		projCategory := ""
		if len(projectEx.Categors) > 0 {
			projCategory = projectEx.Categors[0].Rus
		}

		projectEx.Status = a.ConvertStatus(projectEx.Status, lang)
		colValuesToInsert := []string{
			projectEx.Name,			// name
			projectEx.Organization.Name,		// org
			projectEx.Created.Format(time.RFC3339),
			projCategory,
			strconv.Itoa(int(projectEx.EmployeeCount)),
			strconv.Itoa(projectEx.Finance.Taxes),
			strconv.Itoa(projectEx.LandArea),
			strconv.Itoa(projectEx.Cost.WorkingCapitalInvestor),
			strconv.Itoa(projectEx.Step),
			projectEx.Status,
		}

		for colIndex, colValue := range colValuesToInsert {

			// write to the file
			// (row, col) = value
			cell := sheet.Cell(rowIndex + 1, colIndex)
			cell.Value = colValue

		}
	}

	err = excel.Save(absFilepath)
	if err != nil {
		return err
	}

	return nil
}

func (a Analysis) ConvertStatus(status string, lang string) (string) {
	vmap, ok := constants.MapProjectStatusFirstStatusThenLang[status]
	if !ok {
		return ""
	}

	value, ok := vmap[lang]
	if !ok {
		return ""
	}

	return value
}
