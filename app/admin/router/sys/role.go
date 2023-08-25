package sys

import (
	"csf/app/admin/apis/sys"
	"csf/common/middleware"
	"github.com/gin-gonic/gin"
)

func registerRoleRouter(r *gin.RouterGroup) {
	apis := sys.NewSysRoleApi()
	r1 := r.Group("/sys/role").Use(middleware.JWTAuthMiddleware())
	{
		r1.POST("/add", apis.Add)
		r1.GET("/list", apis.List)
		r1.POST("/delete_batch", apis.DeleteBatch)
	}
}
