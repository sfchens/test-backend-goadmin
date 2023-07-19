package sys_apis

import (
	"csf/app/admin/request/sys_request"
	"csf/app/admin/service/sys_service"
	"csf/common/mysql/model"
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
// @Param raw body     sys_request.AdminAddReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/admin/add [post]
func (c *cSysAdminApi) Add(ctx *gin.Context) {
	var (
		err error
		req sys_request.AdminAddReq
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

// Edit  编辑管理员
// @Summary 编辑管理员
// @Description 编辑管理员
// @Tags 管理员管理
// @Accept application/json
// @Produce application/json
// @Param raw body     sys_request.AdminEditReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/admin/edit [post]
func (c *cSysAdminApi) Edit(ctx *gin.Context) {

	var (
		err error
		req sys_request.AdminEditReq
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err = sys_service.NewSysAdminService(ctx).Edit(req)
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
	adminInfo.Roles = []string{"*"}
	response.SuccessWithData(ctx, adminInfo)
}
