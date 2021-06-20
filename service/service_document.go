package service

import (
	"fmt"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/utils/errormsg"
	"github.com/yerassyldanay/invest/utils/message"
	"os"
	"path/filepath"
	"strconv"
)

// get documents by project id
func (is *InvestService) DocumentGetByProjectId(project_id uint64, stepsRaw []string) message.Msg {
	var document = model.Document{}

	// get documents
	steps := []interface{}{}
	for _, step := range stepsRaw {
		n, _ := strconv.Atoi(step)
		if n < 1 || n > 2 {
			return model.ReturnInvalidParameters("invalid step number")
		}

		steps = append(steps, n)
	}

	if len(steps) == 0 {
		var ganta = model.Ganta{ProjectId: project_id}
		_ = ganta.OnlyGetCurrentStepByProjectId(model.GetDB())
		steps = []interface{}{ganta.Step}
	}

	// get document
	documents, err := document.OnlyGetDocumentsByStepsAndProjectId(project_id, steps, model.GetDB())
	if err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	documentsMap := []map[string]interface{}{}
	for _, document := range documents {
		documentsMap = append(documentsMap, model.Struct_to_map(document))
	}

	var resp = errormsg.NoErrorFineEverthingOk
	resp["info"] = documentsMap

	return model.ReturnNoErrorWithResponseMessage(resp)
}

// remove a file from storage & empty uri
func (is *InvestService) DocumentRemoveDocumentFromProject(document_id uint64) message.Msg {
	var trans = model.GetDB().Begin()
	defer func() {
		if trans != nil {
			trans.Rollback()
		}
	}()

	// get document to delete file from storage
	var document = model.Document{Id: document_id}
	if err := document.OnlyGetDocumentById(trans); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// set uri to ''
	var path = document.Uri
	if err := document.OnlyEmptyUriById(trans); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// get an absolute pather & delete a file from storage
	path, _ = filepath.Abs("./" + path)
	fmt.Println("remove file pather: " + path)
	if err := os.Remove(path); err != nil {
		fmt.Println(err.Error())
	}

	// commit changes
	if err := trans.Commit().Error; err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	trans = nil
	return model.ReturnNoError()
}
