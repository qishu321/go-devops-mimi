package request

// Link 相关请求结构体
// ========================

type LinkAddRequest struct {
	Name  string `json:"name" binding:"required"`          // 链接名称，必填
	Desc  string `json:"desc" binding:"omitempty,max=255"` // 链接描述
	URL   string `json:"url" binding:"required,url"`       // 跳转地址，必填且必须是 URL
	NavID uint   `json:"navId" binding:"required"`         // 所属分类 ID，必填
}

// UpdateLinkRequest 更新一个已有的导航链接
type LinkUpdateRequest struct {
	ID    uint   `json:"id" binding:"required"`            // 要更新的链接 ID
	Name  string `json:"name" binding:"required"`          // 链接名称
	Desc  string `json:"desc" binding:"omitempty,max=255"` // 链接描述
	URL   string `json:"url" binding:"required,url"`       // 跳转地址
	NavID uint   `json:"navId" binding:"required"`         // 分类 ID
}

// DeleteLinkRequest 删除一个导航链接
type LinkDeleteRequest struct {
	ID uint `uri:"id" binding:"required,gt=0"`
}

// GetLinkRequest 获取单个导航链接
type LinkInfoRequest struct {
	ID uint `uri:"id" binding:"required,gt=0"`
}

// ListLinksRequest 列表查询，可按分类或关键词过滤
type LinkListRequest struct {
	Name string `form:"Name"` // 可选：按名称/描述模糊搜索
	// 分页参数
	PageNum  int `json:"pageNum" form:"pageNum"`
	PageSize int `json:"pageSize" form:"pageSize"`
}
