package exec_router

import (
	"go-devops-mimi/server/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type ScriptRouter struct{}

func (s ScriptRouter) InitScriptRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	script := r.Group("/script")
	// 开启jwt认证中间件
	script.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	script.Use(middleware.CasbinMiddleware())

	{
		script.POST("/add_run", ScriptController.Add_Run)
		script.GET("/list", ScriptController.List)
	}

	return r
}
