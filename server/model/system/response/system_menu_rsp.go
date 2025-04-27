package response

import "go-devops-mimi/server/model/system"

type MenuListRsp struct {
	Total int64         `json:"total"`
	Menus []system.Menu `json:"menus"`
}
