package tests

import (
	"github.com/stretchr/testify/require"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/service"
	"github.com/yerassyldanay/invest/utils/constants"
	"testing"
)

var project = &model.Project{}
var TestDefaultDocumentUri = "/documents/docs/test.pdf"

// status update
func TestUpdateProjectStatus(t *testing.T) {
	var project = HelperGetAnyProject(t)
	err := project.GetAndUpdateStatusOfProject(model.GetDB())

	// check
	require.NoError(t, err)
}

func TestServiceUploadDocumentByInvestor(t *testing.T) {
	// get project
	project := HelperGetAnyProject(t)

	// check: it must be the first step
	err := project.GetAndUpdateStatusOfProject(model.GetDB())

	// check
	require.NoError(t, err)
	require.Equal(t, project.Step, 1)

	// investor uploads all documents
	document := model.Document{}
	documents, err := document.OnlyGetDocumentsByProjectId(project.Id, model.GetDB())

	// check
	require.NoError(t, err)
	require.NotZero(t, len(documents))

	//
	var skippedDocId uint64 = 8
	var first = true
	for _, doc := range documents {
		switch {
		case doc.Step == 2:
			// pass
		case !first && doc.Responsible == constants.RoleInvestor:
			// pass
			doc.Uri = TestDefaultDocumentUri
			err := doc.OnlySave(model.GetDB())

			// intermediate check
			require.NoError(t, err)
		case doc.Responsible == constants.RoleInvestor:
			first = false
			skippedDocId = doc.Id
		}
	}

	// test service
	is := service.InvestService{
		BasicInfo: service.BasicInfo{
			UserId: 3, // investor
		},
	}

	// create a document (to update)
	document = model.Document{
		Id:        	skippedDocId,
		ProjectId: 	project.Id,
		Uri: 		TestDefaultDocumentUri,
	}

	// upload a file to a document
	msg := is.Upload_documents_to_project(&document)

	// check
	require.Zero(t, msg.ErrMsg)
}

//// check status
//func TestModelProjectStatusThatPendingAdmin(t *testing.T) {
//	project := GetAnyProject(t)
//	if err := project.GetAndUpdateStatusOfProject(model.GetDB()); err != nil {
//		t.Error(err)
//	}
//
//	// check status
//	if project.Status != constants.ProjectStatusPendingAdmin {
//		t.Error("expected status to be pending admin, but got ", project.Status)
//	}
//
//	// check step
//	if project.Step != 1 {
//		t.Error("step is not 1")
//	}
//}
//
//// assign user by admin
//func TestServiceAssignUserToProject(t *testing.T) {
//	project := GetAnyProject(t)
//	pu := model.ProjectsUsers{
//		ProjectId: project.Id,
//		UserId:    2,
//	}
//
//	// headers
//	is := service.InvestService{
//		BasicInfo: service.BasicInfo{
//			UserId: 1, // admin
//		},
//	}
//
//	// logic
//	msg := is.Assign_user_to_project(pu)
//	if msg.IsThereAnError() {
//		t.Error(msg.ErrMsg)
//	}
//
//	//msg = is.Assign_user_to_project(pu)
//	//if !msg.IsThereAnError() {
//	//	t.Error("expected err, but got nil")
//	//}
//}
//
//// shift status
//func TestServiceChangeProjectStatus(t *testing.T) {
//	project := model.Project{Id: 1}
//	ganta := model.Ganta{ProjectId: project.Id}
//
//	// shift status
//	if err := ganta.OnlyChangeStatusToDoneAndUpdateDeadlineById(model.GetDB()); err != nil {
//		t.Error(err)
//	}
//
//	_ = project.OnlyGetById(model.GetDB())
//	if project.Status != constants.ProjectStatusPendingManager {
//		t.Error("expected " + constants.ProjectStatusPendingManager + " but got " + project.Status)
//	}
//}
//
//// status must be pending_manager
//// check status
//func TestServiceCommentOnProject(t *testing.T) {
//	project := GetAnyProject(t)
//	if err := project.GetAndUpdateStatusOfProject(model.GetDB()); err != nil {
//		t.Error(err)
//	}
//
//	// get all documents
//	var document = model.Document{ProjectId: project.Id}
//	documents, err := document.OnlyGetDocumentsByStepsAndProjectId(project.Id, []interface{}{1}, model.GetDB())
//	if err != nil {
//		t.Error(err)
//	}
//
//	// prepare documents
//	documents[0].Status = constants.ProjectStatusReconsider
//
//	// comment
//	spkComment := model.SpkComment{
//		Comment:   model.Comment{
//			Body:      "test comment",
//			UserId:    2, //manager
//			ProjectId: project.Id,
//			Status:    constants.ProjectStatusReconsider,
//		},
//		Documents: documents,
//	}
//
//	// logic
//	is := service.InvestService{}
//	msg := is.Comment_on_project_documents(spkComment)
//	if msg.IsThereAnError() {
//		t.Error(msg.ErrMsg)
//	}
//
//	// check whether document status changed
//	document = model.Document{Id: documents[0].Id}
//	err = document.OnlyGetDocumentById(model.GetDB())
//
//	if err != nil {
//		t.Error(err)
//	} else if document.Status != constants.ProjectStatusReconsider {
//		t.Error("expected status to be reconsider, but it is " + document.Status)
//	}
//
//	// the status of the project also must chnage
//	_ = project.OnlyGetById(model.GetDB())
//	if project.Status != constants.ProjectStatusPendingInvestor {
//		t.Error("expected status to be pending investor, but got " + project.Status)
//	}
//}