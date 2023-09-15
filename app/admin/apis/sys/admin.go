package sys

import (
	"csf/app/admin/request/sys_req"
	"csf/core/mysql/model"
	"csf/core/query/sys_query"
	"csf/core/service"
	"csf/library/easy_session"
	"csf/library/global"
	"csf/library/response"
	"csf/utils"
	"encoding/json"
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
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body     sys_req.AdminListReq true "请求参数"
// @Success 200 {object} response.Response{data=sys_query.AdminListOut}  "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/admin/list [get]
func (c *cSysAdminApi) List(ctx *gin.Context) {
	var (
		err error
		req sys_req.AdminListReq

		input sys_query.AdminListInput
		res   sys_query.AdminListOut
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(req, &input)
	res, err = service.NewSysServiceGroup().AdminService.List(ctx, input)
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
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body     sys_req.AdminAddOrEditReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/admin/add [post]
func (c *cSysAdminApi) Add(ctx *gin.Context) {
	var (
		err error
		req sys_req.AdminAddOrEditReq

		input sys_query.AdminAddOrEditInput
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(req, &input)
	err = service.NewSysServiceGroup().AdminService.Add(ctx, input)
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
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body     sys_req.AdminSetStatusReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/admin/set_status [post]
func (c *cSysAdminApi) SetStatus(ctx *gin.Context) {
	var (
		err error
		req sys_req.AdminSetStatusReq

		input sys_query.AdminSetStatusInput
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(req, &input)
	err = service.NewSysServiceGroup().AdminService.SetStatus(ctx, input)
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
// @Param Authorization header string true "Bearer 用户令牌"
// @Success 200 {object} response.Response{data=sys_req.AdminInfoRes} "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/admin/get_admin_info [get]
func (c *cSysAdminApi) GetAdminInfo(ctx *gin.Context) {
	var (
		err error

		adminInfo  sys_req.AdminInfoRes
		adminModel model.SysAdmin
	)
	adminModel, err = service.NewSysServiceGroup().AdminService.GetAdminInfo(ctx)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	utils.StructToStruct(adminModel, &adminInfo)
	adminInfo.Roles = []string{"admin"}
	adminInfo.Permissions = global.Permissions

	// 保存session
	sessionStore := easy_session.NewCustomSession(ctx)
	err = sessionStore.Set(global.LoginTypeKey, global.ModuleAdmin)
	if err != nil {
		return
	}
	jsonStr, _ := json.Marshal(adminInfo)
	err = sessionStore.Set(global.UserInfoKey, string(jsonStr))
	if err != nil {
		return
	}
	response.SuccessWithData(ctx, adminInfo)
}

// ResetPwd  重置密码
// @Summary 重置密码
// @Description 重置密码
// @Tags 管理员管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body     sys_req.AdminResetPwdReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/admin/reset_pwd [post]
func (c *cSysAdminApi) ResetPwd(ctx *gin.Context) {
	var (
		err error

		req   sys_req.AdminResetPwdReq
		input sys_query.AdminResetPwdInput
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(req, &input)
	err = service.NewSysServiceGroup().AdminService.ResetPwd(ctx, input)
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
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body sys_req.AdminDeleteBatchReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/admin/delete_batch [post]
func (c *cSysAdminApi) DeleteBatch(ctx *gin.Context) {
	var (
		err error

		req   sys_req.AdminDeleteBatchReq
		input sys_query.AdminDeleteBatchInput
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(req, &input)
	err = service.NewSysServiceGroup().AdminService.DeleteBatch(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// SetRole  设置规则
// @Summary 设置规则
// @Description 设置规则
// @Tags 管理员管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body sys_req.AdminSetRoleReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/admin/set_role [post]
func (c *cSysAdminApi) SetRole(ctx *gin.Context) {
	var (
		err error

		req   sys_req.AdminSetRoleReq
		input sys_query.AdminSetRoleInput
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(req, &input)
	err = service.NewSysServiceGroup().AdminService.SetRole(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}
