package system

import (
	"fmt"

	"go-devops-mimi/server/model/system"
	systemReq "go-devops-mimi/server/model/system/request"

	"go-devops-mimi/server/public/tools"

	"github.com/gin-gonic/gin"
)

type MenuLogic struct{}

// Add 添加数据
func (l MenuLogic) Add(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*systemReq.MenuAddReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	if menuService.Exist(tools.H{"name": r.Name}) {
		return nil, tools.NewMySqlError(fmt.Errorf("菜单名称已存在"))

	}

	// 获取当前用户
	ctxUser, err := userService.GetCurrentLoginUser(c)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取当前登陆用户信息失败"))
	}

	menu := system.Menu{
		Name:       r.Name,
		Title:      r.Title,
		Icon:       r.Icon,
		Path:       r.Path,
		Redirect:   r.Redirect,
		Component:  r.Component,
		Sort:       r.Sort,
		Status:     r.Status,
		Hidden:     r.Hidden,
		NoCache:    r.NoCache,
		AlwaysShow: r.AlwaysShow,
		Breadcrumb: r.Breadcrumb,
		ActiveMenu: r.ActiveMenu,
		ParentId:   r.ParentId,
		Creator:    ctxUser.Username,
	}

	err = menuService.Add(&menu)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("创建记录失败: %s", err.Error()))
	}

	return nil, nil
}

// // List 数据列表
// func (l MenuLogic) List(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
// 	_, ok := req.(*systemReq.MenuListReq)
// 	if !ok {
// 		return nil, tools.ReqAssertErr
// 	}
// 	_ = c

// 	menus, err := menuService.List()
// 	if err != nil {
// 		return nil, tools.NewMySqlError(fmt.Errorf("获取资源列表失败: %s", err.Error()))
// 	}

// 	rets := make([]system.Menu, 0)
// 	for _, menu := range menus {
// 		rets = append(rets, *menu)
// 	}
// 	count, err := menuService.Count()
// 	if err != nil {
// 		return nil, tools.NewMySqlError(fmt.Errorf("获取资源总数失败"))
// 	}

// 	return response.MenuListRsp{
// 		Total: count,
// 		Menus: rets,
// 	}, nil
// }

// Update 更新数据
func (l MenuLogic) Update(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*systemReq.MenuUpdateReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	filter := tools.H{"id": int(r.ID)}
	if !menuService.Exist(filter) {
		return nil, tools.NewMySqlError(fmt.Errorf("该ID对应的记录不存在"))
	}

	// 获取当前登陆用户
	ctxUser, err := userService.GetCurrentLoginUser(c)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取当前登陆用户失败"))
	}

	oldData := new(system.Menu)
	err = menuService.Find(filter, oldData)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取记录失败: %s", err.Error()))
	}

	menu := system.Menu{
		Model:      oldData.Model,
		Name:       r.Name,
		Title:      r.Title,
		Icon:       r.Icon,
		Path:       r.Path,
		Redirect:   r.Redirect,
		Component:  r.Component,
		Sort:       r.Sort,
		Status:     r.Status,
		Hidden:     r.Hidden,
		NoCache:    r.NoCache,
		AlwaysShow: r.AlwaysShow,
		Breadcrumb: r.Breadcrumb,
		ActiveMenu: r.ActiveMenu,
		ParentId:   r.ParentId,
		Creator:    ctxUser.Username,
	}

	err = menuService.Update(&menu)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("更新记录失败: %s", err.Error()))
	}

	return nil, nil
}

// Delete 删除数据
func (l MenuLogic) Delete(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*systemReq.MenuDeleteReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	for _, id := range r.MenuIds {
		filter := tools.H{"id": int(id)}
		if !menuService.Exist(filter) {
			return nil, tools.NewMySqlError(fmt.Errorf("该ID对应的记录不存在"))
		}
	}

	// 删除接口
	err := menuService.Delete(r.MenuIds)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("删除接口失败: %s", err.Error()))
	}
	return nil, nil
}

// GetTree 数据树
func (l MenuLogic) GetTree(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	_, ok := req.(*systemReq.MenuGetTreeReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	menus, err := menuService.List()
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取资源列表失败: " + err.Error()))
	}

	tree := genMenuTree(0, menus)

	return tree, nil
}

// GetAccessTree 获取用户菜单树
func (l MenuLogic) GetAccessTree(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*systemReq.MenuGetAccessTreeReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	// 校验
	filter := tools.H{"id": r.ID}
	if !userService.Exist(filter) {
		return nil, tools.NewValidatorError(fmt.Errorf("该用户不存在"))
	}
	user := new(system.User)
	err := userService.Find(filter, user)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("在MySQL查询用户失败: " + err.Error()))
	}
	var roleIds []uint
	for _, role := range user.Roles {
		roleIds = append(roleIds, role.ID)
	}
	menus, err := menuService.ListUserMenus(roleIds)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取资源列表失败: " + err.Error()))
	}

	tree := genMenuTree(0, menus)

	return tree, nil
}
