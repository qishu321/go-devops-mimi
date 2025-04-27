package exec

import (
	"time"

	"gorm.io/gorm"
)

// 每个全局变量由 Key 与对应的 Value 构成，同时可记录描述信息和创建人。
// type Args struct {
// 	ID      uint   `gorm:"column:id;primaryKey;comment:'唯一标识'" json:"id"` // 唯一标识符
// 	Name    string `gorm:"column:name;size:128;not null;comment:'变量名称'" json:"name"`
// 	Value   string `gorm:"column:value;size:512;comment:'变量对应的值'" json:"value"`    // 变量对应的值
// 	Type    string `gorm:"column:type;size:32;comment:'变量类型'" json:"type"`         // 变量类型
// 	Options string `json:"options" form:"options" gorm:"type:text;comment:'变量内容'"` // 变量内容
// }

// func (m *Args) TableName() string {
// 	return "t_args"
// }

// ScriptLibrary 脚本库中的脚本模板结构体，用于存储可复用的脚本模板
type ScriptLibrary struct {
	gorm.Model
	Name    string `gorm:"column:name;size:128;unique;comment:'模板名称，用于标识和展示模板'" json:"name"` // 模板名称
	Content string `gorm:"column:content;type:text;comment:'脚本内容'" json:"content"`           // 脚本内容
	// Args    string `gorm:"column:args;type:text;comment:'脚本执行时参数信息，建议为JSON格式'" json:"args"`    // 脚本参数信息
	Desc string `gorm:"column:desc;size:512;comment:'模板描述信息'" json:"desc"`                  // 描述信息
	Type string `gorm:"column:type;size:32;comment:'脚本类型，例如 shell、powershell'" json:"type"` // 脚本类型
	// Timeout int    `gorm:"column:timeout;comment:'脚本超时时间（秒）'" json:"timeout"`                  // 脚本超时时间，单位秒
	Creator string `gorm:"column:creator;type:varchar(20);comment:'创建人'" json:"creator"` // 创建人
}

// TableName 指定 ScriptLibrary 模型对应的数据库表名
func (m *ScriptLibrary) TableName() string {
	return "t_script_library"
}

// Script 表示一次脚本执行请求记录，包含脚本具体的执行信息
type Script struct {
	gorm.Model
	Name string `gorm:"column:name;size:128;comment:'任务或脚本执行的名称'" json:"name"`
	// 执行主机ID列表，可以存储多个主机的ID（建议以JSON格式存储或其它约定格式）
	NodesIDs string `gorm:"column:node_ids;comment:'执行主机ID列表'" json:"node_ids"`
	// 任务执行名称
	Status    int8   `gorm:"column:status;comment:'1:成功,2：失败'" json:"status"`
	CmdType   string `gorm:"column:cmd_type;size:32;comment:'类型，如脚本执行日志或命令执行日志'" json:"cmd_type"` // 日志类型
	StartTime string `gorm:"type:varchar(2048);comment:'发起时间'" json:"startTime"`
	EndTime   string `gorm:"type:varchar(2048);comment:'执行结束时间'" json:"endTime"`
	TimeCost  int64  `gorm:"column:time_cost;type:int(6);comment:'执行耗时（毫秒）'" json:"timeCost"` // 执行耗时，单位毫秒
	Desc      string `gorm:"column:desc;size:512;comment:'额外描述或备注'" json:"desc"`              // 描述信息（执行原因、备注）
	Creator   string `gorm:"column:creator;type:varchar(20);comment:'执行人'" json:"creator"`    // 执行人
	// 多对多关联脚本记录（一个日志可能关联多个脚本，适用于组合任务等场景）
	Scripts []*ScriptLog `gorm:"many2many:t_script_s;comment:'关联的脚本记录'" json:"t_script_log_s"`
}

// TableName 指定 Script 模型对应的数据库表名
func (m *Script) TableName() string {
	return "t_script"
}

