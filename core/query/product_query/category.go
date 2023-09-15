package product_query

import (
	"time"
)

type CategoryAddOrEditInput struct {
	Id     int    `json:"id" form:"id" description:"ID"`
	Pid    int    `json:"pid" form:"pid" description:"父级ID"`
	Name   string `json:"name" form:"name" validate:"required" msg:"名称不为空" description:"名称"`
	Sort   int    `json:"sort" form:"sort" validate:"required" msg:"排序不为空" description:"排序"`
	Pic    string `json:"pic" form:"pic" description:"头像"`
	BigPic string `json:"big_pic" form:"big_pic" description:"大头像"`
	IsShow int    `json:"is_show" form:"is_show" description:"是否展示"`
}

type CategoryListInput struct {
	Page     int   `json:"page" form:"page"  default:"1" description:"页码"`
	PageSize int   `json:"page_size" form:"page_size"  default:"20" description:"页数"`
	Pids     []int `json:"pids" form:"pids" description:"名称"`
	IsShow   int   `json:"is_show" form:"is_show" description:"是否显示"`
}
type CategoryListOut struct {
	Total int64                  `json:"total" form:"total" description:"总数"`
	List  []CategoryTreeListItem `json:"list" form:"list" description:"列表"`
}
type CategoryTreeListItem struct {
	Category
	Children []CategoryTreeListItem `json:"children" gorm:"-"`
}
type Category struct {
	Id        int       `json:"id" form:"id" description:"ID"`
	Pid       int       `json:"pid" form:"pid" description:"父级ID"`
	Name      string    `json:"name" form:"name" description:"名称"`
	Label     string    `json:"label" form:"label" description:"名称"`
	Value     int       `json:"value" form:"value" description:"ID"`
	Sort      int       `json:"sort" form:"sort" description:"排序"`
	Pic       string    `json:"pic" form:"pic" description:"头像"`
	BigPic    string    `json:"big_pic" form:"big_pic" description:"大头像"`
	IsShow    int       `json:"is_show" form:"is_show" description:"是否展示"`
	Operator  string    `json:"operator" form:"operator" description:"操作人"`
	CreatedAt time.Time `json:"created_at" form:"created_at" description:"创建时间"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at" description:"最后更新时间"`
	DeletedAt time.Time `json:"deleted_at" form:"deleted_at" description:"删除时间"`
}

type CategoryDeleteBatchInput struct {
	Ids []int `json:"ids" form:"ids"  validate:"required" msg:"参数异常" description:"页码"`
}
