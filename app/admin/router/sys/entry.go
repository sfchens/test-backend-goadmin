package sys

import (
	"csf/library/global"
	"fmt"
)

func InitRouter() {
	fmt.Println("初始化 sys_router路由")
	global.RouterList[global.ModuleAdmin] = append(global.RouterList[global.ModuleAdmin],
		registerLoginRouter,
		registerConfigRouter,
		registerAdminRouter,
		registerMenuRouter,
		registerDeptRouter,
		registerSwitchRouter,
		registerApiRouter,
		registerRoleRouter,
	)
}
