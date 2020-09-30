package model

import "invest/utils"

/*
	Restricted means that you will get steps based on project step (either 1 or 2)
 */
func (g *Ganta) Get_parent_ganta_steps_by_project_id_and_step(project_step interface{}) (utils.Msg) {
	gantas, err := g.OnlyGetParentsByProjectId(project_step, GetDB())
	if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	var gantaMap = []map[string]interface{}{}
	for _, each := range gantas {
		gantaMap = append(gantaMap, Struct_to_map(each))
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = gantaMap

	return ReturnNoErrorWithResponseMessage(resp)
}

/*
	get ganta sub-steps
 */
func (g *Ganta) Get_child_ganta_steps_by_project_id_and_step(project_step int) (utils.Msg) {
	err := g.OnlyGetChildrenByIdAndProjectIdStep(project_step, GetDB())
	if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	var gantasMap = make([]map[string]interface{}, len(g.GantaChildren))
	for _, child := range g.GantaChildren {
		gantasMap = append(gantasMap, Struct_to_map(child))
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = gantasMap

	return ReturnNoErrorWithResponseMessage(resp)
}

/*
	get ganta sub-steps with documents
 */
func (g *Ganta) Get_preloaded_and_restricted_child_ganta_steps(project_step interface{}) (utils.Msg) {
	err := g.OnlyGetPreloadedChildStepsByProjectIdAndStep(project_step, GetDB())
	if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	var gantasMap = []map[string]interface{}{}
	for _, child := range g.GantaChildren {
		gantasMap = append(gantasMap, Struct_to_map(child))
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = gantasMap

	return ReturnNoErrorWithResponseMessage(resp)
}
