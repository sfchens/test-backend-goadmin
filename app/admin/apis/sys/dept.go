package sys

import (
	"csf/app/admin/request/sys_req"
	"csf/app/admin/service/sys_service"
	"csf/library/response"
	"csf/utils"
	"fmt"
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
// @Param raw body     sys.DeptAddOrEditReq true "请求参数"
// @Success 200 {object} response.Response  "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/dept/add [post]
func (c *cSysDeptApi) Add(ctx *gin.Context) {
	var (
		err error
		req sys_req.DeptAddOrEditReq
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	fmt.Printf("req:  %+v\n", req)
	err = sys_service.NewSysDeptService(ctx).AddOrEdit(req)
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
// @Param raw body     sys.DeptAddOrEditReq true "请求参数"
// @Success 200 {object} response.Response  "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/dept/edit [post]
func (c *cSysDeptApi) Edit(ctx *gin.Context) {
	var (
		err error
		req sys_req.DeptAddOrEditReq
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	err = sys_service.NewSysDeptService(ctx).Edit(req)
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
// @Param raw body     sys.DeptDeleteReq true "请求参数"
// @Success 200 {object} response.Response  "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/dept/delete [post]
func (c *cSysDeptApi) Delete(ctx *gin.Context) {
	var (
		err error
		req sys_req.DeptDeleteReq
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	err = sys_service.NewSysDeptService(ctx).Delete(req)
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
// @Param raw body     sys.DeptTreeListReq true "请求参数"
// @Success 200 {object} response.Response{data=sys.DeptTreeListRes}  "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/dept/list [get]
func (c *cSysDeptApi) TreeList(ctx *gin.Context) {
	var (
		err error
		req sys_req.DeptTreeListReq
		res sys_req.DeptTreeListRes
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	res, err = sys_service.NewSysDeptService(ctx).TreeList(req)
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
// @Param raw body     sys.DeptTreeListReq true "请求参数"
// @Success 200 {object} response.Response{data=sys.DeptTreeListRes}  "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/dept/get_one [get]
func (c *cSysDeptApi) GetOne(ctx *gin.Context) {
	var (
		err error
		req sys_req.DeptGetOneReq
		res sys_req.DeptGetOneRes
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	res, err = sys_service.NewSysDeptService(ctx).GetOne(req)
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
// @Param raw body     sys.DeptTreeListReq true "请求参数"
// @Success 200 {object} response.Response{data=sys.DeptTreeListRes}  "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/dept/delete_multi [post]
func (c *cSysDeptApi) DeleteMulti(ctx *gin.Context) {
	var (
		err error
		req sys_req.DeptDeleteMultiReq
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	err = sys_service.NewSysDeptService(ctx).DeleteMulti(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}