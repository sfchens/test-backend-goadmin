package product

import (
	"csf/app/admin/request/product_req"
	"csf/core/query/product_query"
	"csf/core/service"
	"csf/library/response"
	"csf/utils"
	"github.com/gin-gonic/gin"
)

type cRuleApi struct{}

func NewRuleApi() *cRuleApi {
	return &cRuleApi{}
}

// Add  添加商品规格
// @Summary 添加商品规格
// @Description 添加商品规格
// @Tags 商品规格管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query  product_req.RuleAddOrEditReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/product/rule/add [post]
func (c *cRuleApi) Add(ctx *gin.Context) {
	var (
		err error
		req product_req.RuleAddOrEditReq
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	var input product_query.RuleAddOrEditInput
	utils.StructToStruct(req, &input)
	err = service.NewProductServiceGroup().RuleService.Add(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// List  商品规格列表
// @Summary 商品规格列表
// @Description 商品规格列表
// @Tags 商品规格管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query  product_req.RuleListReq true "请求参数"
// @Success 200 {object} response.Response{data=product_req.RuleListRes} "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/product/rule/list [get]
func (c *cRuleApi) List(ctx *gin.Context) {
	var (
		err error
		req product_req.RuleListReq

		res product_req.RuleListRes
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	var (
		input product_query.RuleListInput
		out   product_query.RuleListOut
	)
	utils.StructToStruct(req, &input)
	out, err = service.NewProductServiceGroup().RuleService.List(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	res = product_req.RuleListRes{
		Total: int(out.Total),
		List:  out.List,
	}
	response.SuccessWithData(ctx, res)
}
