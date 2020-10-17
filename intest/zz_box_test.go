package intest

import (
	"invest/model"
	"invest/service"
	"invest/utils"
	"testing"
	"time"
)

var document = model.Document{
	Eng:          "Тестовый документ",
	Uri:          "",
	SetDeadline:  utils.GetCurrentTime().Add(time.Hour * 24 * 3).Unix(),
	Step:         0,
	ProjectId:    0,
	Responsible:  utils.RoleSpk,
}

func TestServiceGanttCreateBox(t *testing.T) {
	// get project
	var project = getAnyProject(t)
	document.Step = project.Step
	document.ProjectId = project.Id

	//
	is := service.InvestService{
		BasicInfo: service.BasicInfo{
			UserId: 2, // manager
		},
	}

	// logic
	msg := is.Add_box_to_upload_document(document)
	if msg.IsThereAnError() {
		t.Error(msg.ErrMsg)
	}
}

// test document box
func TestModelGanttDocumentBox(t *testing.T) {
	var project = getAnyProject(t)
	if project.Id == 0 {
		t.Error("could not get project")
	}

	// set values
	document.ProjectId = project.Id
	document.Step = project.Step
	document.IsAdditional = true
	document.Deadline = utils.GetCurrentTime().Add(time.Hour * 24 * 3)

	// validate document info
	if err := document.Validate(); err != nil {
		t.Error(err)
	}

	switch {
	case document.Eng == "":
		t.Error("invalid document name. eng")
	case document.Rus == "":
		t.Error("invalid doc name. rus")
	case document.Kaz == "":
		t.Error("invalid doc name. kaz")
	case document.Uri != "":
		t.Error("uri is set already")
	}

	// create
	if err := document.OnlyCreate(model.GetDB()); err != nil {
		t.Error(err)
	}

	switch {
	case document.Id < 0:
		t.Error("id is 0")
	case document.Uri != "":
		t.Error("uri is already set")
	}
}
