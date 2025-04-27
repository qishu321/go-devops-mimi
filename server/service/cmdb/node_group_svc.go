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

type NodeGroup_svc struct{}

// List 获取数据列表
func (s NodeGroup_svc) List(req *request.NodeGroupListReq) ([]*cmdb.NodeGroup, error) {
	var list []*cmdb.NodeGroup
	db := common.DB.Model(&cmdb.NodeGroup{}).Order("created_at DESC")

	groupName := strings.TrimSpace(req.GroupName)
	if groupName != "" {
		db = db.Where("groupName LIKE ?", fmt.Sprintf("%%%s%%", groupName))
	}
	pageReq := tools.NewPageOption(req.PageNum, req.PageSize)
	err := db.Offset(pageReq.PageNum).Limit(pageReq.PageSize).Preload("NodeGroups").Find(&list).Error
	// 填充每个分组的 NodeCount
	for _, g := range list {
		g.NodeCount = len(g.NodeGroups)
	}

	return list, err
}

// ListCount 获取数据总数
func (s NodeGroup_svc) ListCount(req *request.NodeGroupListReq) (int64, error) {
	var count int64
	db := common.DB.Model(&cmdb.NodeGroup{}).Order("created_at DESC")

	groupName := strings.TrimSpace(req.GroupName)
	if groupName != "" {
		db = db.Where("groupName LIKE ?", fmt.Sprintf("%%%s%%", groupName))
	}
	err := db.Count(&count).Error
	return count, err
}

// Count 获取数据总数
func (s NodeGroup_svc) Count() (int64, error) {
	var count int64
	db := common.DB.Model(&cmdb.NodeGroup{}).Order("created_at DESC")
	err := db.Count(&count).Error
	return count, err
}

// Add 添加资源
func (s NodeGroup_svc) Add(dataObj *cmdb.NodeGroup) error {
	return common.DB.Create(dataObj).Error
}

// Update 更新资源
func (s NodeGroup_svc) Update(dataObj *cmdb.NodeGroup) error {
	return common.DB.Model(dataObj).Where("id = ?", dataObj.ID).Updates(dataObj).Error
}

// AddNodeToGroup 添加node到分组
func (s NodeGroup_svc) AddNodeToGroup(group *cmdb.NodeGroup, node []cmdb.Nodes) error {
	return common.DB.Model(&group).Association("NodeGroups").Append(node)
}

// RemoveNodeGroup 将node从分组移除
func (s NodeGroup_svc) RemoveNodeGroup(group *cmdb.NodeGroup, node []cmdb.Nodes) error {
	return common.DB.Model(&group).Association("NodeGroups").Delete(node)
}

// Find 获取单个资源
func (s NodeGroup_svc) Find(filter map[string]interface{}, data *cmdb.NodeGroup) error {
	return common.DB.Where(filter).First(&data).Error
}

// 获取指定资源
func (s NodeGroup_svc) Info(id int64) ([]*cmdb.NodeGroup, error) {
	var group []*cmdb.NodeGroup
	err := common.DB.Where("id = ?", id).First(&group).Error
	return group, err

}

// Exist 判断资源是否存在
func (s NodeGroup_svc) Exist(filter map[string]interface{}) bool {
	var dataObj cmdb.NodeGroup
	err := common.DB.Debug().Order("created_at DESC").Where(filter).First(&dataObj).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// Delete 批量删除
func (s NodeGroup_svc) Delete(ids []uint) error {
	// 开启事务
	tx := common.DB.Begin()
	// 删除服务器记录
	if err := tx.Where("id IN (?)", ids).Unscoped().Delete(&cmdb.NodeGroup{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 清空每个服务器记录的关联的 Keys
	err := tx.Exec("DELETE FROM t_node_group_s WHERE node_group_id IN (?)", ids).Error
	if err != nil {
		return err
	}
	// 提交事务
	return tx.Commit().Error
}
