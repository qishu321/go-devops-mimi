package exec_router

import (
	"go-devops-mimi/server/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type CronRouter struct{}

func (s CronRouter) InitCronRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	library := r.Group("/cron")
	// 开启jwt认证中间件
	library.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	library.Use(middleware.CasbinMiddleware())

	{
		library.POST("/add", CronController.Add)
		library.POST("/update", CronController.Update)
		library.POST("/delete", CronController.Delete)
		library.POST("/enable", CronController.Enable)
		library.GET("/list", CronController.List)
		library.GET("/info", CronController.Info)
		library.GET("/log/list", CrondLogController.List)

	}

	return r
}
