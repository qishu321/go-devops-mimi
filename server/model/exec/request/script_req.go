package request

type ScriptListReq struct {
	Name     string `json:"name" form:"name"`
	CmdType  string `json:"cmd_type" form:"cmd_type"` //是命令还是脚本
	PageNum  int    `json:"pageNum" form:"pageNum"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}

type ScriptCmdRunReq struct {
	Name    string `json:"name" form:"name" binding:"required"`
	CmdType string `json:"cmd_type" form:"cmd_type" binding:"required"` //是命令还是脚本
	Type    string `json:"type" form:"type" binding:"required"`         //是shell还是python
	Command string `form:"command" json:"command"  binding:"required"`  //命令或者脚本内容
	Timeout int    `json:"timeout" form:"timeout"`                      //超时时间
	Desc    string ` json:"desc" form:"desc"`                           // 描述信息

	// 执行主机ID列表，可以存储多个主机的ID（建议以JSON格式存储或其它约定格式）
	NodesIDs []int `json:"node_ids" form:"node_ids" binding:"required"` // 执行主机ID列表，可以存储多个主机的ID（建议以JSON格式存储或其它约定格式）
}

// type ScriptLogAddReq struct {
// 	Name    string `json:"name" validate:"required,min=1,max=64"` // 模板名称
// 	Content string `json:"content" validate:"required"`           // 脚本内容
// 	Type    string ` json:"type" validate:"required"`             // 脚本类型
// 	Args    string ` json:"args"`                                 // 脚本参数信息
// 	Desc    string ` json:"desc"`                                 // 描述信息
// 	Timeout int    ` json:"timeout"`                              // 脚本超时时间，单位秒
// }
// type ScriptLogUpdateReq struct {
// 	ID      uint   `json:"id" validate:"required"`
// 	Name    string `json:"name" validate:"required,min=1,max=64"` // 模板名称
// 	Content string `json:"content" validate:"required"`           // 脚本内容
// 	Type    string ` json:"type" validate:"required"`             // 脚本类型
// 	Args    string ` json:"args"`                                 // 脚本参数信息
// 	Desc    string ` json:"desc"`                                 // 描述信息
// 	Timeout int    ` json:"timeout"`                              // 脚本超时时间，单位秒
// }

// type ScriptLogDeleteReq struct {
// 	Ids []uint `json:"ids" validate:"required"`
// }
