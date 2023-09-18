package sys

import (
	"csf/app/admin/request/sys_req"
	"csf/core/query/sys_query"
	"csf/core/service"
	"csf/library/response"
	"csf/utils"
	"github.com/gin-gonic/gin"
)

type cSysApi struct{}

func NewSysApi() *cSysApi {
	return &cSysApi{}
}

// List  接口列表
// @Summary 接口列表
// @Description 接口列表
// @Tags 接口管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query  sys_req.ApiListReq true "请求参数"
// @Success 200 {object} response.Response{data=sys_req.ApiListRes} "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/api/list [get]
func (c *cSysApi) List(ctx *gin.Context) {
	var (
		err error
		req sys_req.ApiListReq
		res sys_req.ApiListRes
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	var (
		input sys_query.ApiListInput
		out   sys_query.ApiListOut
	)
	 utils.StructToStruct(req, &input)
	out, err = service.NewSysServiceGroup().ApiService.List(ctx, input)
	if err != nil {
		response.SuccessWithData(ctx, err.Error())
		return
	}
	response.SuccessWithStruct(ctx, out, &res)
}

// Refresh  刷新接口
// @Summary 刷新接口
// @Description 刷新接口
// @Tags 接口管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/api/refresh [post]
func (c *cSysApi) Refresh(ctx *gin.Context) {
	var (
		err error
	)
	err = service.NewSysServiceGroup().ApiService.Refresh()
	if err != nil {
		response.SuccessWithData(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// Edit  编辑接口
// @Summary 编辑接口
// @Description 编辑接口
// @Tags 接口管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body  sys_req.ApiEditReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/api/edit [post]
func (c *cSysApi) Edit(ctx *gin.Context) {
	var (
		err error

		req   sys_req.ApiEditReq
		input sys_query.ApiEditInput
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(req, &input)
	err = service.NewSysServiceGroup().ApiService.AddOrEdit(ctx, input)
	if err != nil {
		response.SuccessWithData(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// GetTag  接口分类
// @Summary 接口分类
// @Description 接口分类
// @Tags 接口管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query  sys_req.ApiGetTagReq true "请求参数"
// @Success 200 {object} response.Response{data=sys_req.ApiGetTagRes} "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/api/get_tag [get]
func (c *cSysApi) GetTag(ctx *gin.Context) {
	var (
		err error
		req sys_req.ApiGetTagReq
		res sys_req.ApiGetTagRes
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	var (
		input sys_query.ApiGetTagInput
		out   sys_query.ApiGetTagOut
	)
	utils.StructToStruct(req, &input)
	out, err = service.NewSysServiceGroup().ApiService.GetTag(ctx, input)
	if err != nil {
		response.SuccessWithData(ctx, err.Error())
		return
	}
	response.SuccessWithStruct(ctx, out, &res)
}

// DeleteMulti  接口分类
// @Summary 接口分类
// @Description 接口分类
// @Tags 接口管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query  sys_query.ApiDeleteMultiInput true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/api/delete_multi [get]
func (c *cSysApi) DeleteMulti(ctx *gin.Context) {
	var (
		err   error
		input sys_query.ApiDeleteMultiInput
	)

	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err = service.NewSysServiceGroup().ApiService.DeleteMulti(ctx, input)
	if err != nil {
		response.SuccessWithData(ctx, err.Error())
		return
	}
	response.Success(ctx)
}
