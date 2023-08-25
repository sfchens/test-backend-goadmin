package user

import (
	"csf/app/admin/apis/user"
	"csf/common/middleware"
	"github.com/gin-gonic/gin"
)

func registerUserRouter(r *gin.RouterGroup) {
	api := user.NewUserApi()
	r1 := r.Group("/user").Use(middleware.JWTAuthMiddleware())
	{
		r1.POST("/add", api.Add)
		r1.GET("/list", api.List)
	}
}
