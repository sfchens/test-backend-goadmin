package sys_request

import "csf/app/admin/model/sys_model"

type LoginReq struct {
	Username  string `json:"username" form:"username" validate:"required" msg:"required:账号不能为空"`
	Password  string `json:"password" form:"password" validate:"required" msg:"required:密码不能为空"`
	Captcha   string `json:"captcha" form:"captcha"  default:"123456"`
	CaptChaId string `json:"captcha_id" form:"captcha_id"`
}

type LoginRes struct {
	AdminInfo AdminInfo              `json:"userinfo" description:"用户信息"`
	TokenInfo sys_model.TokenInfoOut `json:"token_info" description:"token信息"`
}

type AdminInfo struct {
	Id       int    `json:"id" description:"ID"`
	Username string `json:"username" description:"账号"`
	Realname string `json:"realname"`
	Email    string `json:"email" description:"邮箱"`
	Phone    string `json:"phone" description:"手机号码"`
}
