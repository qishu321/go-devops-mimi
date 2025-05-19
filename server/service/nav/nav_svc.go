package nav

import (
	"go-devops-mimi/server/model/nav"
	"go-devops-mimi/server/model/nav/request"
	"go-devops-mimi/server/public/common"
	"go-devops-mimi/server/public/tools"
)

type Nav_svc struct{}

// List 获取数据列表
func (s Nav_svc) List(req *request.NavListRequest) ([]*nav.Nav, error) {
	var list []*nav.Nav
	db := common.DB.Model(&nav.Nav{}).Order("nav_sort DESC")
	pageReq := tools.NewPageOption(req.PageNum, req.PageSize)
	err := db.Offset(pageReq.PageNum).Limit(pageReq.PageSize).Preload("Links").Find(&list).Error
	return list, err
}

// Count 获取数据总数
func (s Nav_svc) Count() (int64, error) {
	var count int64
	db := common.DB.Model(&nav.Nav{}).Order("created_at DESC")

	err := db.Count(&count).Error
	return count, err
}

// Add 添加资源
func (s Nav_svc) Add(dataObj *nav.Nav) error {
	if err := common.DB.Create(dataObj).Error; err != nil {
		return err
	}
	return nil

}

// Update 更新资源
func (s Nav_svc) Update(dataObj *nav.Nav) error {
	return common.DB.Model(dataObj).Where("id = ?", dataObj.ID).Updates(dataObj).Error
}

func (s Nav_svc) Info(id uint) (*nav.Nav, error) {
	var node *nav.Nav
	err := common.DB.Where("id = ?", id).First(&node).Error
	return node, err

}

// GetUserByIds 根据ID获取node排序最小值
func (s Nav_svc) GetUserByIds(ids []uint) ([]nav.Nav, error) {
	var nodeList []nav.Nav
	err := common.DB.Where("id IN (?)", ids).Find(&nodeList).Error
	return nodeList, err
}

// 只删除分类
func (s Nav_svc) Delete(id uint) error {
	err := common.DB.Where("id = ?", id).Unscoped().Delete(&nav.Nav{}).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteAll 删除分类和关联的链接
func (s Nav_svc) DeleteAll(id uint) error {
	// 开启事务
	tx := common.DB.Begin()
	// 获取关联的链接
	var linkIDs []uint
	if err := tx.Table("t_link_s").Where("nav_id = ?", id).Pluck("link_id", &linkIDs).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 删除关联链接
	if len(linkIDs) > 0 {
		if err := tx.Where("id IN ?", linkIDs).Unscoped().Delete(&nav.Link{}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	// 清空每个分类记录的关联的链接
	err := tx.Exec("DELETE FROM t_link_s WHERE Nav_id IN (?)", id).Error
	if err != nil {
		return err
	}
	// 删除分类记录
	if err := tx.Where("id IN (?)", id).Unscoped().Delete(&nav.Nav{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 提交事务
	return tx.Commit().Error
}
