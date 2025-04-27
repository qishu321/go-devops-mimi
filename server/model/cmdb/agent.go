package cmdb

import (
	"time"

	"gorm.io/gorm"
)

type AgentGroup struct {
	gorm.Model
	GroupName string   `json:"groupName" gorm:"uniqueIndex;column:group_name;type:varchar(50)"`
	Agents    []*Agent `gorm:"many2many:group_agent_s" json:"group_agent_s"`
	Desc      string   `gorm:"type:varchar(50);column:desc;comment:'描述'" json:"desc"`
	Creator   string   `gorm:"type:varchar(20);comment:'创建人'" json:"creator"`
}

func (m *AgentGroup) TableName() string {
	return "t_agent_group"
}

type Agent struct {
	gorm.Model
	Agentid       uint64    `json:"agent_id" gorm:"column:agent_id;type:varchar(150);not null;unique;comment:'Agent唯一标识'"`
	AgentName     string    `json:"agentName" gorm:"uniqueIndex;column:agent_name;type:varchar(150);not null;comment:'Agent名称'"`
	IP            string    `gorm:"column:ip;comment:'IP地址'" json:"ip"`
	Online        bool      `json:"online" gorm:"column:online;default:false;comment:'在线状态'"`
	Version       string    `json:"version" gorm:"type:varchar(20);comment:'Agent版本信息'"`
	OsType        string    `json:"osType" gorm:"type:varchar(20);comment:'操作系统类型'"`
	Label         string    `json:"label" gorm:"column:label;type:varchar(50);comment:'标签'"`
	LastHeartbeat time.Time `json:"lastHeartbeat" gorm:"column:last_heartbeat;comment:'最后心跳时间'"`
}

func (m *Agent) TableName() string {
	return "t_agent"
}

// AgentVersion 表示Agent的版本信息，用于下载或更新Agent
type AgentVersion struct {
	Name       string `json:"name" gorm:"uniqueIndex;comment:'Agent版本名称';type:varchar(150)"`
	OsType     string `json:"osType" gorm:"type:varchar(20);not null;comment:'操作系统类型'"`
	Version    string `json:"version" gorm:"type:varchar(20);not null;comment:'版本信息'"`
	UrlAddress string `json:"urlAddress" gorm:"type:varchar(255);not null;comment:'下载地址'"`
}

func (m *AgentVersion) TableName() string {
	return "t_agent_version"
}
