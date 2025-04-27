package exec_logic

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	execReq "go-devops-mimi/server/model/exec/request"

	"go-devops-mimi/server/public/tools"

	"github.com/gin-gonic/gin"
)

// Run 在同一个 RunWebSocket 连接上，为给定节点 ID 列表并发执行 SSH 会话，
// 将所有 stdout/stderr 实时推送给客户端。
//
//	pool := tools.NewPool(3, 20, 10, 5*time.Second)
//	ctx:    上下文用于控制生命周期
//	ws:     已升级的 WebSocket 连接
//	nodeIDs: 逗号分隔的节点 ID，例如 "1,2,3"
func (l TaskManageLogic) RunWebSocket(c *gin.Context, ws *websocket.Conn, req interface{}) (interface{}, interface{}) {
	// 类型断言并获取任务列表
	r, ok := req.(*execReq.TaskManageInfoReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	list, err := TaskManageService.Info(r.ID, r.Name)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据详情失败: %w", err))
	}

	// —— 1. 按 Sort 排序子任务列表 —— (从小到大)
	sort.Slice(list.Tasks, func(i, j int) bool {
		return list.Tasks[i].Sort < list.Tasks[j].Sort
	}) // Go1.8+ 支持 sort.Slice&#8203;:contentReference[oaicite:3]{index=3}

	// —— 2. 创建可取消上下文 & WaitGroup ——
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()
	var wg sync.WaitGroup

	// 协程池（可选，根据实际需求调整大小）
	pool := tools.NewPool(3, 20, 10, 5*time.Second)

	// —— 3. 遍历已排序的子任务和节点 ——
	for _, task := range list.Tasks {
		nodeIDs := strings.Split(task.NodesIDs, ",")
		for _, nidStr := range nodeIDs {
			nID, err := strconv.Atoi(nidStr)
			if err != nil {
				ws.WriteMessage(websocket.TextMessage,
					[]byte(fmt.Sprintf("非法节点 ID: %q，略过\n", nidStr)))
				continue
			}
			wg.Add(1)
			if err := pool.Exec(func() {
				defer wg.Done()

				// 获取节点信息
				info, err := nodeService.Info(nID)
				if err != nil {
					ws.WriteMessage(websocket.TextMessage,
						[]byte(fmt.Sprintf("节点 %d 信息获取失败: %v\n", nID, err)))
					return
				}

				// 建立 SSH 客户端
				sshCfg := &tools.SSHClientConfig{
					Timeout:    time.Second * time.Duration(5+task.Timeout),
					UserName:   info.Username,
					AuthModel:  info.AuthModel,
					Password:   tools.DecodeStrFromBase64(info.Password),
					PrivateKey: tools.DecodeStrFromBase64(info.PrivateKey),
					Port:       info.SSHPort,
					PublicIP:   info.PublicIP,
				}
				client, err := tools.NewSSHClient(sshCfg)
				if err != nil {
					ws.WriteMessage(websocket.TextMessage,
						[]byte(fmt.Sprintf("[%s] SSH 连接失败: %v\n", info.NodeName, err)))
					return
				}
				defer client.Close()

				// 创建 Turn 会话（启动一个远端 Shell）
				turn, err := tools.NewTurn(ws, client)
				if err != nil {
					ws.WriteMessage(websocket.TextMessage,
						[]byte(fmt.Sprintf("[%s] 会话创建失败: %v\n", info.NodeName, err)))
					return
				}
				defer turn.Close()

				// 并发读取前端输入，可中途 cancel
				go func(t *tools.Turn) {
					t.LoopRead(ctx)
					cancel()
				}(turn)

				// —— 4. 注入脚本内容并执行 ——
				// 写入用户定义的脚本内容（务必以换行结束）
				turn.StdinPipe.Write([]byte(task.Content + "\n")) // 多条命令可按需拼接&#8203;:contentReference[oaicite:4]{index=4}
				// 发送 exit 让 Shell 会话结束
				turn.StdinPipe.Write([]byte("exit\n")) // 退出远端 shell&#8203;:contentReference[oaicite:5]{index=5}

				// 等待脚本执行完成
				if err := turn.SessionWait(); err != nil {
					ws.WriteMessage(websocket.TextMessage,
						[]byte(fmt.Sprintf("[%s] 脚本执行异常: %v\n", info.NodeName, err)))
				}
			}); err != nil {
				// 提交协程池失败
				wg.Done()
				ws.WriteMessage(websocket.TextMessage,
					[]byte(fmt.Sprintf("任务提交失败: %v\n", err)))
			}
		}
	}

	// —— 5. 等待所有节点完成 ——
	wg.Wait()
	// 全部完成通知
	ws.WriteMessage(websocket.TextMessage, []byte("🎉 全部节点执行完毕\n"))
	return nil, nil
}
