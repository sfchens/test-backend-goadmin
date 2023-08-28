package sys

import (
	"csf/app/admin/request/sys_req"
	"csf/app/admin/service/sys_service"
	"csf/library/response"
	"csf/utils"
	"github.com/gin-gonic/gin"
)

type cSysRoleApi struct{}

func NewSysRoleApi() *cSysRoleApi {
	return &cSysRoleApi{}
}

// Add  添加角色
// @Summary 添加角色
// @Description 添加角色
// @Tags 角色管理
// @Accept application/json
// @Produce application/json
// @Param raw body     sys.RoleAddOrEditReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/role/add [post]
func (c *cSysRoleApi) Add(ctx *gin.Context) {
	var (
		err error

		req sys_req.RoleAddOrEditReq
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	err = sys_service.NewSysRoleService(ctx).AddOrEdit(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.Success(ctx)
}

// List  角色列表
// @Summary 角色列表
// @Description 角色列表
// @Tags 角色管理
// @Accept application/json
// @Produce application/json
// @Param raw body     sys.RoleListReq true "请求参数"
// @Success 200 {object} response.Response{data=sys.RoleListRes} "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/role/list [get]
func (c *cSysRoleApi) List(ctx *gin.Context) {
	var (
		err error

		req sys_req.RoleListReq
		res sys_req.RoleListRes
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	res, err = sys_service.NewSysRoleService(ctx).List(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.SuccessWithData(ctx, res)
}

// DeleteBatch  批量删除角色
// @Summary 批量删除角色
// @Description 批量删除角色
// @Tags 角色管理
// @Accept application/json
// @Produce application/json
// @Param raw body     sys.RoleDeleteBatchReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/role/delete_batch [post]
func (c *cSysRoleApi) DeleteBatch(ctx *gin.Context) {
	var (
		err error

		req sys_req.RoleDeleteBatchReq
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	err = sys_service.NewSysRoleService(ctx).DeleteBatch(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.Success(ctx)
}

// Delete  删除角色
// @Summary 删除角色
// @Description 删除角色
// @Tags 角色管理
// @Accept application/json
// @Produce application/json
// @Param raw body     sys.RoleDeleteReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/role/delete[post]
func (c *cSysRoleApi) Delete(ctx *gin.Context) {
	var (
		err error

		req sys_req.RoleDeleteReq
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	err = sys_service.NewSysRoleService(ctx).Delete(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.Success(ctx)
}