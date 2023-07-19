package sys_model

type LoginInput struct {
	Username string `json:"username" description:"账号"`
	Password string `json:"password" description:"密码"`
}

type LoginOut struct {
	Id       int    `json:"id" description:"ID"`
	Username string `json:"username" description:"账号"`
	Email    string `json:"email" description:"邮箱"`
	Password string `json:"password" description:"密码"`
	Realname string `json:"realname" description:"真名"`
}

type TokenInfoOut struct {
	Token     string `json:"token" description:"token"`
	ExpiresAt int64  `json:"expires_at" description:"有效期"`
}
