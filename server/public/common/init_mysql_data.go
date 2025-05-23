package common

import (
	"errors"

	"go-devops-mimi/server/config"
	"go-devops-mimi/server/model/system"
	"go-devops-mimi/server/public/tools"

	"github.com/thoas/go-funk"
	"gorm.io/gorm"
)

// 初始化mysql数据
func InitData() {
	// 是否初始化数据
	if !config.Conf.System.InitData {
		return
	}

	// 1.写入角色数据
	newRoles := make([]*system.Role, 0)
	roles := []*system.Role{
		{
			Model:   gorm.Model{ID: 1},
			Name:    "管理员",
			Keyword: "admin",
			Remark:  "",
			Sort:    1,
			Status:  1,
			Creator: "系统",
		},
		{
			Model:   gorm.Model{ID: 2},
			Name:    "普通用户",
			Keyword: "user",
			Remark:  "",
			Sort:    3,
			Status:  1,
			Creator: "系统",
		},
		{
			Model:   gorm.Model{ID: 3},
			Name:    "访客",
			Keyword: "guest",
			Remark:  "",
			Sort:    5,
			Status:  1,
			Creator: "系统",
		},
	}

	for _, role := range roles {
		err := DB.First(&role, role.ID).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newRoles = append(newRoles, role)
		}
	}

	if len(newRoles) > 0 {
		err := DB.Create(&newRoles).Error
		if err != nil {
			Log.Errorf("写入系统角色数据失败：%v", err)
		}
	}

	// 2写入菜单
	newMenus := make([]system.Menu, 0)
	var uint0 uint = 0
	var uint1 uint = 1
	var uint7 uint = 7
	var uint9 uint = 9
	menus := []system.Menu{
		{
			Model:     gorm.Model{ID: 1},
			Name:      "System",
			Title:     "系统管理",
			Icon:      "system",
			Path:      "/system",
			Component: "Layout",
			Redirect:  "/system/user",
			Sort:      1,
			ParentId:  uint0,
			Roles:     roles[:1],
			Creator:   "系统",
		},
		{
			Model:     gorm.Model{ID: 2},
			Name:      "User",
			Title:     "用户管理",
			Icon:      "user",
			Path:      "user",
			Component: "/system/user/index",
			Sort:      2,
			ParentId:  uint1,
			Roles:     roles[:1],
			Creator:   "系统",
		},
		{
			Model:     gorm.Model{ID: 3},
			Name:      "Group",
			Title:     "分组管理",
			Icon:      "peoples",
			Path:      "group",
			Component: "/system/group/index",
			Sort:      3,
			ParentId:  uint1,
			NoCache:   1,
			Roles:     roles[:1],
			Creator:   "系统",
		},
		{
			Model:     gorm.Model{ID: 4},
			Name:      "Role",
			Title:     "角色管理",
			Icon:      "eye-open",
			Path:      "role",
			Component: "/system/role/index",
			Sort:      4,
			ParentId:  uint1,
			Roles:     roles[:1],
			Creator:   "系统",
		},
		{
			Model:     gorm.Model{ID: 5},
			Name:      "Menu",
			Title:     "菜单管理",
			Icon:      "tree-table",
			Path:      "menu",
			Component: "/system/menu/index",
			Sort:      5,
			ParentId:  uint1,
			Roles:     roles[:1],
			Creator:   "系统",
		},
		{
			Model:     gorm.Model{ID: 6},
			Name:      "Api",
			Title:     "接口管理",
			Icon:      "tree",
			Path:      "api",
			Component: "/system/api/index",
			Sort:      6,
			ParentId:  uint1,
			Roles:     roles[:1],
			Creator:   "系统",
		},
		{
			Model:     gorm.Model{ID: 7},
			Name:      "Log",
			Title:     "日志管理",
			Icon:      "log",
			Path:      "log",
			Component: "/system/log/index",
			Redirect:  "/system/log/operationLog",
			Sort:      7,
			ParentId:  uint1,
			Roles:     roles[:2],
			Creator:   "系统",
		},
		{
			Model:     gorm.Model{ID: 8},
			Name:      "OperationLog",
			Title:     "操作日志",
			Icon:      "documentation",
			Path:      "operationLog",
			Component: "/system/log/operationLog/index",
			Sort:      8,
			ParentId:  uint7,
			Roles:     roles[:2],
			Creator:   "系统",
		},
		{
			Model:     gorm.Model{ID: 9},
			Name:      "Example",
			Title:     "示例模块",
			Icon:      "example",
			Path:      "/example",
			Component: "Layout",
			Sort:      9,
			ParentId:  uint0,
			Roles:     roles[:2],
			Creator:   "系统",
		},
		{
			Model:     gorm.Model{ID: 10},
			Name:      "CloudAccount",
			Title:     "云账户",
			Icon:      "peoples",
			Path:      "CloudAccount",
			Component: "/example/cloudAccount/index",
			Sort:      10,
			ParentId:  uint9,
			Roles:     roles[:2],
			Creator:   "系统",
		},
	}
	for _, menu := range menus {
		err := DB.First(&menu, menu.ID).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newMenus = append(newMenus, menu)
		}
	}
	if len(newMenus) > 0 {
		err := DB.Create(&newMenus).Error
		if err != nil {
			Log.Errorf("写入系统菜单数据失败：%v", err)
		}
	}

	// 3.写入用户
	newUsers := make([]*system.User, 0)
	users := []*system.User{
		{
			Model:         gorm.Model{ID: 1},
			Username:      "admin",
			Password:      tools.NewGenPasswd("123456"),
			Nickname:      "管理员",
			GivenName:     "最强后台",
			Mail:          "admin@eryajf.net",
			JobNumber:     "0000",
			Mobile:        "18888888888",
			Avatar:        "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
			PostalAddress: "中国河南省南阳市",
			Departments:   "ops",
			Position:      "系统管理员",
			Introduction:  "最强后台的管理员",
			Status:        1,
			Creator:       "系统",
			Roles:         roles[:1],
		},
	}

	for _, user := range users {
		err := DB.First(&user, user.ID).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newUsers = append(newUsers, user)
		}
	}

	if len(newUsers) > 0 {
		err := DB.Create(&newUsers).Error
		if err != nil {
			Log.Errorf("写入用户数据失败：%v", err)
		}
	}

	// 4.写入api
	apis := []system.Api{
		{
			Method:   "POST",
			Path:     "/system/base/login",
			Category: "base",
			Remark:   "用户登录",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/system/base/logout",
			Category: "base",
			Remark:   "用户登出",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/system/base/refreshToken",
			Category: "base",
			Remark:   "刷新JWT令牌",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/system/user/info",
			Category: "user",
			Remark:   "获取当前登录用户信息",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/system/user/list",
			Category: "user",
			Remark:   "获取用户列表",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/system/user/changePwd",
			Category: "user",
			Remark:   "更新用户登录密码",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/system/user/add",
			Category: "user",
			Remark:   "创建用户",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/system/user/update",
			Category: "user",
			Remark:   "更新用户",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/system/user/delete",
			Category: "user",
			Remark:   "批量删除用户",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/system/user/changeUserStatus",
			Category: "user",
			Remark:   "更改用户状态",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/system/group/list",
			Category: "group",
			Remark:   "获取分组列表",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/system/group/tree",
			Category: "group",
			Remark:   "获取分组列表树",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/system/group/add",
			Category: "group",
			Remark:   "创建分组",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/system/group/update",
			Category: "group",
			Remark:   "更新分组",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/system/group/delete",
			Category: "group",
			Remark:   "批量删除分组",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/system/group/adduser",
			Category: "group",
			Remark:   "添加用户到分组",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/system/group/removeuser",
			Category: "group",
			Remark:   "将用户从分组移出",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/system/group/useringroup",
			Category: "group",
			Remark:   "获取在分组内的用户列表",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/system/group/usernoingroup",
			Category: "group",
			Remark:   "获取不在分组内的用户列表",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/system/role/list",
			Category: "role",
			Remark:   "获取角色列表",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/system/role/add",
			Category: "role",
			Remark:   "创建角色",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/system/role/update",
			Category: "role",
			Remark:   "更新角色",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/system/role/getmenulist",
			Category: "role",
			Remark:   "获取角色的权限菜单",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/system/role/updatemenus",
			Category: "role",
			Remark:   "更新角色的权限菜单",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/system/role/getapilist",
			Category: "role",
			Remark:   "获取角色的权限接口",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/system/role/updateapis",
			Category: "role",
			Remark:   "更新角色的权限接口",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/system/role/delete",
			Category: "role",
			Remark:   "批量删除角色",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/system/menu/list",
			Category: "menu",
			Remark:   "获取菜单列表",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/system/menu/tree",
			Category: "menu",
			Remark:   "获取菜单树",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/system/menu/access/tree",
			Category: "menu",
			Remark:   "获取用户菜单树",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/system/menu/add",
			Category: "menu",
			Remark:   "创建菜单",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/system/menu/update",
			Category: "menu",
			Remark:   "更新菜单",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/system/menu/delete",
			Category: "menu",
			Remark:   "批量删除菜单",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/system/api/list",
			Category: "api",
			Remark:   "获取接口列表",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/system/api/tree",
			Category: "api",
			Remark:   "获取接口树",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/system/api/add",
			Category: "api",
			Remark:   "创建接口",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/system/api/update",
			Category: "api",
			Remark:   "更新接口",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/system/api/delete",
			Category: "api",
			Remark:   "批量删除接口",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/system/log/operation/list",
			Category: "log",
			Remark:   "获取操作日志列表",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/system/log/operation/delete",
			Category: "log",
			Remark:   "批量删除操作日志",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/example/cloudaccount/list",
			Category: "cloudAccount",
			Remark:   "获取云账户列表",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/example/cloudaccount/add",
			Category: "cloudAccount",
			Remark:   "添加云账户",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/example/cloudaccount/update",
			Category: "cloudAccount",
			Remark:   "更新云账户信息",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/example/cloudaccount/delete",
			Category: "cloudAccount",
			Remark:   "批量删除云账户",
			Creator:  "系统",
		},
	}

	// 5. 将角色绑定给菜单
	newApi := make([]system.Api, 0)
	newRoleCasbin := make([]system.RoleCasbin, 0)
	for i, api := range apis {
		api.ID = uint(i + 1)
		err := DB.First(&api, api.ID).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newApi = append(newApi, api)

			// 管理员拥有所有API权限
			newRoleCasbin = append(newRoleCasbin, system.RoleCasbin{
				Keyword: roles[0].Keyword,
				Path:    api.Path,
				Method:  api.Method,
			})

			// 非管理员拥有基础权限
			basePaths := []string{
				"/system/base/login",
				"/system/base/logout",
				"/system/base/refreshToken",
				"/system/base/changePwd",
				"/system/base/dashboard",
				"/system/user/info",
				"/system/user/changePwd",
				"/system/menu/access/tree",
				"/system/log/operation/list",
			}

			if funk.ContainsString(basePaths, api.Path) {
				newRoleCasbin = append(newRoleCasbin, system.RoleCasbin{
					Keyword: roles[1].Keyword,
					Path:    api.Path,
					Method:  api.Method,
				})
			}
		}
	}

	if len(newApi) > 0 {
		if err := DB.Create(&newApi).Error; err != nil {
			Log.Errorf("写入api数据失败：%v", err)
		}
	}

	if len(newRoleCasbin) > 0 {
		rules := make([][]string, 0)
		for _, c := range newRoleCasbin {
			rules = append(rules, []string{
				c.Keyword, c.Path, c.Method,
			})
		}
		isAdd, err := CasbinEnforcer.AddPolicies(rules)
		if !isAdd {
			Log.Errorf("写入casbin数据失败：%v", err)
		}
	}

	// 6.写入分组
	newGroups := make([]system.Group, 0)
	groups := []system.Group{
		{
			Model:     gorm.Model{ID: 1},
			GroupName: "root",
			Remark:    "根部门",
			Creator:   "system",
			ParentId:  0,
			Source:    "system",
		},
		{
			Model:     gorm.Model{ID: 2},
			GroupName: "backend",
			Remark:    "后端部",
			Creator:   "system",
			ParentId:  1,
			Source:    "system",
		},
		{
			Model:     gorm.Model{ID: 3},
			GroupName: "test",
			Remark:    "测试部",
			Creator:   "system",
			ParentId:  1,
			Source:    "system",
		},
		{
			Model:     gorm.Model{ID: 4},
			GroupName: "ops",
			Remark:    "运维部",
			Creator:   "system",
			ParentId:  1,
			Source:    "system",
		},
	}

	for _, group := range groups {
		err := DB.First(&group, group.ID).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newGroups = append(newGroups, group)
		}
	}
	if len(newGroups) > 0 {
		err := DB.Create(&newGroups).Error
		if err != nil {
			Log.Errorf("写入分组数据失败：%v", err)
		}
	}
}
