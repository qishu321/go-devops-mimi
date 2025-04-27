package common

import (
	"go-devops-mimi/server/model/cmdb"
	"log"
	"time"
)

func InitCronJobs() {
	startAgentStatusChecker()
}

// 后台监控：定时检测 agent 是否超时（未收到心跳）
func startAgentStatusChecker() {
	ticker := time.NewTicker(30 * time.Second)
	go func() {
		for range ticker.C {
			var agents []cmdb.Agent
			// 查询所有在线的 Agent
			if err := DB.Where("online = ?", true).Find(&agents).Error; err != nil {
				log.Printf("查询在线 Agent 失败: %v", err)
				continue
			}
			for _, agent := range agents {
				if time.Since(agent.LastHeartbeat) > 60*time.Second {
					// 超过阈值则标记为离线
					DB.Model(&cmdb.Agent{}).Where("id = ?", agent.ID).Update("online", false)
					log.Printf("Agent %s 超时，标记为离线", agent.Agentid)
				}
			}
		}
	}()
}
