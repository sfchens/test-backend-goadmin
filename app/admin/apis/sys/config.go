package sys

import (
	"csf/app/admin/request/sys_req"
	"csf/core/query/config_query"
	"csf/core/service"
	"csf/library/response"
	"csf/utils"
	"github.com/gin-gonic/gin"
)

type cSysConfigApi struct {
}

func NewSysConfigApi() *cSysConfigApi {
	return &cSysConfigApi{}
}

// List  配置列表
// @Summary 配置列表
// @Description 配置列表
// @Tags 配置管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query     sys_req.ConfigListReq true "请求参数"
// @Success 200 {object} response.Response{data=sys_req.ConfigListRes} "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/config/list [get]
func (c cSysConfigApi) List(ctx *gin.Context) {
	var (
		err error
		req sys_req.ConfigListReq
		res sys_req.ConfigListRes
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	var (
		input config_query.ConfigListInput
		out   config_query.ConfigListOut
	)
	utils.StructToStruct(req, &input)
	out, err = service.NewConfigServiceGroup().ConfigService.List(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.SuccessWithStruct(ctx, out, &res)
}

// Add  添加配置
// @Summary 添加配置
// @Description 添加配置
// @Tags 配置管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body     sys_req.ConfigAddReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/config/add [post]
func (c cSysConfigApi) Add(ctx *gin.Context) {
	var (
		err error
		req sys_req.ConfigAddReq

		input config_query.ConfigAddInput
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(req, &input)
	err = service.NewConfigServiceGroup().ConfigService.Add(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// GetOne  获取一条配置
// @Summary 获取一条配置
// @Description 获取一条配置
// @Tags 配置管理
// @Accept application/json
// @Produce application/json
// @Param object query     sys_req.ConfigGetOneReq true "请求参数"
// @Success 200 {object} response.Response{data=sys_req.ConfigGetOneRes} "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/config/get_one [get]
func (c cSysConfigApi) GetOne(ctx *gin.Context) {
	var (
		err error
		req sys_req.ConfigGetOneReq
		res sys_req.ConfigGetOneRes
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	var (
		input config_query.ConfigGetOneInput
		out   config_query.ConfigGetOneOut
	)
	utils.StructToStruct(req, &input)
	out, err = service.NewConfigServiceGroup().ConfigService.GetOne(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.SuccessWithStruct(ctx, out, &res)
}

// Edit  编辑配置
// @Summary 编辑配置
// @Description 编辑配置
// @Tags 配置管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body     sys_req.ConfigEditReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/config/edit [post]
func (c cSysConfigApi) Edit(ctx *gin.Context) {

	var (
		err error
		req sys_req.ConfigEditReq

		input config_query.ConfigEditInput
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(req, &input)
	err = service.NewConfigServiceGroup().ConfigService.Edit(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// Delete  删除配置
// @Summary 删除配置
// @Description 删除配置
// @Tags 配置管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body     sys_req.ConfigDeleteReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/config/delete [post]
func (c cSysConfigApi) Delete(ctx *gin.Context) {

	var (
		err error
		req sys_req.ConfigDeleteReq

		input config_query.ConfigDeleteInput
	)

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(req, &input)
	err = service.NewConfigServiceGroup().ConfigService.Delete(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// SetStatus  设置状态
// @Summary 设置状态
// @Description 设置状态
// @Tags 配置管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body     sys_req.ConfigSetStatusReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/config/delete [post]
func (c cSysConfigApi) SetStatus(ctx *gin.Context) {

	var (
		err error
		req sys_req.ConfigSetStatusReq

		input config_query.ConfigSetStatusInput
	)

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(req, &input)
	err = service.NewConfigServiceGroup().ConfigService.SetStatus(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}
