package middleware

import (
	"csf/library/easy_auth"
	"csf/library/easy_session"
	"csf/library/global"
	"csf/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(ctx *gin.Context) {
		token := utils.GetToken(ctx)
		if token == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 499,
				"msg":  "非法请求",
			})
			ctx.Abort()
			return
		}
		mc, err := easy_auth.NewJWT().ParseToken(token)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 499,
				"msg":  "登录已过期",
			})
			ctx.Abort()
			return
		}
		err = easy_session.NewCustomSession(ctx).Set(global.UserInfoKey, mc)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 499,
				"msg":  "token异常",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
