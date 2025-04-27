package response

import "go-devops-mimi/server/model/system"

type UserListRsp struct {
	Total int           `json:"total"`
	Users []system.User `json:"users"`
}
