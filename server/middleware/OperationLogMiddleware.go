package middleware

import (
	"fmt"
	"strings"
	"time"

	"go-devops-mimi/server/config"
	"go-devops-mimi/server/model/system"
	"go-devops-mimi/server/public/tools"
	"go-devops-mimi/server/service"

	"github.com/gin-gonic/gin"
)

// 操作日志channel
var OperationLogChan = make(chan *system.OperationLog, 30)

func OperationLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行耗时
		timeCost := endTime.Sub(startTime).Milliseconds()

		// 获取当前登录用户
		var username string
		ctxUser, _ := c.Get("user")
		user, ok := ctxUser.(system.User)
		if !ok {
			username = "未登录"
		} else {
			username = user.Username
		}

		// 获取访问路径
		path := strings.TrimPrefix(c.FullPath(), "/"+config.Conf.System.UrlPathPrefix)
		// path := c.Request.URL.Path
		// 请求方式
		method := c.Request.Method

		// 获取接口描述
		api := new(system.Api)
		_ = service.ServiceGroupApp.SystemServiceGroup.ApiService.Find(tools.H{"path": path, "method": method}, api)

		operationLog := system.OperationLog{
			Username:   username,
			Ip:         c.ClientIP(),
			IpLocation: "",
			Method:     method,
			Path:       path,
			Remark:     api.Remark,
			Status:     c.Writer.Status(),
			StartTime:  fmt.Sprintf("%v", startTime),
			TimeCost:   timeCost,
			UserAgent:  c.Request.UserAgent(),
		}

		// 最好是将日志发送到rabbitmq或者kafka中
		// 这里是发送到channel中，开启3个goroutine处理
		OperationLogChan <- &operationLog
	}
}
