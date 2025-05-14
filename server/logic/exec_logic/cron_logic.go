package exec_logic

import (
	"encoding/json"
	"fmt"
	"time"

	"go-devops-mimi/server/model/exec"
	execReq "go-devops-mimi/server/model/exec/request"
	"go-devops-mimi/server/model/exec/response"

	"go-devops-mimi/server/public/tools"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CronLogic struct {
}

func (l CronLogic) Add(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*execReq.CronAddReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	// 获取当前用户
	ctxUser, err := userService.GetCurrentLoginUser(c)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取当前登陆用户信息失败"))
	}
	var onceTimePtr *time.Time // 声明为指针变量
	switch r.CronType {
	case "cron":
		if err := tools.ValidateCronExpr(r.Cronession); err != nil {
			return nil, tools.NewMySqlError(fmt.Errorf("cron表达式不合法: %s", err.Error()))
		}

	case "once":
		if r.OnceTime == "" {
			return nil, tools.NewMySqlError(fmt.Errorf("OnceTime 不能为空"))
		}
		onceTime := tools.String2Time(r.OnceTime)
		onceTimePtr = &onceTime
	case "interval":
		// 处理间隔任务逻辑（如果需要）
		if r.Interval <= 0 {
			return nil, tools.NewMySqlError(fmt.Errorf("间隔时间必须大于0"))
		}
	default:
		return nil, tools.NewMySqlError(fmt.Errorf("未知的任务类型: %s", r.CronType))
	}
	// 序列化节点 ID 列表
	nodeidsJSON, _ := json.Marshal(r.NodesIDs)
	cron := &exec.Cron{
		Name:       r.Name,
		CronType:   r.CronType,
		Cronession: r.Cronession,
		Interval:   r.Interval,
		OnceTime:   onceTimePtr,
		NodesIDs:   string(nodeidsJSON),
		CmdType:    r.CmdType,
		Type:       r.Type,
		Content:    r.Content,
		Timeout:    r.Timeout,
		Desc:       r.Desc,
		Creator:    ctxUser.Username,
	}
	// 添加数据
	err = CronService.Add(cron)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("添加数据失败: %s", err.Error()))
	}
	// 添加成功
	return nil, nil
}

func (l CronLogic) Update(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*execReq.CronUpdateReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	list, err := CronService.Info(r.ID)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据详情失败: %s", err.Error()))
	}
	if list.Enable == 1 {
		return nil, tools.NewMySqlError(fmt.Errorf("当前定时任务已启用，请先禁用后再修改"))
	}
	// 获取当前用户
	ctxUser, err := userService.GetCurrentLoginUser(c)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取当前登陆用户信息失败"))
	}
	var onceTimePtr *time.Time // 声明为指针变量
	switch r.CronType {
	case "cron":
		if err := tools.ValidateCronExpr(r.Cronession); err != nil {
			return nil, tools.NewMySqlError(fmt.Errorf("cron表达式不合法: %s", err.Error()))
		}

	case "once":
		if r.OnceTime == "" {
			return nil, tools.NewMySqlError(fmt.Errorf("OnceTime 不能为空"))
		}
		onceTime := tools.String2Time(r.OnceTime)
		onceTimePtr = &onceTime
	case "interval":
		// 处理间隔任务逻辑（如果需要）
		if r.Interval <= 0 {
			return nil, tools.NewMySqlError(fmt.Errorf("间隔时间必须大于0"))
		}
	default:
		return nil, tools.NewMySqlError(fmt.Errorf("未知的任务类型: %s", r.CronType))
	}
	// 序列化节点 ID 列表
	nodeidsJSON, _ := json.Marshal(r.NodesIDs)
	cron := &exec.Cron{
		Model:      gorm.Model{ID: r.ID},
		Name:       r.Name,
		CronType:   r.CronType,
		Cronession: r.Cronession,
		Interval:   r.Interval,
		OnceTime:   onceTimePtr,
		NodesIDs:   string(nodeidsJSON),
		CmdType:    r.CmdType,
		Type:       r.Type,
		Content:    r.Content,
		Timeout:    r.Timeout,
		Desc:       r.Desc,
		Creator:    ctxUser.Username,
	}

	// 更新数据
	err = CronService.Update(cron)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("更新数据失败: %s", err.Error()))
	}
	return nil, nil
}

func (l CronLogic) List(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*execReq.CronListReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	// 获取数据列表
	list, err := CronService.List(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据列表失败: %s", err.Error()))
	}

	rets := make([]exec.Cron, 0)
	for _, nodes := range list {
		rets = append(rets, *nodes)
	}
	count, err := CronService.ListCount(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据总数失败"))
	}

	return response.CronListRsp{
		Total: count,
		Crons: rets,
	}, nil
}

// Delete 删除数据
func (l CronLogic) Delete(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*execReq.CronDeleteReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	// 删除数据
	err := CronService.Delete(r.Ids)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("删除数据失败: %s", err.Error()))
	}
	return nil, nil
}

// / Info 查看数据详情
func (s CronLogic) Info(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*execReq.CronInfoReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	list, err := CronService.Info(r.ID)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据详情失败: %s", err.Error()))
	}
	return list, nil
}

// Enable 用于启用或禁用定时任务
func (s CronLogic) Enable(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*execReq.CronEnableReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}

	// 先更新数据库里的 enable 字段
	err := CronService.Enable(r.ID, r.Enable)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("更新任务状态失败: %w", err))
	}

	// 查询最新的任务配置
	cronCfg, err := CronService.Info(r.ID)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取任务详情失败: %w", err))
	}

	// 根据 enable 字段决定是添加还是移除调度
	if r.Enable == 0 {
		// 禁用：移除已经注册的定时任务
		switch cronCfg.CronType {
		case "once":
			ts := cronCfg.OnceTime.Unix()
			tools.CronJob.RemoveAbsTimeJob(ts)
			logrus.Infof("一次性任务 id=%d 已移除调度", cronCfg.ID)
		case "interval":
			tools.CronJob.RemoveIntervalJob(cronCfg.Interval)
			logrus.Infof("间隔任务 id=%d 每 %d 秒 已移除调度", cronCfg.ID, cronCfg.Interval)
		case "cron":
			tools.CronJob.RemoveCronExprJob(cronCfg.Cronession)
			logrus.Infof("Cron 表达式任务 id=%d expr=%s 已移除调度", cronCfg.ID, cronCfg.Cronession)
		}
	} else {
		// 启用：重新注册
		SendCron(cronCfg)
		logrus.Infof("任务 id=%d 重新注册调度", cronCfg.ID)
	}

	return nil, nil
}
