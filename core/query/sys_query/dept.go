package sys_query

import (
	"csf/core/mysql/model"
	"time"
)

type DeptAddOrEditInput struct {
	Id       int    `json:"id" form:"id"  description:"ID"`
	ParentId int    `json:"parent_id" form:"parent_id"  description:"上级部门"`
	Name     string `json:"name" form:"name" validate:"required" msg:"required:部门名称必填" description:"部门名称"`
	Leader   string `json:"leader" form:"leader" validate:"required" msg:"required:负责人必填" description:"负责人"`
	Sort     int    `json:"sort" form:"sort" default:"1" description:"排序"`
	Phone    string `json:"phone" form:"phone" description:"手机号码"`
	Email    string `json:"email" form:"email" validate:"email" msg:"email:邮箱格式有误" description:"邮箱"`
	Status   int8   `json:"status" form:"status" validate:"required" msg:"required:状态必选" description:"状态"`
}

type DeptDeleteInput struct {
	Id int `json:"id" form:"id"  description:"ID"`
}

type DeptTreeListInput struct {
	Page     int    `json:"page" form:"page"  default:"1" description:"页码"`
	PageSize int    `json:"page_size" form:"page_size"  default:"20" description:"页数"`
	ParentId int    `json:"parent_id" form:"parent_id"  description:"名称"`
	Name     string `json:"name" form:"name"  description:"名称"`
	Status   int    `json:"status" form:"status" default:"" description:"状态"`
	Order    string `json:"order" form:"order" default:"id ASC" description:"排序"`
}
type DeptTreeListOut struct {
	Total int64              `json:"total" form:"total" description:"总数"`
	List  []DeptTreeListItem `json:"list" form:"list" description:"列表"`
}

type DeptGetOneInput struct {
	Id int `json:"id" form:"id" description:"ID"`
}
type DeptGetOneOut struct {
	model.SysDept
}

type DeptDeleteMultiInput struct {
	Ids []int `json:"ids" form:"ids" description:"ID"`
}

type DeptTreeListItem struct {
	SysDept
	Children []DeptTreeListItem `json:"children" gorm:"-"`
}
type SysDept struct {
	ID        int       `json:"id"`
	ParentId  int       `json:"parent_id"`
	Label     string    `json:"label"`
	Name      string    `json:"name"`
	Leader    string    `json:"leader"`
	Sort      int       `json:"sort"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Status    int8      `json:"status"`
	Operator  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"-"`
}
