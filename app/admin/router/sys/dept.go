package sys

import (
	"csf/app/admin/apis/sys"
	"csf/common/middleware"
	"github.com/gin-gonic/gin"
)

func registerDeptRouter(r *gin.RouterGroup) {
	api := sys.NewSysDeptApi()
	r1 := r.Group("/sys/dept").Use(middleware.JWTAuthMiddleware())
	{
		r1.POST("/add", api.Add)
		r1.POST("/edit", api.Edit)
		r1.POST("/delete", api.Delete)
		r1.POST("/delete_multi", api.DeleteMulti)
		r1.GET("/get_one", api.GetOne)
		r1.GET("/list", api.TreeList)
		r1.GET("/tree_list", api.TreeList)
	}
}
