package response

import "go-devops-mimi/server/model/example"

type CloudAccountListRsp struct {
	Total         int64                  `json:"total"`
	CloudAccounts []example.CloudAccount `json:"cloudAccounts"`
}
