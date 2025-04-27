package exec

import (
	Req "go-devops-mimi/server/model/exec/request"
	"go-devops-mimi/server/public/tools"

	"github.com/gin-gonic/gin"
)

type ScriptlibraryController struct{}

// 创建脚本库
func (m *ScriptlibraryController) Add(c *gin.Context) {
	req := new(Req.ScriptLibraryAddReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return ScriptLibraryLogic.Add(c, req)
	})
}

// 更新脚本库
func (m *ScriptlibraryController) Update(c *gin.Context) {
	req := new(Req.ScriptLibraryUpdateReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return ScriptLibraryLogic.Update(c, req)
	})
}

// 显示脚本库列表
func (m *ScriptlibraryController) List(c *gin.Context) {
	req := new(Req.ScriptLibraryListReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return ScriptLibraryLogic.List(c, req)
	})
}

// 显示指定脚本库
func (m *ScriptlibraryController) Info(c *gin.Context) {
	req := new(Req.ScriptLibraryInfoReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return ScriptLibraryLogic.Info(c, req)
	})
}
func (m *ScriptlibraryController) Delete(c *gin.Context) {
	req := new(Req.ScriptLibraryDeleteReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return ScriptLibraryLogic.Delete(c, req)
	})
}
