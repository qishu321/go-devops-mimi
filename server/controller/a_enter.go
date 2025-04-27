package controller

import (
	"go-devops-mimi/server/controller/cmdb"
	"go-devops-mimi/server/controller/example"
	"go-devops-mimi/server/controller/exec"
	"go-devops-mimi/server/controller/system"
	"go-devops-mimi/server/public/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerGroup struct {
	SystemControllerGroup  system.ControllerGroup
	ExampleControllerGroup example.ControllerGroup
	CmdbControllerGroup    cmdb.ControllerGroup
	ExecControllerGroup    exec.ControllerGroup
}

var ControllerGroupApp = new(ControllerGroup)

func Demo(c *gin.Context) {
	CodeDebug()
	c.JSON(http.StatusOK, tools.H{"code": 200, "msg": "ok", "data": "pong"})
}

func CodeDebug() {
}
