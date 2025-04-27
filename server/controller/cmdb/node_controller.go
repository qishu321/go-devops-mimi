package cmdb

import (
	cmdbReq "go-devops-mimi/server/model/cmdb/request"
	"go-devops-mimi/server/public/tools"

	"github.com/gin-gonic/gin"
)

type NodeController struct{}

// List 记录列表
func (m *NodeController) List(c *gin.Context) {
	req := new(cmdbReq.NodesListReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return NodeLogic.List(c, req)
	})
}

// Add 新建记录
func (m *NodeController) Add(c *gin.Context) {
	req := new(cmdbReq.NodesCreateReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return NodeLogic.Add(c, req)
	})
}

// Update 更新记录
func (m *NodeController) Update(c *gin.Context) {
	req := new(cmdbReq.NodesUpdateReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return NodeLogic.Update(c, req)
	})
}

// Delete 删除记录
func (m *NodeController) Delete(c *gin.Context) {
	req := new(cmdbReq.NodesDeleteReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return NodeLogic.Delete(c, req)
	})
}

func (m *NodeController) AddNodesGroup(c *gin.Context) {
	req := new(cmdbReq.AddNodesGroupReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return NodeLogic.AddNodesGroup(c, req)
	})
}
