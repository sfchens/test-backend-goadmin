package sys_router

import (
	"csf/app/admin/apis/sys_apis"
	"csf/common/middleware"
	"github.com/gin-gonic/gin"
)

func registerSwitchRouter(r *gin.RouterGroup) {
	apis := sys_apis.NewSwitchApi()
	r1 := r.Group("/sys/switch").Use(middleware.JWTAuthMiddleware())
	{
		r1.GET("/list", apis.List)
		r1.POST("/add", apis.Add)
		r1.POST("/edit", apis.Edit)
	}
}
