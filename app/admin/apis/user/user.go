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

// Add  添加用户
// @Summary 添加用户
// @Description 添加用户
// @Tags 用户管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body  user_req.UserAddOrEditReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/user/add [post]
func (c *cUserApi) Add(ctx *gin.Context) {
	var (
		err error
		req user_req.UserAddOrEditReq

		input user_query.UserAddOrEditInput
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(req, &input)
	err = service.NewUserServiceGroup().UserService.AddOrEdit(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// List  用户列表
// @Summary 用户列表
// @Description 用户列表
// @Tags 用户管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body  user_req.UserListReq true "请求参数"
// @Success 200 {object} response.Response{data=user_req.UserListRes} "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/user/list [get]
func (c *cUserApi) List(ctx *gin.Context) {

	var (
		err error

		req user_req.UserListReq
		res user_req.UserListRes
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	var (
		input user_query.UserListInput
		out   user_query.UserListOut
	)
	utils.StructToStruct(req, &input)
	out, err = service.NewUserServiceGroup().UserService.List(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.SuccessWithStruct(ctx, out, &res)
}

// ResetPwd  密码重置
// @Summary 密码重置
// @Description 密码重置
// @Tags 用户管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body  user_req.UserResetPwdReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/user/reset_pwd [post]
func (c *cUserApi) ResetPwd(ctx *gin.Context) {
	var (
		err error
		req user_req.UserResetPwdReq

		input user_query.UserResetPwdInput
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	utils.StructToStruct(req, &input)

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

// SetStatus  设置状态
// @Summary 设置状态
// @Description 设置状态
// @Tags 用户管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body  user_req.UserSetStatusReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/user/set_status [post]
func (c *cUserApi) SetStatus(ctx *gin.Context) {
	var (
		err error
		req user_req.UserSetStatusReq

		input user_query.UserSetStatusInput
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(req, &input)
	err = service.NewUserServiceGroup().UserService.SetStatus(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// GetInfo  用户信息
// @Summary 用户信息
// @Description 用户信息
// @Tags 用户管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query  user_req.UserGetInfoReq true "请求参数"
// @Success 200 {object} response.Response{data=user_query.UserListItem} "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/user/set_status [get]
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
