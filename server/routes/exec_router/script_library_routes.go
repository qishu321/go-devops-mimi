package exec_router

import (
	"go-devops-mimi/server/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type ScriptLibraryRouter struct{}

func (s ScriptLibraryRouter) InitScriptLibraryRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	library := r.Group("/script_library")
	// 开启jwt认证中间件
	library.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	library.Use(middleware.CasbinMiddleware())

	{
		library.POST("/add", ScriptLibraryController.Add)
		library.POST("/update", ScriptLibraryController.Update)
		library.POST("/delete", ScriptLibraryController.Delete)
		library.GET("/list", ScriptLibraryController.List)
		library.GET("/info", ScriptLibraryController.Info)

	}

	return r
}
