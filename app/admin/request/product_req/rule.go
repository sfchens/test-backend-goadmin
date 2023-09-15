package product_req

type RuleAddOrEditReq struct {
	Id    int                   `json:"id" form:"id" description:"id"`
	Name  string                `json:"name" form:"name"  validate:"required" msg:"名称不为空" description:"名称"`
	Value []map[string][]string `json:"value" form:"value" validate:"required" msg:"规格不为空" description:"规格值"`
}

type RuleListReq struct {
	Page     int    `json:"page" form:"page"  default:"1" description:"页码"`
	PageSize int    `json:"page_size" form:"page_size"  default:"20" description:"页数"`
	Name     string `json:"name" form:"name" description:"名称"`
}
type RuleListRes struct {
	Total int         `json:"total" form:"total"   description:"总数"`
	List  interface{} `json:"list" form:"list"   description:"数组"`
}
