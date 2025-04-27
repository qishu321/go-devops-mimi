package cmdb

import (
	cmdbReq "go-devops-mimi/server/model/cmdb/request"
	"go-devops-mimi/server/public/tools"

	"github.com/gin-gonic/gin"
)

type AgentController struct{}

// List 记录列表
func (m *AgentController) List(c *gin.Context) {
	req := new(cmdbReq.AgentListReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return AgentLogic.List(c, req)
	})
}

// Add 新建记录
func (m *AgentController) Add(c *gin.Context) {
	req := new(cmdbReq.CreateAgentReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return AgentLogic.Add(c, req)
	})
}

// Update 更新记录
func (m *AgentController) Update(c *gin.Context) {
	req := new(cmdbReq.UpdateAgentReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return AgentLogic.Update(c, req)
	})
}
func (m *AgentController) AgentHeartbeat(c *gin.Context) {
	req := new(cmdbReq.AgentHeartbeatReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return AgentLogic.AgentHeartbeat(c, req)
	})
}

// Delete 删除记录
func (m *AgentController) Delete(c *gin.Context) {
	req := new(cmdbReq.DeleteAgentReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return AgentLogic.Delete(c, req)
	})
}
