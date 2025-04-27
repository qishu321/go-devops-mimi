package response

import "go-devops-mimi/server/model/cmdb"

type AgentListRsp struct {
	Total  int64        `json:"total"`
	Agents []cmdb.Agent `json:"agents"`
}

type AgentGroupListRsp struct {
	Total       int64             `json:"total"`
	AgentGroups []cmdb.AgentGroup `json:"agentGroups"`
}

type AgentVersionListRsp struct {
	Total         int64               `json:"total"`
	AgentVersions []cmdb.AgentVersion `json:"agentVersions"`
}
