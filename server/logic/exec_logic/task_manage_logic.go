package exec_logic

import (
	"fmt"

	"go-devops-mimi/server/model/exec"
	execReq "go-devops-mimi/server/model/exec/request"
	"go-devops-mimi/server/model/exec/response"

	"go-devops-mimi/server/public/tools"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ManageLogic struct {
}

func (l ManageLogic) Add(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*execReq.TaskManageAddReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	// 获取当前用户
	ctxUser, err := userService.GetCurrentLoginUser(c)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取当前登陆用户信息失败"))
	}
	var tasks []*exec.Task
	for _, r := range r.Tasks {
		t := &exec.Task{
			Name:     r.Name,
			Type:     r.Type,
			Content:  r.Content,
			Sort:     r.Sort,
			Timeout:  r.Timeout,
			NodesIDs: r.NodeIDs,
			Creator:  ctxUser.Username,
		}
		tasks = append(tasks, t)
	}

	taskManage := &exec.TaskManage{
		Name:    r.Name,
		Args:    string(r.Args),
		Tasks:   tasks,
		Desc:    r.Desc,
		Creator: ctxUser.Username,
	}
	// 添加数据
	err = TaskManageService.Add(taskManage)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("添加数据失败: %s", err.Error()))
	}
	// 添加成功
	return nil, nil
}

func (l ManageLogic) Update(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*execReq.TaskManageUpdateReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	// 获取当前用户
	ctxUser, err := userService.GetCurrentLoginUser(c)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取当前登陆用户信息失败"))
	}
	var tasks []*exec.Task
	for _, r := range r.Tasks {
		t := &exec.Task{
			Model:    gorm.Model{ID: r.ID},
			Name:     r.Name,
			Type:     r.Type,
			Content:  r.Content,
			Sort:     r.Sort,
			Timeout:  r.Timeout,
			NodesIDs: r.NodeIDs,
			Creator:  ctxUser.Username,
		}
		tasks = append(tasks, t)
	}
	taskManage := &exec.TaskManage{
		Model:   gorm.Model{ID: r.ID},
		Name:    r.Name,
		Args:    r.Args,
		Tasks:   tasks,
		Desc:    r.Desc,
		Creator: ctxUser.Username,
	}
	// 更新数据
	err = TaskManageService.Update(taskManage)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("更新数据失败: %s", err.Error()))
	}
	return nil, nil
}

func (l ManageLogic) List(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*execReq.TaskManageListReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	// 获取数据列表
	list, err := TaskManageService.List(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据列表失败: %s", err.Error()))
	}

	rets := make([]exec.TaskManage, 0)
	for _, nodes := range list {
		rets = append(rets, *nodes)
	}
	count, err := TaskManageService.ListCount(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据总数失败"))
	}

	return response.TaskManageListRsp{
		Total:       count,
		TaskManages: rets,
	}, nil
}

// Delete 删除数据
func (l ManageLogic) Delete(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*execReq.TaskManageDeleteReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	// 删除数据
	err := TaskManageService.Delete(r.Ids)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("删除数据失败: %s", err.Error()))
	}
	return nil, nil
}

// / Info 查看数据详情
func (s ManageLogic) Info(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*execReq.TaskManageInfoReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	list, err := TaskManageService.Info(r.ID, r.Name)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据详情失败: %s", err.Error()))
	}
	return list, nil
}
