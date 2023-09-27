package middleware

import (
	"bytes"
	"csf/library/easy_logger"
	"csf/library/response"
	"csf/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
)

// OperateLogger 记录日志
func OperateLogger(config *easy_logger.LoggerConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body, _ := ioutil.ReadAll(ctx.Request.Body)

		// 设置响应数据
		responseWriter := utils.NewResponseWriter(ctx)
		ctx.Writer = responseWriter
		ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
		// 处理请求
		ctx.Next()

		ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
		logObjKey := config.LogObjKey
		if logObjKey == "" {
			logObjKey = easy_logger.GetLogModulesName(ctx)
		}
		if config.LogZap {
			easy_logger.OperateLogger(ctx, logObjKey)
		}
		if config.LogZinc {

		}
	}
}

// RecoveryLogger 捕获错误
func RecoveryLogger(config *easy_logger.LoggerConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				zapLogObj := easy_logger.ZapLog(config.LogObjKey)
				httpRequest, _ := httputil.DumpRequest(ctx.Request, false)

				if isClient(err) {
					zapLogObj.Error(
						ctx.Request.URL.Path,
						zap.Any("error", err),
						zap.String("httpRequest", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					v, _ := err.(error)
					response.FailWithMessage(ctx, v.Error())
					ctx.Abort()
					return
				}

				// 这里可以选择全部打印出来不必要分割然后循环输出
				request := strings.Split(string(httpRequest), "\r\n")
				split := strings.Split(string(debug.Stack()), "\n\t")
				zapLogObj.Error("[Recovery from panic]", zap.Any("error", err))
				for _, str := range request {
					zapLogObj.Panic("[Recovery from request panic]", zap.String("request", str))
				}
				for _, str := range split {
					zapLogObj.Panic("[Recovery from Stack panic]", zap.String("stack", str))
				}
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		ctx.Next()
	}
}

func isClient(err any) bool {
	// 判断是否因客户端原因
	var brokenPipe bool
	if ne, ok := err.(*net.OpError); ok {
		if se, ok := ne.Err.(*os.SyscallError); ok {
			if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
				strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
				brokenPipe = true
			}
		}
	}
	return brokenPipe
}
