package response

import (
	"go-devops-mimi/server/model/exec"
)

type TaskManageListRsp struct {
	Total       int64             `json:"total"`
	TaskManages []exec.TaskManage `json:"task_manage_s"`
}
