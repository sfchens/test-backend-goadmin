package api

import (
	"csf/core/middleware"
	"csf/library/easy_logger"
	"csf/library/global"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func initApiRouter() {
	r := global.GinEngine
	if r == nil {
		r = gin.New()
		global.GinEngine = r
	}
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("SESSIONID", store))
	r.Use(middleware.TraceKeyMiddleware(global.TraceIdKey))
	r.Use(middleware.CORSMiddleware())
	// 全局操作日志
	r.Use(middleware.OperateLogger(&easy_logger.LoggerConfig{
		LogZap: true,
	}))
	r.Use(middleware.RecoveryLogger(&easy_logger.LoggerConfig{
		LogObjKey: easy_logger.LogFileAppKey,
		LogZap:    true,
	}))

	//r.Use(common.Sentinel()).
	//	Use(common.RequestId(pkg.TrafficKey)).
	//	Use(api.SetRequestLogger)
	//
	//common.InitMiddleware(r)
}
