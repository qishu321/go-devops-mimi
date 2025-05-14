package exec

import (
	Req "go-devops-mimi/server/model/exec/request"
	"go-devops-mimi/server/public/tools"

	"github.com/gin-gonic/gin"
)

type CronController struct{}

// 创建脚本库
func (m *CronController) Add(c *gin.Context) {
	req := new(Req.CronAddReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return CronLogic.Add(c, req)
	})
}

// 更新脚本库
func (m *CronController) Update(c *gin.Context) {
	req := new(Req.CronUpdateReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return CronLogic.Update(c, req)
	})
}

// 显示脚本库列表
func (m *CronController) List(c *gin.Context) {
	req := new(Req.CronListReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return CronLogic.List(c, req)
	})
}

// 显示指定脚本库
func (m *CronController) Info(c *gin.Context) {
	req := new(Req.CronInfoReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return CronLogic.Info(c, req)
	})
}
func (m *CronController) Enable(c *gin.Context) {
	req := new(Req.CronEnableReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return CronLogic.Enable(c, req)
	})
}

func (m *CronController) Delete(c *gin.Context) {
	req := new(Req.CronDeleteReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return CronLogic.Delete(c, req)
	})
}
