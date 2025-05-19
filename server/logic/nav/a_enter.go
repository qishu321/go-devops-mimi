package nav

import "go-devops-mimi/server/service"

type LogicGroup struct {
	NavLogic
	LinkLogic
}

// 初始化 service
var (
	NavService  = service.ServiceGroupApp.NavServiceGroup.Nav_svc
	LinkService = service.ServiceGroupApp.NavServiceGroup.Link_svc
)
