package response

import (
	"go-devops-mimi/server/model/exec"
)

type TaskManageListRsp struct {
	Total       int64             `json:"total"`
	TaskManages []exec.TaskManage `json:"task_manage_s"`
}
type ManageLogListRsp struct {
	Total      int64            `json:"total"`
	ManageLogs []exec.ManageLog `json:"manage_log_s"`
}
