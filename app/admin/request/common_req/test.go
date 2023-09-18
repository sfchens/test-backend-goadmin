package common_req

type TestIndexReq struct {
	Id   int    `json:"id" form:"id" `
	Name string `json:"name" form:"name" validate:"required" msg:"required:名称是必填"`
	Test []Test `json:"test" form:"test" `
}

type TestIndexReq1 struct {
	Id   int    `json:"id" form:"id" default:"999" `
	Name string `json:"name" form:"name"`
	Test []Test `json:"test" form:"test"  default:"333"`
}

type Test struct {
	Age int    `json:"age" form:"age"`
	Sex string `json:"sex" form:"sex" default:"2"`
}
