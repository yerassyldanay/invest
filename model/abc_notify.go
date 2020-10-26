package model

import (
	"context"
	"fmt"
	"invest/utils"
	"time"
)

// notify users about the deadline
func onlyNotifyAboutGanttDeadlineHelper() {
	type newStruct struct {
		Id uint64 `json:"id"`
	}

	var projectIds = []newStruct{}

	currTime := utils.GetCurrentTime()
	day := time.Hour * 24

	// get projects which has a deadline coming with 3-4 days
	if err := GetDB().Raw("select distinct project_id as id from gantas " +
		"where is_done = false and ((deadline between ? and ?) or (deadline between ? and ?)) " +
		" and status not in (?);", currTime.Add(day * 3), currTime.Add(day * 4),
		currTime.Add(day * 7), currTime.Add(day * 8), []string{
			utils.ProjectStatusAgreement, utils.ProjectStatusReject,
	}).Scan(&projectIds).Error; err != nil {
		fmt.Println("could not send notifications")
		return
	}

	for _, project := range projectIds {
		// create a notification struct
		ngd := NotifyGantaDeadline{
			ProjectId: project.Id,
			Project: Project{
				Id: project.Id,
			},
		}

		fmt.Println("sending notification to project with id: ", project.Id)

		// this will handle everything else
		GetMailerQueue().NotificationChannel <- &ngd
	}
}

func onlyNotifyAboutDocumentDeadline() {
	// get list of documents (only ids & project id)
	document := Document{}
	documents, err := document.OnlyGetEmptyDocumentsWithComingDeadline()

	switch {
	case err != nil:
		fmt.Println("[ERROR] notify document deadline. err: ", err)
		return
	case len(documents) < 1:
		fmt.Println("[WARN] there is no any deadline in documents")
		return
	}

	for _, document = range documents {
		ndd := NotifyDocDeadline{
			DocumentId: document.Id,
		}

		GetMailerQueue().NotificationChannel <- &ndd
	}
}

// this will be used as a scheduler
func OnlyNotifyAboutGantaDeadline(cnt context.Context) {
	for {
		select {
		case <- cnt.Done():
			// quit the function
			return
		case <- time.Tick(time.Hour * 24):
			// once in 24 hours
			// notify all users about deadline
			onlyNotifyAboutGanttDeadlineHelper()
		}
	}
}

func OnlyNotifyAboutDocumentDeadline(cnt context.Context) {
	for {
		select {
		case <- cnt.Done():
			// quit the function
			return
		case <- time.Tick(time.Hour * 24):
			// once in 24 hours
			onlyNotifyAboutDocumentDeadline()
		}
	}
}

