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

type AgentGroup_svc struct{}

// List 获取数据列表
func (s AgentGroup_svc) List(req *request.AgentGroupListReq) ([]*cmdb.AgentGroup, error) {
	var list []*cmdb.AgentGroup
	db := common.DB.Model(&cmdb.AgentGroup{}).Order("created_at DESC")

	groupName := strings.TrimSpace(req.GroupName)
	if groupName != "" {
		db = db.Where("groupName LIKE ?", fmt.Sprintf("%%%s%%", groupName))
	}
	pageReq := tools.NewPageOption(req.PageNum, req.PageSize)
	err := db.Offset(pageReq.PageNum).Limit(pageReq.PageSize).Preload("group_agent_s").Find(&list).Error
	return list, err
}

// ListCount 获取数据总数
func (s AgentGroup_svc) ListCount(req *request.AgentGroupListReq) (int64, error) {
	var count int64
	db := common.DB.Model(&cmdb.AgentGroup{}).Order("created_at DESC")

	groupName := strings.TrimSpace(req.GroupName)
	if groupName != "" {
		db = db.Where("groupName LIKE ?", fmt.Sprintf("%%%s%%", groupName))
	}
	err := db.Count(&count).Error
	return count, err
}

// Count 获取数据总数
func (s AgentGroup_svc) Count() (int64, error) {
	var count int64
	db := common.DB.Model(&cmdb.AgentGroup{}).Order("created_at DESC")
	err := db.Count(&count).Error
	return count, err
}

// Add 添加资源
func (s AgentGroup_svc) Add(dataObj *cmdb.AgentGroup) error {
	return common.DB.Create(dataObj).Error
}

// Update 更新资源
func (s AgentGroup_svc) Update(dataObj *cmdb.AgentGroup) error {
	return common.DB.Model(dataObj).Where("id = ?", dataObj.ID).Updates(dataObj).Error
}

// AddNodeToGroup 添加node到分组
func (s AgentGroup_svc) AddAgentToGroup(group *cmdb.AgentGroup, agent []cmdb.Agent) error {
	return common.DB.Model(&group).Association("Agents").Append(agent)
}

// RemoveAgentGroup 将node从分组移除
func (s AgentGroup_svc) RemoveAgentGroup(group *cmdb.AgentGroup, agent []cmdb.Agent) error {
	return common.DB.Model(&group).Association("Agents").Delete(agent)
}

// Find 获取单个资源
func (s AgentGroup_svc) Find(filter map[string]interface{}, data *cmdb.AgentGroup) error {
	return common.DB.Where(filter).First(&data).Error
}

// Exist 判断资源是否存在
func (s AgentGroup_svc) Exist(filter map[string]interface{}) bool {
	var dataObj cmdb.AgentGroup
	err := common.DB.Debug().Order("created_at DESC").Where(filter).First(&dataObj).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// Delete 批量删除
func (s AgentGroup_svc) Delete(ids []uint) error {
	// 开启事务
	tx := common.DB.Begin()
	// 删除服务器记录
	if err := tx.Where("id IN (?)", ids).Unscoped().Delete(&cmdb.AgentGroup{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 清空每个服务器记录的关联的 Keys
	err := tx.Exec("DELETE FROM group_agent_s WHERE agent_group_id IN (?)", ids).Error
	if err != nil {
		return err
	}
	// 提交事务
	return tx.Commit().Error
}
