package exec_svc

import (
	"fmt"
	"strings"

	"go-devops-mimi/server/model/exec"
	"go-devops-mimi/server/model/exec/request"
	"go-devops-mimi/server/public/common"
	"go-devops-mimi/server/public/tools"
)

type ScriptService struct{}

// List 获取数据列表
func (s ScriptService) List(req *request.ScriptListReq) ([]*exec.Script, error) {
	var list []*exec.Script
	db := common.DB.Model(&exec.Script{}).Order("created_at DESC")

	name := strings.TrimSpace(req.Name)
	if name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	cmd_type := strings.TrimSpace(req.CmdType)
	if cmd_type != "" {
		db = db.Where("cmd_type LIKE ?", fmt.Sprintf("%%%s%%", cmd_type))
	}

	pageReq := tools.NewPageOption(req.PageNum, req.PageSize)
	err := db.Offset(pageReq.PageNum).Limit(pageReq.PageSize).Preload("Scripts").Find(&list).Error
	return list, err
}

// ListCount 获取数据总数
func (s ScriptService) ListCount(req *request.ScriptListReq) (int64, error) {
	var count int64
	db := common.DB.Model(&exec.Script{}).Order("created_at DESC")

	name := strings.TrimSpace(req.Name)
	if name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	err := db.Count(&count).Error
	return count, err
}

// Add 添加资源
func (s ScriptService) Add(dataObj *exec.Script) error {
	return common.DB.Preload("Scripts").Create(dataObj).Error
}
