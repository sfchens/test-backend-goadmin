package sys_req

import (
	"csf/core/query/sys_query"
)

// MenuListReq 列表数据
type MenuListReq struct {
	Page     int    `json:"page" form:"page"  default:"1" description:"页码"`
	PageSize int    `json:"page_size" form:"page_size"  default:"20" description:"页数"`
	MenuType string `json:"menu_type" form:"menu_type" default:"M"  description:"菜单类型"`
	Key      string `json:"key" form:"key"  description:"名称"`
	Visible  int    `json:"visible" form:"visible" default:"-1" description:"是否显示"`
}

type MenuListRes struct {
	Total int64                       `json:"total" form:"total" description:"总数"`
	List  []sys_query.SysMenuListItem `json:"list" form:"list" description:"列表"`
}

// MenuTreeListReq 菜单数据
type MenuTreeListReq struct {
	Page     int    `json:"page" form:"page"  default:"1" description:"页码"`
	PageSize int    `json:"page_size" form:"page_size"  default:"20" description:"页数"`
	MenuType string `json:"menu_type" form:"menu_type" default:"M"  description:"菜单类型"`
	Key      string `json:"key" form:"key"  description:"名称"`
	Visible  int    `json:"visible" form:"visible" default:"1" description:"是否显示"`
}

type MenuTreeListRes struct {
	Total int64                    `json:"total" form:"total" description:"总数"`
	List  []sys_query.MenuListItem `json:"list" form:"list" description:"列表"`
}

// MenuTreeRoleListReq 选择数据
type MenuTreeRoleListReq struct {
	Page     int    `json:"page" form:"page"  default:"1" description:"页码"`
	PageSize int    `json:"page_size" form:"page_size"  default:"20" description:"页数"`
	Key      string `json:"key" form:"key"  description:"名称"`
	IsShow   int    `json:"is_show" form:"is_show" default:"-1" description:"是否显示"`
}

type MenuTreeRoleListRes struct {
	Total int64                       `json:"total" form:"total" description:"总数"`
	List  []sys_query.SysMenuListItem `json:"list" form:"list" description:"列表"`
}

type MenuAddOrEditReq struct {
	Id         int    `json:"id" form:"id" description:"id"`
	ParentId   int    `json:"parent_id" form:"parent_id"  description:"父级分类"`
	Title      string `json:"title" form:"title" validate:"required" msg:"required:标题必填" description:"标题"`
	MenuName   string `json:"menu_name" form:"menu_name" description:"菜单唯一名"`
	MenuType   string `json:"menu_type" form:"menu_type" validate:"required" msg:"required:菜单类型必选"  description:"菜单类型"`
	Sort       int    `json:"sort" form:"sort" default:"1" description:"排序"`
	Icon       string `json:"icon" form:"icon"  description:"图标"`
	Permission string `json:"permission" form:"permission" description:"权限唯一标识"`
	IsFrame    int    `json:"is_frame" form:"is_frame" description:"是否框架"`
	IsShow     int    `json:"is_show" form:"is_show"  description:"是否显示"`
	Path       string `json:"path" form:"path" description:"路径"`
	Component  string `json:"component" form:"component" description:"路径"`
	Visible    int    `json:"visible" form:"visible" description:"路径"`
	ApisId     []int  `json:"apis_id" form:"apis_id" description:"api"`
}
