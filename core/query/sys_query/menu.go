package sys_query

import (
	"csf/core/mysql/model"
	"time"
)

type MenuTreeListInput struct {
	Page     int    `json:"page" form:"page"  default:"1" description:"页码"`
	PageSize int    `json:"page_size" form:"page_size"  default:"20" description:"页数"`
	MenuType string `json:"menu_type" form:"menu_type" default:"M"  description:"菜单类型"`
	Key      string `json:"key" form:"key"  description:"名称"`
	Visible  int    `json:"visible" form:"visible" default:"-1" description:"是否显示"`
}
type MenuTreeListOut struct {
	Total int64          `json:"total" form:"total" description:"总数"`
	List  []MenuListItem `json:"list" form:"list" description:"列表"`
}

type MenuListItem struct {
	model.SysMenu
	Children []MenuListItem `json:"children" gorm:"-"`
}
type SysMenuListItem struct {
	ID         int64             `json:"id"`         // ID
	MenuName   string            `json:"menu_name"`  // 菜单名称
	Label      string            `json:"label"`      // 菜单名称
	Title      string            `json:"title"`      // 标题
	Icon       string            `json:"icon"`       // 图标
	Path       string            `json:"path"`       // 前端路径
	ParentId   int               `json:"parent_id"`  // 父级
	ParentIds  string            `json:"parent_ids"` // 父级类型
	MenuType   string            `json:"menu_type"`  // 菜单类型，M目录 C菜单，F按钮
	Permission string            `json:"permission"` // 权限标识
	Component  string            `json:"component"`  // 组件
	ApisId     []int             `json:"apis_id"`    // 组件
	Sort       int               `json:"sort"`       // 排序
	Visible    int               `json:"visible"`    // 是否启用，1启用
	IsFrame    int               `json:"is_frame"`   // 是否框架，1
	Operator   string            `json:"operator"`   // 操作人
	CreatedAt  time.Time         `json:"created_at"` // 创建时间
	UpdatedAt  time.Time         `json:"updated_at"` // 最后更新时间
	DeletedAt  time.Time         `json:"deleted_at"` // 删除时间
	Children   []SysMenuListItem `json:"children" gorm:"-"`
	SysApis    []model.SysAPI    `json:"sys_apis" gorm:"-"`
}

type MenuListInput struct {
	Page     int    `json:"page" form:"page"  default:"1" description:"页码"`
	PageSize int    `json:"page_size" form:"page_size"  default:"20" description:"页数"`
	MenuType string `json:"menu_type" form:"menu_type" default:"M"  description:"菜单类型"`
	Key      string `json:"key" form:"key"  description:"名称"`
	Visible  int    `json:"visible" form:"visible" default:"-1" description:"是否显示"`
}
type MenuListOut struct {
	Total int64             `json:"total" form:"total" description:"总数"`
	List  []SysMenuListItem `json:"list" form:"list" description:"列表"`
}

type GetApisByMenuIdOut struct {
	ApisId   []int          `json:"apis_id" descriptions:"Api ID"`
	ApisList []model.SysAPI `json:"apis_list" description:"apis_list"`
}

type MenuAddOrEditInput struct {
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
