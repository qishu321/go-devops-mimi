package cmdb

import (
	"gorm.io/gorm"
)

type NodeGroup struct {
	gorm.Model
	GroupName  string   `json:"groupName" gorm:"type:varchar(50);uniqueIndex;column:group_name"`
	NodeGroups []*Nodes `gorm:"many2many:t_node_group_s" json:"t_node_group_s"`
	Desc       string   `gorm:"type:varchar(50);column:desc;comment:'描述'" json:"desc"`
	Creator    string   `gorm:"type:varchar(20);comment:'创建人'" json:"creator"`
	// 新增：每组的节点数
	NodeCount int `gorm:"-" json:"nodeCount"`
}

func (m *NodeGroup) TableName() string {
	return "t_node_group"
}

type Nodes struct {
	gorm.Model
	NodeName   string       `gorm:"type:varchar(50);column:node_name;uniqueIndex;comment:'CMDB名称'" json:"nodeName"` // 添加 uniqueIndex
	Username   string       `gorm:"column:username;comment:'服务器用户名称'" json:"username"`
	PublicIP   string       `gorm:"column:public_ip;comment:'IP地址'" json:"publicIP"`
	SSHPort    int          `gorm:"column:ssh_port;comment:'SSH端口号'" json:"sshPort"`
	AuthModel  string       `gorm:"column:authmodel;comment:'连接服务器所使用的是密钥还是密码'" json:"authmodel" `
	Password   string       `gorm:"column:password;comment:'password'" json:"password"`
	PrivateKey string       `gorm:"column:private_key;comment:'私钥'" json:"private_key"`
	Status     int8         `gorm:"column:status;comment:'1:SSH连接成功,2：SSH连接失败'" json:"status"`
	Groups     []*NodeGroup `gorm:"many2many:t_node_group_s" json:"t_node_group_s"`
	Timeout    int          `gorm:"column:timeout;comment:'超时时间'" form:"timeout" json:"timeout"` //超时时间
	Label      string       `gorm:"column:label;comment:'标签'" json:"label"`
	Creator    string       `gorm:"type:varchar(20);comment:'创建人'" json:"creator"`
}

func (m *Nodes) TableName() string {
	return "t_nodes"
}
