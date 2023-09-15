package common

import (
	"csf/app/admin/request/common_req"
	"csf/library/response"
	"csf/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type cTestApi struct {
}

func NewTestApi() *cTestApi {
	return &cTestApi{}
}

func (c *cTestApi) Index(ctx *gin.Context) {
	var (
		err error
		req common_req.TestIndexReq
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	fmt.Printf("req: %+v\n", strings.Trim(req.Name, " "))
	//for _, item := range req {
	//	err = my_validator.MyValidator().Validate(item)
	//	if err != nil {
	//		break
	//	}
	//	err = utils.SetDefault(item)
	//	if err != nil {
	//		break
	//	}
	//}
	//if err != nil {
	//	response.FailWithMessage(ctx, err.Error())
	//	return
	//}
	response.Success(ctx)
}

func (c *cTestApi) Index2(ctx *gin.Context) {
	var (
		err error
		req common_req.TestIndexReq
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

}
