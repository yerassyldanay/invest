package model

import (
	"errors"
	"invest/utils"
)

func (p * Project) Get_project_with_documents() (map[string]interface{}, error) {
	if p.Id == 0 {
		return utils.ErrorInvalidParameters, errors.New("invalid project id. project get info")
	}

	if err := GetDB().Table(Document{}.TableName()).Where("id=?", p.Id).First(p).Error; err != nil {
		return utils.ErrorInternalDbError, err
	}

	var docs = []Document{}
	rows, err := GetDB().Table(Document{}.TableName()).Where("project_id=?", p.Id).Rows()
	if err == nil {
		var tdoc = Document{}
		for rows.Next() {
			if err := GetDB().ScanRows(rows, &tdoc); err != nil {
				continue
			}

			docs = append(docs, tdoc)
		}
		defer rows.Close()
	}

	p.Documents = docs

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(*p)

	return resp, nil
}

