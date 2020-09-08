package model

import (
	"invest/templates"
	"invest/utils"
)

/*
	prepare notify users message
*/
func (t *Template) Template_prepare_notify_users_about_changes_in_project(lang string, project_name string, who string) (SendgridMessage) {
	var langs = []interface{}{"kaz", "rus"}
	if yes := utils.Does_a_slice_contain_element(langs, lang); !yes {
		lang = "eng"
	}

	var sm = SendgridMessage{
		Subject:   		templates.Base_message_map_3_changes_made_to_project[utils.KeyEmailSubject][lang],
		PlainText: 		templates.Base_message_map_3_changes_made_to_project[utils.KeyEmailHtml][lang],
		HTML:      		templates.Base_message_map_3_changes_made_to_project[utils.KeyEmailHtml][lang],
		Created:   		utils.GetCurrentTime(),
	}

	return sm
}
