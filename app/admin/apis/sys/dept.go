package sys

import (
	"csf/app/admin/request/sys_req"
	"csf/core/query/sys_query"
	"csf/core/service"
	"csf/library/response"
	"csf/utils"
	"github.com/gin-gonic/gin"
)

type cSysDeptApi struct{}

func NewSysDeptApi() *cSysDeptApi {
	return &cSysDeptApi{}
}

// Add  添加部门
// @Summary 添加部门
// @Description 添加部门
// @Tags 部门管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body     sys_req.DeptAddOrEditReq true "请求参数"
// @Success 200 {object} response.Response  "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/dept/add [post]
func (c *cSysDeptApi) Add(ctx *gin.Context) {
	var (
		err   error
		req   sys_req.DeptAddOrEditReq
		input sys_query.DeptAddOrEditInput
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(req, &input)
	err = service.NewSysServiceGroup().DeptService.AddOrEdit(ctx, input)
	if err != nil {
		response.SuccessWithData(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// Edit  编辑部门
// @Summary 编辑部门
// @Description 编辑部门
// @Tags 部门管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body     sys_req.DeptAddOrEditReq true "请求参数"
// @Success 200 {object} response.Response  "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/dept/edit [post]
func (c *cSysDeptApi) Edit(ctx *gin.Context) {
	var (
		err   error
		req   sys_req.DeptAddOrEditReq
		input sys_query.DeptAddOrEditInput
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(req, &input)
	err = service.NewSysServiceGroup().DeptService.Edit(ctx, input)
	if err != nil {
		response.SuccessWithData(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// Delete  删除部门
// @Summary 删除部门
// @Description 删除部门
// @Tags 部门管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body     sys_req.DeptDeleteReq true "请求参数"
// @Success 200 {object} response.Response  "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/dept/delete [post]
func (c *cSysDeptApi) Delete(ctx *gin.Context) {
	var (
		err error
		req sys_req.DeptDeleteReq

		input sys_query.DeptDeleteInput
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(req, &input)
	err = service.NewSysServiceGroup().DeptService.Delete(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// TreeList  部门tree列表
// @Summary 部门tree列表
// @Description 部门tree列表
// @Tags 部门管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query     sys_req.DeptTreeListReq true "请求参数"
// @Success 200 {object} response.Response{data=sys_query.DeptTreeListOut}  "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/dept/list [get]
func (c *cSysDeptApi) TreeList(ctx *gin.Context) {
	var (
		err error
		req sys_req.DeptTreeListReq

		input sys_query.DeptTreeListInput
		res   sys_query.DeptTreeListOut
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(req, &input)
	res, err = service.NewSysServiceGroup().DeptService.TreeList(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.SuccessWithData(ctx, res)
}

// GetOne  一条部门信息
// @Summary 一条部门信息
// @Description 一条部门信息
// @Tags 部门管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query     sys_req.DeptGetOneReq true "请求参数"
// @Success 200 {object} response.Response{data=sys_query.DeptGetOneOut}  "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/dept/get_one [get]
func (c *cSysDeptApi) GetOne(ctx *gin.Context) {
	var (
		err error
		req sys_req.DeptGetOneReq

		input sys_query.DeptGetOneInput
		res   sys_query.DeptGetOneOut
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(req, &input)
	res, err = service.NewSysServiceGroup().DeptService.GetOne(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.SuccessWithData(ctx, res)
}

// DeleteMulti  批量删除
// @Summary 批量删除
// @Description 批量删除
// @Tags 部门管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body     sys_req.DeptDeleteMultiReq true "请求参数"
// @Success 200 {object} response.Response  "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/dept/delete_multi [post]
func (c *cSysDeptApi) DeleteMulti(ctx *gin.Context) {
	var (
		err error
		req sys_req.DeptDeleteMultiReq

		input sys_query.DeptDeleteMultiInput
	)

	err = utils.BindParams(ctx, &req, &input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err = service.NewSysServiceGroup().DeptService.DeleteMulti(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}
