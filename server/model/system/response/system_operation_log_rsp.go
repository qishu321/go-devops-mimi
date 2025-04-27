package response

import "go-devops-mimi/server/model/system"

type LogListRsp struct {
	Total int64                 `json:"total"`
	Logs  []system.OperationLog `json:"logs"`
}
