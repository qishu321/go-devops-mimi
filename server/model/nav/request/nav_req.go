package request

// —— Nav（分类）相关请求结构体 ——

// NavAddRequest 新增一个导航分类
type NavAddRequest struct {
	Name    string `json:"name" binding:"required"`            // 分类名称，必填
	NavSort int    `json:"nav_sort" binding:"omitempty,min=0"` // 可选：分类排序
}

// UpdateNavRequest 更新一个已有的导航分类
type NavUpdateRequest struct {
	ID      uint   `json:"id" binding:"required"`              // 要更新的分类 ID
	Name    string `json:"name" binding:"required"`            // 分类名称
	NavSort int    `json:"nav_sort" binding:"omitempty,min=0"` // 排序
}
type NavInfoRequest struct {
	ID uint `json:"id" binding:"required"`
}

// DeleteNavRequest 删除一个导航分类
type NavDeleteRequest struct {
	ID uint `json:"id" binding:"required"` // 要删除的分类 ID
}

// GetNavListRequest 获取分类列表（分页/搜索）
type NavListRequest struct {
	PageNum  int `json:"pageNum" form:"pageNum"`
	PageSize int `json:"pageSize" form:"pageSize"`
}
