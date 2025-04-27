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

type NodeLogic struct{}

// Add 添加数据
func (l NodeLogic) Add(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*cmdbReq.NodesCreateReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	// 获取当前用户
	ctxUser, err := userService.GetCurrentLoginUser(c)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取当前登陆用户信息失败"))
	}
	// list, err := NodeGroupService.Info(r.GroupID)
	// if err != nil {
	// 	return nil, tools.NewMySqlError(fmt.Errorf("获取服务器组信息失败"))
	// }
	var status int8
	//判断ssh是否成功
	sshconfig := &tools.SSHClientConfig{
		Timeout:    time.Second * time.Duration(5+r.Timeout),
		UserName:   r.UserName,
		AuthModel:  r.AuthModel,
		Password:   r.Password,
		PrivateKey: r.PrivateKey,
		Port:       r.SSHPort,
		PublicIP:   r.PublicIP,
	}
	_, err = tools.SshCommand(sshconfig, "hostname")
	if err != nil {
		status = 2
	} else {
		status = 1
	}

	Nodes := cmdb.Nodes{
		NodeName:   r.NodeName,
		Username:   r.UserName,
		PublicIP:   r.PublicIP,
		SSHPort:    r.SSHPort,
		AuthModel:  r.AuthModel,
		Password:   tools.EncodeStr2Base64(r.Password),
		PrivateKey: tools.EncodeStr2Base64(r.PrivateKey),
		Timeout:    r.Timeout,
		Label:      r.Label,
		Status:     status,
		Creator:    ctxUser.Username,
	}

	// 创建数据
	err = NodeService.Add(&Nodes, r.GroupID)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("创建数据失败: %s", err.Error()))
	}

	return nil, nil
}

// List 数据列表
func (l NodeLogic) List(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*cmdbReq.NodesListReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	// 获取数据列表
	list, err := NodeService.List(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据列表失败: %s", err.Error()))
	}

	rets := make([]cmdb.Nodes, 0)
	for _, nodes := range list {
		// nodes.PrivateKey = "******"
		// nodes.Password = "******"
		rets = append(rets, *nodes)
	}
	count, err := NodeService.ListCount(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据总数失败"))
	}

	return response.NodeListRsp{
		Total:     count,
		NodeLists: rets,
	}, nil
}

// Update 更新数据
func (l NodeLogic) Update(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*cmdbReq.NodesUpdateReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	filter := tools.H{"id": int(r.ID)}
	if !NodeService.Exist(filter) {
		return nil, tools.NewMySqlError(fmt.Errorf("数据不存在"))
	}
	list, err := NodeService.Info(r.ID)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("数据不存在"))
	}
	var status int8
	switch r.AuthModel {
	case "password":
		if r.Password != list.Password {
			r.Password = tools.EncodeStr2Base64(r.Password)
		} else {
			r.Password = list.Password
		}
	case "private_key":
		if r.PrivateKey != list.PrivateKey {
			r.PrivateKey = tools.EncodeStr2Base64(r.PrivateKey)
		} else {
			r.PrivateKey = list.PrivateKey
		}
	}

	//判断ssh是否成功
	sshconfig := &tools.SSHClientConfig{
		Timeout:    time.Second * time.Duration(5+r.Timeout),
		UserName:   r.UserName,
		AuthModel:  r.AuthModel,
		Password:   tools.DecodeStrFromBase64(r.Password),
		PrivateKey: tools.DecodeStrFromBase64(r.PrivateKey),
		Port:       r.SSHPort,
		PublicIP:   r.PublicIP,
	}
	_, err = tools.SshCommand(sshconfig, "hostname")
	if err != nil {
		status = 2
	} else {
		status = 1
	}

	Nodes := cmdb.Nodes{
		Model:      gorm.Model{ID: uint(r.ID)},
		NodeName:   r.NodeName,
		Username:   r.UserName,
		PublicIP:   r.PublicIP,
		SSHPort:    r.SSHPort,
		AuthModel:  r.AuthModel,
		Password:   r.Password,
		PrivateKey: r.PrivateKey,
		Timeout:    r.Timeout,
		Label:      r.Label,
		Status:     status,
	}
	err = NodeService.Update(&Nodes, r.GroupID)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("更新数据失败: %s", err.Error()))
	}
	return nil, nil
}

// Delete 删除数据
func (l NodeLogic) Delete(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*cmdbReq.NodesDeleteReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	// 删除数据
	err := NodeService.Delete(r.Ids)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("删除数据失败: %s", err.Error()))
	}
	return nil, nil
}
func (l NodeLogic) AddNodesGroup(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*cmdbReq.AddNodesGroupReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	// 删除数据
	err := NodeService.AddNodeToGroup(r.NodeIDs, r.GroupIDs)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("更新数据失败: %s", err.Error()))
	}
	return nil, nil
}
