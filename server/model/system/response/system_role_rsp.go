package response

import "go-devops-mimi/server/model/system"

type RoleListRsp struct {
	Total int64         `json:"total"`
	Roles []system.Role `json:"roles"`
}
