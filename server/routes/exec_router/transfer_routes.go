package exec_router

import (
	"go-devops-mimi/server/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type TransferRouter struct{}

func (s TransferRouter) InitTransferRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	Transfer := r.Group("/transfer")
	// 开启jwt认证中间件
	Transfer.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	Transfer.Use(middleware.CasbinMiddleware())

	{
		Transfer.POST("/add_run", TransferController.Add)
		Transfer.GET("/list", TransferController.List)
		Transfer.GET("/info", TransferController.Info)
		Transfer.POST("/upload", TransferController.UploadFile)
	}

	return r
}
