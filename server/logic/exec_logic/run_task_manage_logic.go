package exec_logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"gorm.io/gorm"

	"go-devops-mimi/server/model/exec"
	execReq "go-devops-mimi/server/model/exec/request"
	"go-devops-mimi/server/model/exec/response"

	"go-devops-mimi/server/public/common"
	"go-devops-mimi/server/public/tools"

	"github.com/gin-gonic/gin"
)

// 全局 map：runID -> 输出 channel
var runChans = struct {
	sync.RWMutex
	m map[uint]chan []byte
}{m: make(map[uint]chan []byte)}

// 获取或创建一个 channel
func getOrCreateRunChan(runID uint) chan []byte {
	runChans.RLock()
	ch, ok := runChans.m[runID]
	runChans.RUnlock()
	if ok {
		return ch
	}
	runChans.Lock()
	defer runChans.Unlock()
	ch = make(chan []byte, 100)
	runChans.m[runID] = ch
	return ch
}

// 执行完毕后，关闭并删除 channel
func closeRunChan(runID uint) {
	runChans.Lock()
	if ch, ok := runChans.m[runID]; ok {
		close(ch)
		delete(runChans.m, runID)
	}
	runChans.Unlock()
}

// multiChanWriter 同时把 Write 的内容发到多个 chan []byte
type multiChanWriter struct {
	chans []chan []byte
}

func (w *multiChanWriter) Write(p []byte) (n int, err error) {
	// 按需防止阻塞、丢弃或缓冲
	for _, ch := range w.chans {
		select {
		case ch <- append([]byte{}, p...): // 复制一份，避免竞态
		default:
		}
	}
	return len(p), nil
}

// Run 触发任务执行，立即返回 runID
func (l *ManageLogic) Run(c *gin.Context, req *execReq.TaskManageRunReq) (interface{}, interface{}) {
	// 1. 创建 ManageLog 记录
	runName := req.Name + "-" + time.Now().Format("20060102-150405")
	mlog := &exec.ManageLog{
		Name:      runName,
		Args:      fmt.Sprintf("%v", req.EnvParams),
		Desc:      req.Desc,
		Status:    "running",
		StartTime: time.Now().Format(time.RFC3339),
	}
	if err := ManageLogService.Add(mlog); err != nil {
		return nil, gin.H{"error": err.Error()}
	}
	runID := mlog.ID

	// 2. 获取 channel 并启动后台执行
	ch := getOrCreateRunChan(runID)
	go func() {
		defer closeRunChan(runID)
		l.doRun(c.Request.Context(), runID, req, ch)
	}()

	// 3. 返回 runID 给客户端
	return gin.H{"runID": runID}, nil
}

