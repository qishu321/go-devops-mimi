package response

import (
	"go-devops-mimi/server/model/exec"
)

type CronListRsp struct {
	Total int64       `json:"total"`
	Crons []exec.Cron `json:"cron_s"`
}
type CronLogListRsp struct {
	Total    int64          `json:"total"`
	CronLogs []exec.CronLog `json:"cron_log_s"`
}
