package model

type BasePage struct {
	PageNum  int `form:"pageNum"`
	PageSize int `form:"pageSize"`
}
