package system

import (
	systemReq "go-devops-mimi/server/model/system/request"
	"go-devops-mimi/server/public/tools"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

// Dashboard 系统首页展示数据
func (m *BaseController) Dashboard(c *gin.Context) {
	req := new(systemReq.BaseDashboardReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return BaseLogic.Dashboard(c, req)
	})
}

// GetPasswd 生成加密密码
func (m *BaseController) GetPasswd(c *gin.Context) {
	req := new(systemReq.GetPasswdReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return BaseLogic.GetPasswd(c, req)
	})
}
