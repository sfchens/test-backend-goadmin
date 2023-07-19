package common_request

type TestIndexReq struct {
	Id   int    `json:"id" form:"id" validate:"required" msg:"required:id是必填"`
	Name string `json:"name" form:"name" validate:"required" msg:"required:名称是必填"`
	Test string `json:"test" form:"test"  default:"333"`
}
