package sys_router

import (
	"csf/app/admin/apis/sys_apis"
	"csf/common/middleware"
	"github.com/gin-gonic/gin"
)

func registerAdminRouter(r *gin.RouterGroup) {

	api := sys_apis.NewSysAdminApi()

	r1 := r.Group("/sys/admin").Use(middleware.JWTAuthMiddleware())
	{
		r1.POST("/add", api.Add)
		r1.POST("/edit", api.Edit)
		r1.GET("/list", api.List)
		r1.GET("/get_admin_info", api.GetAdminInfo)
	}
}
