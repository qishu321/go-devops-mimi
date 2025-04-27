package cmdb

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type AgentRouter struct{}

func (s AgentRouter) InitAgentRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	agentRouter := r.Group("/agent")
	// 开启jwt认证中间件
	// agentRouter.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	// agentRouter.Use(middleware.CasbinMiddleware())

	{
		agentRouter.POST("/add", AgentController.Add)
		agentRouter.GET("/list", AgentController.List)
		agentRouter.POST("/update", AgentController.Update)
		agentRouter.POST("/delete", AgentController.Delete)
		agentRouter.POST("/heartbeat", AgentController.AgentHeartbeat)
	}

	return r
}
