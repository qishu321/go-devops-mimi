package exec_svc

import (
	"errors"
	"fmt"
	"strings"

	"go-devops-mimi/server/model/exec"
	"go-devops-mimi/server/model/exec/request"
	"go-devops-mimi/server/public/common"
	"go-devops-mimi/server/public/tools"

	"gorm.io/gorm"
)

type TransferService struct{}

// List 获取数据列表
func (s TransferService) List(req *request.TransferListReq) ([]*exec.Transfer, error) {
	var list []*exec.Transfer
	db := common.DB.Model(&exec.Transfer{}).Order("created_at DESC")

	name := strings.TrimSpace(req.Name)
	if name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	pageReq := tools.NewPageOption(req.PageNum, req.PageSize)
	err := db.Offset(pageReq.PageNum).Limit(pageReq.PageSize).Find(&list).Error
	return list, err
}

// ListCount 获取数据总数
func (s TransferService) ListCount(req *request.TransferListReq) (int64, error) {
	var count int64
	db := common.DB.Model(&exec.Transfer{}).Order("created_at DESC")

	name := strings.TrimSpace(req.Name)
	if name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	err := db.Count(&count).Error
	return count, err
}

// Add 添加资源
func (s TransferService) Add(dataObj *exec.Transfer) error {
	return common.DB.Create(dataObj).Error
}

// info获取指定name的脚本库信息
func (s TransferService) Info(id uint, name string) (*exec.Transfer, error) {
	var server *exec.Transfer
	err := common.DB.Where("id = ? and name = ?", id, name).First(&server).Error
	return server, err

}

// Find 获取单个资源
func (s TransferService) Find(filter map[string]interface{}, data *exec.Transfer) error {
	return common.DB.Where(filter).First(&data).Error
}

// Exist 判断资源是否存在
func (s TransferService) Exist(filter map[string]interface{}) bool {
	var dataObj exec.Transfer
	err := common.DB.Debug().Order("created_at DESC").Where(filter).First(&dataObj).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// Delete 批量删除
func (s TransferService) Delete(ids []uint) error {
	// 开启事务
	tx := common.DB.Begin()
	// 删除服务器记录
	if err := tx.Where("id IN (?)", ids).Unscoped().Delete(&exec.Transfer{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 提交事务
	return tx.Commit().Error

}
