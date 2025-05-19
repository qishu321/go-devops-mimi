package logic

import (
	"go-devops-mimi/server/logic/cmdb"
	"go-devops-mimi/server/logic/example"
	"go-devops-mimi/server/logic/exec_logic"
	"go-devops-mimi/server/logic/nav"
	"go-devops-mimi/server/logic/system"
)

type LogicGroup struct {
	SystemLogicGroup  system.LogicGroup
	ExampleLogicGroup example.LogicGroup
	CmdbLogicGroup    cmdb.LogicGroup
	ExecLogicGroup    exec_logic.LogicGroup
	NavLogicGroup     nav.LogicGroup
}

var LogicGroupApp = new(LogicGroup)
