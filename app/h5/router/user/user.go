package user

import (
	"csf/app/h5/apis/user"
	"csf/core/middleware"
	"github.com/gin-gonic/gin"
)

func registerUserRouter(r *gin.RouterGroup) {
	api := user.NewUserApi()
	r1 := r.Group("/user")
	{
		r1.POST("/register", api.Register)
	}

	r2 := r1.Use(middleware.JWTAuthMiddleware())
	{
		r2.GET("/get_info", api.GetInfo)
	}
}
