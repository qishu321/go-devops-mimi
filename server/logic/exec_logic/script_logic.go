package exec_logic

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"go-devops-mimi/server/model/exec"
	execReq "go-devops-mimi/server/model/exec/request"
	"go-devops-mimi/server/model/exec/response"

	"go-devops-mimi/server/public/tools"

	"github.com/gin-gonic/gin"
)

type ScriptLogic struct {
}

func (l ScriptLogic) Add_Run(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*execReq.ScriptCmdRunReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}

	// 获取当前用户
	ctxUser, err := userService.GetCurrentLoginUser(c)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取当前登陆用户信息失败"))
	}

	startTime := time.Now()
	var (
		wg            sync.WaitGroup
		mu            sync.Mutex
		scriptLogs    []*exec.ScriptLog
		overallStatus int8 = 1 // 默认成功状态
	)

	// 创建协程池：最少 3 个，最多 20 个，任务队列大小设置为足够大（例如：len(r.NodesIDs)），空闲超时可设置为 5 秒
	pool := tools.NewPool(3, 20, len(r.NodesIDs), 5*time.Second)

	for _, nodeid := range r.NodesIDs {
		wg.Add(1)
		// 捕获 nodeid 的副本
		nID := nodeid

		// 使用协程池提交任务
		err := pool.Exec(func() {
			defer wg.Done()

			// 记录任务起始时间
			startNode := time.Now()

			// 获取节点信息
			list, err := nodeService.Info(nID)
			var taskStatus int8
			var output string
			if err != nil {
				output = fmt.Sprintf("获取节点信息失败: %v", err)
				taskStatus = 2
			} else {
				// 构造 SSH 配置
				sshConfig := &tools.SSHClientConfig{
					Timeout:    time.Second * time.Duration(5+r.Timeout),
					UserName:   list.Username,
					AuthModel:  list.AuthModel,
					Password:   tools.DecodeStrFromBase64(list.Password),
					PrivateKey: tools.DecodeStrFromBase64(list.PrivateKey),
					Port:       list.SSHPort,
					PublicIP:   list.PublicIP,
				}
				// 根据命令类型执行任务
				if r.CmdType == "command" {
					output, err = tools.SshCommand(sshConfig, r.Command)
				} else {
					output, err = tools.CreateFileOnRemoteServer(sshConfig, r.Name+"-"+list.NodeName, r.Type, r.Command)
				}
				if err != nil {
					taskStatus = 2
					output = fmt.Sprintf("执行失败: %v", err)
				} else {
					taskStatus = 1
				}
			}

			endNode := time.Now()
			nodeLog := &exec.ScriptLog{
				Name: r.Name + "-" + func() string {
					if list != nil {
						return list.NodeName
					}
					return "unknown"
				}(),
				NodeName: func() string {
					if list != nil {
						return list.NodeName
					}
					return "unknown"
				}(),
				Type:      r.Type,
				Content:   r.Command,
				Status:    taskStatus,
				Timeout:   r.Timeout,
				RunLog:    output,
				StartTime: startNode.Format(time.RFC3339),
				EndTime:   endNode.Format(time.RFC3339),
				TimeCost:  endNode.Sub(startNode).Milliseconds(),
			}

			// 如果任何一个任务失败，则总体状态设为失败（2）
			if taskStatus == 2 {
				mu.Lock()
				overallStatus = 2
				mu.Unlock()
			}

			// 收集任务日志
			mu.Lock()
			scriptLogs = append(scriptLogs, nodeLog)
			mu.Unlock()
		})

		// 如果任务提交失败（例如任务队列满等情况），则直接记录错误并调用 Done
		if err != nil {
			wg.Done()
			mu.Lock()
			overallStatus = 2
			scriptLogs = append(scriptLogs, &exec.ScriptLog{
				Name:      r.Name + "-unknown",
				NodeName:  "unknown",
				Type:      r.Type,
				Content:   r.Command,
				Status:    2,
				Timeout:   r.Timeout,
				RunLog:    fmt.Sprintf("任务提交失败: %v", err),
				StartTime: time.Now().Format(time.RFC3339),
				EndTime:   time.Now().Format(time.RFC3339),
				TimeCost:  0,
			})
			mu.Unlock()
		}
	}

	wg.Wait() // 等待所有节点任务完成
	nodeids, _ := json.Marshal(r.NodesIDs)
	// 记录总体执行信息
	script := exec.Script{
		Name:      r.Name + "-" + time.Now().Format("20060102-15:04:05"),
		NodesIDs:  string(nodeids),
		Status:    overallStatus,
		CmdType:   r.CmdType,
		StartTime: startTime.Format(time.RFC3339),
		EndTime:   time.Now().Format(time.RFC3339),
		TimeCost:  time.Now().Sub(startTime).Milliseconds(),
		Creator:   ctxUser.Username,
		Scripts:   scriptLogs,
		Desc:      r.Desc,
	}

	// 持久化总体记录
	err = ScriptService.Add(&script)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("创建数据失败: %s", err.Error()))
	}
	return script, nil
}

func (l ScriptLogic) List(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*execReq.ScriptListReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	// 获取数据列表
	list, err := ScriptService.List(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据列表失败: %s", err.Error()))
	}

	rets := make([]exec.Script, 0)
	for _, nodes := range list {
		rets = append(rets, *nodes)
	}
	count, err := ScriptService.ListCount(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据总数失败"))
	}

	return response.ScriptListRsp{
		Total:   count,
		Scripts: rets,
	}, nil
}
