package cmdb

import (
	"fmt"

	"go-devops-mimi/server/model/cmdb"
	cmdbReq "go-devops-mimi/server/model/cmdb/request"
	"go-devops-mimi/server/model/cmdb/response"
	"go-devops-mimi/server/public/tools"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NodeGroupLogic struct{}

// Add 添加数据
func (l NodeGroupLogic) Add(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*cmdbReq.NodeGroupCreateReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	// 获取当前用户
	ctxUser, err := userService.GetCurrentLoginUser(c)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取当前登陆用户信息失败"))
	}
	list := cmdb.NodeGroup{
		GroupName: r.GroupName,
		Desc:      r.Desc,
		Creator:   ctxUser.Username,
	}

	// 创建数据
	err = NodeGroupService.Add(&list)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("创建数据失败: %s", err.Error()))
	}

	return nil, nil
}

// List 数据列表
func (l NodeGroupLogic) List(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*cmdbReq.NodeGroupListReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	// 获取数据列表
	list, err := NodeGroupService.List(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据列表失败: %s", err.Error()))
	}

	rets := make([]cmdb.NodeGroup, 0)
	for _, nodes := range list {
		rets = append(rets, *nodes)
	}
	count, err := NodeGroupService.ListCount(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据总数失败"))
	}

	return response.NodeGroupListRsp{
		Total:          count,
		NodeGroupLists: rets,
	}, nil
}

// Update 更新数据
func (l NodeGroupLogic) Update(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*cmdbReq.NodeGroupUpdateReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	filter := tools.H{"id": int(r.ID)}
	if !NodeGroupService.Exist(filter) {
		return nil, tools.NewMySqlError(fmt.Errorf("数据不存在"))
	}
	list := cmdb.NodeGroup{
		Model:     gorm.Model{ID: uint(r.ID)},
		GroupName: r.GroupName,
		Desc:      r.Desc,
	}
	err := NodeGroupService.Update(&list)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("更新数据失败: %s", err.Error()))
	}
	return nil, nil
}

// Delete 删除数据
func (l NodeGroupLogic) Delete(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*cmdbReq.NodeGroupDeleteReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	// 删除数据
	err := NodeGroupService.Delete(r.Ids)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("删除数据失败: %s", err.Error()))
	}
	return nil, nil
}

// 添加服务器到分组
func (l NodeGroupLogic) AddNodeToGroup(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*cmdbReq.AddNodesToGroupReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	filter := tools.H{"id": r.GroupID}

	if !NodeGroupService.Exist(filter) {
		return nil, tools.NewMySqlError(fmt.Errorf("分组不存在"))
	}

	nodelist, err := NodeService.GetUserByIds(r.NodeIDs)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取node列表失败: %s", err.Error()))
	}

	group := new(cmdb.NodeGroup)
	err = NodeGroupService.Find(filter, group)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取分组失败: %s", err.Error()))
	}

	// 先添加到MySQL
	err = NodeGroupService.AddNodeToGroup(group, nodelist)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("添加node到分组失败: %s", err.Error()))
	}

	return nil, nil
}

// 从分组移除服务器
func (l NodeGroupLogic) RemoveNodeGroup(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*cmdbReq.RemoveNodeGroupReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	filter := tools.H{"id": r.GroupID}

	if !NodeGroupService.Exist(filter) {
		return nil, tools.NewMySqlError(fmt.Errorf("分组不存在"))
	}

	nodelist, err := NodeService.GetUserByIds(r.NodeIDs)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取node列表失败: %s", err.Error()))
	}

	group := new(cmdb.NodeGroup)
	err = NodeGroupService.Find(filter, group)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取分组失败: %s", err.Error()))
	}

	// 先添加到MySQL
	err = NodeGroupService.RemoveNodeGroup(group, nodelist)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("从分组移除服务器失败: %s", err.Error()))
	}

	return nil, nil
}
