package login_query

type AdminLoginInput struct {
	Username  string `json:"username" form:"username" validate:"required" msg:"required:账号不能为空"`
	Password  string `json:"password" form:"password" validate:"required" msg:"required:密码不能为空"`
	Captcha   string `json:"captcha" form:"captcha"  default:"123456"`
	CaptChaId string `json:"captcha_id" form:"captcha_id"`
}
type AdminLoginOut struct {
	AdminInfo AdminInfo    `json:"userinfo" description:"用户信息"`
	TokenInfo TokenInfoOut `json:"token_info" description:"token信息"`
}
type AdminInfo struct {
	Id       int    `json:"id" description:"ID"`
	Username string `json:"username" description:"账号"`
	Realname string `json:"realname"`
	Email    string `json:"email" description:"邮箱"`
	Phone    string `json:"phone" description:"手机号码"`
}
type TokenInfoOut struct {
	Token     string `json:"token" description:"token"`
	ExpiresAt int64  `json:"expires_at" description:"有效期"`
}
