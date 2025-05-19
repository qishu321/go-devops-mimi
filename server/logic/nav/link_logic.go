package nav

import (
	"fmt"

	"go-devops-mimi/server/model/nav"
	navReq "go-devops-mimi/server/model/nav/request"
	"go-devops-mimi/server/public/tools"

	"github.com/gin-gonic/gin"
)

type LinkLogic struct{}

// Add 添加数据
func (l LinkLogic) Add(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*navReq.LinkAddRequest)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	list := nav.Link{
		Name: r.Name,
		Desc: r.Desc,
		URL:  r.URL,
	}

	// 创建数据
	err := LinkService.Add(&list, r.NavID)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("创建数据失败: %s", err.Error()))
	}

	return nil, nil
}

// Update 更新数据
func (l LinkLogic) Update(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*navReq.LinkUpdateRequest)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	list := nav.Link{
		ID:   r.ID,
		Name: r.Name,
		Desc: r.Desc,
		URL:  r.URL,
	}
	err := LinkService.Update(&list)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("更新数据失败: %s", err.Error()))
	}
	return nil, nil
}

// / Info 查看数据详情
func (s LinkLogic) Info(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*navReq.LinkInfoRequest)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	list, err := LinkService.Info(r.ID)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据详情失败: %s", err.Error()))
	}
	return list, nil
}

// Delete 只删除分类
func (l LinkLogic) Delete(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*navReq.LinkDeleteRequest)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	// 删除数据
	err := LinkService.Delete(r.ID)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("删除数据失败: %s", err.Error()))
	}
	return nil, nil
}
