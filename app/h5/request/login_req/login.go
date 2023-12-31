package login_req

import (
	"csf/core/query/login_query"
)

type LoginReq struct {
	Username  string `json:"username" form:"username" validate:"required" msg:"required:账号不能为空"`
	Password  string `json:"password" form:"password" validate:"required" msg:"required:密码不能为空"`
	Captcha   string `json:"captcha" form:"captcha"  default:"123456"`
	CaptChaId string `json:"captcha_id" form:"captcha_id"`
}

type LoginRes struct {
	AdminInfo login_query.H5UserInfo   `json:"userinfo" description:"用户信息"`
	TokenInfo login_query.TokenInfoOut `json:"token_info" description:"token信息"`
}
