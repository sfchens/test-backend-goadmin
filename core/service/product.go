package service

import (
	"csf/core/query/product_query"
	"github.com/gin-gonic/gin"
)

var localProductService productServiceGroup

func NewProductServiceGroup() productServiceGroup {
	return localProductService
}

type productServiceGroup struct {
	CategoryService iProductCategoryService
	RuleService     iProductRuleService
}

type (
	iProductCategoryService interface {
		AddOrEdit(ctx *gin.Context, input product_query.CategoryAddOrEditInput) (err error)
		List(ctx *gin.Context, input product_query.CategoryListInput) (out product_query.CategoryListOut, err error)
		DeleteBatch(ctx *gin.Context, input product_query.CategoryDeleteBatchInput) (err error)
	}

	iProductRuleService interface {
		Add(ctx *gin.Context, input product_query.RuleAddOrEditInput) (err error)
		List(ctx *gin.Context, input product_query.RuleListInput) (out product_query.RuleListOut, err error)
	}
)

func RegisterNewProductCategory(i iProductCategoryService) {
	localProductService.CategoryService = i
}

func RegisterNewProductRule(i iProductRuleService) {
	localProductService.RuleService = i
}
