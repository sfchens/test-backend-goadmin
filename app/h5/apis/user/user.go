package user

import (
	"csf/app/h5/request/user_req"
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

func (c *cUserApi) Register(ctx *gin.Context) {
	var (
		err error

		req user_req.UserRegisterReq

		input user_query.UserAddOrEditInput
	)

	err = utils.BindParams(ctx, &req, &input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err = service.NewUserServiceGroup().UserService.AddOrEdit(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

func (c *cUserApi) GetInfo(ctx *gin.Context) {
	var (
		err error
		req user_req.UserGetInfoReq
		res user_query.UserListItem
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	res, err = service.NewUserServiceGroup().UserService.GetInfo(ctx, req.Id)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.SuccessWithData(ctx, res)
}
