package exec

import (
	Req "go-devops-mimi/server/model/exec/request"
	"go-devops-mimi/server/public/tools"

	"github.com/gin-gonic/gin"
)

type TaskManageController struct{}

// 创建
func (m *TaskManageController) Add(c *gin.Context) {
	req := new(Req.TaskManageAddReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return ManageLogic.Add(c, req)
	})
}

// 更新
func (m *TaskManageController) Update(c *gin.Context) {
	req := new(Req.TaskManageUpdateReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return ManageLogic.Update(c, req)
	})
}

// 显示列表
func (m *TaskManageController) List(c *gin.Context) {
	req := new(Req.TaskManageListReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return ManageLogic.List(c, req)
	})
}

// 显示指定
func (m *TaskManageController) Info(c *gin.Context) {
	req := new(Req.TaskManageInfoReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return ManageLogic.Info(c, req)
	})
}
func (m *TaskManageController) Delete(c *gin.Context) {
	req := new(Req.TaskManageDeleteReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return ManageLogic.Delete(c, req)
	})
}
