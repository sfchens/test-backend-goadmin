package middleware

import (
	"csf/library/easy_logger"
	"csf/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

// OperateLogger 记录日志
func OperateLogger(config *easy_logger.LoggerConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 设置响应数据
		responseWriter := utils.NewResponseWriter(ctx)
		ctx.Writer = responseWriter

		// 处理请求
		ctx.Next()

		if config.LogZap {
			easy_logger.OperateLogger(ctx, config.LogObjKey)
		}
		if config.LogZinc {

		}
	}
}

// RecoveryLogger 捕获错误
func RecoveryLogger(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("232323")
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				//var brokenPipe bool
				//if ne, ok := err.(*net.OpError); ok {
				//	if se, ok := ne.Err.(*os.SyscallError); ok {
				//		if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
				//			strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
				//			brokenPipe = true
				//		}
				//	}
				//}
				//
				//httpRequest, _ := httputil.DumpRequest(c.Request, false)
				//if brokenPipe {
				//	zap.L().Error(c.Request.URL.Path,
				//		zap.Any("error", err),
				//		zap.String("httpRequest", string(httpRequest)),
				//	)
				//	// If the connection is dead, we can't write a status to it.
				//	c.Error(err.(error)) // nolint: errcheck
				//	c.Abort()
				//	return
				//}
				//// 这里可以选择全部打印出来不必要分割然后循环输出
				//request := strings.Split(string(httpRequest), "\r\n")
				//split := strings.Split(string(debug.Stack()), "\n\t")
				//if stack {
				//	zap.L().Error("[Recovery from panic]",
				//		zap.Any("error", err))
				//	for _, str := range request {
				//		zap.L().Error("[Recovery from request panic]", zap.String("request", str))
				//	}
				//	for _, str := range split {
				//		zap.L().Error("[Recovery from Stack panic]", zap.String("stack", str))
				//	}
				//} else {
				//	zap.L().Error("[Recovery from panic]",
				//		zap.Any("error", err))
				//	for _, str := range request {
				//		zap.L().Error("[Recovery from request panic]", zap.String("request", str))
				//	}
				//}
				//c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
