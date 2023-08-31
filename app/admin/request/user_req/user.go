package user_req

import "csf/core/query/user_query"

type UserAddOrEditReq struct {
	Id       int    `json:"id" form:"id"  description:"序号"`
	Username string `json:"username" form:"username" validate:"required" msg:"required:账号必填" description:"账号"`
	Realname string `json:"realname" form:"realname" validate:"required" msg:"required:姓名必填" description:"姓名"`
	Email    string `json:"email" form:"email"  description:"邮箱"`
	Phone    string `json:"phone" form:"phone"  description:"电话号码"`
	Password string `json:"password" form:"password" validate:"required" msg:"required:密码必填"  description:"密码"`
	Status   int    `json:"status" form:"status" validate:"required" msg:"required:状态必填"  description:"状态"`
}

type UserResetPwdReq struct {
	Id       int    `json:"id" form:"id" validate:"required" msg:"required:参数异常"  description:"序号"`
	Password string `json:"password" form:"password" description:"密码"`
}

type UserSetStatusReq struct {
	Id     int `json:"id" form:"id"  validate:"required" msg:"required:参数异常"    description:"序号"`
	Status int `json:"status" form:"status" validate:"required" msg:"required:参数异常"   description:"状态"`
}

type UserListReq struct {
	Page     int    `json:"page" form:"page"  default:"1" description:"页码"`
	PageSize int    `json:"page_size" form:"page_size"  default:"20" description:"页数"`
	Username string `json:"username" form:"username"  description:"名称"`
	Realname string `json:"realname" form:"realname" description:"真名"`
	Email    string `json:"email" form:"email" description:"邮箱"`
	Phone    string `json:"phone" form:"phone" description:"手机号码"`
	Status   int    `json:"status" form:"status" default:"-1" description:"状态"`
}
type UserListRes struct {
	Total int64                     `json:"total" form:"total" description:"总数"`
	List  []user_query.UserListItem `json:"list" form:"list" description:"列表"`
}

type UserGetInfoReq struct {
	Id int `json:"id" form:"id"  validate:"required" msg:"required:参数异常"    description:"序号"`
}
