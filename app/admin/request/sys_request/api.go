package sys_request

import "csf/common/mysql/model"

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
	List  []model.SysApi `json:"list" form:"list" description:"列表"`
}

type ApiEditReq struct {
	Id     int    `json:"id" form:"id" validate:"required" msg:"required:参数异常" description:"ID"`
	Tags   string `json:"tag" form:"tag"  validate:"required" msg:"required:分类必填" description:"标题"`
	Title  string `json:"title" form:"title" validate:"required" msg:"required:标题必填"  description:"标题"`
	Path   string `json:"path" form:"path" validate:"required" msg:"required:路径必填"  description:"标题"`
	Method string `json:"method" form:"method" validate:"required" msg:"required:方法类型必填"  description:"提交类型"`
}
