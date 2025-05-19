package nav

import (
	"go-devops-mimi/server/model/nav"
	"go-devops-mimi/server/public/common"
)

type Link_svc struct{}

// Add 添加资源
func (s Link_svc) Add(dataObj *nav.Link, navid uint) error {
	if err := common.DB.Create(dataObj).Error; err != nil {
		return err
	}
	// 关联分类
	if err := common.DB.Model(&nav.Nav{ID: navid}).Association("Links").Append(dataObj); err != nil {
		return err
	}

	return nil

}

// Update 更新资源
func (s Link_svc) Update(dataObj *nav.Link) error {
	return common.DB.Model(dataObj).Where("id = ?", dataObj.ID).Save(dataObj).Error
}

func (s Link_svc) Info(id uint) (*nav.Link, error) {
	var node *nav.Link
	err := common.DB.Where("id = ?", id).First(&node).Error
	return node, err

}

// GetUserByIds 根据ID获取node排序最小值
func (s Link_svc) GetUserByIds(ids []uint) ([]nav.Link, error) {
	var nodeList []nav.Link
	err := common.DB.Where("id IN (?)", ids).Find(&nodeList).Error
	return nodeList, err
}

// 删除链接
func (s Link_svc) Delete(id uint) error {
	err := common.DB.Where("id = ?", id).Unscoped().Delete(&nav.Link{}).Error
	if err != nil {
		return err
	}
	return nil
}
