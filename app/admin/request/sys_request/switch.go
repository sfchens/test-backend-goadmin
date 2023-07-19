package sys_request

import "csf/common/mysql/model"

type SwitchAddReq struct {
	Name    string `json:"name" form:"name" validate:"required" msg:"required:名称必填" description:"名称"`
	TypeKey string `json:"type_key" form:"type_key" validate:"required" msg:"required:键名必填" description:"键名"`
	Status  int8   `json:"status" form:"status" validate:"required" msg:"required:状态必选" description:"状态"`
	Remark  string `json:"remark" form:"remark" validate:"required" msg:"required:备注必填" description:"备注"`
}

type SwitchAddOrEditReq struct {
	Id      int    `json:"id" form:"id"  description:"名称"`
	Name    string `json:"name" form:"name" validate:"required" msg:"required:名称必填" description:"名称"`
	TypeKey string `json:"type_key" form:"type_key" validate:"required" msg:"required:键名必填" description:"键名"`
	Status  int8   `json:"status" form:"status" validate:"required" msg:"required:状态必选" description:"状态"`
	Remark  string `json:"remark" form:"remark" validate:"required" msg:"required:备注必填" description:"备注"`
}

type SwitchListReq struct {
	Page     int    `json:"page" form:"page"  default:"1" description:"页码"`
	PageSize int    `json:"page_size" form:"page_size"  default:"20" description:"页数"`
	Name     string `json:"name" form:"name"  description:"名称"`
	TypeKey  string `json:"type_key" form:"type_key" description:"键名"`
}

type SwitchListRes struct {
	Total int64             `json:"total" form:"total" description:"总数"`
	List  []model.SysSwitch `json:"list" form:"list" description:"列表"`
}