// doRun 真正的执行逻辑：并发跑子任务，所有输出写到 runChan 和各自 subChan
func (l *ManageLogic) doRun(ctx context.Context, runID uint, req *execReq.TaskManageRunReq, runChan chan []byte) {
	// 0. 读取任务清单
	list, err := TaskManageService.Info(req.ID, req.Name)
	if err != nil {
		runChan <- []byte(fmt.Sprintf("获取任务失败: %v\n", err))
		return
	}

	envRaw := req.EnvParams
	var wg sync.WaitGroup
	pool := tools.NewPool(3, 20, 10, 5*time.Second)

	sort.Slice(list.Tasks, func(i, j int) bool {
		return list.Tasks[i].Sort < list.Tasks[j].Sort
	})

	// 保存所有子任务组 ID，结束时统一关闭
	var subtaskIDs []uint

	for _, task := range list.Tasks {
		taskCopy := task

		// 创建子任务组记录
		subGroup := &exec.TaskManageLog{
			Name:      taskCopy.Name,
			Status:    "running",
			StartTime: time.Now().Format(time.RFC3339),
		}
		if err := ManageLogService.AddTaskManageLog(subGroup); err != nil {
			runChan <- []byte(fmt.Sprintf("子任务 %s 记录失败: %v\n", taskCopy.Name, err))
			continue
		}
		ManageLogService.AddManageToTaskLog(&exec.ManageLog{Model: gorm.Model{ID: runID}}, subGroup)

		// 为子任务组创建独立通道
		subChan := getOrCreateRunChan(subGroup.ID)
		subtaskIDs = append(subtaskIDs, subGroup.ID)

		// 并发执行此子任务组的所有节点
		for _, nidStr := range strings.Split(taskCopy.NodesIDs, ",") {
			nid, err := strconv.Atoi(nidStr)
			if err != nil {
				msg := fmt.Sprintf("[%s] 非法节点 ID: %q\n", taskCopy.Name, nidStr)
				runChan <- []byte(msg)
				subChan <- []byte(msg)
				continue
			}

			wg.Add(1)
			pool.Exec(func() {
				defer wg.Done()

				// 1) 获取节点信息
				info, err := nodeService.Info(nid)
				if err != nil {
					msg := fmt.Sprintf("节点 %d 信息获取失败: %v\n", nid, err)
					runChan <- []byte(msg)
					subChan <- []byte(msg)
					return
				}

				// 2) 写 TaskLog 记录
				taskLog := &exec.TaskLog{
					Name:      taskCopy.Name,
					Content:   taskCopy.Content,
					Sort:      taskCopy.Sort,
					NodeName:  info.NodeName,
					Status:    "running",
					StartTime: time.Now().Format(time.RFC3339),
				}
				if err := TaskLogService.Add(taskLog); err != nil {
					msg := fmt.Sprintf("[%s] 写日志失败: %v\n", info.NodeName, err)
					runChan <- []byte(msg)
					subChan <- []byte(msg)
					return
				}
				TaskLogService.AddTaskLogToManage(subGroup, taskLog)

				// 3) 通过 SSH 执行脚本
				sshCfg := &tools.SSHClientConfig{
					Timeout:    5*time.Second + time.Duration(taskCopy.Timeout)*time.Second,
					UserName:   info.Username,
					AuthModel:  info.AuthModel,
					Password:   tools.DecodeStrFromBase64(info.Password),
					PrivateKey: tools.DecodeStrFromBase64(info.PrivateKey),
					PublicIP:   info.PublicIP,
					Port:       info.SSHPort,
				}
				client, err := tools.NewSSHClient(sshCfg)
				if err != nil {
					msg := fmt.Sprintf("[%s] SSH 连接失败: %v\n", info.NodeName, err)
					runChan <- []byte(msg)
					subChan <- []byte(msg)
					return
				}
				defer client.Close()

				// 新增 fanChan 扇出逻辑
				fanChan := make(chan []byte, 100)
				fanWriter := tools.NewChanWriter(fanChan)
				go func() {
					for msg := range fanChan {
						runChan <- msg
						subChan <- msg
					}
				}()

				// LogWriter 收集完整日志
				var buf bytes.Buffer
				turn, err := tools.NewTurn(nil, client, fanWriter, info.NodeName)
				if err != nil {
					msg := fmt.Sprintf("[%s@%s] 会话创建失败: %v\n", taskCopy.Name, info.NodeName, err)
					runChan <- []byte(msg)
					subChan <- []byte(msg)
					return
				}
				defer turn.Close()
				turn.LogWriter = &buf

				// 取消回显、启动读取
				turn.StdinPipe.Write([]byte("stty -echo; export PS1='' PS2=''\n"))
				go turn.LoopRead(ctx)

				// 5) 写入脚本并等待
				script := taskCopy.Content + "\nexit\n"
				if len(envRaw) > 0 {
					script = tools.BuildScriptWithEnv(envRaw, taskCopy.Content)
				}
				turn.StdinPipe.Write([]byte(script))

				start := time.Now()
				if err := turn.SessionWait(); err != nil {
					taskLog.Status = "failed"
					msg := fmt.Sprintf("[%s@%s] 执行异常: %v\n", taskCopy.Name, info.NodeName, err)
					runChan <- []byte(msg)
					subChan <- []byte(msg)
				} else {
					taskLog.Status = "success"
				}

				// 6) 保存 TaskLog 中的 RunLog
				taskLog.RunLog = buf.String()
				taskLog.EndTime = time.Now().Format(time.RFC3339)
				taskLog.TimeCost = time.Since(start).Milliseconds()
				TaskLogService.Save(taskLog)
			})
		}
	}

	// 等待所有节点完成
	wg.Wait()

	// 7) 推送“全部执行完”消息并关闭子任务组通道
	doneMsg := []byte("🎉 全部节点执行完毕\n")
	for _, id := range subtaskIDs {
		if c := getOrCreateRunChan(id); c != nil {
			c <- doneMsg
			closeRunChan(id)
		}
	}

	// 8) 推送全局完成
	runChan <- doneMsg

	// 9) 最后更新所有状态
	if err := ManageLogService.UpdateAllStatusesByManageLogID(runID); err != nil {
		runChan <- []byte(fmt.Sprintf("更新任务组状态失败: %v\n", err))
	}
}

