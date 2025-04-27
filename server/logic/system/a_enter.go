package system

import (
	"go-devops-mimi/server/model/system"
	"go-devops-mimi/server/service"
)

type LogicGroup struct {
	ApiLogic
	BaseLogic
	GroupLogic
	MenuLogic
	OperationLogLogic
	RoleLogic
	UserLogic
}

// 初始化 service
var (
	apiService          = service.ServiceGroupApp.SystemServiceGroup.ApiService
	groupService        = service.ServiceGroupApp.SystemServiceGroup.GroupService
	menuService         = service.ServiceGroupApp.SystemServiceGroup.MenuService
	operationLogService = service.ServiceGroupApp.SystemServiceGroup.OperationLogService
	roleService         = service.ServiceGroupApp.SystemServiceGroup.RoleService
	userService         = service.ServiceGroupApp.SystemServiceGroup.UserService
)

// genMenuTree 生成菜单树
func genMenuTree(parentId uint, menus []*system.Menu) []*system.Menu {
	tree := make([]*system.Menu, 0)

	for _, m := range menus {
		if m.ParentId == parentId {
			children := genMenuTree(m.ID, menus)
			m.Children = children
			tree = append(tree, m)
		}
	}
	return tree
}

// genGroupTree 生成分组树
func genGroupTree(parentId uint, groups []*system.Group) []*system.Group {
	tree := make([]*system.Group, 0)

	for _, g := range groups {
		if g.ParentId == parentId {
			children := genGroupTree(g.ID, groups)
			g.Children = children
			tree = append(tree, g)
		}
	}
	return tree
}
