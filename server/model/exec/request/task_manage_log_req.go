package request

type ManageLogListReq struct {
	Name     string `json:"name" form:"name"`
	PageNum  int    `json:"pageNum" form:"pageNum"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}
