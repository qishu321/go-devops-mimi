package nav

import (
	navReq "go-devops-mimi/server/model/nav/request"
	"go-devops-mimi/server/public/tools"

	"github.com/gin-gonic/gin"
)

type LinkController struct{}

func (m *LinkController) Info(c *gin.Context) {
	req := new(navReq.LinkInfoRequest)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return LinkLogic.Info(c, req)
	})
}

// Add 新建记录
func (m *LinkController) Add(c *gin.Context) {
	req := new(navReq.LinkAddRequest)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return LinkLogic.Add(c, req)
	})
}

// Update 更新记录
func (m *LinkController) Update(c *gin.Context) {
	req := new(navReq.LinkUpdateRequest)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return LinkLogic.Update(c, req)
	})
}

// Delete 删除记录
func (m *LinkController) Delete(c *gin.Context) {
	req := new(navReq.LinkDeleteRequest)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return LinkLogic.Delete(c, req)
	})
}
