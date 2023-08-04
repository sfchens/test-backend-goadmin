package sys_apis

import (
	"csf/app/admin/request/sys_request"
	"csf/app/admin/service/sys_service"
	"csf/library/response"
	"csf/utils"
	"github.com/gin-gonic/gin"
)

type cSysMenuApi struct{}

func NewSysMenuApi() *cSysMenuApi {
	return &cSysMenuApi{}
}

// TreeList  菜单列表
// @Summary 菜单列表
// @Description 菜单列表
// @Tags 菜单管理
// @Accept application/json
// @Produce application/json
// @Param raw body     sys_request.MenuListReq true "请求参数"
// @Success 200 {object} response.Response{data=sys_request.MenuListRes} "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/menu/tree_list [get]
func (c *cSysMenuApi) TreeList(ctx *gin.Context) {

	var (
		err error
		req sys_request.MenuListReq
		res sys_request.MenuListRes
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	res, err = sys_service.NewSysMenuService(ctx).TreeList(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.SuccessWithData(ctx, res)
}

// TreeListAll  菜单列表
// @Summary 菜单列表
// @Description 菜单列表
// @Tags 菜单管理
// @Accept application/json
// @Produce application/json
// @Param raw body     sys_request.MenuListReq true "请求参数"
// @Success 200 {object} response.Response{data=sys_request.MenuListRes} "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/menu/tree_list_all [get]
func (c *cSysMenuApi) TreeListAll(ctx *gin.Context) {

	var (
		err error
		req sys_request.MenuListReq
		res sys_request.MenuListRes
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	res, err = sys_service.NewSysMenuService(ctx).TreeListAll(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.SuccessWithData(ctx, res)
}

// TreeRoleList  权限菜单
// @Summary 权限菜单
// @Description 权限菜单
// @Tags 菜单管理
// @Accept application/json
// @Produce application/json
// @Param raw body     sys_request.MenuListReq true "请求参数"
// @Success 200 {object} response.Response{data=sys_request.MenuListRes} "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/menu/tree_role_list [get]
func (c *cSysMenuApi) TreeRoleList(ctx *gin.Context) {

	var (
		err error
		req sys_request.MenuListReq
		res sys_request.MenuListRes
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	res, err = sys_service.NewSysMenuService(ctx).TreeList(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.SuccessWithData(ctx, res)
}

// Add  添加菜单
// @Summary 添加菜单
// @Description 添加菜单
// @Tags 菜单管理
// @Accept application/json
// @Produce application/json
// @Param raw body     sys_request.MenuAddOrEditReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/menu/add [post]
func (c *cSysMenuApi) Add(ctx *gin.Context) {
	var (
		err error
		req sys_request.MenuAddOrEditReq
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err = sys_service.NewSysMenuService(ctx).Add(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// Edit  编辑菜单
// @Summary 编辑菜单
// @Description 编辑菜单
// @Tags 菜单管理
// @Accept application/json
// @Produce application/json
// @Param raw body     sys_request.MenuAddOrEditReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/menu/edit [post]
func (c *cSysMenuApi) Edit(ctx *gin.Context) {
	var (
		err error
		req sys_request.MenuAddOrEditReq
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	err = sys_service.NewSysMenuService(ctx).Edit(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.Success(ctx)
}
