package sys_router

import (
	"csf/app/admin/apis/sys_apis"
	"csf/common/middleware"
	"github.com/gin-gonic/gin"
)

func registerLoginRouter(r *gin.RouterGroup) {
	loginApi := sys_apis.NewSysLogin()
	r1 := r.Group("/sys")
	{
		r1.POST("/login", loginApi.Login)
		r1.GET("/login_info", loginApi.LoginInfo)
	}

	r2 := r.Group("/sys").Use(middleware.JWTAuthMiddleware())
	{
		r2.POST("/logout", loginApi.Logout)
	}
}
