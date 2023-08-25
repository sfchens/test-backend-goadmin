package common_req

type GetOneRes struct {
	Id   string `json:"id"  description:"验证码ID"`
	Path string `json:"path"  description:"验证码地址"`
}
