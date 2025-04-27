package request

type TransferListReq struct {
	Name     string `json:"name" form:"name"`
	PageNum  int    `json:"pageNum" form:"pageNum"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}
type TransferAddReq struct {
	Name       string `json:"name" form:"name" validate:"required,min=1,max=64"`            // 分发任务名称
	SourcePath string `json:"source_path" form:"source_path" validate:"required"`           // 本地源文件路径
	TargetPath string `json:"target_path" form:"target_path" validate:"required"`           // 目标主机路径
	NodesIDs   []int  `json:"node_ids" form:"node_ids" validate:"required,min=1,dive,gt=0"` // 目标主机 ID 列表
}
type TransferInfoReq struct {
	ID   uint   `json:"id" form:"id" validate:"required"`
	Name string `json:"name" form:"name"  validate:"required,min=1,max=64"` // 模板名称
}
