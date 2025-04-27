package cmdb

import (
	"go-devops-mimi/server/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type NodeGroupRouter struct{}

func (s NodeGroupRouter) InitNodeGroupRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	group := r.Group("/node_group")
	// 开启jwt认证中间件
	group.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	group.Use(middleware.CasbinMiddleware())

	{
		group.POST("/add", NodeGroupController.Add)
		group.GET("/list", NodeGroupController.List)
		group.POST("/update", NodeGroupController.Update)
		group.POST("/delete", NodeGroupController.Delete)
		group.POST("/add_node_to_group", NodeGroupController.AddNodeToGroup)
		group.POST("/remonv_node_to_group", NodeGroupController.RemoveNodeGroup)

	}

	return r
}
