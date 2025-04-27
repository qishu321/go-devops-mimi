package example

import "go-devops-mimi/server/logic"

type ControllerGroup struct {
	CloudAccountController
}

var (
	CloudAccountLogic = logic.LogicGroupApp.ExampleLogicGroup.CloudAccountLogic
)
