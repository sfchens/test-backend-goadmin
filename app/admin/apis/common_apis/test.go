package common_apis

import (
	"csf/app/admin/request/common_request"
	"csf/library/response"
	"csf/utils"
	"github.com/gin-gonic/gin"
)

type cTestApi struct {
}

func NewTestApi() *cTestApi {
	return &cTestApi{}
}

func (c *cTestApi) Index(ctx *gin.Context) {
	var (
		err error
		req []common_request.TestIndexReq
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
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
		req common_request.TestIndexReq
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

}
