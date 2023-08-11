package middleware

import (
	"csf/library/easy_auth"
	"csf/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(ctx *gin.Context) {
		//loginTypeTmp := custom_session.NewCustomSession(ctx).Get(global.LoginTypeKey)
		//loginType, ok := loginTypeTmp.(string)
		//if !ok || loginType == "" {
		//	ctx.JSON(http.StatusOK, gin.H{
		//		"code": 499,
		//		"msg":  "登录已过期",
		//	})
		//	ctx.Abort()
		//	return
		//}

		token := utils.GetAuthorization(ctx)
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
		// 将当前请求的username信息保存到请求的上下文c上
		ctx.Set("username", mc.Username)
		ctx.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
