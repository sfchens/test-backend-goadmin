package sys_request

import "csf/common/mysql/model"

type DeptAddOrEditReq struct {
	Id       int    `json:"id" form:"id"  description:"ID"`
	ParentId int    `json:"parent_id" form:"parent_id" validate:"required" msg:"required:上级部门必选" description:"上级部门"`
	Name     string `json:"name" form:"name" validate:"required" msg:"required:部门名称必填" description:"部门名称"`
	Leader   string `json:"leader" form:"leader" validate:"required" msg:"required:负责人必填" description:"负责人"`
	Sort     int    `json:"sort" form:"sort" default:"1" description:"排序"`
	Phone    string `json:"phone" form:"phone" description:"手机号码"`
	Email    string `json:"email" form:"email" validate:"email" msg:"email:邮箱格式有误" description:"邮箱"`
	Status   int8   `json:"status" form:"status" validate:"required" msg:"required:状态必选" description:"状态"`
}

type DeptDeleteReq struct {
	Id int `json:"id" form:"id"  description:"ID"`
}

type DeptListReq struct {
	Page     int    `json:"page" form:"page"  default:"1" description:"页码"`
	PageSize int    `json:"page_size" form:"page_size"  default:"20" description:"页数"`
	Name     string `json:"name" form:"name"  description:"名称"`
	Status   int    `json:"status" form:"status" default:"-1" description:"状态"`
}

type DeptListRes struct {
	Total int64           `json:"total" form:"total" description:"总数"`
	List  []model.SysDept `json:"list" form:"list" description:"列表"`
}
