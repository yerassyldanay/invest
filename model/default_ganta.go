package model

import (
	"github.com/jinzhu/gorm"
	"invest/utils"

	"time"
)

func (p *Project) OnlyGoCreate(ganta *Ganta, tx *gorm.DB) error {
	//defer wg.Done()

	return tx.Create(ganta).Error
	//if err != nil {
	//	select {
	//	case errorChan <- err:
	//		// passes error to the channel
	//	default:
	//		// in case it is full, passes through
	//	}
	//}
}

/*
	provide ganta steps, which are ready
 */
func (p *Project) Create_ganta_parent_with_its_children(start_date time.Time, ganta_parent_steps []Ganta, trans *gorm.DB) (from_date time.Time, err error) {
	var days_to_add = time.Duration(0)

	/*
		prepare parent ganta step
	 */
	var duration time.Duration
	for i, _ := range ganta_parent_steps {
		ganta_parent_steps[i].ProjectId = p.Id

		// how many days
		duration = ganta_parent_steps[i].DurationInDays

		// start <-> deadline (end)
		ganta_parent_steps[i].StartDate = start_date.Add(time.Hour * 24 * days_to_add)
		ganta_parent_steps[i].Deadline = ganta_parent_steps[i].StartDate.Add(time.Hour * 24 * duration)

		// for next step must start after this step
		days_to_add += duration

		// additional gantt steps will be added
		ganta_parent_steps[i].IsAdditional = false

		// prepare child gantt steps
		for j, _ := range ganta_parent_steps[i].GantaChildren {
			ganta_parent_steps[i].GantaChildren[j].ProjectId = p.Id

			ganta_parent_steps[i].GantaChildren[j].StartDate = ganta_parent_steps[i].StartDate
			ganta_parent_steps[i].GantaChildren[j].Deadline = ganta_parent_steps[i].Deadline

			ganta_parent_steps[i].GantaChildren[j].DurationInDays = 3
			ganta_parent_steps[i].GantaChildren[j].IsAdditional = false
		}
	}

	//var errorChan = make(chan error, 1)

	/*
		cannot use goroutines with transaction
			because if it is blocked one goroutine, at the level of database,
			transaction will be rolled back
	 */
	//var wg = sync.WaitGroup{}
	for _, parent_step := range ganta_parent_steps {
		parent_step := parent_step

		/*
			if an error occurs then return
		*/
		if err := parent_step.OnlyCreate(trans); err != nil {
			return start_date, err
		}

		for i, _ := range parent_step.GantaChildren {
			//wg.Add(1)
			//p.OnlyGoCreate(&parent_step.GantaChildren[i], trans)
			child := parent_step.GantaChildren[i]
			child.GantaParentId = parent_step.Id
			err := child.OnlyCreate(trans)
			if err != nil {
				return start_date, err
			}
		}
	}
	//wg.Wait()

	//select {
	//case err = <- errorChan:
	//	return start_date, err
	//default:
	//	break
	//}

	start_date = start_date.Add(days_to_add * time.Hour * 24)
	return start_date, nil
}

/*
	CREATE A GANTA TABLE FOR THE PROJECT

	provide:
		project_id
 */
func (p *Project) Create_ganta_table_for_this_project(trans *gorm.DB) (utils.Msg) {

	_ = Update_sequence_id_thus_avoid_duplicate_primary_key_error(trans, "gantas")

	/*
		in case everything is ok, commit changes to db
	 */
	var start_date = utils.GetCurrentTime()
	start_date, err := p.Create_ganta_parent_with_its_children(start_date, DefaultGantaParentsOfStep1, trans)
	if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	_, err = p.Create_ganta_parent_with_its_children(start_date, DefaultGantaParentsOfStep2, trans)
	if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	return ReturnNoError()
}


