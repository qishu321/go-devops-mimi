package cmdb

import (
	"errors"
	"fmt"
	"strings"

	"go-devops-mimi/server/model/cmdb"
	"go-devops-mimi/server/model/cmdb/request"
	"go-devops-mimi/server/public/common"
	"go-devops-mimi/server/public/tools"

	"gorm.io/gorm"
)

type Agent_svc struct{}

// List 获取数据列表
func (s Agent_svc) List(req *request.AgentListReq) ([]*cmdb.Agent, error) {
	var list []*cmdb.Agent
	db := common.DB.Model(&cmdb.Agent{}).Order("created_at DESC")

	agentName := strings.TrimSpace(req.AgentName)
	if agentName != "" {
		db = db.Where("agentName LIKE ?", fmt.Sprintf("%%%s%%", agentName))
	}
	ip := strings.TrimSpace(req.IP)
	if ip != "" {
		db = db.Where("ip LIKE ?", fmt.Sprintf("%%%s%%", ip))
	}

	pageReq := tools.NewPageOption(req.PageNum, req.PageSize)
	err := db.Offset(pageReq.PageNum).Limit(pageReq.PageSize).Find(&list).Error
	return list, err
}

// ListCount 获取数据总数
func (s Agent_svc) ListCount(req *request.AgentListReq) (int64, error) {
	var count int64
	db := common.DB.Model(&cmdb.Agent{}).Order("created_at DESC")

	agentName := strings.TrimSpace(req.AgentName)
	if agentName != "" {
		db = db.Where("agentName LIKE ?", fmt.Sprintf("%%%s%%", agentName))
	}
	ip := strings.TrimSpace(req.IP)
	if ip != "" {
		db = db.Where("ip LIKE ?", fmt.Sprintf("%%%s%%", ip))
	}
	err := db.Count(&count).Error
	return count, err
}

// Count 获取数据总数
func (s Agent_svc) Count() (int64, error) {
	var count int64
	db := common.DB.Model(&cmdb.Agent{}).Order("created_at DESC")

	err := db.Count(&count).Error
	return count, err
}

// Add 添加资源
func (s Agent_svc) Add(dataObj *cmdb.Agent) error {
	return common.DB.Create(dataObj).Error
}

// Update 更新资源
func (s Agent_svc) Update(dataObj *cmdb.Agent) error {
	return common.DB.Model(dataObj).Where("id = ?", dataObj.ID).Updates(dataObj).Error
}

// Find 获取单个资源
func (s Agent_svc) Find(filter map[string]interface{}, data *cmdb.Agent) error {
	return common.DB.Where(filter).First(&data).Error
}
func (s Agent_svc) AgentHeartbeat(dataObj *cmdb.Agent) error {
	return common.DB.Model(dataObj).Where("agent_id  = ?", dataObj.Agentid).Updates(dataObj).Error
}

// GetUserByIds 根据ID获取排序最小值
func (s Agent_svc) GetUserByIds(ids []uint) ([]cmdb.Agent, error) {
	var List []cmdb.Agent
	err := common.DB.Where("id IN (?)", ids).Find(&List).Error
	return List, err
}

// 当为200时，表示agent不存在
func (s Agent_svc) Check(agentid uint64) (code int) {
	var agent *cmdb.Agent
	common.DB.Select("id").Where("agent_id = ?", agentid).First(&agent)
	if agent.ID > 0 {
		return 1
	}
	return 200

}

// Exist 判断资源是否存在
func (s Agent_svc) Exist(filter map[string]interface{}) bool {
	var dataObj cmdb.Agent
	err := common.DB.Debug().Order("created_at DESC").Where(filter).First(&dataObj).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// Delete 批量删除
func (s Agent_svc) Delete(ids []uint) error {
	// 开启事务
	tx := common.DB.Begin()
	// 删除服务器记录
	if err := tx.Where("id IN (?)", ids).Unscoped().Delete(&cmdb.Agent{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 清空每个服务器记录的关联的 Keys
	err := tx.Exec("DELETE FROM agent_groups WHERE agent_id IN (?)", ids).Error
	if err != nil {
		return err
	}
	// 提交事务
	return tx.Commit().Error
}
