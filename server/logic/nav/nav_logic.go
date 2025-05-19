package nav

import (
	"fmt"

	"go-devops-mimi/server/model/nav"
	navReq "go-devops-mimi/server/model/nav/request"
	"go-devops-mimi/server/model/nav/response"
	"go-devops-mimi/server/public/tools"

	"github.com/gin-gonic/gin"
)

type NavLogic struct{}

// Add 添加数据
func (l NavLogic) Add(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*navReq.NavAddRequest)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	list := nav.Nav{
		Name:    r.Name,
		NavSort: r.NavSort,
	}

	// 创建数据
	err := NavService.Add(&list)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("创建数据失败: %s", err.Error()))
	}

	return nil, nil
}

// List 数据列表
func (l NavLogic) List(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*navReq.NavListRequest)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	// 获取数据列表
	list, err := NavService.List(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据列表失败: %s", err.Error()))
	}

	rets := make([]nav.Nav, 0)
	for _, n := range list {
		rets = append(rets, *n)
	}
	count, err := NavService.Count()
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据总数失败"))
	}

	return response.NavListRsp{
		Total:    count,
		NavLists: rets,
	}, nil
}

// Update 更新数据
func (l NavLogic) Update(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*navReq.NavUpdateRequest)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	list := nav.Nav{
		ID:      r.ID,
		Name:    r.Name,
		NavSort: r.NavSort,
	}
	err := NavService.Update(&list)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("更新数据失败: %s", err.Error()))
	}
	return nil, nil
}

// / Info 查看数据详情
func (s NavLogic) Info(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*navReq.NavInfoRequest)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	list, err := NavService.Info(r.ID)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据详情失败: %s", err.Error()))
	}
	return list, nil
}

// Delete 只删除分类
func (l NavLogic) Delete(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*navReq.NavDeleteRequest)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	// 删除数据
	err := NavService.Delete(r.ID)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("删除数据失败: %s", err.Error()))
	}
	return nil, nil
}

// DeleteAll 删除分类和关联的链接
func (l NavLogic) DeleteAll(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*navReq.NavDeleteRequest)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	// 删除数据
	err := NavService.DeleteAll(r.ID)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("删除数据失败: %s", err.Error()))
	}
	return nil, nil
}
