package system

import "go-devops-mimi/server/controller"

type RouterGroup struct {
	ApiRouter
	BaseRouter
	GroupRouter
	MenuRouter
	OperationLogRouter
	RoleRouter
	UserRouter
}

var (
	ApiController          = controller.ControllerGroupApp.SystemControllerGroup.ApiController
	BaseController         = controller.ControllerGroupApp.SystemControllerGroup.BaseController
	GroupController        = controller.ControllerGroupApp.SystemControllerGroup.GroupController
	MenuController         = controller.ControllerGroupApp.SystemControllerGroup.MenuController
	OperationLogController = controller.ControllerGroupApp.SystemControllerGroup.OperationLogController
	RoleController         = controller.ControllerGroupApp.SystemControllerGroup.RoleController
	UserController         = controller.ControllerGroupApp.SystemControllerGroup.UserController
)
