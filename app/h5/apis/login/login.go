package login

import (
	"csf/app/h5/request/login_req"
	"csf/core/query/login_query"
	"csf/core/service"
	"csf/library/response"
	"csf/utils"
	"github.com/gin-gonic/gin"
)

type cLogin struct{}

func NewLoginApi() *cLogin {
	return &cLogin{}
}

func (c *cLogin) Login(ctx *gin.Context) {
	var (
		err error

		req login_req.LoginReq
		res login_query.H5LoginOut

		input login_query.H5LoginInput
	)

	err = utils.BindParams(ctx, &req, &input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	res, err = service.NewLoginServiceGroup().H5LoginService.Login(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.SuccessWithData(ctx, res)
}

func (c *cLogin) Logout(ctx *gin.Context) {
	var (
		err error
	)
	err = service.NewLoginServiceGroup().H5LoginService.Logout(ctx)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
	}
	response.SuccessWithMessage(ctx, "退出成功")
}
