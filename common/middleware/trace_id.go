package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func TraceKeyMiddleware(trafficKey string) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == http.MethodOptions {
			ctx.Next()
			return
		}

		requestId := uuid.New().String()
		ctx.Request.Header.Set(trafficKey, requestId)
		ctx.Next()
	}
}
