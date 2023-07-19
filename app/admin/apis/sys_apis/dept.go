package sys_apis

import (
	"csf/app/admin/request/sys_request"
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
// @Param raw body     sys_request.DeptAddOrEditReq true "请求参数"
// @Success 200 {object} response.Response  "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/dept/add [post]
func (c *cSysDeptApi) Add(ctx *gin.Context) {
	var (
		err error
		req sys_request.DeptAddOrEditReq
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	fmt.Printf("req:  %+v\n", req)
	err = sys_service.NewSysDeptService(ctx).Add(req)
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
// @Param raw body     sys_request.DeptAddOrEditReq true "请求参数"
// @Success 200 {object} response.Response  "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/dept/edit [post]
func (c *cSysDeptApi) Edit(ctx *gin.Context) {
	var (
		err error
		req sys_request.DeptAddOrEditReq
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
// @Param raw body     sys_request.DeptDeleteReq true "请求参数"
// @Success 200 {object} response.Response  "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/dept/delete [post]
func (c *cSysDeptApi) Delete(ctx *gin.Context) {
	var (
		err error
		req sys_request.DeptDeleteReq
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

// List  部门列表
// @Summary 部门列表
// @Description 部门列表
// @Tags 部门管理
// @Accept application/json
// @Produce application/json
// @Param raw body     sys_request.DeptListReq true "请求参数"
// @Success 200 {object} response.Response{data=sys_request.DeptListRes}  "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/dept/list [get]
func (c *cSysDeptApi) List(ctx *gin.Context) {
	var (
		err error
		req sys_request.DeptListReq
		res sys_request.DeptListRes
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	res, err = sys_service.NewSysDeptService(ctx).List(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.SuccessWithData(ctx, res)
}
