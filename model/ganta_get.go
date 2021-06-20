package model

import (
	"github.com/yerassyldanay/invest/utils/errormsg"
	"github.com/yerassyldanay/invest/utils/message"
)

/*
	Restricted means that you will get steps based on project step (either 1 or 2)
*/
func (g *Ganta) GetParentGantaStepsByProjectIdAndStep(project_step interface{}) message.Msg {
	gantas, err := g.OnlyGetParentsByProjectId(project_step, GetDB())
	if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	var gantaMap = []map[string]interface{}{}
	for _, each := range gantas {
		gantaMap = append(gantaMap, Struct_to_map(each))
	}

	var resp = errormsg.NoErrorFineEverthingOk
	resp["info"] = gantaMap

	return ReturnNoErrorWithResponseMessage(resp)
}

// get ganta sub-steps
func (g *Ganta) GetChildGantaStepsByProjectIdAndStep(project_step int) message.Msg {
	err := g.OnlyGetChildrenByIdAndProjectIdStep(project_step, GetDB())
	if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	var gantasMap = make([]map[string]interface{}, len(g.GantaChildren))
	for _, child := range g.GantaChildren {
		gantasMap = append(gantasMap, Struct_to_map(child))
	}

	var resp = errormsg.NoErrorFineEverthingOk
	resp["info"] = gantasMap

	return ReturnNoErrorWithResponseMessage(resp)
}
