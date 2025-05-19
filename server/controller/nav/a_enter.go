package nav

import "go-devops-mimi/server/logic"

type ControllerGroup struct {
	NavController
	LinkController
}

var (
	NavLogic  = logic.LogicGroupApp.NavLogicGroup.NavLogic
	LinkLogic = logic.LogicGroupApp.NavLogicGroup.LinkLogic
)
