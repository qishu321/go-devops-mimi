package exec

import (
	Req "go-devops-mimi/server/model/exec/request"
	"go-devops-mimi/server/public/tools"

	"github.com/gin-gonic/gin"
)

type ScriptController struct{}

// // List 记录列表
// func (m *ScriptController) List(c *gin.Context) {
// 	req := new(cmdbReq.NodeGroupListReq)
// 	tools.Run(c, req, func() (interface{}, interface{}) {
// 		return NodeGroupLogic.List(c, req)
// 	})
// }

// 执行脚本
func (m *ScriptController) Add_Run(c *gin.Context) {
	req := new(Req.ScriptCmdRunReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return ScriptLogic.Add_Run(c, req)
	})
}
func (m *ScriptController) List(c *gin.Context) {
	req := new(Req.ScriptListReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return ScriptLogic.List(c, req)
	})
}
