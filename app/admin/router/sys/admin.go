package sys

import (
	"csf/app/admin/apis/sys"
	"csf/common/middleware"
	"github.com/gin-gonic/gin"
)

func registerAdminRouter(r *gin.RouterGroup) {

	api := sys.NewSysAdminApi()

	r1 := r.Group("/sys/admin").Use(middleware.JWTAuthMiddleware())
	{
		r1.POST("/add", api.Add)
		r1.GET("/list", api.List)
		r1.POST("/reset_pwd", api.ResetPwd)
		r1.POST("/set_status", api.SetStatus)
		r1.POST("/set_role", api.SetRole)
		r1.POST("/delete_batch", api.DeleteBatch)
		r1.GET("/get_admin_info", api.GetAdminInfo)
	}
}
