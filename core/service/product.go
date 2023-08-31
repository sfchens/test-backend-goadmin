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
	CategoryService iCategoryService
}

type (
	iCategoryService interface {
		AddOrEdit(ctx *gin.Context, input product_query.CategoryAddOrEditInput) (err error)
		List(ctx *gin.Context, input product_query.CategoryListInput) (out product_query.CategoryListOut, err error)
		DeleteBatch(ctx *gin.Context, input product_query.CategoryDeleteBatchInput) (err error)
	}
)

func RegisterNewCategory(i iCategoryService) {
	localProductService.CategoryService = i
}
