package exec_logic

import (
	"go-devops-mimi/server/service"
)

type LogicGroup struct {
	ScriptLogic
	ScriptLibraryLogic
	TransferLogic
	TaskManageLogic
}

// 初始化 service
var (
	ScriptService        = service.ServiceGroupApp.ExecServiceGroup.ScriptService
	ScriptLogService     = service.ServiceGroupApp.ExecServiceGroup.ScriptLogService
	ScriptLibraryService = service.ServiceGroupApp.ExecServiceGroup.ScriptLibraryService
	TransferService      = service.ServiceGroupApp.ExecServiceGroup.TransferService
	userService          = service.ServiceGroupApp.SystemServiceGroup.UserService
	nodeService          = service.ServiceGroupApp.CmdbServiceGroup.Node_svc
	TaskManageService    = service.ServiceGroupApp.ExecServiceGroup.TaskManageService
)
