package sys_router

import (
	"csf/app/admin/apis/sys_apis"
	"csf/common/middleware"
	"github.com/gin-gonic/gin"
)

func registerApiRouter(r *gin.RouterGroup) {
	apis := sys_apis.NewSysApi()
	r1 := r.Group("/sys/api").Use(middleware.JWTAuthMiddleware())
	{
		r1.GET("/list", apis.List)
		r1.POST("/refresh", apis.Refresh)
		r1.POST("/edit", apis.Edit)
	}
}
