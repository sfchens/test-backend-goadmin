package sys

import (
	"csf/app/admin/apis/sys"
	"csf/core/middleware"
	"github.com/gin-gonic/gin"
)

func registerApiRouter(r *gin.RouterGroup) {
	apis := sys.NewSysApi()
	r1 := r.Group("/sys/api").
		Use(middleware.JWTAuthMiddleware())
	{
		r1.GET("/list", apis.List)
		r1.POST("/refresh", apis.Refresh)
		r1.POST("/edit", apis.Edit)
		r1.POST("/delete_multi", apis.DeleteMulti)
		r1.GET("/get_tag", apis.GetTag)
	}
}
