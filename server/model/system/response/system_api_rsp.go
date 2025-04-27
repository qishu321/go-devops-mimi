package response

import "go-devops-mimi/server/model/system"

type ApiTreeRsp struct {
	ID       int           `json:"ID"`
	Remark   string        `json:"remark"`
	Category string        `json:"category"`
	Children []*system.Api `json:"children"`
}

type ApiListRsp struct {
	Total int64        `json:"total"`
	Apis  []system.Api `json:"apis"`
}
