package user_query

import "time"

type UserAddOrEditInput struct {
	Id       int    `json:"id" form:"id" description:"序号"`
	Username string `json:"username" form:"username" validate:"required" msg:"required:账号必填" description:"账号"`
	Realname string `json:"realname" form:"realname" validate:"required" msg:"required:姓名必填" description:"姓名"`
	Email    string `json:"email" form:"email"  description:"邮箱"`
	Phone    string `json:"phone" form:"phone"  description:"电话号码"`
	Password string `json:"password" form:"password" validate:"required" msg:"required:密码必填"  description:"密码"`
}

type UserResetPwdInput struct {
	Id       int    `json:"id" form:"id" validate:"required" msg:"required:参数异常"  description:"序号"`
	Password string `json:"password" form:"password" validate:"required" msg:"密码不能为空" description:"密码"`
}

type UserSetStatusInput struct {
	Id     int `json:"id" form:"id" validate:"required" msg:"required:参数异常" description:"序号"`
	Status int `json:"status" form:"status" validate:"required" msg:"required:参数异常" description:"状态"`
}

type UserListInput struct {
	Page     int    `json:"page" form:"page" default:"-1" validate:"required" msg:"required:页码不为空" description:"页码"`
	PageSize int    `json:"page_size" form:"page_size" default:"20" validate:"required" msg:"required:页数不为空" description:"页数"`
	Username string `json:"username" form:"username"  description:"名称"`
	Realname string `json:"realname" form:"realname" description:"真名"`
	Email    string `json:"email" form:"email" description:"邮箱"`
	Phone    string `json:"phone" form:"phone" description:"手机号码"`
	Status   int    `json:"status" form:"status" default:"-1" validate:"required" msg:"required:状态不能为空" description:"状态"`
}
type UserListOut struct {
	Total int64          `json:"total" form:"total" description:"总数"`
	List  []UserListItem `json:"list" form:"list" description:"列表"`
}
type UserListItem struct {
	ID         uint      `json:"id"`          // ID
	Username   string    `json:"username"`    // 用户名
	Realname   string    `json:"realname"`    // 真实姓名
	Email      string    `json:"email"`       // 邮箱
	Phone      string    `json:"phone"`       // 电话
	HeadPic    string    `json:"head_pic"`    // 头像
	LastIP     string    `json:"last_ip"`     // 最后登录IP
	LastTime   int       `json:"last_time"`   // 最后登录时间
	LoginCount int       `json:"login_count"` // 登录次数
	Status     int       `json:"status"`      // 状态
	CreatedAt  time.Time `json:"created_at"`  // 创建时间
	UpdatedAt  time.Time `json:"updated_at"`  // 更新时间
}
