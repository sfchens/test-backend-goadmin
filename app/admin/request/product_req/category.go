package product_req

import "csf/core/query/sys_query"

type CategoryAddOrEditReq struct {
	Id     int    `json:"id" form:"id" description:"ID"`
	Pid    int    `json:"pid" form:"pid" description:"父级ID"`
	Name   string `json:"name" form:"name" validate:"required" msg:"名称不为空" description:"名称"`
	Sort   int    `json:"sort" form:"sort" validate:"required" msg:"排序不为空" description:"排序"`
	Pic    string `json:"pic" form:"pic" description:"头像"`
	BigPic string `json:"big_pic" form:"big_pic" description:"大头像"`
	IsShow int    `json:"is_show" form:"is_show" description:"是否展示"`
}

type CategoryListReq struct {
	Page     int `json:"page" form:"page"  default:"1" description:"页码"`
	PageSize int `json:"page_size" form:"page_size"  default:"20" description:"页数"`
}
type CategoryListRes struct {
	Total int64                        `json:"total" form:"total" description:"总数"`
	List  []sys_query.DeptTreeListItem `json:"list" form:"list" description:"列表"`
}

type CategoryDeleteBatchReq struct {
	Ids []int `json:"ids" form:"ids"  validate:"required" msg:"参数异常" description:"页码"`
}
