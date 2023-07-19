package api

import (
	"csf/app/admin/service/sys_service"
	"csf/common/middleware"
	"csf/library/global"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var EngineRouter *gin.Engine

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
	r.Use(sessions.Sessions("mysession", store))
	r.Use(middleware.TraceKeyMiddleware(global.TraceIdKey))
	r.Use(middleware.CORSMiddleware())
	//r.Use(common.Sentinel()).
	//	Use(common.RequestId(pkg.TrafficKey)).
	//	Use(api.SetRequestLogger)
	//
	//common.InitMiddleware(r)
}

func initRegisterRouter() {

	ctx := &gin.Context{}
	err := sys_service.NewSysApiService(ctx).Refresh()
	if err != nil {
		println("初始化Api数据失败")
	}
}
