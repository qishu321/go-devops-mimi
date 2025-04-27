package exec

import (
	Req "go-devops-mimi/server/model/exec/request"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/gin-gonic/gin"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 创建
func (m *TaskManageController) Run(c *gin.Context) {
	// 1. 绑定 & 验证查询参数
	var req Req.TaskManageInfoReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数绑定失败"}) // 仅 HTTP 路径使用
		return
	}
	// 2. 升级 WebSocket（劫持连接）
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	defer ws.Close()
	// 3. 调用业务逻辑
	TaskManageLogic.RunWebSocket(c, ws, &req)

}
