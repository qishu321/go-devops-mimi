package exec

import "go-devops-mimi/server/logic"

type ControllerGroup struct {
	ScriptController
	ScriptlibraryController
	TransferController
	TaskManageController
}

var (
	ScriptLogic        = logic.LogicGroupApp.ExecLogicGroup.ScriptLogic
	ScriptLibraryLogic = logic.LogicGroupApp.ExecLogicGroup.ScriptLibraryLogic
	TransferLogic      = logic.LogicGroupApp.ExecLogicGroup.TransferLogic
	TaskManageLogic    = logic.LogicGroupApp.ExecLogicGroup.TaskManageLogic
)
