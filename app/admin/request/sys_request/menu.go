package sys_request

import "csf/app/admin/model/sys_model"

type MenuListReq struct {
	Page     int    `json:"page" form:"page"  default:"1" description:"页码"`
	PageSize int    `json:"page_size" form:"page_size"  default:"20" description:"页数"`
	Title    string `json:"name" form:"name"  description:"名称"`
	IsShow   int    `json:"is_show" form:"is_show" default:"-1" description:"是否显示"`
}

type MenuListRes struct {
	Total int64                    `json:"total" form:"total" description:"总数"`
	List  []sys_model.MenuListItem `json:"list" form:"list" description:"列表"`
}

type MenuAddOrEditReq struct {
	Id         int      `json:"id" form:"id" description:"id"`
	ParentIds  []string `json:"parent_ids" form:"parent_ids" description:"父级分类"`
	Title      string   `json:"title" form:"title" validate:"required" msg:"required:标题必填" description:"标题"`
	Sort       int      `json:"sort" form:"sort" default:"99" description:"排序"`
	Icon       string   `json:"icon" form:"icon"  description:"图标"`
	MenuType   uint8    `json:"menu_type" form:"menu_type" validate:"required" msg:"required:菜单类型必选"  description:"菜单类型"`
	UniqueName string   `json:"unique_name" form:"unique_name" description:"菜单唯一名称"`
	UniqueAuth string   `json:"unique_auth" form:"unique_auth" description:"权限唯一标识"`
	IsFrame    int      `json:"is_frame" form:"is_frame" msg:"required:框架类型必选" description:"是否框架"`
	IsShow     int      `json:"is_show" form:"is_show" default:"1" description:"是否显示"`
	Path       string   `json:"path" form:"path" validate:"required" msg:"required:路径必填" description:"路径"`
}
