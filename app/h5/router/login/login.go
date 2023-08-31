package login

import (
	"csf/app/h5/apis/login"
	"github.com/gin-gonic/gin"
)

func registerLoginRouter(r *gin.RouterGroup) {
	api := login.NewLoginApi()

	{
		r.POST("/login", api.Login)
		r.GET("/logout", api.Logout)
	}
}
