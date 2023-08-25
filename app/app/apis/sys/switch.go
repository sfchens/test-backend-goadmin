package sys

import (
	"csf/app/admin/request/sys_req"
	"csf/app/admin/service/sys_service"
	"csf/library/response"
	"csf/utils"
	"github.com/gin-gonic/gin"
)

type cSwitchApi struct{}

func NewSwitchApi() *cSwitchApi {
	return &cSwitchApi{}
}

// Add  添加开关配置
// @Summary 添加开关配置
// @Description 添加开关配置
// @Tags 开关管理
// @Accept application/json
// @Produce application/json
// @Param raw body     sys.SwitchAddOrEditReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/switch/add [post]
func (c *cSwitchApi) Add(ctx *gin.Context) {
	var (
		err error
		req sys_req.SwitchAddOrEditReq
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err = sys_service.NewSwitchService(ctx).AddOrEdit(req)
	response.Success(ctx)
}

// Edit  编辑开关配置
// @Summary 编辑开关配置
// @Description 编辑开关配置
// @Tags 开关管理
// @Accept application/json
// @Produce application/json
// @Param raw body     sys.SwitchAddOrEditReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/switch/edit [post]
func (c *cSwitchApi) Edit(ctx *gin.Context) {
	var (
		err error
		req sys_req.SwitchAddOrEditReq
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err = sys_service.NewSwitchService(ctx).AddOrEdit(req)
	response.Success(ctx)
}

// List  开关配置列表
// @Summary 开关配置列表
// @Description 开关配置列表
// @Tags 开关管理
// @Accept application/json
// @Produce application/json
// @Param raw body     sys.SwitchAddOrEditReq true "请求参数"
// @Success 200 {object} response.Response{data=sys.SwitchListRes} "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/switch/list [get]
func (c *cSwitchApi) List(ctx *gin.Context) {
	var (
		err error
		req sys_req.SwitchListReq
		res sys_req.SwitchListRes
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	res, err = sys_service.NewSwitchService(ctx).List(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.SuccessWithData(ctx, res)
}

// Delete  删除开关
// @Summary 删除开关
// @Description 删除开关
// @Tags 开关管理
// @Accept application/json
// @Produce application/json
// @Param raw body     sys.SwitchDeleteReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/switch/delete [post]
func (c *cSwitchApi) Delete(ctx *gin.Context) {
	var (
		err error
		req sys_req.SwitchDeleteReq
	)
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err = sys_service.NewSwitchService(ctx).Delete(req.Ids)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// SetStatus  设置状态
// @Summary 设置状态
// @Description 设置状态
// @Tags 开关管理
// @Accept application/json
// @Produce application/json
// @Param raw body     sys.SwitchDeleteReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/switch/set_status [post]
func (c *cSwitchApi) SetStatus(ctx *gin.Context) {
	var (
		err error
		req sys_req.SwitchSetStatusReq
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err = sys_service.NewSwitchService(ctx).SetStatus(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}
