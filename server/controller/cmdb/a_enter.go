package cmdb

import "go-devops-mimi/server/logic"

type ControllerGroup struct {
	AgentController
	NodeController
	NodeGroupController
}

var (
	AgentLogic     = logic.LogicGroupApp.CmdbLogicGroup.AgentLogic
	NodeLogic      = logic.LogicGroupApp.CmdbLogicGroup.NodeLogic
	NodeGroupLogic = logic.LogicGroupApp.CmdbLogicGroup.NodeGroupLogic
)
