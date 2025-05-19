package nav

import (
	navReq "go-devops-mimi/server/model/nav/request"
	"go-devops-mimi/server/public/tools"

	"github.com/gin-gonic/gin"
)

type NavController struct{}

// List 记录列表
func (m *NavController) List(c *gin.Context) {
	req := new(navReq.NavListRequest)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return NavLogic.List(c, req)
	})
}
func (m *NavController) Info(c *gin.Context) {
	req := new(navReq.NavInfoRequest)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return NavLogic.Info(c, req)
	})
}

// Add 新建记录
func (m *NavController) Add(c *gin.Context) {
	req := new(navReq.NavAddRequest)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return NavLogic.Add(c, req)
	})
}

// Update 更新记录
func (m *NavController) Update(c *gin.Context) {
	req := new(navReq.NavUpdateRequest)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return NavLogic.Update(c, req)
	})
}

// Delete 删除记录
func (m *NavController) Delete(c *gin.Context) {
	req := new(navReq.NavDeleteRequest)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return NavLogic.Delete(c, req)
	})
}

// DeleteAll 删除分类和关联的链接
func (m *NavController) DeleteAll(c *gin.Context) {
	req := new(navReq.NavDeleteRequest)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return NavLogic.DeleteAll(c, req)
	})
}
