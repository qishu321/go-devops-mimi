package exec_router

import (
	"go-devops-mimi/server/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type TaskManageRouter struct{}

func (s TaskManageRouter) InitTaskManageRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	library := r.Group("/task_manage")
	// 开启jwt认证中间件
	library.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	library.Use(middleware.CasbinMiddleware())

	{
		library.POST("/add", TaskManageController.Add)
		library.POST("/update", TaskManageController.Update)
		library.POST("/delete", TaskManageController.Delete)
		library.GET("/list", TaskManageController.List)
		library.GET("/info", TaskManageController.Info)

	}
	run := r.Group("/run_task_manage")

	{
		run.GET("/run", TaskManageController.Run)
		run.GET("/run_info", TaskManageController.InfoRun)
		run.GET("/run_list", TaskManageController.ListRun)
		run.GET("/run_info_webSocket", TaskManageController.RunInfoWebSocket)
		run.POST("add_run", TaskManageController.AddRun)
	}
	return r
}
