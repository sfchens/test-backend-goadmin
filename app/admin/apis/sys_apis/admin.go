package sys_apis

import (
	"csf/app/admin/request/sys_request"
	"csf/app/admin/service/sys_service"
	"csf/common/mysql/model"
	"csf/library/global"
	"csf/library/response"
	"csf/utils"
	"github.com/gin-gonic/gin"
)

type cSysAdminApi struct {
}

func NewSysAdminApi() *cSysAdminApi {
	return &cSysAdminApi{}
}

// List 管理员列表
// @Summary 管理员列表
// @Description 管理员列表
// @Tags 管理员管理
// @Accept application/json
// @Produce application/json
// @Param raw body     sys_request.AdminListReq true "请求参数"
// @Success 200 {object} response.Response{data=sys_request.AdminListRes}  "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/admin/list [get]
func (c *cSysAdminApi) List(ctx *gin.Context) {
	var (
		err error
		req sys_request.AdminListReq
		res sys_request.AdminListRes
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	res, err = sys_service.NewSysAdminService(ctx).List(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.SuccessWithData(ctx, res)
}

// Add  添加管理员
// @Summary 添加管理员
// @Description 添加管理员
// @Tags 管理员管理
// @Accept application/json
// @Produce application/json
// @Param raw body     sys_request.AdminAddOrEditReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/admin/add [post]
func (c *cSysAdminApi) Add(ctx *gin.Context) {
	var (
		err error
		req sys_request.AdminAddOrEditReq
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err = sys_service.NewSysAdminService(ctx).Add(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// SetStatus  设置状态
// @Summary 设置状态
// @Description 设置状态
// @Tags 管理员管理
// @Accept application/json
// @Produce application/json
// @Param raw body     sys_request.AdminSetStatusReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/admin/set_status [post]
func (c *cSysAdminApi) SetStatus(ctx *gin.Context) {
	var (
		err error
		req sys_request.AdminSetStatusReq
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err = sys_service.NewSysAdminService(ctx).SetStatus(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// GetAdminInfo  管理员信息
// @Summary 管理员信息
// @Description 管理员信息
// @Tags 管理员管理
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=sys_request.AdminListItem} "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/admin/get_admin_info [get]
func (c *cSysAdminApi) GetAdminInfo(ctx *gin.Context) {
	var (
		err error

		adminInfo  sys_request.AdminInfoRes
		adminModel model.SysAdmin
	)
	adminModel, err = sys_service.NewSysAdminService(ctx).GetAdminInfo()
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	utils.StructToStruct(adminModel, &adminInfo)
	adminInfo.Roles = global.Permissions
	adminInfo.Permissions = global.Permissions
	response.SuccessWithData(ctx, adminInfo)
}

// ResetPwd  重置密码
// @Summary 重置密码
// @Description 重置密码
// @Tags 管理员管理
// @Accept application/json
// @Produce application/json
// @Param raw body     sys_request.AdminResetPwd true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/admin/reset_pwd [post]
func (c *cSysAdminApi) ResetPwd(ctx *gin.Context) {
	var (
		err error

		req sys_request.AdminResetPwdReq
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err = sys_service.NewSysAdminService(ctx).ResetPwd(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// DeleteBatch  批量删除
// @Summary 批量删除
// @Description 批量删除
// @Tags 管理员管理
// @Accept application/json
// @Produce application/json
// @Param raw body     sys_request.AdminDeleteBatchReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/admin/delete_batch [post]
func (c *cSysAdminApi) DeleteBatch(ctx *gin.Context) {
	var (
		err error

		req sys_request.AdminDeleteBatchReq
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err = sys_service.NewSysAdminService(ctx).DeleteBatch(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

func (c *cSysAdminApi) SetRole(ctx *gin.Context) {
	var (
		err error

		req sys_request.AdminSetRoleReq
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	err = sys_service.NewSysAdminService(ctx).SetRole(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}
