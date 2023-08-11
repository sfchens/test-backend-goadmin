package response

import (
	"csf/library/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code        int         `json:"code"`
	Msg         string      `json:"msg"`
	RedirectUrl string      `json:"redirect_url"`
	TraceId     string      `json:"trace_id"`
	Data        interface{} `json:"data"`
}

const (
	SUCCESS = 200
	ERROR   = http.StatusCreated
)

func Result(ctx *gin.Context, code int, data interface{}, msg, redirectUrl string) {
	traceId := ctx.Request.Header.Get(global.TraceIdKey)
	ctx.Set(global.ErrorLogKey, msg)
	// 开始时间
	ctx.JSON(code, Response{
		code,
		msg,
		redirectUrl,
		traceId,
		data,
	})
}

func Success(ctx *gin.Context) {
	Result(ctx, SUCCESS, map[string]interface{}{}, "操作成功", "")
}

func SuccessWithMessage(ctx *gin.Context, message string) {
	Result(ctx, SUCCESS, map[string]interface{}{}, message, "")
}

func SuccessWithData(ctx *gin.Context, data interface{}) {
	Result(ctx, SUCCESS, data, "操作成功", "")
}

func SuccessWithDetailed(ctx *gin.Context, data interface{}, message string) {
	Result(ctx, SUCCESS, data, message, "")
}

func Fail(ctx *gin.Context) {
	Result(ctx, ERROR, map[string]interface{}{}, "操作失败", "")
}

func FailWithMessage(ctx *gin.Context, message string) {
	Result(ctx, ERROR, map[string]interface{}{}, message, "")
}

func FailWithDetailed(ctx *gin.Context, data interface{}, message string) {
	Result(ctx, ERROR, data, message, "")
}

func SuccessRedirect(ctx *gin.Context, redirectUrl string) {
	Result(ctx, SUCCESS, map[string]interface{}{}, "操作成功", redirectUrl)
}

func SuccessRedirectWithMessage(ctx *gin.Context, message, redirectUrl string) {
	Result(ctx, SUCCESS, map[string]interface{}{}, message, redirectUrl)
}
