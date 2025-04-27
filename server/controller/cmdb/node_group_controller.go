package cmdb

import (
	cmdbReq "go-devops-mimi/server/model/cmdb/request"
	"go-devops-mimi/server/public/tools"

	"github.com/gin-gonic/gin"
)

type NodeGroupController struct{}

// List 记录列表
func (m *NodeGroupController) List(c *gin.Context) {
	req := new(cmdbReq.NodeGroupListReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return NodeGroupLogic.List(c, req)
	})
}

// Add 新建记录
func (m *NodeGroupController) Add(c *gin.Context) {
	req := new(cmdbReq.NodeGroupCreateReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return NodeGroupLogic.Add(c, req)
	})
}

// Update 更新记录
func (m *NodeGroupController) Update(c *gin.Context) {
	req := new(cmdbReq.NodeGroupUpdateReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return NodeGroupLogic.Update(c, req)
	})
}

// Delete 删除记录
func (m *NodeGroupController) Delete(c *gin.Context) {
	req := new(cmdbReq.NodeGroupDeleteReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return NodeGroupLogic.Delete(c, req)
	})
}

func (m *NodeGroupController) AddNodeToGroup(c *gin.Context) {
	req := new(cmdbReq.AddNodesToGroupReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return NodeGroupLogic.AddNodeToGroup(c, req)
	})
}
func (m *NodeGroupController) RemoveNodeGroup(c *gin.Context) {
	req := new(cmdbReq.RemoveNodeGroupReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return NodeGroupLogic.RemoveNodeGroup(c, req)
	})
}
