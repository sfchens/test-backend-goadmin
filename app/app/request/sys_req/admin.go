package sys_req

import (
	"csf/core/mysql/model"
	"time"
)

type AdminAddOrEditReq struct {
	Id       int    `json:"id" form:"id" validate:"required" msg:"required:参数异常" description:"ID"`
	Username string `json:"username" form:"username" validate:"required" msg:"required:账号必填" description:"账号"`
	Realname string `json:"realname" form:"realname" validate:"required" msg:"required:姓名必填" description:"姓名"`
	Email    string `json:"email" form:"email"  description:"邮箱"`
	Phone    string `json:"phone" form:"phone"  description:"电话号码"`
	Remark   string `json:"remark" form:"remark"  description:"电话号码"`
	Sex      int    `json:"sex" form:"sex" validate:"required" msg:"required:性别必选"   description:"电话号码"`
	DeptId   int    `json:"dept_id" form:"dept_id"   description:"密码"`
	RoleIds  []int  `json:"role_ids" form:"role_ids" validate:"required" msg:"required:角色必选"  description:"密码"`
	Password string `json:"password" form:"password"  description:"密码"`
	Status   int    `json:"status" form:"status" validate:"required" msg:"required:状态必选"  description:"密码"`
}

type AdminEditReq struct {
	Username string `json:"username" form:"username" validate:"required" msg:"required:账号必填" description:"账号"`
	Realname string `json:"realname" form:"realname" validate:"required" msg:"required:姓名必填" description:"姓名"`
	Email    string `json:"email" form:"email"  description:"邮箱"`
	Phone    string `json:"phone" form:"phone"  description:"电话号码"`
	Password string `json:"password" form:"password" validate:"required" msg:"required:密码必填"  description:"密码"`
}

type AdminSetStatusReq struct {
	Id     int   `json:"id" form:"id" validate:"required" msg:"required:参数异常" description:"ID"`
	Status uint8 `json:"status" form:"status" validate:"required" msg:"required:状态异常" description:"状态"`
}

type AdminListReq struct {
	Page     int    `json:"page" form:"page"  default:"1" description:"页码"`
	PageSize int    `json:"page_size" form:"page_size"  default:"20" description:"页数"`
	Username string `json:"username" form:"username"  description:"名称"`
	Realname string `json:"realname" form:"realname" description:"真名"`
	Email    string `json:"email" form:"email" description:"邮箱"`
	Phone    string `json:"phone" form:"phone" description:"手机号码"`
	Status   int    `json:"status" form:"status" default:"-1" description:"状态"`
}

type AdminListRes struct {
	Total int64           `json:"total" form:"total" description:"总数"`
	List  []AdminListItem `json:"list" form:"list" description:"列表"`
}

type AdminListItem struct {
	ID          uint          `json:"id"`            // ID
	Username    string        `json:"username"`      // 用户名
	Realname    string        `json:"realname"`      // 真实姓名
	Email       string        `json:"email"`         // 邮箱
	Phone       string        `json:"phone"`         // 电话
	RoleIds     []int         `json:"role_ids"`      // 电话
	RoleIdsText string        `json:"role_ids_text"` // 电话
	HeadPic     string        `json:"head_pic"`      // 头像
	Password    string        `json:"password"`      // 密码
	LastIP      string        `json:"last_ip"`       // 最后登录IP
	LastTime    int           `json:"last_time"`     // 最后登录时间
	LoginCount  int           `json:"login_count"`   // 登录次数
	Status      int           `json:"status"`        // 状态
	Operator    string        `json:"operator"`      // 操作人
	CreatedAt   time.Time     `json:"created_at"`    // 创建时间
	UpdatedAt   time.Time     `json:"updated_at"`    // 更新时间
	DeptID      int           `json:"dept_id"`       // 部门ID
	Sex         int           `json:"sex"`           // 性别
	Remark      string        `json:"remark"`        // 备注
	DeptInfo    model.SysDept `json:"dept_info" gorm:"-"`
}

type AdminInfoRes struct {
	Id           int      `json:"id" form:"id" description:"ID"`
	Username     string   `json:"username" form:"username"  description:"名称"`
	Realname     string   `json:"realname" form:"realname" description:"真名"`
	Email        string   `json:"email" form:"email" description:"邮箱"`
	Phone        string   `json:"phone" form:"phone" description:"手机号码"`
	Roles        []string `json:"roles" form:"roles" description:"手机号码"`
	Permissions  []string `json:"permissions" form:"permissions" description:"权限"`
	Avatar       string   `json:"avatar" form:"avatar" description:"头像"`
	Introduction string   `json:"introduction" form:"introduction" description:"介绍"`
}

type AdminResetPwdReq struct {
	Id int `json:"id" form:"id" description:"ID"`
}

type AdminDeleteBatchReq struct {
	Ids []int `json:"ids" form:"ids" description:"ID"`
}

type AdminSetRoleReq struct {
	Id      int    `json:"id" form:"id" description:"ID"`
	RoleIds []int  `json:"role_ids" form:"role_ids" validate:"required" msg:"required:角色必选"  description:"密码"`
	Remark  string `json:"remark" form:"remark"  description:"备注"`
}
