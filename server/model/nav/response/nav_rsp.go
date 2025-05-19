package response

import "go-devops-mimi/server/model/nav"

type NavListRsp struct {
	Total    int64     `json:"total"`
	NavLists []nav.Nav `json:"navLists"`
}

type LInkListRsp struct {
	Total     int64      `json:"total"`
	LInkLists []nav.Link `json:"linkLists"`
}
