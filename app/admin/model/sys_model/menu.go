package sys_model

import (
	"csf/common/mysql/model"
	"time"
)

type MenuListItem struct {
	model.SysMenu
	Children []MenuListItem `json:"children" gorm:"-"`
}

type SysMenuListItem struct {
	Id         int               `json:"id"`         // ID
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
	SysApis    []model.SysApi    `json:"sys_apis" gorm:"-"`
}

type GetApisByMenuIdOut struct {
	ApisId   []int          `json:"apis_id" descriptions:"Api ID"`
	ApisList []model.SysApi `json:"apis_list" description:"apis_list"`
}
