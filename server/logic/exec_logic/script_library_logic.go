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

type ScriptLibraryLogic struct {
}

func (l ScriptLibraryLogic) Add(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*execReq.ScriptLibraryAddReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	// 获取当前用户
	ctxUser, err := userService.GetCurrentLoginUser(c)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取当前登陆用户信息失败"))
	}
	library := &exec.ScriptLibrary{
		Name:    r.Name,
		Type:    r.Type,
		Content: r.Content,
		Desc:    r.Desc,
		Creator: ctxUser.Username,
	}
	// 添加数据
	err = ScriptLibraryService.Add(library)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("添加数据失败: %s", err.Error()))
	}
	// 添加成功
	return nil, nil
}

func (l ScriptLibraryLogic) Update(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*execReq.ScriptLibraryUpdateReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	// 获取当前用户
	ctxUser, err := userService.GetCurrentLoginUser(c)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取当前登陆用户信息失败"))
	}
	library := &exec.ScriptLibrary{
		Model:   gorm.Model{ID: r.ID},
		Name:    r.Name,
		Type:    r.Type,
		Content: r.Content,
		Desc:    r.Desc,
		Creator: ctxUser.Username,
	}
	// 更新数据
	err = ScriptLibraryService.Update(library)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("更新数据失败: %s", err.Error()))
	}
	return nil, nil
}

func (l ScriptLibraryLogic) List(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*execReq.ScriptLibraryListReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	// 获取数据列表
	list, err := ScriptLibraryService.List(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据列表失败: %s", err.Error()))
	}

	rets := make([]exec.ScriptLibrary, 0)
	for _, nodes := range list {
		rets = append(rets, *nodes)
	}
	count, err := ScriptLibraryService.ListCount(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据总数失败"))
	}

	return response.ScriptLibraryListRsp{
		Total:          count,
		ScriptLibrarys: rets,
	}, nil
}

// Delete 删除数据
func (l ScriptLibraryLogic) Delete(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*execReq.ScriptLibraryDeleteReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	// 删除数据
	err := ScriptLibraryService.Delete(r.Ids)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("删除数据失败: %s", err.Error()))
	}
	return nil, nil
}

// / Info 查看数据详情
func (s ScriptLibraryLogic) Info(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*execReq.ScriptLibraryInfoReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	list, err := ScriptLibraryService.Info(r.ID, r.Name)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据详情失败: %s", err.Error()))
	}
	return list, nil
}
