package sys

import (
	"csf/app/admin/request/sys_req"
	"csf/core/query/sys_query"
	"csf/core/service"
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
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query     sys_req.MenuTreeListReq true "请求参数"
// @Success 200 {object} response.Response{data=sys_query.MenuTreeListOut} "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/menu/tree_list [get]
func (c *cSysMenuApi) TreeList(ctx *gin.Context) {

	var (
		err error
		req sys_req.MenuTreeListReq

		input sys_query.MenuTreeListInput
		res   sys_query.MenuTreeListOut
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(req, &input)
	res, err = service.NewSysServiceGroup().MenuService.TreeList(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.SuccessWithData(ctx, res)
}

// List  菜单列表
// @Summary 菜单列表
// @Description 菜单列表
// @Tags 菜单管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query     sys_req.MenuListReq true "请求参数"
// @Success 200 {object} response.Response{data=sys_query.MenuListOut} "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/menu/list [get]
func (c *cSysMenuApi) List(ctx *gin.Context) {

	var (
		err error
		req sys_req.MenuListReq

		input sys_query.MenuListInput
		res   sys_query.MenuListOut
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(req, &input)

	res, err = service.NewSysServiceGroup().MenuService.List(ctx, input)
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
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body     sys_req.MenuAddOrEditReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/menu/add [post]
func (c *cSysMenuApi) Add(ctx *gin.Context) {
	var (
		err   error
		req   sys_req.MenuAddOrEditReq
		input sys_query.MenuAddOrEditInput
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(req, &input)
	err = service.NewSysServiceGroup().MenuService.Add(ctx, input)
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
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body     sys_req.MenuAddOrEditReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/menu/edit [post]
func (c *cSysMenuApi) Edit(ctx *gin.Context) {
	var (
		err   error
		req   sys_req.MenuAddOrEditReq
		input sys_query.MenuAddOrEditInput
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(req, &input)
	err = service.NewSysServiceGroup().MenuService.Edit(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.Success(ctx)
}
