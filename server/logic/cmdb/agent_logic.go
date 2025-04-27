package cmdb

import (
	"fmt"
	"time"

	"go-devops-mimi/server/model/cmdb"
	cmdbReq "go-devops-mimi/server/model/cmdb/request"
	"go-devops-mimi/server/model/cmdb/response"
	"go-devops-mimi/server/public/tools"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AgentLogic struct{}

// Add 添加数据
func (l AgentLogic) Add(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*cmdbReq.CreateAgentReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	if !tools.CheckSafeKey(r.AgentSafeParams.Token) {
		return
	}
	agentid := tools.IPNameToNum(r.AgentName, r.IP)
	code := AgentService.Check(agentid)
	if code != 200 {
		return tools.DataObj(fmt.Sprintf("Agent已存在，agent为:%d", agentid)), nil
	}
	list := cmdb.Agent{
		Agentid:       agentid,
		AgentName:     r.AgentName,
		IP:            r.IP,
		Online:        true,
		Version:       r.Version,
		OsType:        r.OsType,
		Label:         r.Label,
		LastHeartbeat: time.Now(),
	}
	// 创建数据
	err := AgentService.Add(&list)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("创建数据失败: %s", err.Error()))
	}

	return nil, nil
}

// List 数据列表
func (l AgentLogic) List(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*cmdbReq.AgentListReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	// 获取数据列表
	list, err := AgentService.List(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据列表失败: %s", err.Error()))
	}

	rets := make([]cmdb.Agent, 0)
	for _, list := range list {
		rets = append(rets, *list)
	}
	count, err := AgentService.ListCount(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据总数失败"))
	}

	return response.AgentListRsp{
		Total:  count,
		Agents: rets,
	}, nil
}

// Update 更新数据只能更新标签
func (l AgentLogic) Update(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*cmdbReq.UpdateAgentReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	if !tools.CheckSafeKey(r.AgentSafeParams.Token) {
		return
	}

	filter := tools.H{"id": int(r.ID)}
	if !AgentService.Exist(filter) {
		return nil, tools.NewMySqlError(fmt.Errorf("数据不存在"))
	}
	list := cmdb.Agent{
		Model: gorm.Model{ID: uint(r.ID)},
		Label: r.Label,
	}
	err := AgentService.Update(&list)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("更新数据失败: %s", err.Error()))
	}
	return nil, nil
}

func (l AgentLogic) AgentHeartbeat(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*cmdbReq.AgentHeartbeatReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	if !tools.CheckSafeKey(r.AgentSafeParams.Token) {
		return
	}
	list := cmdb.Agent{
		Agentid:       r.AgentID,
		Version:       r.Version,
		OsType:        r.OsType,
		LastHeartbeat: time.Now(),
	}
	err := AgentService.AgentHeartbeat(&list)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("更新数据失败: %s", err.Error()))
	}
	return nil, nil
}

// Delete 删除数据
func (l AgentLogic) Delete(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*cmdbReq.DeleteAgentReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	// 删除数据
	err := AgentService.Delete(r.Ids)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("删除数据失败: %s", err.Error()))
	}
	return nil, nil
}

// 添加服务器到分组
func (l AgentLogic) AddAgentToGroup(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*cmdbReq.AddAgentGroupReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	filter := tools.H{"id": r.GroupID}

	if !AgentService.Exist(filter) {
		return nil, tools.NewMySqlError(fmt.Errorf("分组不存在"))
	}

	list, err := AgentService.GetUserByIds(r.AgentIDs)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取agent列表失败: %s", err.Error()))
	}

	group := new(cmdb.AgentGroup)
	err = AgentGroupService.Find(filter, group)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取分组失败: %s", err.Error()))
	}

	// 先添加到MySQL
	err = AgentGroupService.AddAgentToGroup(group, list)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("添加agent到分组失败: %s", err.Error()))
	}

	return nil, nil
}

// 从分组移除服务器
func (l AgentLogic) RemoveAgentGroup(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*cmdbReq.RemoveAgentGroupReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	filter := tools.H{"id": r.GroupID}

	if !AgentService.Exist(filter) {
		return nil, tools.NewMySqlError(fmt.Errorf("分组不存在"))
	}

	list, err := AgentService.GetUserByIds(r.AgentIDs)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取agent列表失败: %s", err.Error()))
	}

	group := new(cmdb.AgentGroup)
	err = AgentGroupService.Find(filter, group)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取分组失败: %s", err.Error()))
	}

	// 先添加到MySQL
	err = AgentGroupService.RemoveAgentGroup(group, list)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("从分组移除服务器失败: %s", err.Error()))
	}

	return nil, nil
}
