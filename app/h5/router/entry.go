package router

import (
	"csf/app/h5/router/login"
	"csf/app/h5/router/user"
	"csf/library/global"
	"os"
)

func InitRouter() {

	var r = global.GinEngine
	if r == nil {
		os.Exit(-1)
	}
	login.InitRouter()
	// 用户管理路由
	user.InitRouter()
	v1 := r.Group("/h5/v1")
	for _, f := range global.RouterList[global.ModuleH5] {
		f(v1)
	}
}
