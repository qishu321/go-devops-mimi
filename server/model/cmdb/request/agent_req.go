package request

type AddAgentGroupReq struct {
	GroupID  uint   `json:"groupId" binding:"required"`        // 必须指定分组ID
	AgentIDs []uint `json:"agentIds" binding:"required,min=1"` // 至少一个Agent
}
type RemoveAgentGroupReq struct {
	GroupID  uint   `json:"groupId" binding:"required"`        // 必须指定分组ID
	AgentIDs []uint `json:"agentIds" binding:"required,min=1"` // 至少一个Agent
}
type AgentSafeParams struct {
	Token string `json:"token" validate:"required"`
}

// agent通用删除请求结构体
type DeleteAgentReq struct {
	Ids []uint `json:"ids" binding:"required"` // 支持批量删除
}
type CreateAgentGroupReq struct {
	GroupName string `json:"groupName" binding:"required,min=2,max=50"`
	Desc      string `json:"desc" binding:"max=50"`
	Creator   string `json:"creator" binding:"required,min=1,max=20"`
}
type UpdateAgentGroupReq struct {
	ID        uint   `json:"id" binding:"required"`
	GroupName string `json:"groupName" binding:"omitempty,min=2,max=50"` // 指针类型支持空值更新
	Desc      string `json:"desc" binding:"omitempty,max=50"`
}
type AgentGroupListReq struct {
	GroupName string `json:"groupName" form:"groupName"` // 支持模糊查询
	// 分页参数
	PageNum  int `json:"pageNum" form:"pageNum"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

type CreateAgentReq struct {
	AgentSafeParams
	AgentName string `json:"agentName" binding:"required,min=2,max=150"`
	IP        string `json:"ip" form:"ip"`
	Version   string `json:"version"`
	OsType    string `json:"osType" binding:"required,oneof=windows linux macos"`
	Label     string `json:"label" binding:"max=50"`
}

type UpdateAgentReq struct {
	AgentSafeParams
	ID uint `json:"id" binding:"required" form:"id"`
	// AgentName string `json:"agentName" binding:"required,min=2,max=150"`
	// AgentID   string `json:"agentId" `
	// IP        string `json:"ip" form:"ip"`
	// Version   string `json:"Version"`
	// OsType    string `json:"osType" binding:"required,oneof=windows linux macos"`
	Label string `json:"label" binding:"max=50"`
}
type AgentHeartbeatReq struct {
	AgentSafeParams
	AgentID uint64 `json:"agent_id" binding:"required"`
	Version string `json:"Version"`
	OsType  string `json:"osType" binding:"required,oneof=windows linux macos"`
}

type AgentListReq struct {
	AgentName string `json:"agentName"`
	// 可按 IP 地址查询
	IP string `json:"ip" form:"ip"`
	// 分页参数
	PageNum  int `json:"pageNum" form:"pageNum"`
	PageSize int `json:"pageSize" form:"pageSize"`
}
type AgentVersionListReq struct {
	Name string `json:"Name"`
	// 分页参数
	PageNum  int `json:"pageNum" form:"pageNum"`
	PageSize int `json:"pageSize" form:"pageSize"`
}
type CreateAgentVersion struct {
	Name       string `json:"name"`
	OsType     string `json:"osType"`
	Version    string `json:"version"`
	UrlAddress string `json:"urlAddress"`
}
