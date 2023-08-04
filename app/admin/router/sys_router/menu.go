package sys_router

import (
	"csf/app/admin/apis/sys_apis"
	"csf/common/middleware"
	"github.com/gin-gonic/gin"
)

func registerMenuRouter(r *gin.RouterGroup) {
	apis := sys_apis.NewSysMenuApi()
	r1 := r.Group("/sys/menu").Use(middleware.JWTAuthMiddleware())
	{
		r1.GET("/tree_list", apis.TreeList)
		r1.GET("/tree_role_list", apis.TreeRoleList)
		r1.GET("/tree_list_all", apis.TreeListAll)
		r1.POST("/add", apis.Add)
		r1.POST("/edit", apis.Edit)
	}
}
