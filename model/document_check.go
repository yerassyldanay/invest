package model

/*
	check whether this user is investor, who created this project
 */
func (d *Document) Is_it_investor() bool {
	var project = Project{}
	if err := GetDB().First(&project, "offered_by_id = ?", d.ChangesMadeById).Error;
		err != nil {
			return false
	}

	return project.Id == d.ProjectId
}
