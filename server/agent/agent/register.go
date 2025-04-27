package agent

import (
	"fmt"
	"go-devops-mimi/server/agent/config"
	"go-devops-mimi/server/public/tools"
	"log"
	"time"
)

// 启动时注册 Agent
func registerAgent() {
	// 构造请求数据，从配置文件中读取参数
	agentid := tools.IPNameToNum(config.Cfg.AgentName, config.Cfg.IP)

	reqData := AgentRegisterReq{
		Token:     config.Cfg.Token,
		AgentID:   agentid, // 假设 config.Cfg.AgentID 为 string 类型，如果为数值类型，请转换为 string
		AgentName: config.Cfg.AgentName,
		Version:   config.Cfg.Version, // 可在 config.ini 中定义版本信息
		OsType:    GetOSType(),        // 请确保在 config.ini 中配置 osType
		IP:        config.Cfg.IP,
		Label:     config.Cfg.Label,
	}

	// 拼接注册接口的 URL，假设接口路径为 "/add"
	url := fmt.Sprintf("%s/add", config.Cfg.ServerURL)

	resp, err := httpPost[AgentResponse](reqData, url)
	if err != nil {
		log.Panicf("注册 Agent 失败: %v", err)
	}
	log.Printf("Agent 注册成功: %s", resp.Msg)
}

// 发送心跳
func sendHeartbeat() {
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()
	agentid := tools.IPNameToNum(config.Cfg.AgentName, config.Cfg.IP)

	for range ticker.C {
		reqData := AgentHeartbeatrReq{
			Token:   config.Cfg.Token,
			AgentID: agentid,            // 假设 config.Cfg.AgentID 为 string 类型，如果为数值类型，请转换为 string
			Version: config.Cfg.Version, // 可在 config.ini 中定义版本信息
			OsType:  GetOSType(),        // 请确保在 config.ini 中配置 osType
		}

		// 拼接注册接口的 URL，假设接口路径为 "/add"
		url := fmt.Sprintf("%s/heartbeat", config.Cfg.ServerURL)
		resp, err := httpPost[AgentResponse](reqData, url)
		if err != nil {
			log.Printf("Agent 发送心跳失败: %s", resp.Msg)
		}
	}
}
