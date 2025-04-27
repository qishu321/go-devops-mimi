package example

import "go-devops-mimi/server/service"

type LogicGroup struct {
	CloudAccountLogic
}

// 初始化 service
var (
	CloudAccountService = service.ServiceGroupApp.ExampleServiceGroup.CloudAccountService
)
