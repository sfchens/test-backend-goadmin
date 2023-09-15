package product

import (
	"csf/app/admin/apis/product"
	"csf/core/middleware"
	"github.com/gin-gonic/gin"
)

func registerRuleRouter(r *gin.RouterGroup) {
	apis := product.NewRuleApi()
	r1 := r.Group("/product/rule").Use(middleware.JWTAuthMiddleware())
	{
		r1.POST("/add", apis.Add)
		r1.GET("/list", apis.List)
	}
}
