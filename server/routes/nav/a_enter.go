package nav

import "go-devops-mimi/server/controller"

type RouterGroup struct {
	Nav
}

// 初始化 cmdb Controller
var (
	NavController  = controller.ControllerGroupApp.NavcontrollerGroup.NavController
	LinkController = controller.ControllerGroupApp.NavcontrollerGroup.LinkController
)
