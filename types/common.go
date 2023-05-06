package types

type BasePage struct {
	PageNum  int `form:"pageNum"`
	PageSize int `form:"pageSize"`
}

// DataListResp 带有总数的Data结构
type DataListResp struct {
	Item  interface{} `json:"item"`
	Total int64       `json:"total"`
}
