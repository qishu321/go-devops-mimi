package exec_logic

import (
	"bytes"
	"context"
	"fmt"
	"go-devops-mimi/server/model/exec"
	execReq "go-devops-mimi/server/model/exec/request"
	"go-devops-mimi/server/public/tools"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Run 在同一个 RunWebSocket 连接上，为给定节点 ID 列表并发执行 SSH 会话，
// 将所有 stdout/stderr 实时推送给客户端。
//
//	pool := tools.NewPool(3, 30, 10, 5*time.Second)
//	ctx:    上下文用于控制生命周期
//	ws:     已升级的 WebSocket 连接
//	nodeIDs: 逗号分隔的节点 ID，例如 "1,2,3"

func (l ManageLogic) RunWebSocket(c *gin.Context, ws *websocket.Conn, req *execReq.TaskManageRunReq) (interface{}, interface{}) {
	// 0. 读取任务
	list, err := TaskManageService.Info(req.ID, req.Name)
	if err != nil {
		writer := tools.NewWSWriter(ws)
		defer writer.Close()
		writer.Send([]byte(fmt.Sprintf("获取任务失败: %v\n", err)))
		return nil, nil
	}

	// 1. env_task_s 参数可能是 map[string]string，也可能是 []string
	//    只有当它是 map 并且非空时，才需要注入到脚本里
	envRaw := req.EnvParams
	// 1. 初始化任务组日志
	runName := req.Name + "-" + time.Now().Format("20060102-15:04:05")
	groupLog := &exec.ManageLog{
		Name:      runName,
		Args:      fmt.Sprintf("%v", envRaw),
		Desc:      list.Desc,
		Status:    "running",
		StartTime: time.Now().Format(time.RFC3339),
	}
	if err := ManageLogService.Add(groupLog); err != nil {
		writer := tools.NewWSWriter(ws)
		defer writer.Close()
		writer.Send([]byte(fmt.Sprintf("写入任务组日志失败: %v\n", err)))
		return nil, nil
	}

	// 2. 启动单写协程
	writer := tools.NewWSWriter(ws)
	defer writer.Close()

	// 3. 排序
	sort.Slice(list.Tasks, func(i, j int) bool {
		return list.Tasks[i].Sort < list.Tasks[j].Sort
	})

	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()
	var wg sync.WaitGroup
	pool := tools.NewPool(3, 30, 10, 5*time.Second)

	// 4. 并发执行每个子任务的每台节点
	for _, task := range list.Tasks {
		// 2.1 为每个子任务创建 TaskManageLog
		subGroup := &exec.TaskManageLog{
			Name:      task.Name,
			Status:    "running",
			StartTime: time.Now().Format(time.RFC3339),
		}
		if err := ManageLogService.AddTaskManageLog(subGroup); err != nil {
			writer.Send([]byte(fmt.Sprintf("子任务 %s 记录失败: %v\n", task.Name, err)))
			continue
		}
		// 关联到任务组
		if err := ManageLogService.AddManageToTaskLog(groupLog, subGroup); err != nil {
			writer.Send([]byte(fmt.Sprintf("关联子任务 %s 失败: %v\n", task.Name, err)))
		}
		for _, nidStr := range strings.Split(task.NodesIDs, ",") {
			nid, err := strconv.Atoi(nidStr)
			if err != nil {
				writer.Send([]byte(fmt.Sprintf("非法节点 ID: %q，略过\n", nidStr)))
				continue
			}
			wg.Add(1)
			pool.Exec(func() {
				defer wg.Done()

				// A. 节点信息
				info, err := nodeService.Info(nid)
				if err != nil {
					writer.Send([]byte(fmt.Sprintf("节点 %d 信息获取失败: %v\n", nid, err)))
					return
				}
				// B. 插入子任务日志 (running)
				taskLog := &exec.TaskLog{
					Name:      task.Name,
					Content:   task.Content,
					Sort:      task.Sort,
					NodeName:  info.NodeName,
					Status:    "running",
					StartTime: time.Now().Format(time.RFC3339),
				}
				if err := TaskLogService.Add(taskLog); err != nil {
					writer.Send([]byte(fmt.Sprintf("[%s] 写入子任务日志失败: %v\n", info.NodeName, err)))
					return
				}
				if err := TaskLogService.AddTaskLogToManage(subGroup, taskLog); err != nil {
					writer.Send([]byte(fmt.Sprintf("[%s] 关联子任务日志失败: %v\n", info.NodeName, err)))
					return
				}
				sshCfg := &tools.SSHClientConfig{
					Timeout:    time.Second * time.Duration(5+task.Timeout),
					UserName:   info.Username,
					AuthModel:  info.AuthModel,
					Password:   tools.DecodeStrFromBase64(info.Password),
					PrivateKey: tools.DecodeStrFromBase64(info.PrivateKey),
					PublicIP:   info.PublicIP,
					Port:       info.SSHPort,
				}
				client, err := tools.NewSSHClient(sshCfg)
				if err != nil {
					writer.Send([]byte(fmt.Sprintf("[%s] SSH 连接失败: %v\n", info.NodeName, err)))
					return
				}
				defer client.Close()

				// B. 启动 Turn
				var buf bytes.Buffer
				turn, err := tools.NewTurn(ws, client, writer, info.NodeName)
				if err != nil {
					writer.Send([]byte(fmt.Sprintf("[%s] 会话创建失败: %v\n", info.NodeName, err)))
					return
				}
				defer turn.Close()
				turn.StdinPipe.Write([]byte("stty -echo; export PS1='' PS2=''\n"))
				// 关键——绑定日志缓冲
				turn.LogWriter = &buf
				// C. 并发读取前端输入
				go func(t *tools.Turn) {
					t.LoopRead(ctx)
					cancel()
				}(turn)
				// 执行完之后，buf.String() 就是完整的 stdout+stderr
				taskLog.RunLog = buf.String()

				// D. 构造要执行的脚本

				var script string
				if len(envRaw) > 0 {
					// 有参数，注入关联数组
					script = tools.BuildScriptWithEnv(envRaw, task.Content)
				} else {
					// 空 map，忽略注入
					script = task.Content + "\nexit\n"
				}
				// E. 写入并执行
				if _, err := turn.StdinPipe.Write([]byte(script)); err != nil {
					writer.Send([]byte(fmt.Sprintf("[%s] 脚本注入失败: %v\n", info.NodeName, err)))
					return
				}
				start := time.Now()
				if err := turn.SessionWait(); err != nil {
					writer.Send([]byte(fmt.Sprintf("[%s] 脚本执行异常: %v\n", info.NodeName, err)))
					taskLog.Status = "failed"
				} else {
					taskLog.Status = "success"
				}
				// E. 更新子任务日志
				duration := time.Since(start)
				taskLog.RunLog = buf.String()
				taskLog.EndTime = time.Now().Format(time.RFC3339)
				taskLog.TimeCost = duration.Milliseconds()
				if err := TaskLogService.Save(taskLog); err != nil {
					writer.Send([]byte(fmt.Sprintf("[%s] 更新子任务日志失败: %v\n", info.NodeName, err)))
				}
			})
		}
	}

	// 5. 等待所有任务完成
	wg.Wait()
	writer.Send([]byte("🎉 全部节点执行完毕\n"))
	// 4. 更新子任务组状态（根据其所有 TaskLog）
	if err := ManageLogService.UpdateAllStatusesByManageLogID(groupLog.ID); err != nil {
		writer.Send([]byte(fmt.Sprintf("更新任务组状态失败: %v", err)))
	}
	return nil, nil
}

// func (l ManageLogic) RunWebSocket(c *gin.Context, ws *websocket.Conn, req *execReq.TaskManageRunReq) (interface{}, interface{}) {
// 	// 0. 读取任务
// 	list, err := TaskManageService.Info(req.ID, req.Name)
// 	if err != nil {
// 		writer := tools.NewWSWriter(ws)
// 		defer writer.Close()
// 		writer.Send([]byte(fmt.Sprintf("获取任务失败: %v\n", err)))
// 		return nil, nil
// 	}

// 	// 1. env_task_s 参数可能是 map[string]string，也可能是 []string
// 	//    只有当它是 map 并且非空时，才需要注入到脚本里
// 	envRaw := req.EnvParams

// 	// 2. 启动单写协程
// 	writer := tools.NewWSWriter(ws)
// 	defer writer.Close()

// 	// 3. 排序
// 	sort.Slice(list.Tasks, func(i, j int) bool {
// 		return list.Tasks[i].Sort < list.Tasks[j].Sort
// 	})

// 	ctx, cancel := context.WithCancel(c.Request.Context())
// 	defer cancel()
// 	var wg sync.WaitGroup
// 	pool := tools.NewPool(3, 30, 10, 5*time.Second)

// 	// 4. 并发执行每个子任务的每台节点
// 	for _, task := range list.Tasks {
// 		for _, nidStr := range strings.Split(task.NodesIDs, ",") {
// 			nid, err := strconv.Atoi(nidStr)
// 			if err != nil {
// 				writer.Send([]byte(fmt.Sprintf("非法节点 ID: %q，略过\n", nidStr)))
// 				continue
// 			}
// 			wg.Add(1)
// 			pool.Exec(func() {
// 				defer wg.Done()

// 				// A. 节点信息
// 				info, err := nodeService.Info(nid)
// 				if err != nil {
// 					writer.Send([]byte(fmt.Sprintf("节点 %d 信息获取失败: %v\n", nid, err)))
// 					return
// 				}
// 				sshCfg := &tools.SSHClientConfig{
// 					Timeout:    time.Second * time.Duration(5+task.Timeout),
// 					UserName:   info.Username,
// 					AuthModel:  info.AuthModel,
// 					Password:   tools.DecodeStrFromBase64(info.Password),
// 					PrivateKey: tools.DecodeStrFromBase64(info.PrivateKey),
// 					PublicIP:   info.PublicIP,
// 					Port:       info.SSHPort,
// 				}
// 				client, err := tools.NewSSHClient(sshCfg)
// 				if err != nil {
// 					writer.Send([]byte(fmt.Sprintf("[%s] SSH 连接失败: %v\n", info.NodeName, err)))
// 					return
// 				}
// 				defer client.Close()

// 				// B. 启动 Turn
// 				turn, err := tools.NewTurn(ws, client, writer, info.NodeName)
// 				if err != nil {
// 					writer.Send([]byte(fmt.Sprintf("[%s] 会话创建失败: %v\n", info.NodeName, err)))
// 					return
// 				}
// 				defer turn.Close()

// 				// C. 并发读取前端输入
// 				go func(t *tools.Turn) {
// 					t.LoopRead(ctx)
// 					cancel()
// 				}(turn)

// 				// D. 构造要执行的脚本
// 				var script string
// 				if len(envRaw) > 0 {
// 					// 有参数，注入关联数组
// 					script = tools.BuildScriptWithEnv(envRaw, task.Content)
// 				} else {
// 					// 空 map，忽略注入
// 					script = task.Content + "\nexit\n"
// 				}
// 				// E. 写入并执行
// 				if _, err := turn.StdinPipe.Write([]byte(script)); err != nil {
// 					writer.Send([]byte(fmt.Sprintf("[%s] 脚本注入失败: %v\n", info.NodeName, err)))
// 					return
// 				}
// 				if err := turn.SessionWait(); err != nil {
// 					writer.Send([]byte(fmt.Sprintf("[%s] 脚本执行异常: %v\n", info.NodeName, err)))
// 				}
// 			})
// 		}
// 	}

// 	// 5. 等待所有任务完成
// 	wg.Wait()
// 	writer.Send([]byte("🎉 全部节点执行完毕\n"))
// 	return nil, nil
// }