// ScriptLog 脚本执行历史日志记录结构体，用于记录每一次脚本执行的详细信息
type ScriptLog struct {
	gorm.Model
	Name      string `gorm:"column:name;size:128;comment:'脚本名称，用于标识和展示'" json:"name"`          // 脚本名称
	NodeName  string `gorm:"column:node_name;size:128;comment:'节点名称，标识执行节点'" json:"node_name"` // 节点名称
	Type      string `gorm:"column:type;size:32;comment:'脚本类型，如 shell、python 等'" json:"type"`  // 脚本类型
	Content   string `gorm:"column:content;type:text;comment:'脚本内容'" json:"content"`           // 脚本内容
	Status    int8   `gorm:"column:status;comment:'1:成功,2：失败'" json:"status"`
	Timeout   int    `gorm:"column:timeout;comment:'脚本超时时间（秒）'" json:"timeout"`
	RunLog    string `gorm:"column:run_log;type:longtext;comment:'执行结果'" json:"run_log"`
	StartTime string `gorm:"type:varchar(2048);comment:'发起时间'" json:"startTime"`
	EndTime   string `gorm:"type:varchar(2048);comment:'执行结束时间'" json:"endTime"`
	TimeCost  int64  `gorm:"column:time_cost;type:int(6);comment:'执行耗时（毫秒）'" json:"timeCost"` // 执行耗时，单位毫秒
}

// TableName 指定 ScriptLog 模型对应的数据库表名
func (m *ScriptLog) TableName() string {
	return "t_script_log"
}

// Transfer 文件分发记录
type Transfer struct {
	gorm.Model
	Name       string `gorm:"column:name;size:128;comment:'分发名称，用于标识和展示'" json:"name"`
	SourcePath string `gorm:"size:256;not null" json:"source_path"`
	// 目标路径（目标主机上的路径），可以是单个或JSON数组
	TargetPath string `gorm:"size:256;not null" json:"target_path"`
	// 目标主机或主机组
	NodesIDs string `gorm:"column:node_ids;comment:'目标主机'" json:"node_ids"`
	// 分发状态，如 pending、running、success、failed
	Status  int8   `gorm:"column:status;comment:'1:成功,2：失败'" json:"status"`
	RunLog  string `gorm:"column:run_log;type:text;comment:'执行结果'" json:"run_log"`
	Creator string `gorm:"column:creator;type:varchar(20);comment:'创建人'" json:"creator"` // 创建人

}

func (m *Transfer) TableName() string {
	return "t_transfer"
}

// Task 子任务，属于一个任务组
type Task struct {
	gorm.Model
	// 子任务名称或类型
	Name string `gorm:"column:name;size:128;not null;comment:'子任务名称或类型'" json:"name"`
	// 脚本类型，如 shell、python 等
	Type string `gorm:"column:type;size:32;comment:'脚本类型，如 shell、python 等'" json:"type"`
	// 执行命令或脚本内容
	Content string `gorm:"column:content;type:text;comment:'执行命令或脚本内容'" json:"content"`
	// 执行顺序(1-999)，用于确定子任务的执行先后
	Sort int `gorm:"column:sort;type:int;default:999;comment:'执行顺序(1-999)'" json:"sort"`
	// 脚本超时时间，单位为秒
	Timeout int `gorm:"column:timeout;comment:'脚本超时时间（秒）'" json:"timeout"`
	// 执行主机ID列表，可以存储多个主机的ID（建议以JSON格式存储或其它约定格式）
	NodesIDs string `gorm:"column:node_ids;comment:'执行主机ID列表'" json:"node_ids"`
	// 创建人，记录该子任务的创建者
	Creator string `gorm:"column:creator;type:varchar(20);comment:'创建人'" json:"creator"`
}

func (m *Task) TableName() string {
	return "t_task"
}

// TaskManage 任务管理（任务组或大任务）
type TaskManage struct {
	gorm.Model
	// 任务组名称
	Name string `gorm:"column:name;size:128;not null;comment:'任务组名称'" json:"name"`
	// 执行时传入的参数或全局变量，建议以JSON格式存储
	Args string `gorm:"column:args;type:text;comment:'执行时传入的参数或全局变量，建议为JSON格式'" json:"args"`
	// 额外描述或备注，记录任务组的描述信息
	Desc string `gorm:"column:desc;size:512;comment:'任务组描述信息或备注'" json:"desc"`
	// 关联的子任务，通过一对多关系关联
	Tasks []*Task `gorm:"many2many:t_task_s;comment:'关联的子任务'" json:"t_task_s"`
	// 任务组创建人
	Creator string `gorm:"column:creator;type:varchar(20);comment:'创建人'" json:"creator"`
}

