package intest

import (
	"fmt"
	"invest/model"
	"testing"
)

func TestDocumentGet(t *testing.T) {
	var project = getAnyProject(t)
	var document = model.Document{ProjectId: project.Id, Step: project.Step}

	documents, err := document.OnlyGetDocumentsByStepsAndProjectId(project.Id, []interface{}{1}, model.GetDB())
	switch {
	case err != nil:
		t.Error("could not get a list of documents. err: ", err)
	case len(documents) < 1:
		t.Error("len of docs is 0")
	default:
		fmt.Println("step1: got ", len(documents), " number of documents")
	}

	step1 := len(documents)

	documents, err = document.OnlyGetDocumentsByStepsAndProjectId(project.Id, []interface{}{2}, model.GetDB())
	switch {
	case err != nil:
		t.Error("expected no error, but got ", err)
	case len(documents) < 1:
		t.Error("len of docs is 0")
	default:
		fmt.Println("step2: got ", len(documents), " number of documents")
	}

	step2 := len(documents)

	documents, err = document.OnlyGetDocumentsByProjectId(project.Id, model.GetDB())
	switch {
	case err != nil:
		t.Error(err)
	case len(documents) != step1 + step2:
		t.Error("get documents separately works different from getting documents together")
	default:
		//fmt.Println()
	}
}

