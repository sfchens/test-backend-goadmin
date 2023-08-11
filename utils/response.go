package utils

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"time"
)

var ResponseWriter *responseWriter

type responseWriter struct {
	gin.ResponseWriter
	Body      *bytes.Buffer
	StartTime time.Time
}

// 重写 Write 方法，以便记录响应数据
func (w responseWriter) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

func NewResponseWriter(ctx *gin.Context) *responseWriter {
	bodyRewrite := &responseWriter{
		ResponseWriter: ctx.Writer,
		Body:           bytes.NewBufferString(""), // 初始化响应数据
		StartTime:      time.Now(),
	}
	ResponseWriter = bodyRewrite
	return bodyRewrite
}
