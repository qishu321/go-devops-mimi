package exec

import (
	Req "go-devops-mimi/server/model/exec/request"
	"go-devops-mimi/server/public/tools"

	"github.com/gin-gonic/gin"
)

type CronLogController struct{}

func (m *CronLogController) List(c *gin.Context) {
	req := new(Req.CronListReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return CronLogLogic.List(c, req)
	})
}
