package model

import (
	"fmt"
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

	//var a = templates.Base_message_map_3_changes_made_to_project[utils.KeyEmailSubject]
	//var b = a[lang]
	//fmt.Println(a, b)

	var sm = SendgridMessage{
		Subject:   		fmt.Sprintf(templates.Base_message_map_3_changes_made_to_project[utils.KeyEmailSubject][lang]),
		PlainText: 		fmt.Sprintf(templates.Base_message_map_3_changes_made_to_project[utils.KeyEmailHtml][lang], project_name, who),
		HTML:      		fmt.Sprintf(templates.Base_message_map_3_changes_made_to_project[utils.KeyEmailHtml][lang], project_name, who),
		Created:   		utils.GetCurrentTime(),
	}

	return sm
}
