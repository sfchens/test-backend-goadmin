package router

import (
	"csf/app/admin/router/common_router"
	"csf/app/admin/router/live_router"
	"csf/app/admin/router/sys_router"
	"csf/library/global"
	"os"
)

func InitRouter() {

	var r = global.GinEngine
	if r == nil {
		os.Exit(-1)
	}

	// 加载公共路由
	common_router.InitRouter()
	// 加载路由
	sys_router.InitRouter()
	// 直播路由
	live_router.InitRouter()

	v1 := r.Group("/api/v1")
	for _, f := range global.RouterList {
		f(v1)
	}
}
