package intest

import (
	"invest/model"
	"testing"
)

var testProjectId = uint64(1000)

func TestAssignExpertsToProject(t *testing.T) {
	pu := model.ProjectsUsers{}
	if err := pu.OnlyAssignExpertsToProject(testProjectId, model.GetDB()); err != nil {
		t.Error(err)
		return
	}
}

func TestAssignRemoveRelation(t *testing.T) {
	pu := model.ProjectsUsers{ProjectId: testProjectId}
	if err := pu.OnlyDeleteByProjectId(testProjectId, model.GetDB()); err != nil {
		t.Error(err)
		return
	}
}

