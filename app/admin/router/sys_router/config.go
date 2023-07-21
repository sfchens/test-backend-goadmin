package sys_router

import (
	"csf/app/admin/apis/sys_apis"
	"csf/common/middleware"
	"github.com/gin-gonic/gin"
)

func registerConfigRouter(r *gin.RouterGroup) {

	api := sys_apis.NewSysConfigApi()

	r1 := r.Group("/sys/config").Use(middleware.JWTAuthMiddleware())
	{
		r1.POST("/add", api.Add)
		r1.POST("/edit", api.Edit)
		r1.POST("/delete", api.Delete)
		r1.POST("/set_status", api.SetStatus)
		r1.GET("/list", api.List)
		r1.GET("/get_one", api.GetOne)
	}
}
