package exec

import "go-devops-mimi/server/logic"

type ControllerGroup struct {
	ScriptController
	ScriptlibraryController
	TransferController
	TaskManageController
	CronController
	CronLogController
}

var (
	ScriptLogic        = logic.LogicGroupApp.ExecLogicGroup.ScriptLogic
	ScriptLibraryLogic = logic.LogicGroupApp.ExecLogicGroup.ScriptLibraryLogic
	TransferLogic      = logic.LogicGroupApp.ExecLogicGroup.TransferLogic
	ManageLogic        = logic.LogicGroupApp.ExecLogicGroup.ManageLogic
	CronLogic          = logic.LogicGroupApp.ExecLogicGroup.CronLogic
	CronLogLogic       = logic.LogicGroupApp.ExecLogicGroup.CronLogLogic
)
