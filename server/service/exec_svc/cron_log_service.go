package exec_svc

import (
	"fmt"
	"strings"

	"go-devops-mimi/server/model/exec"
	"go-devops-mimi/server/model/exec/request"
	"go-devops-mimi/server/public/common"
	"go-devops-mimi/server/public/tools"
)

type CronLogService struct{}

// List 获取数据列表
func (s CronLogService) List(req *request.CronListReq) ([]*exec.CronLog, error) {
	var list []*exec.CronLog
	db := common.DB.Model(&exec.CronLog{}).Order("created_at DESC")
	name := strings.TrimSpace(req.Name)
	if name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	cronType := strings.TrimSpace(req.CronType)
	if cronType != "" {
		db = db.Where("cron_type LIKE ?", fmt.Sprintf("%%%s%%", cronType))
	}
	pageReq := tools.NewPageOption(req.PageNum, req.PageSize)
	err := db.Offset(pageReq.PageNum).Limit(pageReq.PageSize).Preload("ScriptLogs").Find(&list).Error
	return list, err
}

// ListCount 获取数据总数
func (s CronLogService) ListCount(req *request.CronListReq) (int64, error) {
	var count int64
	db := common.DB.Model(&exec.CronLog{}).Order("created_at DESC")

	name := strings.TrimSpace(req.Name)
	if name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	err := db.Count(&count).Error
	return count, err
}

// Add 添加资源
func (s CronLogService) Add(dataObj *exec.CronLog) error {
	return common.DB.Preload("ScriptLogs").Create(dataObj).Error
}
