package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// AgentInfo 用于接收 agent 注册信息
type AgentInfo struct {
	AgentID   string `json:"agent_id"`
	Host      string `json:"host"`
	Timestamp int64  `json:"timestamp"`
}

// Command 命令结构
type Command struct {
	Command   string `json:"command"`
	Timestamp int64  `json:"timestamp"`
}

// 全局 map 存储每个 Agent 的命令队列（用 slice 模拟队列）
var (
	agentCommands = make(map[string][]Command)
	agentMutex    sync.RWMutex
)

func registerAgent(c *gin.Context) {
	var info AgentInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	info.Timestamp = time.Now().Unix()
	// 注册后初始化该 agent 的命令队列为空
	agentMutex.Lock()
	agentCommands[info.AgentID] = []Command{}
	agentMutex.Unlock()
	log.Printf("Agent %s 注册成功，Host=%s", info.AgentID, info.Host)
	c.JSON(http.StatusOK, gin.H{"status": "registered", "agent_id": info.AgentID})
}

func getCommands(c *gin.Context) {
	agentID := c.Param("agentid")
	agentMutex.Lock()
	commands, exists := agentCommands[agentID]
	if !exists {
		commands = []Command{}
	}
	// 下发后清空队列
	agentCommands[agentID] = []Command{}
	agentMutex.Unlock()
	c.JSON(http.StatusOK, commands)
}

func addCommand(c *gin.Context) {
	agentID := c.Param("agentid")
	var cmd Command
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cmd.Timestamp = time.Now().Unix()
	agentMutex.Lock()
	agentCommands[agentID] = append(agentCommands[agentID], cmd)
	agentMutex.Unlock()
	log.Printf("下发命令给 Agent %s: %s", agentID, cmd.Command)
	c.JSON(http.StatusOK, gin.H{"status": "command added"})
}

func main() {
	router := gin.Default()

	// Agent 注册接口
	router.POST("/agents/register", registerAgent)
	// Agent 轮询获取命令
	router.GET("/agents/:agentid/commands", getCommands)
	// 外部下发命令接口（例如运维人员通过这个接口下发命令）
	router.POST("/agents/:agentid/command", addCommand)

	log.Println("Server 正在监听 :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Gin 启动失败: %v", err)
	}
}
