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

type ScriptLibraryService struct{}

// List 获取数据列表
func (s ScriptLibraryService) List(req *request.ScriptLibraryListReq) ([]*exec.ScriptLibrary, error) {
	var list []*exec.ScriptLibrary
	db := common.DB.Model(&exec.ScriptLibrary{}).Order("created_at DESC")

	name := strings.TrimSpace(req.Name)
	if name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	pageReq := tools.NewPageOption(req.PageNum, req.PageSize)
	err := db.Offset(pageReq.PageNum).Limit(pageReq.PageSize).Find(&list).Error
	return list, err
}

// ListCount 获取数据总数
func (s ScriptLibraryService) ListCount(req *request.ScriptLibraryListReq) (int64, error) {
	var count int64
	db := common.DB.Model(&exec.ScriptLibrary{}).Order("created_at DESC")

	name := strings.TrimSpace(req.Name)
	if name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	err := db.Count(&count).Error
	return count, err
}

// Add 添加资源
func (s ScriptLibraryService) Add(dataObj *exec.ScriptLibrary) error {
	return common.DB.Create(dataObj).Error
}

// Update 更新资源
func (s ScriptLibraryService) Update(dataObj *exec.ScriptLibrary) error {
	return common.DB.Model(dataObj).Where("id = ?", dataObj.ID).Updates(dataObj).Error
}

// info获取指定name的脚本库信息
func (s ScriptLibraryService) Info(id uint, name string) (*exec.ScriptLibrary, error) {
	var server *exec.ScriptLibrary
	err := common.DB.Where("id = ? and name = ?", id, name).First(&server).Error
	return server, err

}

// Find 获取单个资源
func (s ScriptLibraryService) Find(filter map[string]interface{}, data *exec.ScriptLibrary) error {
	return common.DB.Where(filter).First(&data).Error
}

// Exist 判断资源是否存在
func (s ScriptLibraryService) Exist(filter map[string]interface{}) bool {
	var dataObj exec.ScriptLibrary
	err := common.DB.Debug().Order("created_at DESC").Where(filter).First(&dataObj).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// Delete 批量删除
func (s ScriptLibraryService) Delete(ids []uint) error {
	// 开启事务
	tx := common.DB.Begin()
	// 删除服务器记录
	if err := tx.Where("id IN (?)", ids).Unscoped().Delete(&exec.ScriptLibrary{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 提交事务
	return tx.Commit().Error

}
