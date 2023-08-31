package product

import (
	"csf/app/admin/apis/product"
	"csf/core/middleware"
	"github.com/gin-gonic/gin"
)

func registerCategoryRouter(r *gin.RouterGroup) {
	apis := product.NewCategoryApi()
	r1 := r.Group("/product/category").Use(middleware.JWTAuthMiddleware())
	{
		r1.POST("/add", apis.Add)
		r1.POST("/edit", apis.Edit)
		r1.GET("/list", apis.List)
		r1.POST("/delete_multi", apis.DeleteBatch)
	}
}
