package request

import "encoding/json"

type TaskManageListReq struct {
	Name     string `json:"name" form:"name"`
	PageNum  int    `json:"pageNum" form:"pageNum"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}
type TaskManageAddReq struct {
	Name  string          `json:"name" validate:"required,min=1,max=64"`
	Args  json.RawMessage `json:"args"`  // 接收任意结构，作为 JSON 保存
	Desc  string          ` json:"desc"` // 描述信息
	Tasks []*TaskReq      `json:"t_task_s"`
}

// TaskReq 子任务请求结构体，用于 Add/Update 时接收前端数据
type TaskReq struct {
	// 如果是更新已有子任务，前端可以传 ID；新增时可省略或置 0
	ID uint `json:"id,omitempty"`
	// 子任务名称或类型，必填，长度 1–128
	Name string `json:"name" validate:"required,min=1,max=128"`
	// 脚本类型，如 shell、python 等，必填，长度 1–32
	Type string `json:"type" validate:"required,min=1,max=32"`
	// 执行命令或脚本内容，必填
	Content string `json:"content" validate:"required"`
	// 执行顺序 (1–999)，默认可设为 999
	Sort int `json:"sort" validate:"gte=1,lte=999"`
	// 脚本超时时间（秒），非负
	Timeout int `json:"timeout" validate:"gte=0"`
	// 执行主机 ID 列表，建议前端传 JSON 数组字符串或逗号分隔
	NodeIDs string `json:"node_ids" validate:"required"`
}

type TaskManageUpdateReq struct {
	ID    uint       `json:"id" validate:"required"`
	Name  string     `json:"name" validate:"required,min=1,max=64"` // 模板名称
	Tasks []*TaskReq `json:"t_task_s"`
	Args  string     `json:"args" ` // 执行时传入的参数或全局变量，建议以JSON格式存储
	Desc  string     ` json:"desc"` // 描述信息
}
type TaskManageInfoReq struct {
	ID   uint   `json:"id" form:"id" validate:"required"`
	Name string `json:"name" form:"name"  validate:"required,min=1,max=64"` // 模板名称
}

type TaskManageRunReq struct {
	ID        uint              `json:"id" form:"id" validate:"required"`
	Name      string            `json:"name" form:"name"  validate:"required,min=1,max=64"` // 模板名称
	EnvParams map[string]string `form:"env_task_s" json:"env_task_s"`                       // <-- 这里，QueryMap 会填充
	Desc      string            ` json:"desc"`                                              // 描述信息
}
type TaskManageRunInfoReq struct {
	RunID uint `json:"run_id" form:"run_id" validate:"required"`
}
type TaskManageRunInfoWebsocketReq struct {
	TaskID uint `json:"task_id" form:"task_id" validate:"required"`
}

type TaskManageDeleteReq struct {
	Ids []uint `json:"ids" validate:"required"`
}
