package exec_svc

import (
	"go-devops-mimi/server/model/exec"
	"go-devops-mimi/server/public/common"
)

type ScriptLogService struct{}

// // List 获取数据列表
// func (s ScriptLogService) List(req *request.ScriptLogListReq) ([]*exec.ScriptLog, error) {
// 	var list []*exec.ScriptLog
// 	db := common.DB.Model(&exec.ScriptLog{}).Order("created_at DESC")

// 	name := strings.TrimSpace(req.Name)
// 	if name != "" {
// 		db = db.Where("cloud_name LIKE ?", fmt.Sprintf("%%%s%%", name))
// 	}
// 	pageReq := tools.NewPageOption(req.PageNum, req.PageSize)
// 	err := db.Offset(pageReq.PageNum).Limit(pageReq.PageSize).Find(&list).Error
// 	return list, err
// }

// // ListCount 获取数据总数
// func (s ScriptLogService) ListCount(req *request.ScriptListReq) (int64, error) {
// 	var count int64
// 	db := common.DB.Model(&exec.ScriptLog{}).Order("created_at DESC")

// 	name := strings.TrimSpace(req.Name)
// 	if name != "" {
// 		db = db.Where("cloud_name LIKE ?", fmt.Sprintf("%%%s%%", name))
// 	}
// 	err := db.Count(&count).Error
// 	return count, err
// }

// Add 添加资源
func (s ScriptLogService) Add(dataObj *exec.ScriptLog) error {
	return common.DB.Create(dataObj).Error
}