func (m *TaskManage) TableName() string {
	return "t_task_manage"
}

// TaskManageLog 任务组执行记录
type TaskManageLog struct {
	gorm.Model
	// 任务组执行名称
	Name string `gorm:"column:name;size:128;not null;comment:'任务组执行名称'" json:"name"`
	// 执行时传入的参数或全局变量，建议以JSON格式存储
	Args string `gorm:"column:args;type:text;comment:'执行时传入的参数或全局变量，建议为JSON格式'" json:"args"`
	// 任务组描述或备注信息
	Description string `gorm:"column:description;size:512;comment:'任务组描述或备注信息'" json:"description"`
	// 关联的子任务日志，通过多对多关系关联到 TaskLog
	Tasklogs []*TaskLog `gorm:"many2many:t_task_log_s;comment:'关联的子任务日志'" json:"t_task_log_s"`
	// 任务状态，如 pending、running、success、failed
	Status string `gorm:"column:status;size:32;default:'draft';comment:'任务状态（pending、running、success、failed）'" json:"status"`
	// 发起执行时间
	StartTime string `gorm:"type:varchar(2048);comment:'发起时间'" json:"startTime"`
	EndTime   string `gorm:"type:varchar(2048);comment:'执行结束时间'" json:"endTime"`
	TimeCost  int64  `gorm:"column:time_cost;type:int(6);comment:'执行耗时（毫秒）'" json:"timeCost"` // 执行耗时，单位毫秒
}

func (m *TaskManageLog) TableName() string {
	return "t_task_manage_log"
}

// TaskLog 任务执行历史记录
type TaskLog struct {
	gorm.Model
	// 子任务名称
	Name string `gorm:"column:name;size:128;not null;comment:'子任务名称'" json:"name"`
	// 执行命令或脚本内容
	Content string `gorm:"column:content;type:text;comment:'执行命令或脚本内容'" json:"content"`
	// 执行顺序(1-999)，用于排序
	Sort int `gorm:"column:sort;type:int;default:999;comment:'执行顺序(1-999)'" json:"sort"`
	// 节点名称，用于标识执行节点
	NodeName string `gorm:"column:node_name;size:128;comment:'节点名称，标识执行节点'" json:"node_name"`
	// 执行结果日志，记录详细的执行输出或错误信息
	RunLog string `gorm:"column:run_log;type:text;comment:'执行结果日志'" json:"run_log"`
	// 子任务状态，如 pending、running、success、failed
	Status string `gorm:"column:status;size:32;default:'pending';comment:'子任务执行状态（pending、running、success、failed）'" json:"status"`
	// 发起执行时间
	StartTime string `gorm:"type:varchar(2048);comment:'发起时间'" json:"startTime"`
	EndTime   string `gorm:"type:varchar(2048);comment:'执行结束时间'" json:"endTime"`
	TimeCost  int64  `gorm:"column:time_cost;type:int(6);comment:'执行耗时（毫秒）'" json:"timeCost"` // 执行耗时，单位毫秒
}

func (m *TaskLog) TableName() string {
	return "t_task_log"
}

//	全局变量结构体，用于存储全局范围内的变量信息。
//

// Cron 定时任务
type Cron struct {
	gorm.Model
	// 定时表达式，如标准 Cron 表达式
	Expression string `gorm:"size:128;not null" json:"expression"`
	// 定时任务名称或描述
	Name        string `gorm:"size:128;not null" json:"name"`
	Description string `gorm:"size:512" json:"description"`
	// 状态：active、paused、disabled等
	Status string `gorm:"size:32;default:'active'" json:"status"`
	// 下一次执行时间（可选）
	NextRunAt *time.Time `json:"next_run_at"`
	SendTime  *time.Time `gorm:"column:send_at;type:timestamp;default:NULL;comment:'执行发起时间'" json:"sendTime"` // 发起时间
	EndTime   *time.Time `gorm:"column:end_at;type:timestamp;default:NULL;comment:'执行结束时间'" json:"endTime"`   // 结束时间
	TimeCost  int64      `gorm:"column:time_cost;type:int(6);comment:'执行耗时（毫秒）'" json:"timeCost"`             // 执行耗时，单位毫秒
}
