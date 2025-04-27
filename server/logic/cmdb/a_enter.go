package cmdb

import (
	"go-devops-mimi/server/service"
	user_svc "go-devops-mimi/server/service"
)

type LogicGroup struct {
	NodeLogic
	NodeGroupLogic
	AgentLogic
}

// 初始化 service
var (
	NodeService       = service.ServiceGroupApp.CmdbServiceGroup.Node_svc
	NodeGroupService  = service.ServiceGroupApp.CmdbServiceGroup.NodeGroup_svc
	AgentService      = service.ServiceGroupApp.CmdbServiceGroup.Agent_svc
	AgentGroupService = service.ServiceGroupApp.CmdbServiceGroup.AgentGroup_svc

	userService = user_svc.ServiceGroupApp.SystemServiceGroup.UserService
)
