package exec_svc

import (
	"go-devops-mimi/server/model/exec"
	"go-devops-mimi/server/public/common"
)

type TaskLogService struct{}

// Add 添加资源
func (s TaskLogService) Add(dataObj *exec.TaskLog) error {
	return common.DB.Create(dataObj).Error
}

// AddNodeToGroup 添加node到分组
func (s TaskLogService) AddTaskLogToManage(group *exec.TaskManageLog, task *exec.TaskLog) error {
	return common.DB.Model(&group).Association("Tasklogs").Append(task)
}
func (s TaskLogService) Save(dataObj *exec.TaskLog) error {
	return common.DB.Save(dataObj).Error
}
