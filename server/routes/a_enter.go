package routes

import (
	"fmt"
	"time"

	"go-devops-mimi/server/config"
	"go-devops-mimi/server/middleware"
	"go-devops-mimi/server/public/common"
	"go-devops-mimi/server/routes/cmdb"
	"go-devops-mimi/server/routes/example"
	"go-devops-mimi/server/routes/exec_router"
	"go-devops-mimi/server/routes/system"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Cmdb    cmdb.RouterGroup
	Exec    exec_router.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

// 初始化
func InitRoutes() *gin.Engine {
	//设置模式
	gin.SetMode(config.Conf.System.Mode)

	// 创建带有默认中间件的路由:
	// 日志与恢复中间件
	r := gin.Default()
	// 创建不带中间件的路由:
	// r := gin.New()
	// r.Use(gin.Recovery())

	// 启用限流中间件
	// 默认每50毫秒填充一个令牌，最多填充200个
	fillInterval := time.Duration(config.Conf.RateLimit.FillInterval)
	capacity := config.Conf.RateLimit.Capacity
	r.Use(middleware.RateLimitMiddleware(time.Millisecond*fillInterval, capacity))

	// 启用全局跨域中间件
	r.Use(middleware.CORSMiddleware())

	// 启用操作日志中间件
	r.Use(middleware.OperationLogMiddleware())

	// 初始化JWT认证中间件
	authMiddleware, err := middleware.InitAuth()
	if err != nil {
		common.Log.Panicf("初始化JWT中间件失败：%v", err)
		panic(fmt.Sprintf("初始化JWT中间件失败：%v", err))
	}

	// 基础路由分组
	apiGroup := r.Group("/" + config.Conf.System.UrlPathPrefix)

	// 路由分组
	systemApiGroup := apiGroup.Group("/system")
	systemApiGroup.Use(middleware.OperationLogMiddleware())
	// 注册路由
	RouterGroupApp.System.InitBaseRoutes(systemApiGroup, authMiddleware)         // 注册基础路由, 不需要jwt认证中间件,不需要casbin中间件
	RouterGroupApp.System.InitUserRoutes(systemApiGroup, authMiddleware)         // 注册用户路由, jwt认证中间件,casbin鉴权中间件
	RouterGroupApp.System.InitGroupRoutes(systemApiGroup, authMiddleware)        // 注册分组路由, jwt认证中间件,casbin鉴权中间件
	RouterGroupApp.System.InitRoleRoutes(systemApiGroup, authMiddleware)         // 注册角色路由, jwt认证中间件,casbin鉴权中间件
	RouterGroupApp.System.InitMenuRoutes(systemApiGroup, authMiddleware)         // 注册菜单路由, jwt认证中间件,casbin鉴权中间件
	RouterGroupApp.System.InitApiRoutes(systemApiGroup, authMiddleware)          // 注册接口路由, jwt认证中间件,casbin鉴权中间件
	RouterGroupApp.System.InitOperationLogRoutes(systemApiGroup, authMiddleware) // 注册操作日志路由, jwt认证中间件,casbin鉴权中间件

	// 路由分组
	exampleApiGroup := apiGroup.Group("/example")
	RouterGroupApp.Example.InitExamleRoutes(exampleApiGroup, authMiddleware)
	//cmdb路由分组
	cmdbApiGroup := apiGroup.Group("/cmdb")
	RouterGroupApp.Cmdb.InitNodeGroupRoutes(cmdbApiGroup, authMiddleware)
	RouterGroupApp.Cmdb.InitNodeRoutes(cmdbApiGroup, authMiddleware)
	//exec路由分组
	execApiGroup := apiGroup.Group("/exec")
	RouterGroupApp.Exec.InitScriptRoutes(execApiGroup, authMiddleware)
	RouterGroupApp.Exec.InitScriptLibraryRoutes(execApiGroup, authMiddleware)
	RouterGroupApp.Exec.InitTransferRoutes(execApiGroup, authMiddleware)
	RouterGroupApp.Exec.InitTaskManageRoutes(execApiGroup, authMiddleware)
	common.Log.Info("初始化路由完成！")
	return r
}
