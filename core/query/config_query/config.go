package config_query

type ConfigAddInput struct {
	Key    string `json:"key" form:"key" validate:"required" msg:"required:类型必填" description:"配置类型"`
	Name   string `json:"name" form:"name" validate:"required" msg:"required:名称必填" description:"名称"`
	Config string `json:"config" form:"config" validate:"required" msg:"required:配置必填" description:"配置名称"`
	IsOpen int    `json:"is_open" form:"is_open"  description:"是否开启"`
	Remark string `json:"remark" form:"remark" description:"备注"`
}

type ConfigEditInput struct {
	Id     int               `json:"id" form:"id" validate:"required" msg:"required:参数异常" description:"ID"`
	Key    string            `json:"key" form:"key" validate:"required" msg:"required:键不为空" description:"名称"`
	Name   string            `json:"name" form:"name" validate:"required" msg:"required:名称必填" description:"名称"`
	Config map[string]string `json:"config" form:"config" validate:"required" msg:"required:配置必填" description:"配置名称"`
	IsOpen int               `json:"is_open" form:"is_open"  description:"是否开启"`
	Remark string            `json:"remark" form:"remark" description:"备注"`
}

type ConfigGetOneInput struct {
	Id   int      `json:"id" form:"id"  description:"ID"`
	Key  string   `json:"key" form:"key" description:"key"`
	Key2 []string `json:"key2" form:"key2" description:"key"`
}
type ConfigGetOneOut struct {
	ID        uint        `json:"id" form:"id" comment:"ID"`
	Name      string      `json:"name" form:"name" comment:"名称"`
	Key       string      `json:"key" form:"key" comment:"0json配置1基础配置2商城配置3用户配置"`
	Config    interface{} `json:"config" form:"配置" comment:"配置"`
	IsOpen    uint        `json:"is_open" form:"is_open" comment:"是否开启"`
	Remark    string      `json:"remark" form:"remark" comment:"备注"`
	Operator  string      ` json:"operator"  form:"operator" comment:"操作人"`
	CreatedAt string      `json:"created_at" form:"created_at" comment:"创建时间"`
	UpdatedAt string      `json:"updated_at" form:"updated_at" comment:"更新时间"`
}

type ConfigListInput struct {
	Page     int    `json:"page" form:"page"  default:"1" description:"页码"`
	PageSize int    `json:"page_size" form:"page_size"  default:"20" description:"页数"`
	Name     string `json:"name" form:"name"  description:"名称"`
	Key      string `json:"key" form:"key"  description:"名称"`
	KeyName  string `json:"key_name" form:"key_name"  description:"名称"`
	Order    string `json:"order" form:"order" default:"id ASC" description:"排序"`
}
type ConfigListOut struct {
	Total int64             `json:"total" form:"total" description:"总数"`
	List  []ConfigGetOneOut `json:"list" form:"list" description:"列表"`
}

type ConfigDeleteInput struct {
	Ids []int `json:"ids" form:"ids" comment:"ID"`
}

type ConfigSetStatusOut struct {
	Id     int `json:"id" form:"id" description:"ID"`
	IsOpen int `json:"is_open" form:"is_open" description:"是否开启"`
}

type SysConfig struct {
	ID        uint        `json:"id" description:"ID"`
	Name      string      ` json:"name" description:"名称"`
	Type      string      ` json:"type" description:"类型"`
	Config    interface{} `json:"config" description:"配置"`
	Operator  string      `json:"operator" description:"操作人"`
	CreatedAt string      `json:"created_at" description:"创建时间"`
	UpdatedAt string      `json:"updated_at" description:"更新时间"`
}

type ConfigSetStatusInput struct {
	Id     int `json:"id" form:"id" description:"ID"`
	IsOpen int `json:"is_open" form:"is_open" description:"是否开启"`
}
