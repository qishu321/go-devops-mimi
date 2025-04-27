package cmdb

import "go-devops-mimi/server/controller"

type RouterGroup struct {
	AgentRouter
	NodeRouter
	NodeGroupRouter
}

// 初始化 cmdb Controller
var (
	AgentController     = controller.ControllerGroupApp.CmdbControllerGroup.AgentController
	NodeController      = controller.ControllerGroupApp.CmdbControllerGroup.NodeController
	NodeGroupController = controller.ControllerGroupApp.CmdbControllerGroup.NodeGroupController
)
