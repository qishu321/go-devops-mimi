package exec_svc

import (
	"fmt"
	"strings"
	"time"

	"go-devops-mimi/server/model/exec"
	"go-devops-mimi/server/model/exec/request"
	"go-devops-mimi/server/public/common"
	"go-devops-mimi/server/public/tools"
)

type ManageLogService struct{}

// List 获取数据列表
func (s ManageLogService) List(req *request.ManageLogListReq) ([]*exec.ManageLog, error) {
	var list []*exec.ManageLog
	db := common.DB.Model(&exec.ManageLog{}).Order("created_at DESC")

	name := strings.TrimSpace(req.Name)
	if name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}

	pageReq := tools.NewPageOption(req.PageNum, req.PageSize)
	err := db.Offset(pageReq.PageNum).Limit(pageReq.PageSize).Find(&list).Error
	return list, err
}

// ListCount 获取数据总数
func (s ManageLogService) ListCount(req *request.ManageLogListReq) (int64, error) {
	var count int64
	db := common.DB.Model(&exec.ManageLog{}).Order("created_at DESC")

	name := strings.TrimSpace(req.Name)
	if name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	err := db.Count(&count).Error
	return count, err
}

// Add 添加资源
func (s ManageLogService) Add(dataObj *exec.ManageLog) error {
	return common.DB.Preload("TaskManageLogs").Create(dataObj).Error
}

func (s ManageLogService) AddTaskManageLog(dataObj *exec.TaskManageLog) error {
	return common.DB.Preload("Tasklogs").Create(dataObj).Error
}

// AddNodeToGroup 添加node到分组
func (s ManageLogService) AddManageToTaskLog(group *exec.ManageLog, task *exec.TaskManageLog) error {
	return common.DB.Model(&group).Association("TaskManageLogs").Append(task)
}
func (s ManageLogService) SaveTaskManageLog(dataObj *exec.TaskManageLog) error {
	return common.DB.Save(dataObj).Error
}
func (s TaskManageService) InfoManageLog(id uint) (*exec.ManageLog, error) {
	var server *exec.ManageLog
	err := common.DB.Preload("TaskManageLogs").Where("id = ?", id).First(&server).Error
	return server, err
}

// UpdateTaskManageLogsStatus 更新子任务组(TaskManageLog)状态
// UpdateAllStatusesByManageLogID 递归更新：
//  1. 各个 TaskManageLog（子任务组）状态
//  2. ManageLog（大任务组）状态
func (s *ManageLogService) UpdateAllStatusesByManageLogID(manageLogID uint) error {
	var (
		mng exec.ManageLog
		now = time.Now()
	)

	// 1. 把 ManageLog 和它的 TaskManageLogs 一起查出来
	if err := common.DB.
		Preload("TaskManageLogs").
		First(&mng, manageLogID).
		Error; err != nil {
		return fmt.Errorf("加载 ManageLog %d 失败: %w", manageLogID, err)
	}

	overallFailed := false

	// 2. 遍历每个 TaskManageLog，更新它们自己的状态
	for _, tml := range mng.TaskManageLogs {
		// 2.1 统计失败的小任务数
		var failedCount int64
		if err := common.DB.
			Model(&exec.TaskLog{}).
			Joins("JOIN t_task_log_s tls ON tls.task_log_id = t_task_log.id").
			Where("tls.task_manage_log_id = ?", tml.ID).
			Where("t_task_log.status = ?", "failed").
			Count(&failedCount).
			Error; err != nil {
			return fmt.Errorf("统计子任务组 %d 失败数量失败: %w", tml.ID, err)
		}

		// 2.2 决定子任务组的新状态
		tmlStatus := "success"
		if failedCount > 0 {
			tmlStatus = "failed"
			overallFailed = true
		}

		// 2.3 计算耗时
		startTime, err := tools.ParseTime(tml.StartTime)
		if err != nil {
			return fmt.Errorf("解析子任务组 %d StartTime 失败: %w", tml.ID, err)
		}
		timeCost := now.Sub(startTime).Milliseconds()

		// 2.4 更新该 TaskManageLog
		if err := common.DB.
			Model(&exec.TaskManageLog{}).
			Where("id = ?", tml.ID).
			Updates(map[string]interface{}{
				"status":    tmlStatus,
				"end_time":  now.Format(time.RFC3339),
				"time_cost": timeCost,
			}).
			Error; err != nil {
			return fmt.Errorf("更新子任务组 %d 状态失败: %w", tml.ID, err)
		}
	}

	// 3. 更新大任务组本身的状态
	mgStatus := "success"
	if overallFailed {
		mgStatus = "failed"
	}
	mgStart, err := tools.ParseTime(mng.StartTime)
	if err != nil {
		return fmt.Errorf("解析 ManageLog %d StartTime 失败: %w", mng.ID, err)
	}
	mgTimeCost := now.Sub(mgStart).Milliseconds()

	if err := common.DB.
		Model(&exec.ManageLog{}).
		Where("id = ?", mng.ID).
		Updates(map[string]interface{}{
			"status":    mgStatus,
			"end_time":  now.Format(time.RFC3339),
			"time_cost": mgTimeCost,
		}).
		Error; err != nil {
		return fmt.Errorf("更新 ManageLog %d 状态失败: %w", mng.ID, err)
	}

	return nil
}
