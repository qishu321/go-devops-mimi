package test

import (
	"fmt"
	"testing"

	"go-devops-mimi/server/config"
	"go-devops-mimi/server/public/common"
	"go-devops-mimi/server/public/tools"
	"go-devops-mimi/server/service"
)

func InitConfig() {
	// 加载配置文件到全局配置结构体
	config.InitConfig()

	// 初始化日志
	common.InitLogger()

	// 初始化数据库(mysql)
	common.InitDB()

	// 初始化casbin策略管理器
	common.InitCasbinEnforcer()

	// 初始化Validator数据校验
	common.InitValidate()
}

func TestUserExist(t *testing.T) {
	InitConfig()

	filter := tools.H{
		"id": "111",
	}

	if service.ServiceGroupApp.SystemServiceGroup.UserService.Exist(filter) {
		fmt.Println("用户名已存在")
	} else {
		fmt.Println("用户名不存在")
	}
}
