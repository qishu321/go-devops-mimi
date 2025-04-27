package example

import "go-devops-mimi/server/controller"

type RouterGroup struct {
	CloudAccount
}

// 初始化 cmdb Controller
var (
	CloudAccountController = controller.ControllerGroupApp.ExampleControllerGroup.CloudAccountController
)
