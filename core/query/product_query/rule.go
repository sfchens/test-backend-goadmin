package product_query

import "time"

type RuleAddOrEditInput struct {
	Id    int                   `json:"id" form:"id" description:"id"`
	Name  string                `json:"name" form:"name" validate:"required" msg:"名称不为空" description:"名称"`
	Value []map[string][]string `json:"value" form:"value" validate:"required" msg:"规则不为空"  description:"规格值"`
}

type RuleListInput struct {
	Page     int    `json:"page" form:"page"  default:"1" description:"页码"`
	PageSize int    `json:"page_size" form:"page_size"  default:"20" description:"页数"`
	Name     string `json:"name" form:"name" description:"名称"`
}
type RuleListOut struct {
	Total int64          `json:"total" form:"total" description:"总数"`
	List  []RuleModelOut `json:"list" form:"list" description:"列表"`
}
type RuleModelOut struct {
	ID        int                   `json:"id" description:"自增ID"`
	Name      string                `json:"name" description:"规格名称"`
	Value     []map[string][]string `json:"value" description:"规格值"`
	Operator  string                `json:"operator" description:"操作人"`
	CreatedAt time.Time             `json:"created_at" description:"创建时间"`
	UpdatedAt time.Time             `json:"updated_at" description:"更新时间"`
}
