package exec_svc

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"go-devops-mimi/server/model/exec"
	"go-devops-mimi/server/model/exec/request"
	"go-devops-mimi/server/public/common"
	"go-devops-mimi/server/public/tools"

	"gorm.io/gorm"
)

type CronService struct{}

// List 获取数据列表
func (s CronService) List(req *request.CronListReq) ([]*exec.Cron, error) {
	var list []*exec.Cron
	db := common.DB.Model(&exec.Cron{}).Order("created_at DESC")

	name := strings.TrimSpace(req.Name)
	if name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	pageReq := tools.NewPageOption(req.PageNum, req.PageSize)
	err := db.Offset(pageReq.PageNum).Limit(pageReq.PageSize).Find(&list).Error
	return list, err
}

// ListCount 获取数据总数
func (s CronService) ListCount(req *request.CronListReq) (int64, error) {
	var count int64
	db := common.DB.Model(&exec.Cron{}).Order("created_at DESC")

	name := strings.TrimSpace(req.Name)
	if name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	err := db.Count(&count).Error
	return count, err
}

// Add 添加资源
func (s CronService) Add(dataObj *exec.Cron) error {
	return common.DB.Create(dataObj).Error
}

// Update 更新资源
func (s CronService) Update(dataObj *exec.Cron) error {
	return common.DB.Model(dataObj).Where("id = ?", dataObj.ID).Updates(dataObj).Error
}

// info获取指定name的信息
func (s CronService) Info(id uint) (*exec.Cron, error) {
	var server *exec.Cron
	err := common.DB.Where("id = ?", id).First(&server).Error
	return server, err

}

func (s CronService) Enable(id uint, Enable int8) (err error) {
	var cron exec.Cron
	if err := common.DB.Where("id = ?", id).First(&cron).Error; err != nil {
		return err
	}
	// 只有一次性任务在「启用」时才校验发送时间
	if cron.CronType == "once" && Enable == 1 {
		if cron.OnceTime == nil {
			return errors.New("一次性任务发送时间不能为空")
		}
		if cron.OnceTime.Before(time.Now()) {
			return errors.New("发送时间已过，不可启用")
		}
	}
	cron.Enable = Enable
	if err := common.DB.Model(&exec.Cron{}).Where("id = ?", id).Save(&cron).Error; err != nil {
		return err
	}
	return err
}

// Find 获取单个资源
func (s CronService) Find(filter map[string]interface{}, data *exec.Cron) error {
	return common.DB.Where(filter).First(&data).Error
}

// Exist 判断资源是否存在
func (s CronService) Exist(filter map[string]interface{}) bool {
	var dataObj exec.Cron
	err := common.DB.Debug().Order("created_at DESC").Where(filter).First(&dataObj).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// Delete 批量删除
func (s CronService) Delete(ids []uint) error {
	// 开启事务
	tx := common.DB.Begin()
	// 删除服务器记录
	if err := tx.Where("id IN (?)", ids).Unscoped().Delete(&exec.Cron{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 提交事务
	return tx.Commit().Error

}
