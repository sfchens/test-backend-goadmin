package sys

import (
	"csf/app/admin/model/sys_model"
	"csf/app/admin/request/sys_req"
	"csf/app/admin/service/sys_service"
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
// @Param raw body     sys.ConfigListReq true "请求参数"
// @Success 200 {object} response.Response{data=sys.ConfigListRes} "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/config/list [get]
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

	res, err = sys_service.NewSysConfigService(ctx).List(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.SuccessWithData(ctx, res)
}

// Add  添加配置
// @Summary 添加配置
// @Description 添加配置
// @Tags 配置管理
// @Accept application/json
// @Produce application/json
// @Param raw body     sys.ConfigAddReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/config/add [post]
func (c cSysConfigApi) Add(ctx *gin.Context) {
	var (
		err error
		req sys_req.ConfigAddReq
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err = sys_service.NewSysConfigService(ctx).Add(req)
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
// @Param raw body     sys.ConfigGetOneReq true "请求参数"
// @Success 200 {object} response.Response{data=sys.ConfigGetOneRes} "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/config/get_one [get]
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

	var out sys_model.SysConfig
	out, err = sys_service.NewSysConfigService(ctx).GetOne(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(out, &res)
	response.SuccessWithData(ctx, res)
}

// Edit  编辑配置
// @Summary 编辑配置
// @Description 编辑配置
// @Tags 配置管理
// @Accept application/json
// @Produce application/json
// @Param raw body     sys.ConfigEditReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/config/edit [post]
func (c cSysConfigApi) Edit(ctx *gin.Context) {

	var (
		err error
		req sys_req.ConfigEditReq
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err = sys_service.NewSysConfigService(ctx).Edit(req)
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
// @Param raw body     sys.ConfigDeleteReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/config/delete [post]
func (c cSysConfigApi) Delete(ctx *gin.Context) {

	var (
		err error
		req sys_req.ConfigDeleteReq
	)

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err = sys_service.NewSysConfigService(ctx).Delete(req)
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
// @Param raw body     sys.ConfigEditReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/config/delete [post]
func (c cSysConfigApi) SetStatus(ctx *gin.Context) {

	var (
		err error
		req sys_req.ConfigSetStatusReq
	)

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err = sys_service.NewSysConfigService(ctx).SetStatus(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}
