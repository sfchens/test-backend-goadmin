package router

import (
	"csf/app/admin/router/common"
	"csf/app/admin/router/live"
	"csf/app/admin/router/sys"
	"csf/app/admin/router/user"
	"csf/library/global"
	"os"
)

func InitRouter() {

	var r = global.GinEngine
	if r == nil {
		os.Exit(-1)
	}

	// 加载公共路由
	common.InitRouter()
	// 加载路由
	sys.InitRouter()
	// 直播路由
	live.InitRouter()
	// 用户管理路由
	user.InitRouter()

	v1 := r.Group("/api/v1")
	for _, f := range global.RouterList {
		f(v1)
	}
}
