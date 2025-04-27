package main

import (
	"fmt"
	"go-devops-mimi/server/agent/agent"
	"go-devops-mimi/server/agent/config"
)

func main() {
	err := config.LoadConfig("./config/config.ini")
	if err != nil {
		fmt.Println("配置加载失败:", err)
		return
	}

	agent.Init()
}
