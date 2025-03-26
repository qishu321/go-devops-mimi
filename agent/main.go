package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// AgentInfo 与 Command 结构与 server 中保持一致
type AgentInfo struct {
	AgentID   string `json:"agent_id"`
	Host      string `json:"host"`
	Timestamp int64  `json:"timestamp"`
}

type Command struct {
	Command   string `json:"command"`
	Timestamp int64  `json:"timestamp"`
}

// 注册 agent 到 server
func registerAgent(serverURL, agentID, host string) error {
	info := AgentInfo{
		AgentID: agentID,
		Host:    host,
	}
	data, err := json.Marshal(info)
	if err != nil {
		return err
	}
	resp, err := http.Post(serverURL+"/agents/register", "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("注册失败: %s", body)
	}
	return nil
}

// 定时轮询命令
func pollCommands(serverURL, agentID string) {
	for {
		url := fmt.Sprintf("%s/agents/%s/commands", serverURL, agentID)
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("轮询命令错误: %v", err)
			time.Sleep(10 * time.Second)
			continue
		}
		if resp.StatusCode != http.StatusOK {
			body, _ := ioutil.ReadAll(resp.Body)
			log.Printf("轮询命令失败: %s", body)
			resp.Body.Close()
			time.Sleep(10 * time.Second)
			continue
		}
		var cmds []Command
		err = json.NewDecoder(resp.Body).Decode(&cmds)
		resp.Body.Close()
		if err != nil {
			log.Printf("解析命令错误: %v", err)
		} else {
			for _, cmd := range cmds {
				log.Printf("收到命令: %s", cmd.Command)
			}
		}
		time.Sleep(10 * time.Second)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("请提供 agentID 作为参数，例如：agent myAgent123")
	}
	agentID := os.Args[1]
	serverURL := "http://localhost:8080" // 请根据实际情况修改

	// 假设 host 为本机名称或 IP
	host := "localhost"

	err := registerAgent(serverURL, agentID, host)
	if err != nil {
		log.Fatalf("注册 Agent 失败: %v", err)
	}
	log.Printf("Agent %s 注册成功", agentID)

	// 定时轮询获取命令
	pollCommands(serverURL, agentID)
}
