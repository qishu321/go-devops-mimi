package request

type ScriptLibraryListReq struct {
	Name     string `json:"name" form:"name"`
	PageNum  int    `json:"pageNum" form:"pageNum"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}
type ScriptLibraryAddReq struct {
	Name    string `json:"name" validate:"required,min=1,max=64"` // 模板名称
	Content string `json:"content" validate:"required"`           // 脚本内容
	Type    string ` json:"type" validate:"required"`             // 脚本类型
	Desc    string ` json:"desc"`                                 // 描述信息
}
type ScriptLibraryUpdateReq struct {
	ID      uint   `json:"id" validate:"required"`
	Name    string `json:"name" validate:"required,min=1,max=64"` // 模板名称
	Content string `json:"content" validate:"required"`           // 脚本内容
	Type    string ` json:"type" validate:"required"`             // 脚本类型
	Desc    string ` json:"desc"`                                 // 描述信息
}
type ScriptLibraryInfoReq struct {
	ID   uint   `json:"id" form:"id" validate:"required"`
	Name string `json:"name" form:"name"  validate:"required,min=1,max=64"` // 模板名称
}

type ScriptLibraryDeleteReq struct {
	Ids []uint `json:"ids" validate:"required"`
}
