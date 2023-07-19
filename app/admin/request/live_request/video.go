package live_request

import "csf/common/mysql/model"

type VideoAddOrEditReq struct {
	Id     int    `json:"id" form:"id" description:"ID"`
	Name   string `json:"name" form:"name" validate:"required" msg:"required:标题必填" description:"标题"`
	Type   int    `json:"type" form:"type" validate:"required" msg:"required:类型必选" description:"类型"`
	Status int    `json:"status" form:"status" validate:"required" msg:"required:状态必填" description:"状态"`
	Url    string `json:"url" form:"url" validate:"required" msg:"required:请上传视频" description:"视频地址"`
}

type VideoListReq struct {
	Page     int    `json:"page" form:"page"  default:"1" description:"页码"`
	PageSize int    `json:"page_size" form:"page_size"  default:"20" description:"页数"`
	Name     string `json:"name" form:"name"  description:"名称"`
	Status   string `json:"status" form:"status" description:"状态"`
	Type     string `json:"type" form:"type" description:"类型"`
}

type VideoListRes struct {
	Total int64             `json:"total" form:"total" description:"总数"`
	List  []model.LiveVideo `json:"list" form:"list" description:"列表"`
}
