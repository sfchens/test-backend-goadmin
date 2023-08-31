package user

import (
	"csf/app/admin/apis/user"
	"csf/core/middleware"
	"github.com/gin-gonic/gin"
)

func registerUserRouter(r *gin.RouterGroup) {
	api := user.NewUserApi()
	r1 := r.Group("/user").Use(middleware.JWTAuthMiddleware())
	{
		r1.POST("/add", api.Add)
		r1.GET("/list", api.List)
		r1.GET("/get_info", api.GetInfo)
	}
}
