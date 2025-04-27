package service

import (
	"go-devops-mimi/server/service/cmdb"
	"go-devops-mimi/server/service/example"
	"go-devops-mimi/server/service/exec_svc"
	"go-devops-mimi/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	CmdbServiceGroup    cmdb.ServiceGroup
	ExecServiceGroup    exec_svc.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
