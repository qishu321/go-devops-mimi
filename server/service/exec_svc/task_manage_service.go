package exec_svc

import (
	"fmt"
	"strings"

	"go-devops-mimi/server/model/exec"
	"go-devops-mimi/server/model/exec/request"
	"go-devops-mimi/server/public/common"
	"go-devops-mimi/server/public/tools"
)

type TaskManageService struct{}

// List 获取数据列表
func (s TaskManageService) List(req *request.TaskManageListReq) ([]*exec.TaskManage, error) {
	var list []*exec.TaskManage
	db := common.DB.Model(&exec.TaskManage{}).Order("created_at DESC")

	name := strings.TrimSpace(req.Name)
	if name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	pageReq := tools.NewPageOption(req.PageNum, req.PageSize)
	err := db.Offset(pageReq.PageNum).Limit(pageReq.PageSize).Preload("Tasks").Find(&list).Error
	return list, err
}

// ListCount 获取数据总数
func (s TaskManageService) ListCount(req *request.TaskManageListReq) (int64, error) {
	var count int64
	db := common.DB.Model(&exec.TaskManage{}).Order("created_at DESC")

	name := strings.TrimSpace(req.Name)
	if name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	err := db.Count(&count).Error
	return count, err
}

// Add 添加资源
func (s TaskManageService) Add(dataObj *exec.TaskManage) error {
	return common.DB.Preload("Tasks").Create(dataObj).Error
}

// Update 更新资源
func (s *TaskManageService) Update(input *exec.TaskManage) error {
	// 开启事务，defer 回滚；Commit 成功后回滚会被忽略
	tx := common.DB.Begin()
	defer tx.Rollback()
	err := tx.Exec("DELETE FROM t_task_s WHERE task_manage_id IN (?)", input.ID).Error
	if err != nil {
		return err
	}
	// 1️⃣ 更新主表（只更新可变字段）
	if err := tx.Model(&exec.TaskManage{}).Where("id = ?", input.ID).Updates(map[string]interface{}{
		"name": input.Name,
		"args": input.Args,
		"desc": input.Desc,
		// Creator 一般不变，这里可按需添加
	}).Error; err != nil {
		return fmt.Errorf("update TaskManage failed: %w", err)
	}
	// 2️⃣ 遍历子任务
	for _, t := range input.Tasks {
		if t.ID != 0 {
			// —— 已有子任务：更新属性
			if err := tx.
				Model(&exec.Task{}).
				Where("id = ?", t.ID).
				Updates(map[string]interface{}{
					"name":     t.Name,
					"type":     t.Type,
					"content":  t.Content,
					"sort":     t.Sort,
					"timeout":  t.Timeout,
					"node_ids": t.NodesIDs,
				}).Error; err != nil {
				return fmt.Errorf("update Task[%d] failed: %w", t.ID, err)
			}
			// 再在关联表插入记录
			if err := tx.Exec(
				"INSERT INTO t_task_s (task_manage_id, task_id) VALUES (?, ?)",
				input.ID, t.ID,
			).Error; err != nil {
				return fmt.Errorf("insert t_task_s(%d,%d) failed: %w", input.ID, t.ID, err)
			}

		} else {
			// —— 新子任务：先创建 Task
			t.Creator = input.Creator // 根据需要把创建人沿用下来
			if err := tx.Create(&t).Error; err != nil {
				return fmt.Errorf("create new Task failed: %w", err)
			}
			// 再在关联表插入记录
			if err := tx.Exec(
				"INSERT INTO t_task_s (task_manage_id, task_id) VALUES (?, ?)",
				input.ID, t.ID,
			).Error; err != nil {
				return fmt.Errorf("insert t_task_s(%d,%d) failed: %w", input.ID, t.ID, err)
			}
		}
	}

	// 3️⃣ 提交事务
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("commit transaction failed: %w", err)
	}
	return nil
}

func (s TaskManageService) Info(id uint, name string) (*exec.TaskManage, error) {
	var server *exec.TaskManage
	err := common.DB.Preload("Tasks").Where("id = ? and name = ?", id, name).First(&server).Error
	return server, err
}

// Delete 批量删除
func (s TaskManageService) Delete(ids []uint) error {
	// 开启事务
	tx := common.DB.Begin()
	// 删除服务器记录
	if err := tx.Where("id IN (?)", ids).Unscoped().Delete(&exec.TaskManage{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 清空每个服务器记录的关联的 Keys
	err := tx.Exec("DELETE FROM t_task_s WHERE task_manage_id IN (?)", ids).Error
	if err != nil {
		return err
	}
	// 提交事务
	return tx.Commit().Error
}
