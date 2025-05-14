package exec_router

import "go-devops-mimi/server/controller"

type RouterGroup struct {
	ScriptRouter
	ScriptLibraryRouter
	TransferRouter
	TaskManageRouter
	CronRouter
}

// 初始化 cmdb Controller
var (
	ScriptController        = controller.ControllerGroupApp.ExecControllerGroup.ScriptController
	ScriptLibraryController = controller.ControllerGroupApp.ExecControllerGroup.ScriptlibraryController
	TransferController      = controller.ControllerGroupApp.ExecControllerGroup.TransferController
	TaskManageController    = controller.ControllerGroupApp.ExecControllerGroup.TaskManageController
	CronController          = controller.ControllerGroupApp.ExecControllerGroup.CronController
	CrondLogController      = controller.ControllerGroupApp.ExecControllerGroup.CronLogController
)
