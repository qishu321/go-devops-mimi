package nav

import (
	"go-devops-mimi/server/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type Nav struct{}

func (s Nav) InitNavRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	nav := r.Group("/nav")
	nav.GET("/list", NavController.List)
	// 开启jwt认证中间件
	nav.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	nav.Use(middleware.CasbinMiddleware())

	{
		nav.POST("/add", NavController.Add)
		nav.GET("/info", NavController.Info)
		nav.POST("/update", NavController.Update)
		nav.POST("/delete", NavController.Delete)
		nav.POST("/delete_all", NavController.DeleteAll)
	}
	link := nav.Group("/link")
	// 开启jwt认证中间件
	link.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	link.Use(middleware.CasbinMiddleware())

	{
		link.POST("/add", LinkController.Add)
		link.GET("/info", LinkController.Info)
		link.POST("/update", LinkController.Update)
		link.POST("/delete", LinkController.Delete)
	}
	return r
}
