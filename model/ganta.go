package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"invest/utils"
	"time"
)

var GantaDefaultStepHours time.Duration = 3 * 24

/*
	required:
		project id
		trans - transaction
 */
func (p Project) create_default_parents(trans *gorm.DB) (p1, p2 uint64, err error) {
	/*
		create the first parent
	 */
	var ganta = DefaultGantaParent1
	ganta.ProjectId = p.Id
	if err = GetDB().Create(&ganta).Error; err != nil {
		return 0, 0, err
	}
	p1 = ganta.Id

	/*
		create the second parent
	 */
	ganta = DefaultGantaParent2
	ganta.ProjectId = p.Id
	ganta.Status = utils.ProjectStatusInprogress

	if err = GetDB().Create(&ganta).Error; err != nil {
		return 0, 0, err
	}
	p2 = ganta.Id

	return p1, p2, err
}

func (p *Project) create_child_steps_of_the_ganta_table(p1 uint64, steps []Ganta, trans *gorm.DB) (err error) {
	for _, step := range steps {
		/*
			set project_id & parent_id
		 */
		step.GantaParentId = p1
		step.ProjectId = p.Id

		/*
			using transaction create steps
				in case of an error
				we will use rollback
		 */
		if err = trans.Create(&step).Error; err != nil {
			return err
		}
	}

	return nil
}

/*
	CREATE A GANTA TABLE FOR THE PROJECT

	provide:
		project_id
 */
func (p * Project) Create_ganta_table_for_this_project() (*utils.Msg) {
	if p.Id == 0 {
		return &utils.Msg{utils.ErrorInvalidParameters, 400, "", "invalid project id. id = 0"}
	}

	trans := GetDB().Begin()
	defer func() { if trans != nil {trans.Rollback()} }()

	/*
		create parents
	 */
	p1, p2, err := p.create_default_parents(trans)
	if err != nil {
		return &utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	fmt.Println(p1, p2)

	/*
		create child processes
	 */
	if err = p.create_child_steps_of_the_ganta_table(p1, DefaultGantaChildren1, trans); err != nil {
		return &utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	if err = p.create_child_steps_of_the_ganta_table(p2, DefaultGantaChildren2, trans); err != nil {
		return &utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	/*
		in case everything is ok, commit changes to db
	 */
	err = trans.Commit().Error
	if err != nil {
		return &utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	trans = nil
	return &utils.Msg{utils.NoErrorFineEverthingOk, 200, "", ""}
}


