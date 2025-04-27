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

type Node_svc struct{}

// List 获取数据列表
func (s Node_svc) List(req *request.NodesListReq) ([]*cmdb.Nodes, error) {
	var list []*cmdb.Nodes
	db := common.DB.Model(&cmdb.Nodes{}).Order("created_at DESC")

	nodeName := strings.TrimSpace(req.NodeName)
	if nodeName != "" {
		db = db.Where("node_name LIKE ?", fmt.Sprintf("%%%s%%", nodeName))
	}
	publicIP := strings.TrimSpace(req.PublicIP)
	if publicIP != "" {
		db = db.Where("public_ip LIKE ?", fmt.Sprintf("%%%s%%", publicIP))
	}

	pageReq := tools.NewPageOption(req.PageNum, req.PageSize)
	err := db.Offset(pageReq.PageNum).Limit(pageReq.PageSize).Preload("Groups").Find(&list).Error
	return list, err
}

// ListCount 获取数据总数
func (s Node_svc) ListCount(req *request.NodesListReq) (int64, error) {
	var count int64
	db := common.DB.Model(&cmdb.Nodes{}).Order("created_at DESC")

	nodeName := strings.TrimSpace(req.NodeName)
	if nodeName != "" {
		db = db.Where("node_name LIKE ?", fmt.Sprintf("%%%s%%", nodeName))
	}
	publicIP := strings.TrimSpace(req.PublicIP)
	if publicIP != "" {
		db = db.Where("publicIP LIKE ?", fmt.Sprintf("%%%s%%", publicIP))
	}
	err := db.Count(&count).Error
	return count, err
}

// Count 获取数据总数
func (s Node_svc) Count() (int64, error) {
	var count int64
	db := common.DB.Model(&cmdb.Nodes{}).Order("created_at DESC")

	err := db.Count(&count).Error
	return count, err
}

// Add 添加资源
func (s Node_svc) Add(dataObj *cmdb.Nodes, ids []int64) error {
	// 在数据库中创建服务器组记录
	if err := common.DB.Create(dataObj).Error; err != nil {
		return err
	}
	// 遍历ids数组，将每个ID与服务器组建立关联
	for _, id := range ids {
		// 执行插入操作，如果存在主键冲突则执行更新操作
		err := common.DB.Exec("INSERT INTO t_node_group_s (nodes_id, node_group_id) VALUES (?, ?) ON DUPLICATE KEY UPDATE node_group_id = node_group_id", dataObj.ID, id).Error
		if err != nil {
			return err
		}
	}

	return nil

}

// func (s Node_svc) AddNodeToGroup(node *cmdb.Nodes, group []cmdb.NodeGroup) error {
// 	return common.DB.Model(&node).Association("Groups").Append(group)
// }

// RemoveNodeGroup 将node从分组移除

func (s Node_svc) RemoveNodeGroup(node *cmdb.Nodes, group []cmdb.NodeGroup) error {
	return common.DB.Model(&node).Association("Groups").Delete(group)
}

// Update 更新资源
//
//	func (s Node_svc) Update(dataObj *cmdb.Nodes) error {
//		return common.DB.Model(dataObj).Where("id = ?", dataObj.ID).Updates(dataObj).Error
//	}
func (s Node_svc) AddNodeToGroup(nodeids []int64, ids []int64) error {
	tx := common.DB.Begin()
	//先清空每个服务器关联的分组
	for _, nodeid := range nodeids {
		// 执行插入操作，如果存在主键冲突则执行更新操作
		err := tx.Exec("DELETE FROM t_node_group_s WHERE nodes_id IN (?)", nodeid).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	// 遍历 nodeids 与 ids 数组，将每个节点和主机组建立关联
	for _, nodeid := range nodeids {
		for _, groupid := range ids {
			err := tx.Exec("INSERT INTO t_node_group_s (nodes_id, node_group_id) VALUES (?, ?)", nodeid, groupid).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}
func (s Node_svc) Update(nodes *cmdb.Nodes, ids []int64) error {
	// 开启事务
	tx := common.DB.Begin()
	err := tx.Exec("DELETE FROM t_node_group_s WHERE nodes_id IN (?)", nodes.ID).Error
	if err != nil {
		return err
	}
	if err := tx.Model(&cmdb.Nodes{}).Where("id = ? ", nodes.ID).Updates(nodes).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to Update servers: %w", err)
	}
	// 遍历ids数组，将每个ID与服务器组建立关联
	for _, id := range ids {
		// 执行插入操作，如果存在主键冲突则执行更新操作
		err = tx.Exec("INSERT INTO t_node_group_s (nodes_id, node_group_id) VALUES (?, ?) ON DUPLICATE KEY UPDATE nodes_id = nodes_id", nodes.ID, id).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// 提交事务
	return tx.Commit().Error

}

// Find 获取单个资源
func (s Node_svc) Find(filter map[string]interface{}, data *cmdb.Nodes) error {
	return common.DB.Where(filter).First(&data).Error
}
func (s Node_svc) Info(id int) (*cmdb.Nodes, error) {
	var node *cmdb.Nodes
	err := common.DB.Where("id = ?", id).First(&node).Error
	return node, err

}

// GetUserByIds 根据ID获取node排序最小值
func (s Node_svc) GetUserByIds(ids []uint) ([]cmdb.Nodes, error) {
	var nodeList []cmdb.Nodes
	err := common.DB.Where("id IN (?)", ids).Find(&nodeList).Error
	return nodeList, err
}

// Exist 判断资源是否存在
func (s Node_svc) Exist(filter map[string]interface{}) bool {
	var dataObj cmdb.Nodes
	err := common.DB.Debug().Order("created_at DESC").Where(filter).First(&dataObj).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// Delete 批量删除
func (s Node_svc) Delete(ids []uint) error {
	// 开启事务
	tx := common.DB.Begin()
	// 删除服务器记录
	if err := tx.Where("id IN (?)", ids).Unscoped().Delete(&cmdb.Nodes{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 清空每个服务器记录的关联的 Keys
	err := tx.Exec("DELETE FROM t_node_group_s WHERE nodes_id IN (?)", ids).Error
	if err != nil {
		return err
	}
	// 提交事务
	return tx.Commit().Error
}
