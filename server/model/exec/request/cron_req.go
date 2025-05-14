package request

type CronListReq struct {
	Name     string `json:"name" form:"name"`
	CronType string `json:"cronType" form:"cronType"`
	PageNum  int    `json:"pageNum" form:"pageNum"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}

type CronAddReq struct {
	Name       string `json:"name" binding:"required,min=1,max=128"`                  // 任务名，长度 1–128 :contentReference[oaicite:1]{index=1}
	Desc       string `json:"desc" binding:"max=512"`                                 // 描述，可选，最长 512 :contentReference[oaicite:2]{index=2}
	CronType   string `json:"cronType" binding:"required,oneof=interval once cron"`   // 类型：interval/once/cron :contentReference[oaicite:3]{index=3}
	Cronession string `json:"cronession" binding:"required_if=CronType cron"`         // Cron 表达式（仅 Cron 类型必填） :contentReference[oaicite:4]{index=4}
	Interval   int    `json:"interval" binding:"required_if=CronType interval,gte=1"` // 间隔秒数（仅 Interval 类型必填，>=1） :contentReference[oaicite:5]{index=5}
	OnceTime   string `json:"onceTime" binding:"required_if=CronType once"`           // 一次性任务执行时间 :contentReference[oaicite:6]{index=6}
	CmdType    string `json:"cmd_type" form:"cmd_type" binding:"required"`            //是命令还是脚本
	Type       string `json:"type" form:"type" binding:"required"`                    //是shell还是python
	Content    string `form:"content" json:"content"  binding:"required"`             //内容
	Timeout    int    `json:"timeout" form:"timeout"`                                 //超时时间
	NodesIDs   []int  `json:"node_ids" form:"node_ids" binding:"required"`            // 执行主机ID列表，可以存储多个主机的ID（建议以JSON格式存储或其它约定格式）
}
type CronUpdateReq struct {
	ID         uint   `json:"id" validate:"required"`
	Name       string `json:"name" binding:"required,min=1,max=128"`                  // 任务名，长度 1–128 :contentReference[oaicite:1]{index=1}
	Desc       string `json:"desc" binding:"max=512"`                                 // 描述，可选，最长 512 :contentReference[oaicite:2]{index=2}
	CronType   string `json:"cronType" binding:"required,oneof=interval once cron"`   // 类型：interval/once/cron :contentReference[oaicite:3]{index=3}
	Cronession string `json:"cronession" binding:"required_if=CronType cron"`         // Cron 表达式（仅 Cron 类型必填） :contentReference[oaicite:4]{index=4}
	Interval   int    `json:"interval" binding:"required_if=CronType interval,gte=1"` // 间隔秒数（仅 Interval 类型必填，>=1） :contentReference[oaicite:5]{index=5}
	OnceTime   string `json:"onceTime" binding:"required_if=CronType once"`           // 一次性任务执行时间 :contentReference[oaicite:6]{index=6}
	CmdType    string `json:"cmd_type" form:"cmd_type" binding:"required"`            //是命令还是脚本
	Type       string `json:"type" form:"type" binding:"required"`                    //是shell还是python
	Content    string `form:"content" json:"content"  binding:"required"`             //内容
	Timeout    int    `json:"timeout" form:"timeout"`                                 //超时时间
	NodesIDs   []int  `json:"node_ids" form:"node_ids" binding:"required"`            // 执行主机ID列表，可以存储多个主机的ID（建议以JSON格式存储或其它约定格式）

}
type CronEnableReq struct {
	ID     uint `json:"id" form:"id" validate:"required"`
	Enable int8 `json:"enable" validate:"oneof=0 1"`
}
type CronInfoReq struct {
	ID uint `json:"id" form:"id" validate:"required"`
}

type CronDeleteReq struct {
	Ids []uint `json:"ids" validate:"required"`
}