// RunInfoWebSocket 客户端 connect WS，并传 runID，服务器推送对应 channel 或历史日志
func (l *ManageLogic) RunInfo(c *gin.Context, req *execReq.TaskManageRunInfoReq) (interface{}, interface{}) {
	list, err := TaskManageService.InfoManageLog(req.RunID)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取执行详情失败: %s", err.Error()))
	}
	return list, nil
}
func (l ManageLogic) RunList(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*execReq.ManageLogListReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	// 获取数据列表
	list, err := ManageLogService.List(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据列表失败: %s", err.Error()))
	}

	rets := make([]exec.ManageLog, 0)
	for _, nodes := range list {
		rets = append(rets, *nodes)
	}
	count, err := ManageLogService.ListCount(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据总数失败"))
	}

	return response.ManageLogListRsp{
		Total:      count,
		ManageLogs: rets,
	}, nil
}

// func (l *ManageLogic) RunInfoWebSocket(c *gin.Context, ws *websocket.Conn, req *execReq.TaskManageRunInfoWebsocketReq) (interface{}, interface{}) {
// 	writer := tools.NewWSWriter(ws)
// 	defer writer.Close()

// 	// 1. 优先从内存里按 subtask ID 拿 channel
// 	// runChans.RLock()
// 	// ch, exists := runChans.m[req.TaskID]
// 	// runChans.RUnlock()
// 	// if exists {
// 	// 	// 实时推送
// 	// 	for msg := range ch {
// 	// 		writer.Send(msg)
// 	// 	}
// 	// 	return nil, nil
// 	// }

// 	// 2. 内存里没有，再走历史回放，只查这个子任务组
// 	// 查到 TaskLog 里所有日志，再按行拆出来推送
// 	var tlogs []exec.TaskLog
// 	if err := common.DB.
// 		Joins("JOIN t_task_log_s tls ON tls.task_log_id = t_task_log.id").
// 		Where("tls.task_manage_log_id = ?", req.TaskID).
// 		Order("t_task_log.start_time").
// 		Find(&tlogs).Error; err != nil {
// 		writer.Send([]byte(fmt.Sprintf("读取执行日志失败: %v\n", err)))
// 		return nil, nil
// 	}

// 	for _, tl := range tlogs {
// 		for _, line := range strings.Split(tl.RunLog, "\n") {
// 			if line == "" {
// 				continue
// 			}
// 			msgObj := map[string]string{
// 				"task": tl.Name,     // 子任务组名称
// 				"node": tl.NodeName, // 节点
// 				"data": line + "\n",
// 			}
// 			b, _ := json.Marshal(msgObj)
// 			writer.Send(b)
// 		}
// 	}

// 	writer.Send([]byte("🎉 日志回放结束\n"))
// 	return nil, nil
// }

func (l *ManageLogic) RunInfoWebSocket(c *gin.Context, ws *websocket.Conn, req *execReq.TaskManageRunInfoWebsocketReq) (interface{}, interface{}) {
	writer := tools.NewWSWriter(ws)
	defer writer.Close()
	// 1. 优先从内存里按 subtask ID 拿 channel
	runChans.RLock()
	ch, exists := runChans.m[req.TaskID]
	runChans.RUnlock()
	if exists {
		// 实时推送
		for msg := range ch {
			writer.Send(msg)
		}
		return nil, nil
	}
	// 1. 查询并重放历史日志（和你原来的一样）
	var tlogs []exec.TaskLog
	if err := common.DB.
		Joins("JOIN t_task_log_s tls ON tls.task_log_id = t_task_log.id").
		Where("tls.task_manage_log_id = ?", req.TaskID).
		Order("t_task_log.start_time").
		Find(&tlogs).Error; err != nil {
		writer.Send([]byte(fmt.Sprintf("读取执行日志失败: %v\n", err)))
		return nil, nil
	}

	for _, tl := range tlogs {
		clean := strings.ReplaceAll(strings.ReplaceAll(tl.RunLog, "\r\n", "\n"), "\r", "")
		for _, line := range strings.Split(clean, "\n") {
			if line == "" {
				continue
			}
			msgObj := map[string]string{
				"task": tl.Name,
				"node": tl.NodeName,
				"data": line + "\n",
			}
			b, _ := json.Marshal(msgObj)
			writer.Send(b) // 建议这里用阻塞式 Send
		}
	}

	// 2. 历史回放结束标记
	writer.Send([]byte("🎉 日志回放结束\n"))

	// 3. 等待客户端断开或 60 秒超时
	select {
	case <-c.Request.Context().Done():
		log.Printf("TaskID=%d: 客户端主动断开 WebSocket\n", req.TaskID)
	case <-time.After(10 * time.Second):
		// log.Printf("TaskID=%d: 等待超时 60s，自动结束 WebSocket\n", req.TaskID)
	}

	// 客户端一旦关闭连接，Done() 会解除阻塞，继续执行 defer writer.Close()
	return nil, nil
}
