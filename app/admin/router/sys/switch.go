package sys

import (
	"csf/app/admin/apis/sys"
	"csf/common/middleware"
	"github.com/gin-gonic/gin"
)

func registerSwitchRouter(r *gin.RouterGroup) {
	apis := sys.NewSwitchApi()
	r1 := r.Group("/sys/switch").Use(middleware.JWTAuthMiddleware())
	{
		r1.GET("/list", apis.List)
		r1.POST("/add", apis.Add)
		r1.POST("/edit", apis.Edit)
		r1.POST("/delete", apis.Delete)
		r1.POST("/set_status", apis.SetStatus)
	}
}
