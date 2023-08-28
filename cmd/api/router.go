package api

import (
	"csf/core/middleware"
	"csf/core/service"
	"csf/library/easy_config"
	"csf/library/easy_logger"
	"csf/library/global"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func initApiRouter() {
	//var r *gin.Engine
	r := global.GinEngine
	if r == nil {
		r = gin.New()
		global.GinEngine = r
	}
	//switch h.(type) {
	//case *gin.Engine:
	//	r = h.(*gin.Engine)
	//default:
	//	log.Fatal("not support other engine")
	//	os.Exit(-1)
	//}
	//if config.SslConfig.Enable {
	//	r.Use(handler.TlsHandler())
	//}
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("SESSIONID", store))
	r.Use(middleware.TraceKeyMiddleware(global.TraceIdKey))
	r.Use(middleware.CORSMiddleware())
	// 全局操作日志
	r.Use(middleware.OperateLogger(&easy_logger.LoggerConfig{
		LogObjKey: easy_logger.LogFileAppKey,
		LogZap:    true,
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

func initRegisterRouter() {
	if easy_config.Config.App.IsApiMysql {
		err := service.NewSysServiceGroup().ApiService.Refresh()
		if err != nil {
			println("初始化Api数据失败")
		}
	}
}
