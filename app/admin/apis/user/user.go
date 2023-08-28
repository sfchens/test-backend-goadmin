package user

import (
	"csf/app/admin/request/user_req"
	"csf/core/query/user_query"
	"csf/core/service"
	"csf/library/response"
	"csf/utils"
	"github.com/gin-gonic/gin"
)

type cUserApi struct {
}

func NewUserApi() *cUserApi {
	return &cUserApi{}
}

func (c *cUserApi) Add(ctx *gin.Context) {
	var (
		err error

		req user_req.UserAddOrEditReq

		input user_query.UserAddOrEditInput
	)

	err = utils.BindParams(ctx, &req, &input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	err = service.NewUserServiceGroup().UserService.Add(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

func (c *cUserApi) List(ctx *gin.Context) {

}
