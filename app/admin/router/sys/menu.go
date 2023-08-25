package sys

import (
	"csf/app/admin/apis/sys"
	"csf/common/middleware"
	"github.com/gin-gonic/gin"
)

func registerMenuRouter(r *gin.RouterGroup) {
	apis := sys.NewSysMenuApi()
	r1 := r.Group("/sys/menu").Use(middleware.JWTAuthMiddleware())
	{
		r1.GET("/tree_list", apis.TreeList)
		r1.GET("/list", apis.List)
		r1.POST("/add", apis.Add)
		r1.POST("/edit", apis.Edit)
	}
}
