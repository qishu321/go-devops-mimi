package cmdb

import (
	"go-devops-mimi/server/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type NodeRouter struct{}

func (s NodeRouter) InitNodeRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	node := r.Group("/node")
	// 开启jwt认证中间件
	node.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	node.Use(middleware.CasbinMiddleware())

	{
		node.POST("/add", NodeController.Add)
		node.GET("/list", NodeController.List)
		node.POST("/update", NodeController.Update)
		node.POST("/delete", NodeController.Delete)
		node.POST("/add_nodes_group", NodeController.AddNodesGroup)
	}

	return r
}
