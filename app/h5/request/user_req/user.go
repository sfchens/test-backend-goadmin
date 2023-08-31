package user_req

type UserRegisterReq struct {
	Username string `json:"username" form:"username" validate:"required" msg:"required:账号必填" description:"账号"`
	Realname string `json:"realname" form:"realname" validate:"required" msg:"required:姓名必填" description:"姓名"`
	Email    string `json:"email" form:"email"  description:"邮箱"`
	Phone    string `json:"phone" form:"phone"  description:"电话号码"`
	Password string `json:"password" form:"password" validate:"required" msg:"required:密码必填"  description:"密码"`
}

type UserGetInfoReq struct {
	Id int `json:"id" form:"id"  validate:"required" msg:"required:参数异常"    description:"序号"`
}
