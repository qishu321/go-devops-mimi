package response

import "go-devops-mimi/server/model/cmdb"

type NodeListRsp struct {
	Total     int64        `json:"total"`
	NodeLists []cmdb.Nodes `json:"nodeLists"`
}

type NodeGroupListRsp struct {
	Total          int64            `json:"total"`
	NodeGroupLists []cmdb.NodeGroup `json:"nodeGroupLists"`
}
