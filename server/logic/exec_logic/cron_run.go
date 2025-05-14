package exec_logic

import (
	"encoding/json"
	"fmt"
	"go-devops-mimi/server/model/exec"
	"go-devops-mimi/server/public/tools"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

func SendCron(cron *exec.Cron) {
	switch cron.CronType {
	case "once":
		// 拷贝一份，避免闭包引用带来的潜在问题
		c := *cron
		sendTime := c.OnceTime.Add(2 * time.Second)
		tools.CronJob.AddAbsTimeJob(sendTime, func() error {
			if err := CronRun(&c); err != nil {
				logrus.Errorf("一次性任务 id=%d 执行失败：%v", c.ID, err)
				return err
			}
			logrus.Infof("一次性任务 id=%d 执行成功", c.ID)
			return nil
		})
		logrus.Infof("Cron(id=%d,type=once) 注册完毕，触发时间：%s", c.ID, sendTime.Format(time.DateTime))

	case "interval":
		c := *cron
		tools.CronJob.AddIntervalJob(c.Interval, func() error {
			if err := CronRun(&c); err != nil {
				logrus.Errorf("间隔任务 id=%d 执行失败：%v", c.ID, err)
				return err
			}
			logrus.Infof("间隔任务 id=%d 执行成功", c.ID)
			return nil
		})
		logrus.Infof("Cron(id=%d,type=interval) 注册完毕，每 %d 秒触发一次", c.ID, c.Interval)

	case "cron":
		c := *cron
		expr := c.Cronession
		if err := tools.CronJob.AddCronExprJob(expr, func() error {
			if err := CronRun(&c); err != nil {
				logrus.Errorf("Cron 表达式任务 id=%d 执行失败：%v", c.ID, err)
				return err
			}
			logrus.Infof("Cron 表达式任务 id=%d 执行成功", c.ID)
			return nil
		}); err != nil {
			logrus.Errorf("添加 Cron 表达式任务 id=%d 失败：%v", c.ID, err)
			return
		}
		logrus.Infof("Cron(id=%d,type=cron) 注册完毕，表达式：%s", c.ID, expr)

	default:
		logrus.Warnf("未知 CronType=%q，id=%d，跳过注册", cron.CronType, cron.ID)
	}
}

func CronRun(cron *exec.Cron) (err error) {
	startTime := time.Now()
	var (
		wg            sync.WaitGroup
		mu            sync.Mutex
		scriptLogs    []*exec.ScriptLog
		overallStatus int8    = 1 // 默认成功
		errs          []error     // 收集各节点错误
	)

	// 创建协程池：最少3，最多20，队列长度len(nodes)，空闲超时5秒
	raw := strings.Trim(cron.NodesIDs, "[]") // 去掉 "[" 和 "]"
	nodeStrs := strings.Split(raw, ",")      // 按逗号分隔
	pool := tools.NewPool(3, 20, 10, 5*time.Second)

	for _, nidStr := range nodeStrs {
		nid, convErr := strconv.Atoi(nidStr)
		if convErr != nil {
			logrus.Warnf("非法节点 ID %q，跳过", nidStr)
			continue
		}

		wg.Add(1)
		nID := nid

		// 使用协程池执行
		submitErr := pool.Exec(func() {
			defer wg.Done()

			// 获取节点信息
			list, infoErr := nodeService.Info(nID)
			var (
				taskStatus int8 = 1
				output     string
			)
			startNode := time.Now()

			if infoErr != nil {
				output = fmt.Sprintf("获取节点信息失败: %v", infoErr)
				taskStatus = 2
			} else {
				// SSH 执行
				sshCfg := &tools.SSHClientConfig{
					Timeout:    time.Second * time.Duration(5+cron.Timeout),
					UserName:   list.Username,
					AuthModel:  list.AuthModel,
					Password:   tools.DecodeStrFromBase64(list.Password),
					PrivateKey: tools.DecodeStrFromBase64(list.PrivateKey),
					Port:       list.SSHPort,
					PublicIP:   list.PublicIP,
				}

				if cron.CmdType == "command" {
					output, infoErr = tools.SshCommand(sshCfg, cron.Content)
				} else {
					output, infoErr = tools.CreateFileOnRemoteServer(sshCfg, cron.Name+"-"+list.NodeName, cron.Type, cron.Content)
				}
				if infoErr != nil {
					taskStatus = 2
					output = fmt.Sprintf("节点：%v,执行失败: %v", list.NodeName, infoErr)
				}
			}

			endNode := time.Now()
			nodeName := "unknown"
			if list != nil {
				nodeName = list.NodeName
			}

			nodeLog := &exec.ScriptLog{
				Name:      cron.Name + "-" + nodeName,
				NodeName:  nodeName,
				Type:      cron.Type,
				Content:   cron.Content,
				Status:    taskStatus,
				Timeout:   cron.Timeout,
				RunLog:    output,
				StartTime: startNode.Format(time.RFC3339),
				EndTime:   endNode.Format(time.RFC3339),
				TimeCost:  endNode.Sub(startNode).Milliseconds(),
			}

			// 收集日志与错误
			mu.Lock()
			defer mu.Unlock()

			scriptLogs = append(scriptLogs, nodeLog)
			if taskStatus == 2 {
				overallStatus = 2
			}
			if infoErr != nil {
				errs = append(errs, infoErr)
			}
		})

		if submitErr != nil {
			wg.Done()
			mu.Lock()
			scriptLogs = append(scriptLogs, &exec.ScriptLog{
				Name:      cron.Name + "-unknown",
				NodeName:  "unknown",
				Type:      cron.Type,
				Content:   cron.Content,
				Status:    2,
				Timeout:   cron.Timeout,
				RunLog:    fmt.Sprintf("任务提交失败: %v", submitErr),
				StartTime: time.Now().Format(time.RFC3339),
				EndTime:   time.Now().Format(time.RFC3339),
				TimeCost:  0,
			})
			overallStatus = 2
			errs = append(errs, submitErr)
			mu.Unlock()
		}
	}

	wg.Wait()

	// 组装总体脚本记录
	nodesJSON, _ := json.Marshal(cron.NodesIDs)
	list := exec.CronLog{
		Name:       cron.Name + "-" + startTime.Format("20060102-15:04:05"),
		NodesIDs:   string(nodesJSON),
		Status:     overallStatus,
		Desc:       cron.Desc,
		CronType:   cron.CronType,
		Cronession: cron.Cronession,
		Interval:   cron.Interval,
		OnceTime:   cron.OnceTime,
		StartTime:  startTime.Format(time.RFC3339),
		EndTime:    time.Now().Format(time.RFC3339),
		TimeCost:   time.Now().Sub(startTime).Milliseconds(),
		ScriptLogs: scriptLogs,
	}

	// 持久化
	if persistErr := CronLogService.Add(&list); persistErr != nil {
		// 返回保存错误为主要错误
		return fmt.Errorf("保存脚本记录失败: %w", persistErr)
	}

	// 如果有节点执行出错，则返回合并后的错误
	if len(errs) > 0 {
		return fmt.Errorf("部分节点执行失败: %v", errs)
	}

	return nil
}
