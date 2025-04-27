package exec_logic

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"go-devops-mimi/server/model/exec"
	execReq "go-devops-mimi/server/model/exec/request"
	"go-devops-mimi/server/model/exec/response"

	"go-devops-mimi/server/public/tools"

	"github.com/gin-gonic/gin"
)

type TransferLogic struct {
}

func (l TransferLogic) Add(c *gin.Context, req interface{}) (interface{}, interface{}) {
	// 1. 参数断言
	r, ok := req.(*execReq.TransferAddReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}

	// 2. 获取当前登录用户
	ctxUser, err := userService.GetCurrentLoginUser(c)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取当前登陆用户信息失败: %w", err))
	}

	// 3. 并发控制与日志汇总初始化
	var (
		wg            sync.WaitGroup
		mu            sync.Mutex
		logBuilder    strings.Builder     // 累积所有节点的日志
		overallStatus int8            = 1 // 默认 1=成功
	)

	// 4. 创建协程池：最少3，最多20，队列大小为节点数，空闲超时5秒
	pool := tools.NewPool(3, 20, len(r.NodesIDs), 5*time.Second)

	// 5. 遍历每个目标节点
	for _, nodeid := range r.NodesIDs {
		wg.Add(1)
		nID := nodeid // 捕获变量

		// 提交任务
		if err := pool.Exec(func() {
			defer wg.Done()

			// 5.1 获取节点信息
			info, err := nodeService.Info(nID)
			var taskStatus int8
			var output string

			if err != nil {
				taskStatus = 2
				output = fmt.Sprintf("节点ID %d 获取信息失败: %v\n", nID, err)
			} else {
				// 5.2 构造 SSH 配置
				sshConfig := &tools.SSHClientConfig{
					UserName:   info.Username,
					AuthModel:  info.AuthModel,
					Password:   tools.DecodeStrFromBase64(info.Password),
					PrivateKey: tools.DecodeStrFromBase64(info.PrivateKey),
					PublicIP:   info.PublicIP,
					Port:       info.SSHPort,
					Timeout:    time.Duration(info.Timeout) * time.Second,
				}
				// 5.3 执行文件上传
				remotePath, err := tools.UploadFileToHost(sshConfig, r.SourcePath, r.TargetPath)
				if err != nil {
					taskStatus = 2
					output = fmt.Sprintf("[%s] 上传失败: %v\n", info.NodeName, err)
				} else {
					taskStatus = 1
					output = fmt.Sprintf("[%s] 上传成功，目标路径：%s\n", info.NodeName, remotePath)
				}
			}

			// 5.4 如果有失败，整体状态置为 2
			if taskStatus == 2 {
				mu.Lock()
				overallStatus = 2
				mu.Unlock()
			}

			// 5.5 追加到统一日志
			mu.Lock()
			logBuilder.WriteString(output)
			mu.Unlock()
		}); err != nil {
			// 提交任务失败时也要减少计数
			wg.Done()
			mu.Lock()
			overallStatus = 2
			logBuilder.WriteString(fmt.Sprintf("节点ID %d 任务提交失败: %v\n", nID, err))
			mu.Unlock()
		}
	}

	// 6. 等待所有上传任务完成
	wg.Wait()

	// 7. 序列化节点 ID 列表
	nodeidsJSON, _ := json.Marshal(r.NodesIDs)

	// 8. 构造并保存总体 Transfer 记录
	transfer := exec.Transfer{
		Name:       fmt.Sprintf("%s-%s", r.Name, time.Now().Format("20060102-150405")),
		SourcePath: r.SourcePath,
		TargetPath: r.TargetPath,
		NodesIDs:   string(nodeidsJSON),
		Status:     overallStatus,
		RunLog:     logBuilder.String(),
		Creator:    ctxUser.Username,
	}
	if err := TransferService.Add(&transfer); err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("创建分发记录失败: %w", err))
	}

	return nil, nil
}

func (l TransferLogic) List(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*execReq.TransferListReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	// 获取数据列表
	list, err := TransferService.List(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据列表失败: %s", err.Error()))
	}

	rets := make([]exec.Transfer, 0)
	for _, nodes := range list {
		rets = append(rets, *nodes)
	}
	count, err := TransferService.ListCount(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据总数失败"))
	}

	return response.TransferRsp{
		Total:     count,
		Transfers: rets,
	}, nil
}

// Delete 删除数据
// func (l TransferLogic) Delete(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
// 	r, ok := req.(*execReq.TransferDeleteReq)
// 	if !ok {
// 		return nil, tools.ReqAssertErr
// 	}
// 	_ = c
// 	// 删除数据
// 	err := TransferService.Delete(r.Ids)
// 	if err != nil {
// 		return nil, tools.NewMySqlError(fmt.Errorf("删除数据失败: %s", err.Error()))
// 	}
// 	return nil, nil
// }

// / Info 查看数据详情
func (s TransferLogic) Info(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*execReq.TransferInfoReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	list, err := TransferService.Info(r.ID, r.Name)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据详情失败: %s", err.Error()))
	}
	return list, nil
}
