package sys_req

import (
	"time"
)

type RoleListReq struct {
	Page     int    `json:"page" form:"page"  default:"1" description:"页码"`
	PageSize int    `json:"page_size" form:"page_size"  default:"20" description:"页数"`
	Name     string `json:"name" form:"name"  description:"名称"`
	Key      string `json:"key" form:"key"  description:"权限标识"`
	Status   int    `json:"status" form:"status" description:"是否显示"`
}

type RoleListRes struct {
	Total int64          `json:"total" form:"total" description:"总数"`
	List  []RoleListItem `json:"list" form:"list" description:"列表"`
}

type RoleListItem struct {
	ID        int64     ` json:"id" form:"id" description:"表主键"`
	Name      string    `json:"name" form:"name" description:"名称"`
	Status    int       `json:"status" form:"status" description:"状态"`
	Key       string    `json:"key" form:"key" description:"权限标识"`
	Sort      int64     `json:"sort" form:"sort" description:"排序"`
	MenuIds   []int     `json:"menu_ids" form:"menu_ids" description:"菜单ID"`
	Remark    string    `json:"remark" form:"remark" description:"备注"`
	CreatedAt time.Time `json:"created_at" form:"created_at" description:"创建时间"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at" description:"最后更新时间"`
	DeletedAt time.Time `json:"deleted_at" form:"deleted_at" description:"删除时间"`
}

type RoleAddOrEditReq struct {
	Id      int    `json:"id" form:"id" description:"id"`
	Name    string `json:"name" form:"name" validate:"required" msg:"required:名称必填" description:"名称"`
	Status  int    `json:"status" form:"status" validate:"required" msg:"required:状态必填" description:"标题"`
	MenuIds []int  `json:"menu_ids" form:"menu_ids" description:"标题"`
	Key     string `json:"key" form:"key" validate:"required" msg:"required:状态必填" description:"权限标识必填"`
	Sort    int    `json:"sort" form:"sort"  description:"排序"`
	Remark  string `json:"remark" form:"remark"  description:"备注"`
}

type RoleDeleteBatchReq struct {
	Ids []int `json:"ids" form:"ids" description:"ID"`
}

type RoleDeleteReq struct {
	Id int `json:"id" form:"id" description:"ID"`
}
