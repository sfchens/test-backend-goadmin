package product

import (
	"csf/app/admin/request/product_req"
	"csf/core/query/product_query"
	"csf/core/service"
	"csf/library/response"
	"csf/utils"
	"github.com/gin-gonic/gin"
)

type cCategoryApi struct {
}

func NewCategoryApi() *cCategoryApi {
	return &cCategoryApi{}
}

// Add  添加分类
// @Summary 添加分类
// @Description 添加分类
// @Tags 分类管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body  product_req.CategoryAddOrEditReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/product/category/add [post]
func (c *cCategoryApi) Add(ctx *gin.Context) {

	var (
		err error

		req   product_req.CategoryAddOrEditReq
		input product_query.CategoryAddOrEditInput
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(req, &input)
	err = service.NewProductServiceGroup().CategoryService.AddOrEdit(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.Success(ctx)
}

// Edit  编辑分类
// @Summary 编辑分类
// @Description 编辑分类
// @Tags 分类管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body  product_req.CategoryAddOrEditReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/product/category/edit [post]
func (c *cCategoryApi) Edit(ctx *gin.Context) {

	var (
		err error

		req   product_req.CategoryAddOrEditReq
		input product_query.CategoryAddOrEditInput
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(req, &input)
	err = service.NewProductServiceGroup().CategoryService.AddOrEdit(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.Success(ctx)
}

// List  分类列表
// @Summary 分类列表
// @Description 分类列表
// @Tags 分类管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query  product_req.CategoryListReq true "请求参数"
// @Success 200 {object} response.Response{data=product_req.CategoryListRes} "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/product/category/list [get]
func (c *cCategoryApi) List(ctx *gin.Context) {

	var (
		err error

		req product_req.CategoryListReq
		res product_req.CategoryListRes
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	var (
		input product_query.CategoryListInput
		out   product_query.CategoryListOut
	)
	utils.StructToStruct(req, &input)
	out, err = service.NewProductServiceGroup().CategoryService.List(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.SuccessWithStruct(ctx, out, &res)
}

// DeleteBatch  批量删除
// @Summary 批量删除
// @Description 批量删除
// @Tags 分类管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body  product_req.CategoryDeleteBatchReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/product/category/delete_multi [post]
func (c *cCategoryApi) DeleteBatch(ctx *gin.Context) {
	var (
		err error

		req   product_req.CategoryDeleteBatchReq
		input product_query.CategoryDeleteBatchInput
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	utils.StructToStruct(req, &input)
	err = service.NewProductServiceGroup().CategoryService.DeleteBatch(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.Success(ctx)
}
