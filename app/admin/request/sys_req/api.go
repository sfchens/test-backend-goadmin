package sys_req

import (
	"csf/core/mysql/model"
)

type ApiListReq struct {
	Page     int    `json:"page" form:"page"  default:"1" description:"页码"`
	PageSize int    `json:"page_size" form:"page_size"  default:"20" description:"页数"`
	Tag      string `json:"tag" form:"tag"  description:"分类"`
	Path     string `json:"path" form:"path" description:"路由"`
	Title    string `json:"title" form:"title" description:"标题"`
	Method   string `json:"method" form:"method" description:"提交类型"`
}

type ApiListRes struct {
	Total int64          `json:"total" form:"total" description:"总数"`
	List  []model.SysAPI `json:"list" form:"list" description:"列表"`
}

type ApiEditReq struct {
	Id     int    `json:"id" form:"id"  description:"ID"`
	Tags   string `json:"tags" form:"tags"  validate:"required" msg:"required:分类必填" description:"标题"`
	Title  string `json:"title" form:"title" validate:"required" msg:"required:标题必填"  description:"标题"`
	Path   string `json:"paths" form:"paths" validate:"required" msg:"required:路径必填"  description:"标题"`
	Handle string `json:"handle" form:"handle" validate:"required" msg:"required:路径必填"  description:"标题"`
	Method string `json:"method" form:"method" validate:"required" msg:"required:方法类型必填"  description:"提交类型"`
}

type ApiGetTagReq struct {
	Page     int    `json:"page" form:"page"  default:"1" description:"页码"`
	PageSize int    `json:"page_size" form:"page_size"  default:"20" description:"页数"`
	Tag      string `json:"tag" form:"tag"  description:"分类"`
}

type ApiGetTagRes struct {
	Total int64    `json:"total" form:"total" description:"总数"`
	List  []string `json:"list" form:"list" description:"列表"`
}

type ApiDeleteMultiReq struct {
	Ids []int `json:"ids" form:"ids" description:"页码"`
}
