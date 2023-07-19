package sys_request

type ConfigAddReq struct {
	Key    string `json:"key" form:"key" validate:"required" msg:"required:类型必填" description:"配置类型"`
	Name   string `json:"name" form:"name" validate:"required" msg:"required:名称必填" description:"名称"`
	Config string `json:"config" form:"config" validate:"required" msg:"required:配置必填" description:"配置名称"`
}

type ConfigEditReq struct {
	Id     int    `json:"id" form:"id" validate:"required" msg:"required:参数异常" description:"ID"`
	Name   string `json:"name" form:"name" validate:"required" msg:"required:名称必填" description:"名称"`
	Config string `json:"config" form:"config" validate:"required" msg:"required:配置必填" description:"配置名称"`
}

type ConfigGetOneReq struct {
	Id  int    `json:"id" form:"id"  description:"ID"`
	Key string `json:"key" form:"key" description:"key"`
}

type ConfigGetOneRes struct {
	ID        uint        `json:"id" description:"ID"`
	Name      string      ` json:"name" description:"名称"`
	Key       string      ` json:"key" description:"类型"`
	Config    interface{} `json:"config" description:"配置"`
	Operator  string      `json:"operator" description:"操作人"`
	CreatedAt string      `json:"created_at" description:"创建时间"`
	UpdatedAt string      `json:"updated_at" description:"更新时间"`
}

type ConfigListReq struct {
	Page     int    `json:"page" form:"page"  default:"1" description:"页码"`
	PageSize int    `json:"page_size" form:"page_size"  default:"20" description:"页数"`
	Name     string `json:"name" form:"name"  description:"名称"`
	Types    string `json:"type" form:"type" description:"配置类型"`
}

type ConfigListRes struct {
	Total int64             `json:"total" form:"total" description:"总数"`
	List  []ConfigGetOneRes `json:"list" form:"list" description:"列表"`
}
