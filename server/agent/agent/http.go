package agent

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type AgentRegisterReq struct {
	Token     string `json:"token"`
	AgentID   uint64 `json:"agent_id"`
	AgentName string `json:"agent_name"`
	Version   string `json:"version"`
	OsType    string `json:"osType"`
	IP        string `json:"ip"`
	Label     string `json:"label"`
}

type AgentHeartbeatrReq struct {
	Token   string `json:"token"`
	AgentID uint64 `json:"agent_id"`
	Version string `json:"version"`
	OsType  string `json:"osType"`
}

type AgentResponse struct {
	Code int         `json:"code"` // 响应状态码，0 代表成功，其他代表错误
	Data interface{} `json:"data"` // 响应数据，当前 API 返回 null，可以用 interface{} 占位
	Msg  string      `json:"msg"`  // 响应消息，包含错误信息或 "success"
}

// httpPost 发送 POST 请求，并解析 JSON 响应
func httpPost[T any](data any, url string) (*T, error) {
	// 1. 序列化请求数据
	mjson, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// 2. 创建 HTTP 请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(mjson))
	if err != nil {
		return nil, err
	}

	// 3. 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 4. 使用超时控制 HTTP 客户端
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // 确保资源释放

	// 5. 检查 HTTP 响应码
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, errors.New("HTTP 请求失败，状态码：" + resp.Status)
	}

	// 6. 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 7. 解析 JSON 响应
	stru := new(T)
	if err := json.Unmarshal(body, stru); err != nil {
		return nil, err
	}

	return stru, nil
}
