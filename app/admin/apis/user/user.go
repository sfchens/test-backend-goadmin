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

	err = service.NewUserServiceGroup().UserService.AddOrEdit(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

func (c *cUserApi) List(ctx *gin.Context) {

	var (
		err error

		req user_req.UserListReq
		res user_req.UserListRes

		input user_query.UserListInput
	)
	err = utils.BindParams(ctx, &req, &input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	res.Total, res.List, err = service.NewUserServiceGroup().UserService.List(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.SuccessWithData(ctx, res)
}

func (c *cUserApi) ResetPwd(ctx *gin.Context) {
	var (
		err error
		req user_req.UserResetPwdReq

		input user_query.UserResetPwdInput
	)

	err = utils.BindParams(ctx, &req, &input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	if req.Password == "" {
		input.Password = "123456"
	}
	err = service.NewUserServiceGroup().UserService.ResetPwd(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

func (c *cUserApi) SetStatus(ctx *gin.Context) {
	var (
		err error
		req user_req.UserSetStatusReq

		input user_query.UserSetStatusInput
	)
	err = utils.BindParams(ctx, &req, &input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err = service.NewUserServiceGroup().UserService.SetStatus(ctx, input)
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
