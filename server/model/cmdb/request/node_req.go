package request

type AddNodesToGroupReq struct {
	GroupID uint   `json:"groupId" binding:"required"`       // 必须指定分组ID
	NodeIDs []uint `json:"nodeIds" binding:"required,min=1"` // 至少一个节点
}
type RemoveNodeGroupReq struct {
	GroupID uint   `json:"groupId" binding:"required"`       // 必须指定分组ID
	NodeIDs []uint `json:"nodeIds" binding:"required,min=1"` // 至少一个节点
}
type AddNodesGroupReq struct {
	GroupIDs []int64 `json:"groupIds"`
	NodeIDs  []int64 `json:"nodeIds" binding:"required,min=1"`
}

// NodeGroupCreateReq 创建主机组请求结构体
type NodeGroupCreateReq struct {
	GroupName string `json:"groupName" binding:"required" form:"groupName"` // 主机组名称
	// NodesIDs 用于指定加入此主机组的服务器 ID 数组（前端选择多个服务器后提交的 ID）
	NodesIDs []uint `json:"NodesIds" form:"NodesIds"`
	Desc     string `json:"desc" form:"desc"` // 描述
}

// NodeGroupUpdateReq 更新主机组请求结构体
type NodeGroupUpdateReq struct {
	ID        uint   `json:"id" binding:"required" form:"id"` // 主机组 ID，用于更新指定记录
	GroupName string `json:"groupName" form:"groupName"`      // 主机组名称
	// NodesIDs []uint `json:"NodesIds" form:"NodesIds"`  // 更新时，指定加入此主机组的服务器 ID 数组
	Desc string `json:"desc" form:"desc"` // 描述
}

// NodeGroupDeleteReq 删除主机组请求结构体
type NodeGroupDeleteReq struct {
	Ids []uint `json:"ids" validate:"required"`
}

// NodeGroupListReq 查询主机组列表请求结构体
type NodeGroupListReq struct {
	GroupName string `json:"groupName" form:"groupName"` // 模糊查询组名（可选）
	PageNum   int    `json:"pageNum" form:"pageNum"`     // 当前页码
	PageSize  int    `json:"pageSize" form:"pageSize"`   // 每页数量
}

// NodesCreateReq 创建节点（服务器）请求结构体
type NodesCreateReq struct {
	// 节点名称（必填）
	NodeName string  `json:"nodeName" binding:"required" form:"nodeName"`
	GroupID  []int64 `json:"groupId" ` // 必须指定分组ID
	//服务器用户名称
	UserName string `json:"username" binding:"required" form:"username"`
	// 公网 IP 地址（必填）
	PublicIP string `json:"publicIP" binding:"required,ip" form:"publicIP"`
	// SSH 端口号，默认一般为 22
	SSHPort int `json:"sshPort" binding:"required" form:"sshPort"`
	// 认证方式：例如 "password" 或 "key"，必填
	AuthModel  string `json:"authmodel" binding:"required" form:"authmodel"`
	Password   string `json:"password" form:"password" comment:"密码"`
	PrivateKey string `json:"privateKey" form:"privateKey" comment:"私钥"`
	// 超时时间，单位秒
	Timeout int `json:"timeout" form:"timeout"`
	// 标签，可用于描述节点用途等
	Label     string `json:"label" form:"label"`
	GroupName string `json:"groupName" form:"groupName"` // 模糊查询组名（可选）

}

// NodesUpdateReq 更新节点请求结构体
type NodesUpdateReq struct {
	// 节点记录的 ID（必填），用于更新指定记录
	ID      int     `json:"id" binding:"required" form:"id"`
	GroupID []int64 `json:"groupId" `
	// 以下字段均为可选更新字段
	//服务器用户名称
	UserName   string `json:"username" binding:"required" form:"username"`
	NodeName   string `json:"nodeName" form:"nodeName"`
	PublicIP   string `json:"publicIP" form:"publicIP"`
	SSHPort    int    `json:"sshPort" form:"sshPort"`
	AuthModel  string `json:"authmodel" form:"authmodel"`
	Password   string `json:"password" form:"password" comment:"密码"`
	PrivateKey string `json:"privateKey" form:"privateKey" comment:"私钥"`
	Timeout    int    `json:"timeout" form:"timeout"`
	Label      string `json:"label" form:"label"`
}

// NodesDeleteReq 删除节点请求结构体
type NodesDeleteReq struct {
	// 根据节点记录 ID 删除
	Ids []uint `json:"ids" validate:"required"`
}

// NodesListReq 查询节点列表请求结构体
type NodesListReq struct {
	// 可按节点名称模糊查询
	NodeName string `json:"nodeName" form:"nodeName"`
	// 可按 IP 地址查询
	PublicIP string `json:"publicIP" form:"publicIP"`
	// 分页参数
	PageNum  int `json:"pageNum" form:"pageNum"`
	PageSize int `json:"pageSize" form:"pageSize"`
}
