package tests

import (
	"invest/service"
	"testing"
)

func TestProjectGetById(t *testing.T) {
	var project = getAnyProject(t)

	// check model
	err := project.Get_this_project_by_project_id()

	switch {
	case err != nil:
		t.Error("got ", err)
	case project.Organization.Id < 1:
		// check org
		t.Error("could not get organization")
	case project.CurrentStep.Id < 1:
		// check gantt step
		t.Error("could not get current step")
	case len(project.Users) < 1:
		// check
		t.Error("got no users")
	default:
		//fmt.Println("project get by id: ", project.Name)
	}

	// go through users & check
	for _, user := range project.Users {
		switch {
		case user.Password != "":
			t.Error("you are sending password of a user")
		case user.Email.Address == "":
			t.Error("you are not providing info on email of the user")
		case user.Role.Name == "":
			t.Error("you are not providing info on role of a user")
		}
	}

	// service
	is := service.InvestService{BasicInfo: service.BasicInfo{UserId: 1}}
	msg := is.Project_get_by_id(project.Id)

	// check service
	if msg.IsThereAnError() {
		t.Error("expected no error, but got " + msg.ErrMsg)
	}
}

