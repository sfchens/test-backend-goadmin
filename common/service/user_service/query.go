package user_service

type LoginInfo struct {
	Id       int    `json:"id" form:"id" description:"ID"`
	Username string `json:"username" form:"username" description:"账号"`
	Realname string `json:"realname" form:"realname" description:"姓名"`
	Email    string `json:"email" form:"email" description:"邮箱"`
	Phone    string `json:"phone" form:"phone" description:"手机号码"`
}
