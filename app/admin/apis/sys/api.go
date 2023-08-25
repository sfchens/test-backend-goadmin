package sys

import (
	"csf/app/admin/request/sys_req"
	"csf/app/admin/service/sys_service"
	"csf/library/response"
	"csf/utils"
	"fmt"
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
// @Param raw body     sys.AdminSetStatusReq true "请求参数"
// @Success 200 {object} response.Response{data=sys.ApiListRes} "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/api/list [get]
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

	res, err = sys_service.NewSysApiService(ctx).List(req)
	if err != nil {
		response.SuccessWithData(ctx, err.Error())
		return
	}
	response.SuccessWithData(ctx, res)
}

// Refresh  刷新接口
// @Summary 刷新接口
// @Description 刷新接口
// @Tags 接口管理
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/api/refresh [get]
func (c *cSysApi) Refresh(ctx *gin.Context) {
	var (
		err error
	)

	err = sys_service.NewSysApiService(ctx).Refresh()
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
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/api/edit [post]
func (c *cSysApi) Edit(ctx *gin.Context) {
	var (
		err error

		req sys_req.ApiEditReq
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		fmt.Printf("err.Error():  %+v\n", err.Error())
		response.FailWithMessage(ctx, err.Error())
		return
	}

	err = sys_service.NewSysApiService(ctx).AddOrEdit(req)
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
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/api/get_tag [get]
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
	res, err = sys_service.NewSysApiService(ctx).GetTag(req)
	if err != nil {
		response.SuccessWithData(ctx, err.Error())
		return
	}
	response.SuccessWithData(ctx, res)
}

// DeleteMulti  接口分类
// @Summary 接口分类
// @Description 接口分类
// @Tags 接口管理
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/api/delete_multi [get]
func (c *cSysApi) DeleteMulti(ctx *gin.Context) {
	var (
		err error

		req sys_req.ApiDeleteMultiReq
	)

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err = sys_service.NewSysApiService(ctx).DeleteMulti(req)
	if err != nil {
		response.SuccessWithData(ctx, err.Error())
		return
	}
	response.Success(ctx)
}
